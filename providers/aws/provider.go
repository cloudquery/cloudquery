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
	"gorm.io/gorm/schema"
	"strings"
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
}

type Account struct {
	ID      string `mapstructure:"id"`
	RoleARN string `mapstructure:"role_arn"`
}

type Config struct {
	Regions   []string
	Accounts  []Account `mapstructure:"accounts"`
	Resources []struct {
		Name  string
		Other map[string]interface{} `mapstructure:",remain"`
	}
}

type NewResourceFunc func(session *session.Session, awsConfig *aws.Config, db *gorm.DB, log *zap.Logger,
	accountID string, region string) resource.ClientInterface

var resourceFactory = map[string]NewResourceFunc{
	"ec2":              ec2.NewClient,
	"ecr":              ecr.NewClient,
	"ecs":              ecs.NewClient,
	"autoscaling":      autoscaling.NewClient,
	"efs":              efs.NewClient,
	"elasticbeanstalk": elasticbeanstalk.NewClient,
	"directconnect":    directconnect.NewClient,
	"emr":              emr.NewClient,
	"fsx":              fsx.NewClient,
	"iam":              iam.NewClient,
	"rds":              rds.NewClient,
	"redshift":         redshift.NewClient,
	"s3":               s3.NewClient,
	"elbv2":            elbv2.NewClient,
	"kms":              kms.NewClient,
}

var globalServices = map[string]bool{
	"s3":  true,
	"iam": true,
}

var globalCollectedResources = map[string]bool{}

func NewProvider(db *gorm.DB, log *zap.Logger) (provider.Interface, error) {
	p := Provider{
		db:              db,
		resourceClients: map[string]resource.ClientInterface{},
		log:             log,
	}
	p.db.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "aws_",
	}
	return &p, nil
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

	for _, account := range p.config.Accounts {
		for _, region := range regions {
			sess, err := session.NewSession(&aws.Config{
				Region: aws.String(region)})
			if err != nil {
				return err
			}
			p.region = region
			p.session = sess

			var cred *credentials.Credentials
			if account.ID != "default" {
				cred = stscreds.NewCredentials(sess, account.RoleARN)
			}
			svc := sts.New(p.session, &aws.Config{
				Credentials: cred,
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
			p.cred = cred

			for _, resource := range p.config.Resources {
				err := p.collectResource(resource.Name, resource.Other)
				if err != nil {
					return err
				}
			}
			// TODO: re-enable service cache
			//p.resetClients()
		}
		globalCollectedResources = map[string]bool{}
	}

	return nil
}

func (p *Provider) resetClients() {
	p.resourceClients = map[string]resource.ClientInterface{}
}

func (p *Provider) collectResource(fullResourceName string, config interface{}) error {
	resourcePath := strings.Split(fullResourceName, ".")
	if len(resourcePath) != 2 {
		return fmt.Errorf("resource %s should be in format {service}.{resource}", fullResourceName)
	}
	service := resourcePath[0]
	resourceName := resourcePath[1]

	if resourceFactory[service] == nil {
		return fmt.Errorf("unsupported service %s", service)
	}

	if globalServices[service] {
		if globalCollectedResources[fullResourceName] {
			return nil
		}
		globalCollectedResources[fullResourceName] = true
	}
	// TODO: re-enable service caching
	//if p.resourceClients[service] == nil {
	log := p.log.With(zap.String("account_id", p.accountID), zap.String("region", aws.StringValue(p.session.Config.Region)), zap.String("resource", fullResourceName))
	p.resourceClients[service] = resourceFactory[service](p.session, &aws.Config{Credentials: p.cred},
		p.db, log, p.accountID, p.region)
	//}
	p.db.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("aws_%s_", service),
	}
	return p.resourceClients[service].CollectResource(resourceName, config)
}
