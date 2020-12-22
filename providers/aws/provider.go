package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/cloudquery/cloudquery/providers/aws/autoscaling"
	"github.com/cloudquery/cloudquery/providers/aws/cloudtrail"
	"github.com/cloudquery/cloudquery/providers/aws/directconnect"
	"github.com/cloudquery/cloudquery/providers/aws/ec2"
	"github.com/cloudquery/cloudquery/providers/aws/ecr"
	"github.com/cloudquery/cloudquery/providers/aws/ecs"
	"github.com/cloudquery/cloudquery/providers/aws/efs"
	"github.com/cloudquery/cloudquery/providers/aws/elasticbeanstalk"
	"github.com/cloudquery/cloudquery/providers/aws/elbv2"
	"github.com/cloudquery/cloudquery/providers/aws/emr"
	"github.com/cloudquery/cloudquery/providers/aws/fsx"
	"github.com/cloudquery/cloudquery/providers/aws/iam"
	"github.com/cloudquery/cloudquery/providers/aws/kms"
	"github.com/cloudquery/cloudquery/providers/aws/rds"
	"github.com/cloudquery/cloudquery/providers/aws/redshift"
	"github.com/cloudquery/cloudquery/providers/aws/resource"
	"github.com/cloudquery/cloudquery/providers/aws/s3"
	"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"strings"
	"sync"
)

type Provider struct {
	session         *session.Session
	cred            *credentials.Credentials
	region          string
	db              *gorm.DB
	config          Config
	accountID       string
	resourceClients map[string]resource.ClientInterface
	log             *zap.Logger
	logLevel        aws.LogLevelType
}

type Account struct {
	ID      string `mapstructure:"id"`
	RoleARN string `mapstructure:"role_arn"`
}

type Config struct {
	Regions    []string
	Accounts   []Account `mapstructure:"accounts"`
	LogLevel   *string   `mapstructure:"log_level"`
	MaxRetries *int      `mapstructure:"max_retries"`
	Resources  []struct {
		Name  string
		Other map[string]interface{} `mapstructure:",remain"`
	}
}

var globalCollectedResources = map[string]bool{}

type ServiceNewFunction func(session *session.Session, awsConfig *aws.Config, db *gorm.DB, log *zap.Logger, accountID string, region string) resource.ClientInterface

var globalServices = map[string]ServiceNewFunction{
	"iam": iam.NewClient,
	"s3":  s3.NewClient,
}

var regionalServices = map[string]ServiceNewFunction{
	"autoscaling":      autoscaling.NewClient,
	"cloudtrail":       cloudtrail.NewClient,
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
}

var migrateFunctions = []func(*gorm.DB) error{
	autoscaling.MigrateLaunchConfigurations,
	cloudtrail.MigrateTrails,
	directconnect.MigrateGateways,
	ec2.MigrateByoipCidrs,
	ec2.MigrateCustomerGateways,
	ec2.MigrateFlowLogs,
	ec2.MigrateImages,
	ec2.MigrateInstances,
	ec2.MigrateInternetGateways,
	ec2.MigrateNatGateways,
	ec2.MigrateNetworkAcls,
	ec2.MigrateRouteTables,
	ec2.MigrateSecurityGroups,
	ec2.MigrateSubnets,
	ec2.MigrateVPCPeeringConnections,
	ec2.MigrateVPCs,
	ecr.MigrateImageIdentifiers,
	ecs.MigrateClusters,
	efs.MigrateFileSystems,
	elasticbeanstalk.MigrateEnvironments,
	elbv2.MigrateLoadBalancers,
	emr.MigrateClusters,
	fsx.MigrateBackups,
	iam.MigrateGroups,
	iam.MigratePasswordPolicies,
	iam.MigratePolicies,
	iam.MigrateRoles,
	iam.MigrateUsers,
	kms.MigrateKeys,
	rds.MigrateClusters,
	rds.MigrateCertificates,
	rds.MigrateDBSubnetGroups,
	redshift.MigrateClusters,
	redshift.MigrateClusterSubnetGroups,
	s3.MigrateBuckets,
}

func NewProvider(db *gorm.DB, log *zap.Logger) (provider.Interface, error) {
	p := Provider{
		db:              db,
		resourceClients: map[string]resource.ClientInterface{},
		log:             log,
	}
	log.Info("Creating tables if needed")
	for _, f := range migrateFunctions {
		err := f(db)
		if err != nil {
			return nil, err
		}
	}
	return &p, nil
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

func (p *Provider) Run(config interface{}) error {
	err := mapstructure.Decode(config, &p.config)
	if err != nil {
		return err
	}
	if len(p.config.Resources) == 0 {
		return fmt.Errorf("please specify at least 1 resource in config.yml. see: https://docs.cloudquery.io/aws/tables-reference")
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
		p.session, err = session.NewSession()
		if err != nil {
			return err
		}
		if account.ID != "default" {
			// assume role if different account
			cred := stscreds.NewCredentials(p.session, account.RoleARN)
			p.session, err = session.NewSession(&aws.Config{
				Credentials: cred,
			})
			if err != nil {
				return err
			}
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
			if awsErr.Code() == "AccessDenied" || awsErr.Code() == "AccessDeniedException" {
				p.log.Warn("Skipping resource. Access denied", zap.String("resource", fullResourceName), zap.Error(err))
				return
			}
		}
		log.Fatal(err)
	}
}
