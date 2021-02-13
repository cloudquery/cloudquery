package main

import (
	"fmt"
	"github.com/cloudquery/cloudquery/cqlog"
	"github.com/cloudquery/cloudquery/sdk"
	"log"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/autoscaling"
	"github.com/cloudquery/cq-provider-aws/cloudtrail"
	"github.com/cloudquery/cq-provider-aws/cloudwatch"
	"github.com/cloudquery/cq-provider-aws/cloudwatchlogs"
	"github.com/cloudquery/cq-provider-aws/directconnect"
	"github.com/cloudquery/cq-provider-aws/ec2"
	"github.com/cloudquery/cq-provider-aws/ecr"
	"github.com/cloudquery/cq-provider-aws/ecs"
	"github.com/cloudquery/cq-provider-aws/efs"
	"github.com/cloudquery/cq-provider-aws/elasticbeanstalk"
	"github.com/cloudquery/cq-provider-aws/elbv2"
	"github.com/cloudquery/cq-provider-aws/emr"
	"github.com/cloudquery/cq-provider-aws/fsx"
	"github.com/cloudquery/cq-provider-aws/iam"
	"github.com/cloudquery/cq-provider-aws/kms"
	"github.com/cloudquery/cq-provider-aws/organizations"
	"github.com/cloudquery/cq-provider-aws/rds"
	"github.com/cloudquery/cq-provider-aws/redshift"
	"github.com/cloudquery/cq-provider-aws/resource"
	"github.com/cloudquery/cq-provider-aws/s3"
	"github.com/cloudquery/cq-provider-aws/sns"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Provider struct {
	session         *session.Session
	cred            *credentials.Credentials
	region          string
	db              *database.Database
	config          Config
	accountID       string
	resourceClients map[string]resource.ClientInterface
	log             *zap.Logger
	logLevel        aws.LogLevelType
}

type Account struct {
	ID      string
	RoleARN string
}

type Config struct {
	Regions    []string
	Accounts   []Account
	LogLevel   *string
	MaxRetries *int
	Resources  []struct {
		Name  string
		Other map[string]interface{} `yaml:",inline"`
	}
}

var globalCollectedResources = map[string]bool{}

type ServiceNewFunction func(session *session.Session, awsConfig *aws.Config, db *database.Database, log *zap.Logger, accountID string, region string) resource.ClientInterface

var globalServices = map[string]ServiceNewFunction{
	"iam":           iam.NewClient,
	"s3":            s3.NewClient,
	"organizations": organizations.NewClient,
}

var regionalServices = map[string]ServiceNewFunction{
	"autoscaling":      autoscaling.NewClient,
	"cloudtrail":       cloudtrail.NewClient,
	"cloudwatchlogs":   cloudwatchlogs.NewClient,
	"cloudwatch":       cloudwatch.NewClient,
	"directconnect":    directconnect.NewClient,
	"ec2":              ec2.NewClient,
	"ecr":              ecr.NewClient,
	"ecs":              ecs.NewClient,
	"efs":              efs.NewClient,
	"elasticbeanstalk": elasticbeanstalk.NewClient,
	"elbv2":            elbv2.NewClient,
	"emr":              emr.NewClient,
	"fsx":              fsx.NewClient,
	"kms":              kms.NewClient,
	"rds":              rds.NewClient,
	"redshift":         redshift.NewClient,
	"sns":              sns.NewClient,
}

var tablesArr = [][]interface{}{
	autoscaling.LaunchConfigurationTables,
	cloudtrail.TrailTables,
	cloudwatchlogs.MetricFilterTables,
	cloudwatch.MetricAlarmTables,
	directconnect.GatewayTables,
	ec2.ByoipCidrTables,
	ec2.CustomerGatewayTables,
	ec2.FlowLogsTables,
	ec2.ImageTables,
	ec2.InstanceTables,
	ec2.InternetGatewayTables,
	ec2.NatGatewayTables,
	ec2.NetworkAclTables,
	ec2.RouteTableTables,
	ec2.SecurityGroupTables,
	ec2.SubnetTables,
	ec2.VPCPeeringConnectionTables,
	ec2.VPCTables,
	ecr.ImageTables,
	ecs.ClusterTables,
	efs.FileSystemTables,
	elasticbeanstalk.EnvironmentTables,
	elbv2.LoadBalancerTables,
	elbv2.TargetGroupTables,
	emr.ClusterTables,
	fsx.BackupTables,
	iam.GroupTables,
	iam.PasswordPolicyTables,
	iam.PolicyTables,
	iam.RoleTables,
	iam.UserTables,
	iam.VirtualMFADeviceTables,
	kms.KeyTables,
	organizations.AccountTables,
	rds.ClusterTables,
	rds.CertificateTables,
	rds.DBSubnetGroupTables,
	redshift.ClusterTables,
	redshift.ClusterSubnetGroupTables,
	s3.BucketTables,
	sns.SubscriptionTables,
	sns.TopicTables,
}

func (p *Provider) Init(driver string, dsn string, verbose bool) error {
	var err error
	p.db, err = database.Open(driver, dsn)
	if err != nil {
		return err
	}

	zapLogger, err := cqlog.NewLogger(verbose)
	p.log = zapLogger
	p.resourceClients = map[string]resource.ClientInterface{}
	p.log.Info("Creating tables if needed")
	for _, tables := range tablesArr {
		err := p.db.AutoMigrate(tables...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Provider) GenConfig() (string, error) {
	return configYaml, nil
}

func (p *Provider) parseLogLevel() {
	if p.config.LogLevel == nil {
		return
	}
	switch *p.config.LogLevel {
	case "debug", "debug_with_signing":
		p.logLevel = aws.LogDebug
	case "debug_with_http_body":
		p.logLevel = aws.LogDebugWithSigning
	case "debug_with_request_retries":
		p.logLevel = aws.LogDebugWithRequestRetries
	case "debug_with_request_error":
		p.logLevel = aws.LogDebugWithRequestErrors
	case "debug_with_event_stream_body":
		p.logLevel = aws.LogDebugWithEventStreamBody
	default:
		log.Fatalf("unknown log_level %s", *p.config.LogLevel)
	}
}

func (p *Provider) Fetch(data []byte) error {
	err := yaml.Unmarshal(data, &p.config)
	if err != nil {
		return err
	}

	if len(p.config.Resources) == 0 {
		p.log.Info("no resources specified. See available resources: see: https://docs.cloudquery.io/aws/tables-reference")
		return nil
	}
	regions := p.config.Regions
	if len(regions) == 0 {
		resolver := endpoints.DefaultResolver()
		partitions := resolver.(endpoints.EnumPartitions).Partitions()
		for _, p := range partitions {
			if p.ID() == "aws" {
				for id, _ := range p.Regions() {
					regions = append(regions, id)
				}
			}
		}
		p.log.Info(fmt.Sprintf("No regions specified in config.yml. Assuming all %d regions", len(regions)))
	}

	if len(p.config.Accounts) == 0 {
		p.config.Accounts = append(p.config.Accounts, Account{
			ID:      "default",
			RoleARN: "default",
		})
	}

	p.parseLogLevel()

	for _, account := range p.config.Accounts {

		if account.ID != "default" && account.RoleARN != "" {
			// assume role if specified (SDK takes it from default or env var: AWS_PROFILE)
			p.session, err = session.NewSession()
			cred := stscreds.NewCredentials(p.session, account.RoleARN)
			p.session, err = session.NewSession(&aws.Config{
				Credentials: cred,
			})
			if err != nil {
				return err
			}
		} else if account.ID != "default" {
			p.session, err = session.NewSession(&aws.Config{Credentials: credentials.NewSharedCredentials("", account.ID)})
		} else {
			p.session, err = session.NewSession()
		}
		if err != nil {
			return err
		}

		for _, region := range regions {
			p.region = region

			svc := sts.New(p.session, &aws.Config{
				Region: aws.String(region),
			})
			output, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
			if err != nil {
				if awsErr, ok := err.(awserr.Error); ok {
					if awsErr.Code() == "InvalidClientTokenId" {
						p.log.Debug("Region is disabled (to enable see: https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-enable). skipping...",
							zap.String("account_id", p.accountID),
							zap.String("region", region))
						continue
					}
				}
				return err
			}
			p.accountID = aws.StringValue(output.Account)
			p.initRegionalClients()
			var wg sync.WaitGroup
			for _, resource := range p.config.Resources {
				wg.Add(1)
				go p.collectResource(&wg, resource.Name, resource.Other)
			}
			wg.Wait()
		}
		globalCollectedResources = map[string]bool{}
		p.resourceClients = map[string]resource.ClientInterface{}
	}

	return nil
}

func (p *Provider) initRegionalClients() {
	zapLog := p.log.With(zap.String("account_id", p.accountID), zap.String("region", p.region))
	for serviceName, newFunc := range regionalServices {
		p.resourceClients[serviceName] = newFunc(p.session,
			&aws.Config{
				Region:     aws.String(p.region),
				LogLevel:   &p.logLevel,
				MaxRetries: p.config.MaxRetries,
			},
			p.db, zapLog, p.accountID, p.region)
	}
}

var lock = sync.RWMutex{}

func (p *Provider) collectResource(wg *sync.WaitGroup, fullResourceName string, config interface{}) {
	defer wg.Done()
	resourcePath := strings.Split(fullResourceName, ".")
	if len(resourcePath) != 2 {
		log.Fatalf("resource %s should be in format {service}.{resource}", fullResourceName)
	}
	service := resourcePath[0]
	resourceName := resourcePath[1]

	if globalServices[service] != nil {
		lock.Lock()
		if globalCollectedResources[fullResourceName] {
			lock.Unlock()
			return
		}
		globalCollectedResources[fullResourceName] = true
		if p.resourceClients[service] == nil {
			zapLog := p.log.With(zap.String("account_id", p.accountID))
			p.resourceClients[service] = globalServices[service](p.session,
				&aws.Config{Region: aws.String(p.region),
					LogLevel:   &p.logLevel,
					MaxRetries: p.config.MaxRetries,
				},
				p.db, zapLog, p.accountID, p.region)
		}
		lock.Unlock()
	}

	if p.resourceClients[service] == nil {
		log.Fatalf("unsupported service %s for resource %s", service, resourceName)
	}

	err := p.resourceClients[service].CollectResource(resourceName, config)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation":
				p.log.Info("Skipping resource. Access denied", zap.String("account_id", p.accountID), zap.String("resource", fullResourceName), zap.Error(err))
				return
			}
		}
		log.Fatal(err)
	}
}

func main() {
	sdk.ServePlugin(&Provider{})
}
