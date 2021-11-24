package client

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/guardduty"

	"github.com/aws/smithy-go/logging"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
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
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3control "github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

// Provider Client passed as meta to all table fetchers

var allRegions = []string{
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
	"af-south-1",
	"ap-east-1",
	"ap-south-1",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ca-central-1",
	"cn-north-1",
	"cn-northwest-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"eu-south-1",
	"eu-north-1",
	"me-south-1",
	"sa-east-1",
}

const (
	defaultRegion              = "us-east-1"
	awsFailedToConfigureErrMsg = "failed to configure provider for account %s. AWS Error: %w"
	defaultVar                 = "default"
)

type Services struct {
	Analyzer             AnalyzerClient
	Autoscaling          AutoscalingClient
	Cloudfront           CloudfrontClient
	Cloudtrail           CloudtrailClient
	Cloudwatch           CloudwatchClient
	CloudwatchLogs       CloudwatchLogsClient
	CognitoIdentityPools CognitoIdentityPoolsClient
	CognitoUserPools     CognitoUserPoolsClient
	Directconnect        DirectconnectClient
	DMS                  DatabasemigrationserviceClient
	ECR                  EcrClient
	ECS                  EcsClient
	EC2                  Ec2Client
	EFS                  EfsClient
	Eks                  EksClient
	ElasticBeanstalk     ElasticbeanstalkClient
	ElasticSearch        ElasticSearch
	EMR                  EmrClient
	SNS                  SnsClient
	ELBv1                ElbV1Client
	ELBv2                ElbV2Client
	FSX                  FsxClient
	IAM                  IamClient
	KMS                  KmsClient
	MQ                   MQClient
	Organizations        OrganizationsClient
	Redshift             RedshiftClient
	Route53              Route53Client
	Route53Domains       Route53DomainsClient
	RDS                  RdsClient
	S3                   S3Client
	S3Control            S3ControlClient
	S3Manager            S3ManagerClient
	SQS                  SQSClient
	Apigateway           ApigatewayClient
	Apigatewayv2         Apigatewayv2Client
	Lambda               LambdaClient
	ConfigService        ConfigServiceClient
	Waf                  WafClient
	WafV2                WafV2Client
	GuardDuty            GuardDutyClient
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
		s.services[accountId] = make(map[string]*Services, len(allRegions))
	}
	s.services[accountId][region] = &services
}

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	accounts        []Account
	regions         []string
	logLevel        *string
	maxRetries      int
	maxBackoff      int
	ServicesManager ServicesManager
	logger          hclog.Logger
	// this is set by table clientList
	AccountID string
	Region    string
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

func NewAwsClient(logger hclog.Logger, accounts []Account, regions []string) Client {
	return Client{
		ServicesManager: ServicesManager{
			services: ServicesAccountRegionMap{},
		},
		logger:   logger,
		accounts: accounts,
		regions:  regions,
	}
}
func (c *Client) Logger() hclog.Logger {
	return &awsLogger{c.logger, c.accounts}
}

func (c *Client) Services() *Services {
	return c.ServicesManager.ServicesByAccountAndRegion(c.AccountID, c.Region)
}

func (c *Client) withAccountID(accountID string) *Client {
	return &Client{
		accounts:        c.accounts,
		regions:         c.regions,
		logLevel:        c.logLevel,
		maxRetries:      c.maxRetries,
		maxBackoff:      c.maxBackoff,
		ServicesManager: c.ServicesManager,
		logger:          c.logger.With("account_id", obfuscateAccountId(accountID)),
		AccountID:       accountID,
		Region:          c.Region,
	}
}

func (c *Client) withAccountIDAndRegion(accountID string, region string) *Client {

	return &Client{
		accounts:        c.accounts,
		regions:         c.regions,
		logLevel:        c.logLevel,
		maxRetries:      c.maxRetries,
		maxBackoff:      c.maxBackoff,
		ServicesManager: c.ServicesManager,
		logger:          c.logger.With("account_id", obfuscateAccountId(accountID), "Region", region),
		AccountID:       accountID,
		Region:          region,
	}
}

