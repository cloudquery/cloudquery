package client

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	rdsauth "github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// assumeRoleDuration is how long credentials obtained by assuming
// `aws_iam_auth.role_arn` stay valid before they are refreshed. Every role allows
// for at least one hour.
const assumeRoleDuration = time.Hour

// awsIAMAuthTarget is the connection an authentication token is signed for. AWS
// services differ in what they sign: Amazon RDS binds the token to the host, port
// and database user, while Aurora DSQL binds it to the host and picks the token
// type from the user.
type awsIAMAuthTarget struct {
	Host   string
	Port   uint16
	User   string
	Region string
}

// awsIAMAuthProvider describes how to authenticate to one AWS-managed,
// PostgreSQL-compatible service with IAM credentials.
type awsIAMAuthProvider struct {
	// minTLS is the weakest TLS configuration the service accepts.
	minTLS tlsRequirement
	// signToken mints a short-lived authentication token for a single connection.
	// Signing is a local SigV4 operation, so it adds no network round trip.
	signToken func(ctx context.Context, target awsIAMAuthTarget, creds aws.CredentialsProvider) (string, error)
}

// awsIAMAuthProviders holds one entry per supported `aws_iam_auth.service`. Only
// token signing and the TLS requirement are service-specific; region resolution,
// credential resolution and the connection plumbing are shared. Supporting Aurora
// DSQL, for example, means adding its `spec.AWSIAMAuthService` value plus an entry
// here that signs with `feature/dsql/auth` and requires tlsVerified.
var awsIAMAuthProviders = map[spec.AWSIAMAuthService]awsIAMAuthProvider{
	spec.AWSIAMAuthServiceRDS: {
		minTLS:    tlsEncrypted,
		signToken: signRDSAuthToken,
	},
}

// configureAWSIAMAuth wires up AWS IAM database authentication on the given pool
// config. It resolves the AWS credentials once and installs a BeforeConnect
// callback that signs a fresh short-lived authentication token and uses it as the
// connection password for every new connection.
func configureAWSIAMAuth(ctx context.Context, pgxConfig *pgxpool.Config, authSpec *spec.AWSIAMAuthSpec) error {
	service := authSpec.ServiceOrDefault()
	provider, ok := awsIAMAuthProviders[service]
	if !ok {
		return fmt.Errorf("unsupported `aws_iam_auth.service` %q", service)
	}

	// IAM database authentication requires TLS, and the token is a signed
	// credential, so reject any connection string that could connect without it.
	// Fail fast with a clear error rather than leaking the token over plaintext or
	// failing later in a less obvious way.
	if !checkTLS(pgxConfig.ConnConfig, provider.minTLS) {
		return fmt.Errorf("aws_iam_auth with `service: %s` requires a TLS connection: set %s in `connection_string`", service, provider.minTLS.hint())
	}

	// Parse the endpoint override once, so that a malformed value fails here
	// instead of on every connection attempt.
	endpointHost, endpointPort, err := parseAWSIAMAuthEndpoint(authSpec.Endpoint)
	if err != nil {
		return err
	}

	awsCfg, err := loadAWSIAMAuthConfig(ctx, authSpec)
	if err != nil {
		return err
	}

	// Preserve any previously configured BeforeConnect hook and run it first, so
	// IAM auth composes with other hooks instead of discarding them. The token is
	// set last as it must be the connection password.
	prevBeforeConnect := pgxConfig.BeforeConnect
	pgxConfig.BeforeConnect = func(ctx context.Context, connConfig *pgx.ConnConfig) error {
		if prevBeforeConnect != nil {
			if err := prevBeforeConnect(ctx, connConfig); err != nil {
				return err
			}
		}
		// Sign for the connection actually being made - pgx may be connecting via a
		// fallback - because the service rejects a token that was signed for a
		// different endpoint or user.
		target := awsIAMAuthTarget{
			Host:   connConfig.Host,
			Port:   connConfig.Port,
			User:   connConfig.User,
			Region: awsCfg.Region,
		}
		if endpointHost != "" {
			target.Host = endpointHost
		}
		if endpointPort != 0 {
			target.Port = endpointPort
		}
		token, err := provider.signToken(ctx, target, awsCfg.Credentials)
		if err != nil {
			return fmt.Errorf("failed to sign %s iam authentication token: %w", service, err)
		}
		connConfig.Password = token
		return nil
	}

	// Note that, unlike an OAuth credential, an IAM authentication token is only
	// used to authenticate: a session stays valid once established, even after the
	// token's lifetime has passed, so pooled connections need not be recycled early.
	return nil
}

