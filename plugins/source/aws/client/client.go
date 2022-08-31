package client

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
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
	"github.com/aws/aws-sdk-go-v2/service/fsx"
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
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	logLevel        *string
	maxRetries      int
	maxBackoff      int
	ServicesManager ServicesManager
	logger          hclog.Logger
	// this is set by table clientList
	AccountID            string
	GlobalRegion         string
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
	l hclog.Logger
}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

type Services struct {
	ACM                    ACMClient
	AccessAnalyzer         AnalyzerClient
	Apigateway             ApigatewayClient
	Apigatewayv2           Apigatewayv2Client
	ApplicationAutoscaling ApplicationAutoscalingClient
	AppSync                AppSyncClient
	Athena                 AthenaClient
	Autoscaling            AutoscalingClient
	Backup                 BackupClient
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
	DynamoDB               DynamoDBClient
	EC2                    Ec2Client
	ECR                    EcrClient
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
	FSX                    FsxClient
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

var envVarsToCheck = []string{
	"AWS_PROFILE",
	"AWS_ACCESS_KEY_ID",
	"AWS_SECRET_ACCESS_KEY",
	"AWS_CONFIG_FILE",
	"AWS_ROLE_ARN",
	"AWS_SESSION_TOKEN",
	"AWS_SHARED_CREDENTIALS_FILE",
}

var errInvalidRegion = fmt.Errorf("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}

var (
	_ schema.ClientMeta       = (*Client)(nil)
	_ schema.ClientIdentifier = (*Client)(nil)
)

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

func NewAwsClient(logger hclog.Logger) Client {
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
func (c *Client) Logger() hclog.Logger {
	return &awsLogger{c.logger, c.ServicesManager.services.Accounts()}
}

// Identify the given client
func (c *Client) Identify() string {
	return strings.TrimRight(strings.Join([]string{
		obfuscateAccountId(c.AccountID),
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
		logger:               c.logger.With("account_id", obfuscateAccountId(accountID), "Region", region),
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
		logger:               c.logger.With("account_id", obfuscateAccountId(accountID), "Region", region, "AutoscalingNamespace", namespace),
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
		logger:               c.logger.With("account_id", obfuscateAccountId(accountID), "Region", region, "Scope", scope),
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

func configureAwsClient(ctx context.Context, logger hclog.Logger, awsConfig *Config, account Account, stsClient AssumeRoleAPIClient) (aws.Config, error) {
	var err error
	var awsCfg aws.Config
	configFns := []func(*config.LoadOptions) error{
		config.WithDefaultRegion(defaultRegion),
		config.WithRetryer(newRetryer(logger, awsConfig.MaxRetries, awsConfig.MaxBackoff)),
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
		logger.Error("error loading default config", "err", err)
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
		awsCfg.Logger = AwsLogger{logger.With("accountName", account.AccountName)}
	}

	// Test out retrieving credentials
	if _, err := awsCfg.Credentials.Retrieve(ctx); err != nil {
		logger.Error("error retrieving credentials", "err", err)

		var ae smithy.APIError
		if errors.As(err, &ae) {
			if strings.Contains(ae.ErrorCode(), "InvalidClientTokenId") {
				return awsCfg, diag.FromError(
					err,
					diag.USER,
					diag.WithSummary("Invalid credentials for assuming role"),
					diag.WithDetails("The credentials being used to assume role are invalid. Please check that your credentials are valid in the partition you are using. If you are using a partition other than the AWS commercial region, be sure set the default_region attribute in the cloudquery.yml file."),
					diag.WithSeverity(diag.WARNING),
				)
			}
		}

		return awsCfg, diag.FromError(
			err,
			diag.USER,
			diag.WithSummary("No credentials available"),
			diag.WithDetails("Couldn't find any credentials in environment variables or configuration files."),
			diag.WithSeverity(diag.WARNING),
		)
	}

	return awsCfg, err
}

func Configure(logger hclog.Logger, providerConfig interface{}) (schema.ClientMeta, diag.Diagnostics) {
	var diags diag.Diagnostics

	ctx := context.Background()
	awsConfig := providerConfig.(*Config)
	client := NewAwsClient(logger)
	client.GlobalRegion = awsConfig.GlobalRegion
	var adminAccountSts AssumeRoleAPIClient
	if awsConfig.Organization != nil && len(awsConfig.Accounts) > 0 {
		return nil, diags.Add(diag.FromError(errors.New("specifying accounts via both the Accounts and Org properties is not supported. If you want to do both, you should use multiple provider blocks"), diag.USER))
	}
	if awsConfig.Organization != nil {
		var err error
		awsConfig.Accounts, adminAccountSts, err = loadOrgAccounts(ctx, logger, awsConfig)
		if err != nil {
			logger.Error("error getting child accounts", "err", err)
			var ae smithy.APIError
			if errors.As(err, &ae) {
				if strings.Contains(ae.ErrorCode(), "AccessDenied") {
					return nil, diags.Add(diag.FromError(fmt.Errorf(awsOrgsFailedToFindMembers), diag.ACCESS, diag.WithSeverity(diag.ERROR)))
				}
			}
			return nil, diags.Add(classifyError(err, diag.INTERNAL, nil))
		}
	}
	if len(awsConfig.Accounts) == 0 {
		awsConfig.Accounts = append(awsConfig.Accounts, Account{
			ID: defaultVar,
		})
	}

	for _, account := range awsConfig.Accounts {
		logger.Debug("user defined account", "account", account)
		if account.AccountID != "" {
			return nil, diags.Add(diag.FromError(errors.New("account_id is no longer supported. To specify a profile use `local_profile`. To specify an account alias use `account_name`"), diag.USER))
		}

		if account.AccountName == "" {
			account.AccountName = account.ID
		}

		localRegions := account.Regions
		if len(localRegions) == 0 {
			localRegions = awsConfig.Regions
		}

		if err := verifyRegions(localRegions); err != nil {
			return nil, diags.Add(classifyError(err, diag.USER, nil))
		}

		if isAllRegions(localRegions) {
			logger.Info("All regions specified in `cloudquery.yml`. Assuming all regions")
		}

		awsCfg, err := configureAwsClient(ctx, logger, awsConfig, account, adminAccountSts)
		if err != nil {
			if account.source == "org" {
				logger.Warn("unable to assume role in account")
				principal := "unknown principal"
				// Identify the principal making the request and use it to construct the error message. Any errors can be ignored as they are only for improving the user experience.
				awsAdminCfg, _ := configureAwsClient(ctx, logger, awsConfig, *awsConfig.Organization.AdminAccount, nil)
				output, accountErr := getAccountId(ctx, awsAdminCfg)
				if accountErr == nil {
					principal = *output.Arn
				}

				diags = diags.Add(diag.FromError(err, diag.ACCESS, diag.WithDetails("ensure that %s has access to be able perform `sts:AssumeRole` on %s ", principal, account.RoleARN), diag.WithSeverity(diag.WARNING)))
				continue
			}
			var ae smithy.APIError
			if errors.As(err, &ae) {
				if strings.Contains(ae.ErrorCode(), "AccessDenied") {
					diags = diags.Add(diag.FromError(fmt.Errorf(awsFailedToConfigureErrMsg, account.AccountName, err, checkEnvVariables()), diag.ACCESS, diag.WithSeverity(diag.WARNING)))
					continue
				}
			}

			return nil, diags.Add(diag.FromError(err, diag.ACCESS))
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
			diags = diags.Add(diag.FromError(fmt.Errorf("failed to find disabled regions for account %s. AWS Error: %w", account.AccountName, err), diag.ACCESS, diag.WithSeverity(diag.WARNING)))
			continue
		}
		account.Regions = filterDisabledRegions(localRegions, res.Regions)

		if len(account.Regions) == 0 {
			diags = diags.Add(diag.FromError(fmt.Errorf("no enabled regions provided in config for account %s", account.AccountName), diag.ACCESS, diag.WithSeverity(diag.WARNING)))
			continue
		}
		awsCfg.Region = account.Regions[0]
		output, err := getAccountId(ctx, awsCfg)
		if err != nil {
			// This should only ever fail when there is a network or endpoint issue. There is no way for IAM to deny this call.
			diags = diags.Add(diag.FromError(fmt.Errorf("failed to get caller identity. AWS Error: %w", err), diag.ACCESS, diag.WithSeverity(diag.WARNING)))
			continue
		}
		iamArn, err := arn.Parse(*output.Arn)
		if err != nil {
			return nil, diags.Add(classifyError(err, diag.INTERNAL, nil))
		}

		for _, region := range account.Regions {
			client.ServicesManager.InitServicesForPartitionAccountAndRegion(iamArn.Partition, *output.Account, region, initServices(region, awsCfg))
		}
		client.ServicesManager.InitServicesForPartitionAccountAndScope(iamArn.Partition, *output.Account, initServices(cloudfrontScopeRegion, awsCfg))
	}
	if len(client.ServicesManager.services) == 0 {
		return nil, diags.Add(diag.FromError(errors.New("no accounts instantiated"), diag.USER))
	}
	return &client, diags
}

func initServices(region string, c aws.Config) Services {
	awsCfg := c.Copy()
	awsCfg.Region = region
	return Services{
		ACM:                    acm.NewFromConfig(awsCfg),
		AccessAnalyzer:         accessanalyzer.NewFromConfig(awsCfg),
		Apigateway:             apigateway.NewFromConfig(awsCfg),
		Apigatewayv2:           apigatewayv2.NewFromConfig(awsCfg),
		ApplicationAutoscaling: applicationautoscaling.NewFromConfig(awsCfg),
		AppSync:                appsync.NewFromConfig(awsCfg),
		Athena:                 athena.NewFromConfig(awsCfg),
		Autoscaling:            autoscaling.NewFromConfig(awsCfg),
		Backup:                 backup.NewFromConfig(awsCfg),
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
		DynamoDB:               dynamodb.NewFromConfig(awsCfg),
		EC2:                    ec2.NewFromConfig(awsCfg),
		ECR:                    ecr.NewFromConfig(awsCfg),
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
		FSX:                    fsx.NewFromConfig(awsCfg),
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
		a.l.Warn(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug(fmt.Sprintf(format, v...))
	}
}

func obfuscateAccountId(accountId string) string {
	if len(accountId) <= 4 {
		return accountId
	}
	return accountId[:4] + "xxxxxxxx"
}

// checkEnvVariables checks which aws environment variables are set
func checkEnvVariables() string {
	var result []string
	for _, v := range envVarsToCheck {
		if _, present := os.LookupEnv(v); present {
			result = append(result, v)
		}
	}
	return strings.Join(result, ",")
}