func Configure(logger hclog.Logger, providerConfig interface{}) (schema.ClientMeta, error) {
	ctx := context.Background()
	awsConfig := providerConfig.(*Config)
	client := NewAwsClient(logger, awsConfig.Accounts, awsConfig.Regions)

	if len(client.regions) == 0 {
		client.regions = allRegions
		logger.Info(fmt.Sprintf("No regions specified in config.yml. Assuming all %d regions", len(client.regions)))
	}

	if len(awsConfig.Accounts) == 0 {
		awsConfig.Accounts = append(awsConfig.Accounts, Account{
			ID:        defaultVar,
			AccountID: defaultVar,
			RoleARN:   defaultVar,
		})
	}

	for _, account := range awsConfig.Accounts {
		var err error
		var awsCfg aws.Config

		// account id can be defined in account block label or in block attr
		// we take the block att as default and use the block label if the attr is not defined
		var accountID = account.AccountID
		if accountID == "" {
			accountID = account.ID
		}

		// This is a try to solve https://aws.amazon.com/premiumsupport/knowledge-center/iam-validate-access-credentials/
		// with this https://github.com/aws/aws-sdk-go-v2/issues/515#issuecomment-607387352
		switch {
		case accountID != "default" && account.RoleARN != "":
			// assume role if specified (SDK takes it from default or env var: AWS_PROFILE)
			awsCfg, err = config.LoadDefaultConfig(
				ctx,
				config.WithDefaultRegion(defaultRegion),
				config.WithRetryer(newRetryer(awsConfig.MaxRetries, awsConfig.MaxBackoff)),
			)
			if err != nil {
				return nil, fmt.Errorf(awsFailedToConfigureErrMsg, accountID, err)
			}
			opts := make([]func(*stscreds.AssumeRoleOptions), 0, 1)
			if account.ExternalID != "" {
				opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
					opts.ExternalID = &account.ExternalID
				})
			}
			provider := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(awsCfg), account.RoleARN, opts...)
			awsCfg.Credentials = aws.NewCredentialsCache(provider)
		case accountID != "default":
			awsCfg, err = config.LoadDefaultConfig(
				ctx,
				config.WithDefaultRegion(defaultRegion),
				config.WithSharedConfigProfile(accountID),
				config.WithRetryer(newRetryer(awsConfig.MaxRetries, awsConfig.MaxBackoff)),
			)
		default:
			awsCfg, err = config.LoadDefaultConfig(
				ctx,
				config.WithDefaultRegion(defaultRegion),
				config.WithRetryer(newRetryer(awsConfig.MaxRetries, awsConfig.MaxBackoff)),
			)
		}

		if err != nil {
			return nil, fmt.Errorf(awsFailedToConfigureErrMsg, accountID, err)
		}

		if awsConfig.AWSDebug {
			awsCfg.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
			awsCfg.Logger = AwsLogger{logger.With("account", obfuscateAccountId(accountID))}
		}
		svc := sts.NewFromConfig(awsCfg)
		output, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}, func(o *sts.Options) {
			o.Region = "aws-global"
		})
		if err != nil {
			return nil, fmt.Errorf(awsFailedToConfigureErrMsg, accountID, err)
		}
		// This is a work-around to skip disabled regions
		// https://github.com/aws/aws-sdk-go-v2/issues/1068
		res, err := ec2.NewFromConfig(awsCfg).DescribeRegions(ctx,
			&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)},
			func(o *ec2.Options) {
				o.Region = "us-east-1"
			})
		if err != nil {
			return nil, fmt.Errorf("failed to find disabled regions for account %s. AWS Error: %w", accountID, err)
		}
		client.regions = filterDisabledRegions(client.regions, res.Regions)

		if len(client.regions) == 0 {
			return nil, fmt.Errorf("no enabled regions provided in config for account %s", accountID)
		}

		if client.AccountID == "" {
			// set default
			client.AccountID = *output.Account
			client.Region = client.regions[0]
			client.accounts = append(client.accounts, Account{ID: *output.Account, RoleARN: *output.Arn})
		}
		for _, region := range client.regions {
			client.ServicesManager.InitServicesForAccountAndRegion(*output.Account, region, initServices(region, awsCfg))
		}
	}

	return &client, nil
}

func initServices(region string, c aws.Config) Services {
	awsCfg := c.Copy()
	awsCfg.Region = region
	return Services{
		Analyzer:             accessanalyzer.NewFromConfig(awsCfg),
		Apigateway:           apigateway.NewFromConfig(awsCfg),
		Apigatewayv2:         apigatewayv2.NewFromConfig(awsCfg),
		Autoscaling:          autoscaling.NewFromConfig(awsCfg),
		Cloudfront:           cloudfront.NewFromConfig(awsCfg),
		Cloudtrail:           cloudtrail.NewFromConfig(awsCfg),
		Cloudwatch:           cloudwatch.NewFromConfig(awsCfg),
		CloudwatchLogs:       cloudwatchlogs.NewFromConfig(awsCfg),
		CognitoIdentityPools: cognitoidentity.NewFromConfig(awsCfg),
		CognitoUserPools:     cognitoidentityprovider.NewFromConfig(awsCfg),
		ConfigService:        configservice.NewFromConfig(awsCfg),
		Directconnect:        directconnect.NewFromConfig(awsCfg),
		DMS:                  databasemigrationservice.NewFromConfig(awsCfg),
		EC2:                  ec2.NewFromConfig(awsCfg),
		ECR:                  ecr.NewFromConfig(awsCfg),
		ECS:                  ecs.NewFromConfig(awsCfg),
		EFS:                  efs.NewFromConfig(awsCfg),
		Eks:                  eks.NewFromConfig(awsCfg),
		ElasticBeanstalk:     elasticbeanstalk.NewFromConfig(awsCfg),
		ElasticSearch:        elasticsearchservice.NewFromConfig(awsCfg),
		ELBv1:                elbv1.NewFromConfig(awsCfg),
		ELBv2:                elbv2.NewFromConfig(awsCfg),
		EMR:                  emr.NewFromConfig(awsCfg),
		FSX:                  fsx.NewFromConfig(awsCfg),
		GuardDuty:            guardduty.NewFromConfig(awsCfg),
		IAM:                  iam.NewFromConfig(awsCfg),
		KMS:                  kms.NewFromConfig(awsCfg),
		Lambda:               lambda.NewFromConfig(awsCfg),
		MQ:                   mq.NewFromConfig(awsCfg),
		Organizations:        organizations.NewFromConfig(awsCfg),
		RDS:                  rds.NewFromConfig(awsCfg),
		Redshift:             redshift.NewFromConfig(awsCfg),
		Route53:              route53.NewFromConfig(awsCfg),
		Route53Domains:       route53domains.NewFromConfig(awsCfg),
		S3:                   s3.NewFromConfig(awsCfg),
		S3Control:            s3control.NewFromConfig(awsCfg),
		S3Manager:            newS3ManagerFromConfig(awsCfg),
		SNS:                  sns.NewFromConfig(awsCfg),
		SQS:                  sqs.NewFromConfig(awsCfg),
		Waf:                  waf.NewFromConfig(awsCfg),
		WafV2:                wafv2.NewFromConfig(awsCfg),
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
		regionsMap[*r.RegionName] = true
	}

	var filteredRegions []string
	for _, r := range regions {
		if regionsMap[r] {
			filteredRegions = append(filteredRegions, r)
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
