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
			p.initClients()
			var wg sync.WaitGroup
			for _, resource := range p.config.Resources {
				wg.Add(1)
				go p.collectResource(&wg, resource.Name, resource.Other)
			}
			wg.Wait()
		}
		globalCollectedResources = map[string]bool{}
	}

	return nil
}

func (p *Provider) initClients() {
	zapLog := p.log.With(zap.String("account_id", p.accountID), zap.String("region", p.region))
	p.resourceClients["autoscaling"] = autoscaling.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["cloudtrail"] = cloudtrail.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["ec2"] = ec2.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["ecr"] = ecr.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["ecs"] = ecs.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["efs"] = efs.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["elasticbeanstalk"] = elasticbeanstalk.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["directconnect"] = directconnect.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["emr"] = emr.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["fsx"] = fsx.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["iam"] = iam.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["rds"] = rds.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["redshift"] = redshift.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["s3"] = s3.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["elbv2"] = elbv2.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
	p.resourceClients["kms"] = kms.NewClient(p.session, &aws.Config{Region: aws.String(p.region)},
		p.db, zapLog, p.accountID, p.region)
}

func (p *Provider) collectResource(wg *sync.WaitGroup, fullResourceName string, config interface{}) {
	defer wg.Done()
	resourcePath := strings.Split(fullResourceName, ".")
	if len(resourcePath) != 2 {
		log.Fatalf("resource %s should be in format {service}.{resource}", fullResourceName)
	}
	service := resourcePath[0]
	resourceName := resourcePath[1]

	if p.resourceClients[service] == nil {
		log.Fatalf("unsupported service %s", service)
	}

	if globalServices[service] {
		if globalCollectedResources[fullResourceName] {
			return
		}
		globalCollectedResources[fullResourceName] = true
	}

	err := p.resourceClients[service].CollectResource(resourceName, config)
	if err != nil {
		log.Fatal(err)
	}
}
