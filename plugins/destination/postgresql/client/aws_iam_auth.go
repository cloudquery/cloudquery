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

// Every role allows for at least one hour.
const assumeRoleDuration = time.Hour

// awsIAMAuthTarget is what a token is signed for. RDS binds the token to the host,
// port and user; DSQL binds it to the host and picks the token type from the user.
type awsIAMAuthTarget struct {
	Host   string
	Port   uint16
	User   string
	Region string
}

type awsIAMAuthProvider struct {
	minTLS    tlsRequirement
	signToken func(ctx context.Context, target awsIAMAuthTarget, creds aws.CredentialsProvider) (string, error)
}

// Signing and the TLS requirement are the only service-specific parts; region and
// credential resolution are shared. Aurora DSQL would sign with `feature/dsql/auth`
// and require tlsVerified.
var awsIAMAuthProviders = map[spec.AWSIAMAuthService]awsIAMAuthProvider{
	spec.AWSIAMAuthServiceRDS: {
		minTLS:    tlsEncrypted,
		signToken: signRDSAuthToken,
	},
}

// configureAWSIAMAuth installs a BeforeConnect callback that signs a fresh
// short-lived IAM authentication token and uses it as the connection password.
func configureAWSIAMAuth(ctx context.Context, pgxConfig *pgxpool.Config, authSpec *spec.AWSIAMAuthSpec) error {
	service := authSpec.ServiceOrDefault()
	provider, ok := awsIAMAuthProviders[service]
	if !ok {
		return fmt.Errorf("unsupported `aws_iam_auth.service` %q", service)
	}

	// The token is a signed credential, so fail fast rather than leak it over a
	// connection that could be plaintext.
	if !checkTLS(pgxConfig.ConnConfig, provider.minTLS) {
		return fmt.Errorf("aws_iam_auth with `service: %s` requires a TLS connection: set %s in `connection_string`", service, provider.minTLS.hint())
	}

	// Parse once so a malformed override fails here, not on every connection.
	endpointHost, endpointPort, err := parseAWSIAMAuthEndpoint(authSpec.Endpoint)
	if err != nil {
		return err
	}

	awsCfg, err := loadAWSIAMAuthConfig(ctx, authSpec)
	if err != nil {
		return err
	}

	// Run any existing hook first; the token is set last as it must be the password.
	prevBeforeConnect := pgxConfig.BeforeConnect
	pgxConfig.BeforeConnect = func(ctx context.Context, connConfig *pgx.ConnConfig) error {
		if prevBeforeConnect != nil {
			if err := prevBeforeConnect(ctx, connConfig); err != nil {
				return err
			}
		}
		// Sign for the connection actually being made (pgx may be using a fallback):
		// the service rejects a token signed for a different endpoint or user.
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

	// Unlike a Lakebase OAuth credential, the token only authenticates: the session
	// outlives it, so pooled connections need not be recycled early.
	return nil
}

func signRDSAuthToken(ctx context.Context, target awsIAMAuthTarget, creds aws.CredentialsProvider) (string, error) {
	endpoint := net.JoinHostPort(target.Host, strconv.Itoa(int(target.Port)))
	token, err := rdsauth.BuildAuthToken(ctx, endpoint, target.Region, target.User, creds)
	if err != nil {
		return "", err
	}
	return withTokenPathSeparator(token), nil
}

// remove when https://github.com/aws/aws-sdk-go-v2/issues/3365 is resolved
// withTokenPathSeparator inserts the `/` between the endpoint and the query string.
// `rdsauth.BuildAuthToken` signs a URL whose path is empty, so `net/url` renders the
// token as `host:port?Action=connect&...`. RDS only accepts the `host:port/?Action=...`
// form the AWS CLI emits and rejects the other with `PAM authentication failed`, even
// though both carry the same signature. Passing an endpoint with a trailing slash is
// not an option: BuildAuthToken parses the port off the end of the string.
func withTokenPathSeparator(token string) string {
	endpoint, query, found := strings.Cut(token, "?")
	if !found || strings.HasSuffix(endpoint, "/") {
		return token
	}
	return endpoint + "/?" + query
}

// parseAWSIAMAuthEndpoint splits an `aws_iam_auth.endpoint` override. A zero port
// means the port from `connection_string` is used.
func parseAWSIAMAuthEndpoint(endpoint string) (host string, port uint16, err error) {
	if endpoint == "" {
		return "", 0, nil
	}

	host, portStr, splitErr := net.SplitHostPort(endpoint)
	if splitErr != nil {
		// RDS and DSQL endpoints are hostnames, so a value without a port is the host.
		if strings.Contains(endpoint, ":") {
			return "", 0, fmt.Errorf("invalid `aws_iam_auth.endpoint` %q: %w", endpoint, splitErr)
		}
		return endpoint, 0, nil
	}

	// Reject 0, which would be indistinguishable from an absent port.
	parsedPort, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil || parsedPort == 0 {
		return "", 0, fmt.Errorf("invalid port %q in `aws_iam_auth.endpoint` %q", portStr, endpoint)
	}
	return host, uint16(parsedPort), nil
}

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
	// An unresolved region would only surface as an opaque authentication failure.
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
		// Reuse the credentials across connections instead of calling STS per token.
		awsCfg.Credentials = aws.NewCredentialsCache(provider)
	}

	return awsCfg, nil
}