// signRDSAuthToken signs a token for Amazon RDS and Aurora PostgreSQL. RDS
// validates the token against the endpoint being dialed and the user in the
// startup message, so both are part of the signature.
func signRDSAuthToken(ctx context.Context, target awsIAMAuthTarget, creds aws.CredentialsProvider) (string, error) {
	endpoint := net.JoinHostPort(target.Host, strconv.Itoa(int(target.Port)))
	return rdsauth.BuildAuthToken(ctx, endpoint, target.Region, target.User, creds)
}

// parseAWSIAMAuthEndpoint splits an `aws_iam_auth.endpoint` override into its host
// and port. The port is optional and reported as 0 when absent, in which case the
// port from `connection_string` is used.
func parseAWSIAMAuthEndpoint(endpoint string) (host string, port uint16, err error) {
	if endpoint == "" {
		return "", 0, nil
	}

	host, portStr, splitErr := net.SplitHostPort(endpoint)
	if splitErr != nil {
		// Both RDS and DSQL endpoints are hostnames, so a value without a port is
		// the host itself. Anything else containing a colon is malformed.
		if strings.Contains(endpoint, ":") {
			return "", 0, fmt.Errorf("invalid `aws_iam_auth.endpoint` %q: %w", endpoint, splitErr)
		}
		return endpoint, 0, nil
	}

	// Port 0 is not a valid destination port, and reporting it would be
	// indistinguishable from an endpoint without a port.
	parsedPort, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil || parsedPort == 0 {
		return "", 0, fmt.Errorf("invalid port %q in `aws_iam_auth.endpoint` %q", portStr, endpoint)
	}
	return host, uint16(parsedPort), nil
}

// loadAWSIAMAuthConfig resolves the AWS config used to sign authentication tokens.
// Credentials come from the standard AWS sources, optionally narrowed to a shared
// config profile or exchanged for assumed-role credentials.
func loadAWSIAMAuthConfig(ctx context.Context, authSpec *spec.AWSIAMAuthSpec) (aws.Config, error) {
	loadOpts := make([]func(*config.LoadOptions) error, 0, 2)
	if authSpec.Region != "" {
		loadOpts = append(loadOpts, config.WithRegion(authSpec.Region))
	}
	if authSpec.LocalProfile != "" {
		loadOpts = append(loadOpts, config.WithSharedConfigProfile(authSpec.LocalProfile))
	}

	awsCfg, err := config.LoadDefaultConfig(ctx, loadOpts...)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load aws config: %w", err)
	}
	// The token is signed for a region, so an unresolved one would only surface as
	// an opaque authentication failure from the database.
	if awsCfg.Region == "" {
		return aws.Config{}, errors.New("failed to resolve the aws region: set `aws_iam_auth.region` or the `AWS_REGION` environment variable")
	}

	if authSpec.RoleARN != "" {
		assumeRoleOpts := []func(*stscreds.AssumeRoleOptions){
			func(o *stscreds.AssumeRoleOptions) { o.Duration = assumeRoleDuration },
		}
		if authSpec.RoleSessionName != "" {
			assumeRoleOpts = append(assumeRoleOpts, func(o *stscreds.AssumeRoleOptions) {
				o.RoleSessionName = authSpec.RoleSessionName
			})
		}
		if authSpec.ExternalID != "" {
			assumeRoleOpts = append(assumeRoleOpts, func(o *stscreds.AssumeRoleOptions) {
				o.ExternalID = &authSpec.ExternalID
			})
		}
		provider := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(awsCfg), authSpec.RoleARN, assumeRoleOpts...)
		// Cache the assumed-role credentials so that they are reused across
		// connections instead of calling STS for every token.
		awsCfg.Credentials = aws.NewCredentialsCache(provider)
	}

	return awsCfg, nil
}
