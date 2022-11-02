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
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	logLevel        *string
	maxRetries      int
	maxBackoff      int
	ServicesManager ServicesManager
	logger          zerolog.Logger
	// this is set by table clientList
	AccountID            string
	Region               string
	AutoscalingNamespace string
	WAFScope             wafv2types.Scope
	Partition            string
}

// S3Manager This is needed because https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/feature/s3/manager
// has different structure then all other services (i.e no service but just a function) and we need
// the ability to mock it.
// Also we need to use s3 manager to be able to query the bucket-region https://github.com/aws/aws-sdk-go-v2/pull/1027#issuecomment-759818990
type S3Manager struct {
	s3Client *s3.Client
}

type AwsLogger struct {
	l zerolog.Logger
}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

type Services struct {
	ACM                    ACMClient
	Analyzer               AnalyzerClient
	Apigateway             ApigatewayClient
	Apigatewayv2           Apigatewayv2Client
	ApplicationAutoscaling ApplicationAutoscalingClient
	Apprunner              AppRunnerClient
	AppSync                AppSyncClient
	Athena                 AthenaClient
	Autoscaling            AutoscalingClient
	Backup                 BackupClient
	CloudHSMV2             CloudHSMV2Client
	Cloudformation         CloudFormationClient
	Cloudfront             CloudfrontClient
	Cloudtrail             CloudtrailClient
	Cloudwatch             CloudwatchClient
	CloudwatchLogs         CloudwatchLogsClient
	Codebuild              CodebuildClient
	CodePipeline           CodePipelineClient
	CognitoIdentityPools   CognitoIdentityPoolsClient
	CognitoUserPools       CognitoUserPoolsClient
	ConfigService          ConfigServiceClient
	DAX                    DAXClient
	Directconnect          DirectconnectClient
	DMS                    DatabasemigrationserviceClient
	DocDB                  DocDBClient
	DynamoDB               DynamoDBClient
	EC2                    Ec2Client
	ECR                    EcrClient
	ECRPublic              EcrPublicClient
	ECS                    EcsClient
	EFS                    EfsClient
	Eks                    EksClient
	ElastiCache            ElastiCache
	ElasticBeanstalk       ElasticbeanstalkClient
	ElasticSearch          ElasticSearch
	ELBv1                  ElbV1Client
	ELBv2                  ElbV2Client
	EMR                    EmrClient
	EventBridge            EventBridgeClient
	Firehose               FirehoseClient
	FraudDetector          FraudDetectorClient
	FSX                    FsxClient
	Glacier                GlacierClient
	Glue                   GlueClient
	GuardDuty              GuardDutyClient
	IAM                    IamClient
	Inspector              InspectorClient
	InspectorV2            InspectorV2Client
	IOT                    IOTClient
	Kinesis                KinesisClient
	KMS                    KmsClient
	Lambda                 LambdaClient
	Lightsail              LightsailClient
	MQ                     MQClient
	Neptune                NeptuneClient
	Organizations          OrganizationsClient
	QLDB                   QLDBClient
	RDS                    RdsClient
	Redshift               RedshiftClient
	ResourceGroups         ResourceGroupsClient
	Route53                Route53Client
	Route53Domains         Route53DomainsClient
	S3                     S3Client
	S3Control              S3ControlClient
	S3Manager              S3ManagerClient
	SageMaker              SageMakerClient
	SecretsManager         SecretsManagerClient
	ServiceCatalog         ServiceCatalogClient
	ServiceCatalogAR       ServiceCatalogAppRegistryClient
	SES                    SESClient
	Shield                 ShieldClient
	SNS                    SnsClient
	SQS                    SQSClient
	SSM                    SSMClient
	Transfer               TransferClient
	Waf                    WafClient
	WafRegional            WafRegionalClient
	WafV2                  WafV2Client
	Workspaces             WorkspacesClient
	Xray                   XrayClient
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

var errInvalidRegion = fmt.Errorf("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}

func (s *ServicesManager) ServicesByPartitionAccountAndRegion(partition, accountId, region string) *Services {
	if region == "" {
		region = defaultRegion
	}
	return s.services[partition][accountId][region]
}

func (s *ServicesManager) ServicesByAccountForWAFScope(partition, accountId string) *Services {
	return s.wafScopeServices[partition][accountId]
}

func (s *ServicesManager) InitServicesForPartitionAccountAndRegion(partition, accountId, region string, services Services) {
	if s.services == nil {
		s.services = make(map[string]map[string]map[string]*Services)
	}
	if s.services[partition] == nil {
		s.services[partition] = make(map[string]map[string]*Services)
	}
	if s.services[partition][accountId] == nil {
		s.services[partition][accountId] = make(map[string]*Services)
	}
	s.services[partition][accountId][region] = &services
}

func (s *ServicesManager) InitServicesForPartitionAccountAndScope(partition, accountId string, services Services) {
	if s.wafScopeServices == nil {
		s.wafScopeServices = make(map[string]map[string]*Services)
	}
	if s.wafScopeServices[partition] == nil {
		s.wafScopeServices[partition] = make(map[string]*Services)
	}
	s.wafScopeServices[partition][accountId] = &services
}

func newS3ManagerFromConfig(cfg aws.Config) S3Manager {
	return S3Manager{
		s3Client: s3.NewFromConfig(cfg),
	}
}

func (s3Manager S3Manager) GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error) {
	return manager.GetBucketRegion(ctx, s3Manager.s3Client, bucket, optFns...)
}

func NewAwsClient(logger zerolog.Logger) Client {
	return Client{
		ServicesManager: ServicesManager{
			services: ServicesPartitionAccountRegionMap{},
		},
		logger: logger,
	}
}

func (s ServicesPartitionAccountRegionMap) Accounts() []string {
	accounts := make([]string, 0)
	for partitions := range s {
		for account := range s[partitions] {
			accounts = append(accounts, account)
		}
	}
	return accounts
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

// ARN builds an ARN tied to current client's partition, accountID and region
func (c *Client) ARN(service AWSService, idParts ...string) string {
	return makeARN(service, c.Partition, c.AccountID, c.Region, idParts...).String()
}

// AccountGlobalARN builds an ARN tied to current client's partition and accountID
func (c *Client) AccountGlobalARN(service AWSService, idParts ...string) string {
	return makeARN(service, c.Partition, c.AccountID, "", idParts...).String()
}

// PartitionGlobalARN builds an ARN tied to current client's partition
func (c *Client) PartitionGlobalARN(service AWSService, idParts ...string) string {
	return makeARN(service, c.Partition, "", "", idParts...).String()
}

// RegionGlobalARN builds an ARN tied to current client's partition and accountID
func (c *Client) RegionGlobalARN(service AWSService, idParts ...string) string {
	return makeARN(service, c.Partition, "", c.Region, idParts...).String()
}

func (c *Client) withPartitionAccountIDAndRegion(partition, accountID, region string) *Client {
	return &Client{
		Partition:            partition,
		logLevel:             c.logLevel,
		maxRetries:           c.maxRetries,
		maxBackoff:           c.maxBackoff,
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
		logLevel:             c.logLevel,
		maxRetries:           c.maxRetries,
		maxBackoff:           c.maxBackoff,
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
		logLevel:             c.logLevel,
		maxRetries:           c.maxRetries,
		maxBackoff:           c.maxBackoff,
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
			// return retry.NewAdaptiveMode()
			return retry.NewStandard(func(so *retry.StandardOptions) {
				so.MaxAttempts = maxAttempts
				so.MaxBackoff = time.Duration(maxBackoff) * time.Second
				so.RateLimiter = &NoRateLimiter{}
			})
			// return retry.AddWithMaxAttempts(retry.NewStandard(), 5)
		}),
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

		awsCfg.Credentials = aws.NewCredentialsCache(provider)
	}

	if awsConfig.AWSDebug {
		awsCfg.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
		awsCfg.Logger = AwsLogger{logger.With().Str("accountName", account.AccountName).Logger()}
	}

	// Test out retrieving credentials
	if _, err := awsCfg.Credentials.Retrieve(ctx); err != nil {
		logger.Error().Err(err).Msg("error retrieving credentials")
		return awsCfg, fmt.Errorf("error retrieving credentials: %w", err)
	}

	return awsCfg, err
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
	var awsConfig Spec
	err := spec.UnmarshalSpec(&awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	client := NewAwsClient(logger)
	var adminAccountSts AssumeRoleAPIClient
	if awsConfig.Organization != nil && len(awsConfig.Accounts) > 0 {
		return nil, errors.New("specifying accounts via both the Accounts and Org properties is not supported. If you want to do both, you should use multiple provider blocks")
	}
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

			return nil, err
		}

		// This is a work-around to skip disabled regions
		// https://github.com/aws/aws-sdk-go-v2/issues/1068
		res, err := ec2.NewFromConfig(awsCfg).DescribeRegions(ctx,
			&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)},
			func(o *ec2.Options) {
				o.Region = defaultRegion
				if account.DefaultRegion != "" {
					o.Region = account.DefaultRegion
				}

				if len(localRegions) > 0 && !isAllRegions(localRegions) {
					o.Region = localRegions[0]
				}
			})
		if err != nil {
			logger.Warn().Str("account", account.AccountName).Err(err).Msg("Failed to find disabled regions for account")
			continue
		}
		account.Regions = filterDisabledRegions(localRegions, res.Regions)

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

