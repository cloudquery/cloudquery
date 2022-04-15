package client

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
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
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/smithy-go/logging"
	"github.com/hashicorp/go-hclog"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
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

const (
	defaultRegion              = "us-east-1"
	awsFailedToConfigureErrMsg = "failed to retrieve credentials for account %s. AWS Error: %w, detected aws env variables: %s"
	defaultVar                 = "default"
)

var errInvalidRegion = fmt.Errorf("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}

type Services struct {
	ACM                    ACMClient
	Analyzer               AnalyzerClient
	Apigateway             ApigatewayClient
	Apigatewayv2           Apigatewayv2Client
	ApplicationAutoscaling ApplicationAutoscalingClient
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
	DMS                    DatabasemigrationserviceClient
	Directconnect          DirectconnectClient
	DynamoDB               DynamoDBClient
	EC2                    Ec2Client
	ECR                    EcrClient
	ECS                    EcsClient
	EFS                    EfsClient
	ELBv1                  ElbV1Client
	ELBv2                  ElbV2Client
	EMR                    EmrClient
	Eks                    EksClient
	ElasticBeanstalk       ElasticbeanstalkClient
	ElasticSearch          ElasticSearch
	FSX                    FsxClient
	GuardDuty              GuardDutyClient
	IAM                    IamClient
	IOT                    IOTClient
	KMS                    KmsClient
	Lambda                 LambdaClient
	MQ                     MQClient
	Organizations          OrganizationsClient
	QLDB                   QLDBClient
	RDS                    RdsClient
	Redshift               RedshiftClient
	Route53                Route53Client
	Route53Domains         Route53DomainsClient
	S3                     S3Client
	S3Control              S3ControlClient
	S3Manager              S3ManagerClient
	SNS                    SnsClient
	SQS                    SQSClient
	SSM                    SSMClient
	SageMaker              SageMakerClient
	SecretsManager         SecretsManagerClient
	Waf                    WafClient
	WafV2                  WafV2Client
	WafRegional            WafRegionalClient
	Workspaces             WorkspacesClient
}

type ServicesAccountRegionMap map[string]map[string]*Services

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager struct {
	services ServicesAccountRegionMap
}

func (s *ServicesManager) ServicesByAccountAndRegion(accountId string, region string) *Services {
	if region == "" {
		region = defaultRegion
	}
	return s.services[accountId][region]
}

func (s *ServicesManager) InitServicesForAccountAndRegion(accountId string, region string, services Services) {
	if s.services[accountId] == nil {
		s.services[accountId] = make(map[string]*Services)
	}
	s.services[accountId][region] = &services
}

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	Accounts        []Account
	logLevel        *string
	maxRetries      int
	maxBackoff      int
	ServicesManager ServicesManager
	logger          hclog.Logger
	// this is set by table clientList
	AccountID            string
	Region               string
	AutoscalingNamespace string
	WAFScope             wafv2types.Scope
}

// S3Manager This is needed because https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/feature/s3/manager
// has different structure then all other services (i.e no service but just a function) and we need
// the ability to mock it.
// Also we need to use s3 manager to be able to query the bucket-region https://github.com/aws/aws-sdk-go-v2/pull/1027#issuecomment-759818990
type S3Manager struct {
	s3Client *s3.Client
}

func newS3ManagerFromConfig(cfg aws.Config) S3Manager {
	return S3Manager{
		s3Client: s3.NewFromConfig(cfg),
	}
}

func (s3Manager S3Manager) GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error) {
	return manager.GetBucketRegion(ctx, s3Manager.s3Client, bucket, optFns...)
}

func NewAwsClient(logger hclog.Logger, accounts []Account) Client {
	return Client{
		ServicesManager: ServicesManager{
			services: ServicesAccountRegionMap{},
		},
		logger:   logger,
		Accounts: accounts,
	}
}
func (c *Client) Logger() hclog.Logger {
	return &awsLogger{c.logger, c.Accounts}
}

