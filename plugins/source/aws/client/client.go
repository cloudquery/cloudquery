package client

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	ServicesManager ServicesManager
	logger          zerolog.Logger
	// this is set by table clientList
	AccountID            string
	Region               string
	AutoscalingNamespace string
	WAFScope             wafv2types.Scope
	Partition            string
}

type AwsLogger struct {
	l zerolog.Logger
}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

type ServicesPartitionAccountRegionMap map[string]map[string]map[string]*Services

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager struct {
	services         ServicesPartitionAccountRegionMap
	wafScopeServices map[string]map[string]*Services
}

const (
	defaultRegion              = "us-east-1"
	awsFailedToConfigureErrMsg = "failed to retrieve credentials for account %s. AWS Error: %w, detected aws env variables: %s"
	awsOrgsFailedToFindMembers = "failed to list Org member accounts. Make sure that your credentials have the proper permissions"
	defaultVar                 = "default"
	cloudfrontScopeRegion      = defaultRegion
)

var errInvalidRegion = errors.New("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}
var errRetrievingCredentials = errors.New("error retrieving AWS credentials (see logs for details). Please verify your credentials and try again")

func (s *ServicesManager) ServicesByPartitionAccountAndRegion(partition, accountId, region string) *Services {
	if region == "" {
		region = defaultRegion
	}
	return s.services[partition][accountId][region]
}

func (s *ServicesManager) ServicesByAccountForWAFScope(partition, accountId string) *Services {
	return s.wafScopeServices[partition][accountId]
}

func (s *ServicesManager) InitServicesForPartitionAccountAndRegion(partition, accountId, region string, svcs Services) {
	if s.services == nil {
		s.services = make(map[string]map[string]map[string]*Services)
	}
	if s.services[partition] == nil {
		s.services[partition] = make(map[string]map[string]*Services)
	}
	if s.services[partition][accountId] == nil {
		s.services[partition][accountId] = make(map[string]*Services)
	}
	s.services[partition][accountId][region] = &svcs
}

func (s *ServicesManager) InitServicesForPartitionAccountAndScope(partition, accountId string, svcs Services) {
	if s.wafScopeServices == nil {
		s.wafScopeServices = make(map[string]map[string]*Services)
	}
	if s.wafScopeServices[partition] == nil {
		s.wafScopeServices[partition] = make(map[string]*Services)
	}
	s.wafScopeServices[partition][accountId] = &svcs
}

func NewAwsClient(logger zerolog.Logger) Client {
	return Client{
		ServicesManager: ServicesManager{
			services: ServicesPartitionAccountRegionMap{},
		},
		logger: logger,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return strings.TrimRight(strings.Join([]string{
		c.AccountID,
		c.Region,
		c.AutoscalingNamespace,
		string(c.WAFScope),
	}, ":"), ":")
}

func (c *Client) Services() *Services {
	s := c.ServicesManager.ServicesByPartitionAccountAndRegion(c.Partition, c.AccountID, c.Region)
	if s == nil && c.WAFScope == wafv2types.ScopeCloudfront {
		return c.ServicesManager.ServicesByAccountForWAFScope(c.Partition, c.AccountID)
	}
	return s
}

func (c *Client) withPartitionAccountIDAndRegion(partition, accountID, region string) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		WAFScope:             c.WAFScope,
	}
}

func (c *Client) withPartitionAccountIDRegionAndNamespace(partition, accountID, region, namespace string) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Str("autoscaling_namespace", namespace).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: namespace,
		WAFScope:             c.WAFScope,
	}
}

func (c *Client) withPartitionAccountIDRegionAndScope(partition, accountID, region string, scope wafv2types.Scope) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Str("waf_scope", string(scope)).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		WAFScope:             scope,
	}
}

func verifyRegions(regions []string) error {
	availableRegions, err := getAvailableRegions()
	if err != nil {
		return err
	}

	// validate regions values
	var hasWildcard bool
	for i, region := range regions {
		if region == "*" {
			hasWildcard = true
		}
		if i != 0 && region == "*" {
			return errInvalidRegion
		}
		if i > 0 && hasWildcard {
			return errInvalidRegion
		}
		regionExist := availableRegions[region]
		if !hasWildcard && !regionExist {
			return errUnknownRegion(region)
		}
	}
	return nil
}
func isAllRegions(regions []string) bool {
	// if regions array is not valid return false
	err := verifyRegions(regions)
	if err != nil {
		return false
	}

	wildcardAllRegions := false
	if (len(regions) == 1 && regions[0] == "*") || (len(regions) == 0) {
		wildcardAllRegions = true
	}
	return wildcardAllRegions
}

