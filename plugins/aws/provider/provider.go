package provider

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/resources/autoscaling"
	"github.com/cloudquery/cq-provider-aws/resources/cloudtrail"
	"github.com/cloudquery/cq-provider-aws/resources/cloudwatch"
	"github.com/cloudquery/cq-provider-aws/resources/cloudwatchlogs"
	"github.com/cloudquery/cq-provider-aws/resources/directconnect"
	"github.com/cloudquery/cq-provider-aws/resources/ec2"
	"github.com/cloudquery/cq-provider-aws/resources/ecr"
	"github.com/cloudquery/cq-provider-aws/resources/ecs"
	"github.com/cloudquery/cq-provider-aws/resources/efs"
	"github.com/cloudquery/cq-provider-aws/resources/elasticbeanstalk"
	"github.com/cloudquery/cq-provider-aws/resources/elbv2"
	"github.com/cloudquery/cq-provider-aws/resources/emr"
	"github.com/cloudquery/cq-provider-aws/resources/fsx"
	"github.com/cloudquery/cq-provider-aws/resources/iam"
	"github.com/cloudquery/cq-provider-aws/resources/kms"
	"github.com/cloudquery/cq-provider-aws/resources/organizations"
	"github.com/cloudquery/cq-provider-aws/resources/rds"
	"github.com/cloudquery/cq-provider-aws/resources/redshift"
	"github.com/cloudquery/cq-provider-aws/resources/s3"
	"github.com/cloudquery/cq-provider-aws/resources/sns"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/sync/errgroup"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"gopkg.in/yaml.v3"
)

type Provider struct {
	cfg             aws.Config
	region          string
	db              *database.Database
	config          Config
	accountID       string
	resourceClients map[string]resource.ClientInterface
	Logger          hclog.Logger
}

type Account struct {
	ID      string
	RoleARN string `yaml:"role_arn"`
}

type Config struct {
	Regions    []string  `yaml:"regions"`
	Accounts   []Account `yaml:"accounts"`
	LogLevel   *string   `yaml:"log_level"`
	MaxRetries *int      `yaml:"max_retries"`
	Resources  []struct {
		Name  string
		Other map[string]interface{} `yaml:",inline"`
	}
}

var globalCollectedResources = map[string]bool{}

type ServiceNewFunction func(awsConfig aws.Config, db *database.Database, log hclog.Logger, accountID string, region string) resource.ClientInterface

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

	p.resourceClients = map[string]resource.ClientInterface{}
	p.Logger.Info("Creating tables if needed")
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

func (p *Provider) Fetch(data []byte) error {
	err := yaml.Unmarshal(data, &p.config)
	ctx := context.Background()
	var ae smithy.APIError
	if err != nil {
		return err
	}

	if len(p.config.Resources) == 0 {
		p.Logger.Info("no resources specified. See available resources: see: https://docs.cloudquery.io/aws/tables-reference")
		return nil
	}
	regions := p.config.Regions
	if len(regions) == 0 {
		regions = allRegions
		p.Logger.Info(fmt.Sprintf("No regions specified in config.yml. Assuming all %d regions", len(regions)))
	}

	if len(p.config.Accounts) == 0 {
		p.config.Accounts = append(p.config.Accounts, Account{
			ID:      "default",
			RoleARN: "default",
		})
	}
	retryOpt := config.WithRetryer(func() aws.Retryer {return retry.NewStandard()})
	for _, account := range p.config.Accounts {
		if account.ID != "default" && account.RoleARN != "" {
			// assume role if specified (SDK takes it from default or env var: AWS_PROFILE)
			p.cfg, err = config.LoadDefaultConfig(ctx, retryOpt)
			if err != nil {
				return err
			}
			p.cfg.Credentials = stscreds.NewAssumeRoleProvider(sts.NewFromConfig(p.cfg), account.RoleARN)

		} else if account.ID != "default" {
			p.cfg, err = config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(account.ID), retryOpt)
		} else {
			p.cfg, err = config.LoadDefaultConfig(ctx, retryOpt)
		}
		if err != nil {
			return err
		}
		svc := sts.NewFromConfig(p.cfg)
		output, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}, func(o *sts.Options) {
			o.Region = "us-east-1"
		})
		if err != nil {
			return err
		}
		p.accountID = *output.Account

		for _, region := range regions {
			p.region = region

			// Find a better way in AWS SDK V2 to decide if a region is disabled.
			_, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}, func(o *sts.Options) {
				o.Region = region
			})
			if err != nil {
				if errors.As(err, &ae) && (ae.ErrorCode() == "InvalidClientTokenId" || ae.ErrorCode() == "OptInRequired") {
					p.Logger.Info("region disabled. skipping...", "region", region)
					continue
				}
				return err
			}

			p.initRegionalClients()
			g := errgroup.Group{}
			for _, resource := range p.config.Resources {
				resourceName := resource.Name
				resourceConfig := resource.Other
				g.Go(func() error {
					return p.collectResource(resourceName, resourceConfig)
				})
			}
			if err := g.Wait(); err != nil {
				return err
			}

		}
		globalCollectedResources = map[string]bool{}
		p.resourceClients = map[string]resource.ClientInterface{}
	}

	return nil
}

func (p *Provider) initRegionalClients() {
	innerLog := p.Logger.With("account_id", p.accountID, "region", p.region)
	for serviceName, newFunc := range regionalServices {
		p.resourceClients[serviceName] = newFunc(p.cfg,
			p.db, innerLog, p.accountID, p.region)
	}
}

var lock = sync.RWMutex{}

func (p *Provider) collectResource(fullResourceName string, config interface{}) error {
	resourcePath := strings.Split(fullResourceName, ".")
	if len(resourcePath) != 2 {
		return fmt.Errorf("resource %s should be in format {service}.{resource}", fullResourceName)
	}
	service := resourcePath[0]
	resourceName := resourcePath[1]

	if globalServices[service] != nil {
		lock.Lock()
		if globalCollectedResources[fullResourceName] {
			lock.Unlock()
			return nil
		}
		globalCollectedResources[fullResourceName] = true
		if p.resourceClients[service] == nil {
			innerLogger := p.Logger.With("account_id", p.accountID)
			p.resourceClients[service] = globalServices[service](p.cfg,
				p.db, innerLogger, p.accountID, p.region)
		}
		lock.Unlock()
	}

	if p.resourceClients[service] == nil {
		return fmt.Errorf("unsupported service %s for resource %s", service, resourceName)
	}

	err := p.resourceClients[service].CollectResource(resourceName, config)
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) {
			switch ae.ErrorCode() {
			case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation":
				p.Logger.Info("Skipping resource. Access denied", "account_id", p.accountID, "region", p.region, "resource", fullResourceName, "error", err)
				return nil
			case "OptInRequired", "SubscriptionRequiredException":
				p.Logger.Info("Skipping resource. Service disabled", "account_id", p.accountID, "region", p.region, "resource", fullResourceName, "error", err)

				return nil
			}
		}
		return err
	}
	return nil
}