func (c *Client) Services() *Services {
	return c.ServicesManager.ServicesByAccountAndRegion(c.AccountID, c.Region)
}

func (c *Client) withAccountID(accountID string) *Client {
	return &Client{
		Accounts:             c.Accounts,
		logLevel:             c.logLevel,
		maxRetries:           c.maxRetries,
		maxBackoff:           c.maxBackoff,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With("account_id", obfuscateAccountId(accountID)),
		AccountID:            accountID,
		Region:               c.Region,
		AutoscalingNamespace: c.AutoscalingNamespace,
	}
}

func (c *Client) withAccountIDAndRegion(accountID, region string) *Client {
	return &Client{
		Accounts:             c.Accounts,
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

func (c *Client) withAccountIDRegionAndNamespace(accountID, region, namespace string) *Client {
	return &Client{
		Accounts:             c.Accounts,
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

func (c *Client) withAccountIDRegionAndScope(accountID, region string, scope wafv2types.Scope) *Client {
	return &Client{
		Accounts:             c.Accounts,
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
	return svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}, func(o *sts.Options) {
		o.Region = "aws-global"
	})

}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

func configureAwsClient(ctx context.Context, logger hclog.Logger, awsConfig *Config, account Account, stsClient AssumeRoleAPIClient) (aws.Config, error) {
	var err error
	var awsCfg aws.Config
	configFns := []func(*config.LoadOptions) error{
		config.WithDefaultRegion(defaultRegion),
		config.WithRetryer(newRetryer(awsConfig.MaxRetries, awsConfig.MaxBackoff)),
	}

	if account.LocalProfile != "" {
		configFns = append(configFns, config.WithSharedConfigProfile(account.LocalProfile))
	}

	awsCfg, err = config.LoadDefaultConfig(ctx, configFns...)

	if err != nil {
		logger.Error("error loading default config", "err", err)
		return awsCfg, fmt.Errorf(awsFailedToConfigureErrMsg, account.AccountName, err, checkEnvVariables())
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
		return awsCfg, fmt.Errorf(awsFailedToConfigureErrMsg, account.AccountName, err, checkEnvVariables())
	}

	return awsCfg, err
}

func Configure(logger hclog.Logger, providerConfig interface{}) (schema.ClientMeta, error) {
	ctx := context.Background()
	awsConfig := providerConfig.(*Config)
	client := NewAwsClient(logger, awsConfig.Accounts)
	var adminAccountSts AssumeRoleAPIClient

	if awsConfig.Organization != nil {
		var err error
		awsConfig.Accounts, adminAccountSts, err = loadOrgAccounts(ctx, logger, awsConfig)
		if err != nil {
			logger.Error("error getting child accounts", "err", err)
			return nil, err
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
			return nil, fmt.Errorf("account_id is no longer supported. To specify a profile use `local_profile`. To specify an account alias use `account_name`")
		}

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
			logger.Info("All regions specified in config.yml. Assuming all regions")
		}

		awsCfg, err := configureAwsClient(ctx, logger, awsConfig, account, adminAccountSts)
		if err != nil {
			if account.source == "org" {
				logger.Warn("unable to assume role in account")
				continue

			}
			return nil, err
		}

		// This is a work-around to skip disabled regions
		// https://github.com/aws/aws-sdk-go-v2/issues/1068
		res, err := ec2.NewFromConfig(awsCfg).DescribeRegions(ctx,
			&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)},
			func(o *ec2.Options) {
				o.Region = defaultRegion
				if len(localRegions) > 0 && !isAllRegions(localRegions) {
					o.Region = localRegions[0]
				}
			})
		if err != nil {
			return nil, fmt.Errorf("failed to find disabled regions for account %s. AWS Error: %w", account.AccountName, err)
		}
		account.Regions = filterDisabledRegions(localRegions, res.Regions)

		if len(account.Regions) == 0 {
			return nil, fmt.Errorf("no enabled regions provided in config for account %s", account.AccountName)
		}

		output, err := getAccountId(ctx, awsCfg)
		if err != nil {
			return nil, err
		}
		if client.AccountID == "" {
			// set default
			client.AccountID = *output.Account
			client.Region = account.Regions[0]
			client.Accounts = append(client.Accounts, Account{ID: *output.Account, RoleARN: *output.Arn})
		}
		for _, region := range account.Regions {
			client.ServicesManager.InitServicesForAccountAndRegion(*output.Account, region, initServices(region, awsCfg))
		}
	}
	if len(client.Accounts) == 0 {
		return nil, errors.New("no accounts instantiated")
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
		Autoscaling:            autoscaling.NewFromConfig(awsCfg),
		Backup:                 backup.NewFromConfig(awsCfg),
		Cloudfront:             cloudfront.NewFromConfig(awsCfg),
		Cloudtrail:             cloudtrail.NewFromConfig(awsCfg),
		Cloudwatch:             cloudwatch.NewFromConfig(awsCfg),
		CloudwatchLogs:         cloudwatchlogs.NewFromConfig(awsCfg),
		Cloudformation:         cloudformation.NewFromConfig(awsCfg),
		CognitoIdentityPools:   cognitoidentity.NewFromConfig(awsCfg),
		CognitoUserPools:       cognitoidentityprovider.NewFromConfig(awsCfg),
		Codebuild:              codebuild.NewFromConfig(awsCfg),
		CodePipeline:           codepipeline.NewFromConfig(awsCfg),
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
		ElasticBeanstalk:       elasticbeanstalk.NewFromConfig(awsCfg),
		ElasticSearch:          elasticsearchservice.NewFromConfig(awsCfg),
		ELBv1:                  elbv1.NewFromConfig(awsCfg),
		ELBv2:                  elbv2.NewFromConfig(awsCfg),
		EMR:                    emr.NewFromConfig(awsCfg),
		FSX:                    fsx.NewFromConfig(awsCfg),
		GuardDuty:              guardduty.NewFromConfig(awsCfg),
		IAM:                    iam.NewFromConfig(awsCfg),
		KMS:                    kms.NewFromConfig(awsCfg),
		Lambda:                 lambda.NewFromConfig(awsCfg),
		MQ:                     mq.NewFromConfig(awsCfg),
		Organizations:          organizations.NewFromConfig(awsCfg),
		QLDB:                   qldb.NewFromConfig(awsCfg),
		RDS:                    rds.NewFromConfig(awsCfg),
		Redshift:               redshift.NewFromConfig(awsCfg),
		Route53:                route53.NewFromConfig(awsCfg),
		Route53Domains:         route53domains.NewFromConfig(awsCfg),
		S3:                     s3.NewFromConfig(awsCfg),
		S3Control:              s3control.NewFromConfig(awsCfg),
		S3Manager:              newS3ManagerFromConfig(awsCfg),
		SageMaker:              sagemaker.NewFromConfig(awsCfg),
		SecretsManager:         secretsmanager.NewFromConfig(awsCfg),
		SNS:                    sns.NewFromConfig(awsCfg),
		SSM:                    ssm.NewFromConfig(awsCfg),
		SQS:                    sqs.NewFromConfig(awsCfg),
		Waf:                    waf.NewFromConfig(awsCfg),
		WafV2:                  wafv2.NewFromConfig(awsCfg),
		WafRegional:            wafregional.NewFromConfig(awsCfg),
		Workspaces:             workspaces.NewFromConfig(awsCfg),
		IOT:                    iot.NewFromConfig(awsCfg),
	}
}

func newRetryer(maxRetries int, maxBackoff int) func() aws.Retryer {
	return func() aws.Retryer {
		return retry.NewStandard(func(o *retry.StandardOptions) {
			o.MaxAttempts = maxRetries
			o.MaxBackoff = time.Second * time.Duration(maxBackoff)
		})
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

type AwsLogger struct {
	l hclog.Logger
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