func getAccountId(ctx context.Context, awsCfg aws.Config) (*sts.GetCallerIdentityOutput, error) {
	svc := sts.NewFromConfig(awsCfg)
	return svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
}

func configureAwsClient(ctx context.Context, logger zerolog.Logger, awsConfig *Spec, account Account, stsClient AssumeRoleAPIClient) (aws.Config, error) {
	var err error
	var awsCfg aws.Config

	maxAttempts := 10
	if awsConfig.MaxRetries != nil {
		maxAttempts = *awsConfig.MaxRetries
	}
	maxBackoff := 30
	if awsConfig.MaxBackoff != nil {
		maxBackoff = *awsConfig.MaxBackoff
	}

	configFns := []func(*config.LoadOptions) error{
		config.WithDefaultRegion(defaultRegion),
		// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/retries-timeouts/
		config.WithRetryer(func() aws.Retryer {
			return retry.NewStandard(func(so *retry.StandardOptions) {
				so.MaxAttempts = maxAttempts
				so.MaxBackoff = time.Duration(maxBackoff) * time.Second
				so.RateLimiter = &NoRateLimiter{}
			})
		}),
	}
	if awsConfig.EndpointURL != "" {
		configFns = append(configFns, config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...any) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:               awsConfig.EndpointURL,
					HostnameImmutable: aws.ToBool(awsConfig.HostnameImmutable),
					PartitionID:       awsConfig.PartitionID,
					SigningRegion:     awsConfig.SigningRegion,
				}, nil
			})),
		)
	}

	if account.DefaultRegion != "" {
		// According to the docs: If multiple WithDefaultRegion calls are made, the last call overrides the previous call values
		configFns = append(configFns, config.WithDefaultRegion(account.DefaultRegion))
	}

	if account.LocalProfile != "" {
		configFns = append(configFns, config.WithSharedConfigProfile(account.LocalProfile))
	}

	awsCfg, err = config.LoadDefaultConfig(ctx, configFns...)

	if err != nil {
		logger.Error().Err(err).Msg("error loading default config")
		return awsCfg, err
	}

	if account.RoleARN != "" {
		opts := make([]func(*stscreds.AssumeRoleOptions), 0, 1)
		if account.ExternalID != "" {
			opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
				opts.ExternalID = &account.ExternalID
			})
		}
		if account.RoleSessionName != "" {
			opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
				opts.RoleSessionName = account.RoleSessionName
			})
		}

		if stsClient == nil {
			stsClient = sts.NewFromConfig(awsCfg)
		}
		provider := stscreds.NewAssumeRoleProvider(stsClient, account.RoleARN, opts...)

		awsCfg.Credentials = aws.NewCredentialsCache(provider, func(options *aws.CredentialsCacheOptions) {
			// ExpiryWindow will allow the credentials to trigger refreshing prior to
			// the credentials actually expiring. This is beneficial so race conditions
			// with expiring credentials do not cause requests to fail unexpectedly
			// due to ExpiredToken exceptions.
			//
			// An ExpiryWindow of 5 minute would cause calls to IsExpired() to return true
			// 5 minutes before the credentials are actually expired. This can cause an
			// increased number of requests to refresh the credentials to occur. We balance this with jitter.
			options.ExpiryWindow = 5 * time.Minute
			// Jitter is added to avoid the thundering herd problem of many refresh requests
			// happening all at once.
			options.ExpiryWindowJitterFrac = 0.5
		})
	}

	if awsConfig.AWSDebug {
		awsCfg.ClientLogMode = aws.LogRequestWithBody | aws.LogResponseWithBody | aws.LogRetries
		awsCfg.Logger = AwsLogger{logger.With().Str("accountName", account.AccountName).Logger()}
	}

	// Test out retrieving credentials
	if _, err := awsCfg.Credentials.Retrieve(ctx); err != nil {
		logger.Error().Err(err).Msg("error retrieving credentials")
		return awsCfg, errRetrievingCredentials
	}

	return awsCfg, err
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var awsConfig Spec
	err := spec.UnmarshalSpec(&awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	err = awsConfig.Validate()
	if err != nil {
		return nil, fmt.Errorf("spec validation failed: %w", err)
	}

	client := NewAwsClient(logger)
	var adminAccountSts AssumeRoleAPIClient

	if awsConfig.Organization != nil {
		var err error
		awsConfig.Accounts, adminAccountSts, err = loadOrgAccounts(ctx, logger, &awsConfig)
		if err != nil {
			logger.Error().Err(err).Msg("error getting child accounts")
			return nil, err
		}
	}
	if len(awsConfig.Accounts) == 0 {
		awsConfig.Accounts = append(awsConfig.Accounts, Account{
			ID: defaultVar,
		})
	}

	for _, account := range awsConfig.Accounts {
		if account.AccountName == "" {
			account.AccountName = account.ID
		}

		localRegions := account.Regions
		if len(localRegions) == 0 {
			localRegions = awsConfig.Regions
		}

		if err := verifyRegions(localRegions); err != nil {
			return nil, err
		}

		if isAllRegions(localRegions) {
			logger.Info().Msg("All regions specified in `cloudquery.yml`. Assuming all regions")
		}

		awsCfg, err := configureAwsClient(ctx, logger, &awsConfig, account, adminAccountSts)
		if err != nil {
			if account.source == "org" {
				logger.Warn().Msg("Unable to assume role in account")
				continue
			}
			var ae smithy.APIError
			if errors.As(err, &ae) {
				if strings.Contains(ae.ErrorCode(), "AccessDenied") {
					logger.Warn().Str("account", account.AccountName).Err(err).Msg("Access denied for account")
					continue
				}
			}
			if errors.Is(err, errRetrievingCredentials) {
				logger.Warn().Str("account", account.AccountName).Err(err).Msg("Could not retrieve credentials for account")
				continue
			}

			return nil, err
		}
		account.Regions = findEnabledRegions(ctx, logger, account.AccountName, ec2.NewFromConfig(awsCfg), localRegions, account.DefaultRegion)
		if len(account.Regions) == 0 {
			logger.Warn().Str("account", account.AccountName).Err(err).Msg("No enabled regions provided in config for account")
			continue
		}
		awsCfg.Region = account.Regions[0]
		output, err := getAccountId(ctx, awsCfg)
		if err != nil {
			return nil, err
		}
		iamArn, err := arn.Parse(*output.Arn)
		if err != nil {
			return nil, err
		}

		for _, region := range account.Regions {
			client.ServicesManager.InitServicesForPartitionAccountAndRegion(iamArn.Partition, *output.Account, region, initServices(region, awsCfg))
		}
		client.ServicesManager.InitServicesForPartitionAccountAndScope(iamArn.Partition, *output.Account, initServices(cloudfrontScopeRegion, awsCfg))
	}
	if len(client.ServicesManager.services) == 0 {
		return nil, fmt.Errorf("no enabled accounts instantiated")
	}
	return &client, nil
}

