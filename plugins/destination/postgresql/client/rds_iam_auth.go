package client

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
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
// `rds_iam_auth.role_arn` stay valid before they are refreshed. Every role allows
// for at least one hour.
const assumeRoleDuration = time.Hour

// configureRDSIAMAuth wires up AWS RDS IAM database authentication on the given
// pool config. It resolves the AWS credentials once and installs a BeforeConnect
// callback that signs a fresh short-lived IAM authentication token and uses it as
// the connection password for every new connection.
func configureRDSIAMAuth(ctx context.Context, pgxConfig *pgxpool.Config, rdsSpec *spec.RDSIAMAuthSpec) error {
	// RDS requires TLS for IAM database authentication, and the token is a signed
	// credential, so reject any connection string that could connect without it.
	// Fail fast with a clear error rather than leaking the token over plaintext or
	// failing later in a less obvious way.
	if !requiresTLS(pgxConfig.ConnConfig) {
		return errors.New("rds_iam_auth requires a TLS connection: set `sslmode=require` (or `verify-ca`/`verify-full`) in `connection_string`")
	}

	awsCfg, err := loadRDSIAMAuthAWSConfig(ctx, rdsSpec)
	if err != nil {
		return err
	}

	// Preserve any previously configured BeforeConnect hook and run it first, so
	// RDS IAM auth composes with other hooks instead of discarding them. The token
	// is set last as it must be the connection password.
	prevBeforeConnect := pgxConfig.BeforeConnect
	pgxConfig.BeforeConnect = func(ctx context.Context, connConfig *pgx.ConnConfig) error {
		if prevBeforeConnect != nil {
			if err := prevBeforeConnect(ctx, connConfig); err != nil {
				return err
			}
		}
		// RDS rejects a token that was signed for a different host, port or user, so
		// sign for the endpoint actually being dialed (pgx may be connecting via a
		// fallback) unless the spec pins one, and for the user pgx will send in the
		// startup message.
		endpoint := rdsSpec.Endpoint
		if endpoint == "" {
			endpoint = net.JoinHostPort(connConfig.Host, strconv.Itoa(int(connConfig.Port)))
		}
		token, err := rdsauth.BuildAuthToken(ctx, endpoint, awsCfg.Region, connConfig.User, awsCfg.Credentials)
		if err != nil {
			return fmt.Errorf("failed to build rds iam authentication token: %w", err)
		}
		connConfig.Password = token
		return nil
	}

	// Note that, unlike an OAuth credential, an RDS IAM token is only used to
	// authenticate: a session stays valid once established, even after the token's
	// 15 minute lifetime has passed, so pooled connections need not be recycled
	// early.
	return nil
}

// loadRDSIAMAuthAWSConfig resolves the AWS config used to sign authentication
// tokens. Credentials are resolved from the standard AWS sources, optionally
// narrowed to a shared config profile or exchanged for assumed-role credentials.
func loadRDSIAMAuthAWSConfig(ctx context.Context, rdsSpec *spec.RDSIAMAuthSpec) (aws.Config, error) {
	loadOpts := make([]func(*config.LoadOptions) error, 0, 2)
	if rdsSpec.Region != "" {
		loadOpts = append(loadOpts, config.WithRegion(rdsSpec.Region))
	}
	if rdsSpec.LocalProfile != "" {
		loadOpts = append(loadOpts, config.WithSharedConfigProfile(rdsSpec.LocalProfile))
	}

	awsCfg, err := config.LoadDefaultConfig(ctx, loadOpts...)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load aws config: %w", err)
	}
	// The token is signed for a region, so an unresolved one would only surface as
	// an opaque authentication failure from RDS.
	if awsCfg.Region == "" {
		return aws.Config{}, errors.New("failed to resolve the aws region: set `rds_iam_auth.region` or the `AWS_REGION` environment variable")
	}

	if rdsSpec.RoleARN != "" {
		assumeRoleOpts := []func(*stscreds.AssumeRoleOptions){
			func(o *stscreds.AssumeRoleOptions) { o.Duration = assumeRoleDuration },
		}
		if rdsSpec.RoleSessionName != "" {
			assumeRoleOpts = append(assumeRoleOpts, func(o *stscreds.AssumeRoleOptions) {
				o.RoleSessionName = rdsSpec.RoleSessionName
			})
		}
		if rdsSpec.ExternalID != "" {
			assumeRoleOpts = append(assumeRoleOpts, func(o *stscreds.AssumeRoleOptions) {
				o.ExternalID = &rdsSpec.ExternalID
			})
		}
		provider := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(awsCfg), rdsSpec.RoleARN, assumeRoleOpts...)
		// Cache the assumed-role credentials so that they are reused across
		// connections instead of calling STS for every token.
		awsCfg.Credentials = aws.NewCredentialsCache(provider)
	}

	return awsCfg, nil
}