func initServices(region string, c aws.Config) Services {
	awsCfg := c.Copy()
	awsCfg.Region = region
	return Services{
		ACM:                    acm.NewFromConfig(awsCfg),
		Analyzer:               accessanalyzer.NewFromConfig(awsCfg),
		Apigateway:             apigateway.NewFromConfig(awsCfg),
		Apigatewayv2:           apigatewayv2.NewFromConfig(awsCfg),
		ApplicationAutoscaling: applicationautoscaling.NewFromConfig(awsCfg),
		Apprunner:              apprunner.NewFromConfig(awsCfg),
		AppSync:                appsync.NewFromConfig(awsCfg),
		Athena:                 athena.NewFromConfig(awsCfg),
		Autoscaling:            autoscaling.NewFromConfig(awsCfg),
		Backup:                 backup.NewFromConfig(awsCfg),
		CloudHSMV2:             cloudhsmv2.NewFromConfig(awsCfg),
		Cloudformation:         cloudformation.NewFromConfig(awsCfg),
		Cloudfront:             cloudfront.NewFromConfig(awsCfg),
		Cloudtrail:             cloudtrail.NewFromConfig(awsCfg),
		Cloudwatch:             cloudwatch.NewFromConfig(awsCfg),
		CloudwatchLogs:         cloudwatchlogs.NewFromConfig(awsCfg),
		Codebuild:              codebuild.NewFromConfig(awsCfg),
		CodePipeline:           codepipeline.NewFromConfig(awsCfg),
		CognitoIdentityPools:   cognitoidentity.NewFromConfig(awsCfg),
		CognitoUserPools:       cognitoidentityprovider.NewFromConfig(awsCfg),
		ConfigService:          configservice.NewFromConfig(awsCfg),
		DAX:                    dax.NewFromConfig(awsCfg),
		Directconnect:          directconnect.NewFromConfig(awsCfg),
		DMS:                    databasemigrationservice.NewFromConfig(awsCfg),
		DocDB:                  docdb.NewFromConfig(awsCfg),
		DynamoDB:               dynamodb.NewFromConfig(awsCfg),
		EC2:                    ec2.NewFromConfig(awsCfg),
		ECR:                    ecr.NewFromConfig(awsCfg),
		ECRPublic:              ecrpublic.NewFromConfig(awsCfg),
		ECS:                    ecs.NewFromConfig(awsCfg),
		EFS:                    efs.NewFromConfig(awsCfg),
		Eks:                    eks.NewFromConfig(awsCfg),
		ElastiCache:            elasticache.NewFromConfig(awsCfg),
		ElasticBeanstalk:       elasticbeanstalk.NewFromConfig(awsCfg),
		ElasticSearch:          elasticsearchservice.NewFromConfig(awsCfg),
		ELBv1:                  elbv1.NewFromConfig(awsCfg),
		ELBv2:                  elbv2.NewFromConfig(awsCfg),
		EMR:                    emr.NewFromConfig(awsCfg),
		EventBridge:            eventbridge.NewFromConfig(awsCfg),
		Firehose:               firehose.NewFromConfig(awsCfg),
		FraudDetector:          frauddetector.NewFromConfig(awsCfg),
		FSX:                    fsx.NewFromConfig(awsCfg),
		Glacier:                glacier.NewFromConfig(awsCfg),
		Glue:                   glue.NewFromConfig(awsCfg),
		GuardDuty:              guardduty.NewFromConfig(awsCfg),
		IAM:                    iam.NewFromConfig(awsCfg),
		Inspector:              inspector.NewFromConfig(awsCfg),
		InspectorV2:            inspector2.NewFromConfig(awsCfg),
		IOT:                    iot.NewFromConfig(awsCfg),
		Kinesis:                kinesis.NewFromConfig(awsCfg),
		KMS:                    kms.NewFromConfig(awsCfg),
		Lambda:                 lambda.NewFromConfig(awsCfg),
		Lightsail:              lightsail.NewFromConfig(awsCfg),
		MQ:                     mq.NewFromConfig(awsCfg),
		Neptune:                neptune.NewFromConfig(awsCfg),
		Organizations:          organizations.NewFromConfig(awsCfg),
		QLDB:                   qldb.NewFromConfig(awsCfg),
		RDS:                    rds.NewFromConfig(awsCfg),
		ResourceGroups:         resourcegroups.NewFromConfig(awsCfg),
		Redshift:               redshift.NewFromConfig(awsCfg),
		Route53:                route53.NewFromConfig(awsCfg),
		Route53Domains:         route53domains.NewFromConfig(awsCfg),
		S3:                     s3.NewFromConfig(awsCfg),
		S3Control:              s3control.NewFromConfig(awsCfg),
		S3Manager:              newS3ManagerFromConfig(awsCfg),
		SageMaker:              sagemaker.NewFromConfig(awsCfg),
		SecretsManager:         secretsmanager.NewFromConfig(awsCfg),
		ServiceCatalog:         servicecatalog.NewFromConfig(awsCfg),
		ServiceCatalogAR:       servicecatalogappregistry.NewFromConfig(awsCfg),
		SES:                    sesv2.NewFromConfig(awsCfg),
		Shield:                 shield.NewFromConfig(awsCfg),
		SNS:                    sns.NewFromConfig(awsCfg),
		SQS:                    sqs.NewFromConfig(awsCfg),
		SSM:                    ssm.NewFromConfig(awsCfg),
		Transfer:               transfer.NewFromConfig(awsCfg),
		Waf:                    waf.NewFromConfig(awsCfg),
		WafRegional:            wafregional.NewFromConfig(awsCfg),
		WafV2:                  wafv2.NewFromConfig(awsCfg),
		Workspaces:             workspaces.NewFromConfig(awsCfg),
		Xray:                   xray.NewFromConfig(awsCfg),
	}
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

func (a AwsLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	if classification == logging.Warn {
		a.l.Warn().Msg(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug().Msg(fmt.Sprintf(format, v...))
	}
}