func findEnabledRegions(ctx context.Context, logger zerolog.Logger, accountName string, ec2Client services.Ec2Client, localRegions []string, accountDefaultRegion string) []string {
	// By default we should use the default region (us-east-1)
	regionsToCheck := []string{defaultRegion}
	// If user specifies a Default Region we should use it
	if accountDefaultRegion != "" {
		regionsToCheck = []string{accountDefaultRegion}
		// If no default region and * is not specified we should use all specified regions
	} else if len(localRegions) > 0 && !isAllRegions(localRegions) {
		regionsToCheck = localRegions
	}

	for _, region := range regionsToCheck {
		enabledRegions, err := getEnabledRegions(ctx, ec2Client, region)
		if err != nil {
			logger.Warn().Str("account", accountName).Err(err).Msgf("Failed to find disabled regions for account when checking: %s", region)
			continue
		}
		filteredRegions := filterDisabledRegions(localRegions, enabledRegions)
		if len(filteredRegions) > 0 {
			return filteredRegions
		}
	}
	return []string{}
}

func getEnabledRegions(ctx context.Context, ec2Client services.Ec2Client, region string) ([]types.Region, error) {
	res, err := ec2Client.DescribeRegions(ctx,
		&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)},
		func(o *ec2.Options) {
			o.Region = region
		})
	if err != nil {
		return nil, err
	}
	return res.Regions, nil
}

func filterDisabledRegions(regions []string, enabledRegions []types.Region) []string {
	regionsMap := map[string]bool{}
	for _, r := range enabledRegions {
		if r.RegionName != nil && r.OptInStatus != nil && *r.OptInStatus != "not-opted-in" {
			regionsMap[*r.RegionName] = true
		}
	}

	var filteredRegions []string
	// Our list of regions might not always be the latest and most up to date list
	// if a user specifies all regions via a "*" then they should get the most broad list possible
	if isAllRegions(regions) {
		for region := range regionsMap {
			filteredRegions = append(filteredRegions, region)
		}
	} else {
		for _, r := range regions {
			if regionsMap[r] {
				filteredRegions = append(filteredRegions, r)
			}
		}
	}
	return filteredRegions
}

func (a AwsLogger) Logf(classification logging.Classification, format string, v ...any) {
	if classification == logging.Warn {
		a.l.Warn().Msg(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug().Msg(fmt.Sprintf(format, v...))
	}
}
