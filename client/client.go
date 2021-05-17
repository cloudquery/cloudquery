package client

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sts"
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

type Services struct {
	Autoscaling      AutoscalingClient
	Cloudfront       CloudfrontClient
	Cloudtrail       CloudtrailClient
	Cloudwatch       CloudwatchClient
	CloudwatchLogs   CloudwatchLogsClient
	Directconnect    DirectconnectClient
	ECR              EcrClient
	ECS              EcsClient
	EC2              Ec2Client
	EFS              EfsClient
	Eks              EksClient
	ElasticBeanstalk ElasticbeanstalkClient
	EMR              EmrClient
	SNS              SnsClient
	ELBv2            ElbV2Client
	FSX              FsxClient
	IAM              IamClient
	KMS              KmsClient
	Organizations    OrganizationsClient
	Redshift         RedshiftClient
	Route53          Route53Client
	RDS              RdsClient
	S3               S3Client
	S3Manager        S3ManagerClient
}

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	regions    []string
	logLevel   *string
	maxRetries int
	maxBackoff int
	services   map[string]*Services
	logger     hclog.Logger

	// this is set by table clientList
	AccountID string
	Region    string

	// this is for iam.user specific use-case
	ReportUsers interface{}
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

func NewAwsClient(logger hclog.Logger, regions []string) Client {
	return Client{
		services: map[string]*Services{},
		logger:   logger,
		regions:  regions,
	}
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func (c *Client) Services() *Services {
	return c.services[c.AccountID]
}

func (c *Client) withAccountID(accountID string) *Client {
	return &Client{
		regions:    c.regions,
		logLevel:   c.logLevel,
		maxRetries: c.maxRetries,
		maxBackoff: c.maxBackoff,
		services:   c.services,
		logger:     c.logger.With("account_id", accountID),
		AccountID:  accountID,
		Region:     c.Region,
	}
}

func (c *Client) withAccountIDAndRegion(accountID string, region string) *Client {
	return &Client{
		regions:    c.regions,
		logLevel:   c.logLevel,
		maxRetries: c.maxRetries,
		maxBackoff: c.maxBackoff,
		services:   c.services,
		logger:     c.logger.With("account_id", accountID, "Region", region),
		AccountID:  accountID,
		Region:     region,
	}
}

func (c Client) SetAccountServices(accountId string, s Services) {
	c.services[accountId] = &s
}

func Configure(logger hclog.Logger, providerConfig interface{}) (schema.ClientMeta, error) {
	ctx := context.Background()
	awsConfig := providerConfig.(*Config)
	client := NewAwsClient(logger, awsConfig.Regions)

	if len(client.regions) == 0 {
		client.regions = allRegions
		logger.Info(fmt.Sprintf("No regions specified in config.yml. Assuming all %d regions", len(client.regions)))
	}

	if len(awsConfig.Accounts) == 0 {
		awsConfig.Accounts = append(awsConfig.Accounts, Account{
			ID:      "default",
			RoleARN: "default",
		})
	}

	for _, account := range awsConfig.Accounts {
		var err error
		var awsCfg aws.Config
		// This is a try to solve https://aws.amazon.com/premiumsupport/knowledge-center/iam-validate-access-credentials/
		// with this https://github.com/aws/aws-sdk-go-v2/issues/515#issuecomment-607387352
		defaultRegion := "us-east-1"
		switch {
		case account.ID != "default" && account.RoleARN != "":
			// assume role if specified (SDK takes it from default or env var: AWS_PROFILE)
			awsCfg, err = config.LoadDefaultConfig(ctx, config.WithDefaultRegion(defaultRegion))
			if err != nil {
				return nil, err
			}
			awsCfg.Credentials = stscreds.NewAssumeRoleProvider(sts.NewFromConfig(awsCfg), account.RoleARN)
		case account.ID != "default":
			awsCfg, err = config.LoadDefaultConfig(ctx, config.WithDefaultRegion(defaultRegion), config.WithSharedConfigProfile(account.ID))
		default:
			awsCfg, err = config.LoadDefaultConfig(ctx, config.WithDefaultRegion(defaultRegion))
		}

		if err != nil {
			return nil, err
		}

		if awsConfig.AWSDebug {
			awsCfg.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
		}
		awsCfg.Retryer = newRetryer(awsConfig.MaxRetries, awsConfig.MaxBackoff)
		svc := sts.NewFromConfig(awsCfg)
		output, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}, func(o *sts.Options) {
			o.Region = "aws-global"
		})
		if err != nil {
			return nil, err
		}
		// This is a work-around to skip disabled regions
		// https://github.com/aws/aws-sdk-go-v2/issues/1068
		res, err := ec2.NewFromConfig(awsCfg).DescribeRegions(ctx,
			&ec2.DescribeRegionsInput{AllRegions: false},
			func(o *ec2.Options) {
				o.Region = "us-east-1"
			})
		if err != nil {
			return nil, err
		}
		client.regions = filterDisabledRegions(client.regions, res.Regions)

		if client.AccountID == "" {
			// set default
			client.AccountID = *output.Account
			client.Region = client.regions[0]
		}
		client.SetAccountServices(*output.Account, initServices(awsCfg))
	}

	return &client, nil
}

func initServices(awsCfg aws.Config) Services {
	return Services{
		Autoscaling:      autoscaling.NewFromConfig(awsCfg),
		Cloudfront:       cloudfront.NewFromConfig(awsCfg),
		Cloudtrail:       cloudtrail.NewFromConfig(awsCfg),
		Cloudwatch:       cloudwatch.NewFromConfig(awsCfg),
		CloudwatchLogs:   cloudwatchlogs.NewFromConfig(awsCfg),
		Directconnect:    directconnect.NewFromConfig(awsCfg),
		EC2:              ec2.NewFromConfig(awsCfg),
		ECR:              ecr.NewFromConfig(awsCfg),
		ECS:              ecs.NewFromConfig(awsCfg),
		EFS:              efs.NewFromConfig(awsCfg),
		Eks:              eks.NewFromConfig(awsCfg),
		ElasticBeanstalk: elasticbeanstalk.NewFromConfig(awsCfg),
		EMR:              emr.NewFromConfig(awsCfg),
		FSX:              fsx.NewFromConfig(awsCfg),
		S3:               s3.NewFromConfig(awsCfg),
		SNS:              sns.NewFromConfig(awsCfg),
		ELBv2:            elbv2.NewFromConfig(awsCfg),
		IAM:              iam.NewFromConfig(awsCfg),
		KMS:              kms.NewFromConfig(awsCfg),
		Organizations:    organizations.NewFromConfig(awsCfg),
		RDS:              rds.NewFromConfig(awsCfg),
		Redshift:         redshift.NewFromConfig(awsCfg),
		Route53:          route53.NewFromConfig(awsCfg),
		S3Manager:        newS3ManagerFromConfig(awsCfg),
	}
}

func newRetryer(maxRetries int, maxBackoff int) func() aws.Retryer {
	return func() aws.Retryer {
		return retry.AddWithMaxBackoffDelay(retry.AddWithMaxAttempts(retry.NewStandard(), maxRetries), time.Second*time.Duration(maxBackoff))
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
