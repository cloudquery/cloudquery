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
	"github.com/cloudquery/cq-provider-aws/resources/eks"
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
	"github.com/creasty/defaults"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/sync/errgroup"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"gopkg.in/yaml.v3"
)

type Provider struct {
	db              *database.Database
	config          Config
	Logger          hclog.Logger
	retryOpt        config.LoadOptionsFunc
	regions         []string
}

type Account struct {
	ID      string
	RoleARN string `yaml:"role_arn"`
}

type Config struct {
	Regions    []string  `yaml:"regions"`
	Accounts   []Account `yaml:"accounts"`
	LogLevel   *string   `yaml:"log_level"`
	MaxRetries int       `yaml:"max_retries" default:"5"`
	MaxBackoff int       `yaml:"max_backoff" default:"30"`
	Resources  []struct {
		Name  string
		Other map[string]interface{} `yaml:",inline"`
	}
}

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
	"eks":              eks.NewClient,
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
	eks.ClusterTables,
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

func (p *Provider) validateFetchConfig() error {
	if len(p.config.Resources) == 0 {
		p.Logger.Warn("no resources specified. See available resources: see: https://docs.cloudquery.io/aws/tables-reference")
		return nil
	}

	for _, r := range p.config.Resources {
		resourcePath := strings.Split(r.Name, ".")
		if len(resourcePath) != 2 {
			return fmt.Errorf("resource %s should be in format {service}.{resource}", r.Name)
		}
		service := resourcePath[0]
		resourceName := resourcePath[1]

		if regionalServices[service] == nil && globalServices[service] == nil {
			return fmt.Errorf("unsupported service %s for resource %s", service, resourceName)
		}
	}

	return nil
}

func (p *Provider) fetchAccount(accountID string, awsCfg aws.Config, svc *sts.Client) error {
	var ae smithy.APIError
	ctx := context.Background()
	resourceClients := map[string]resource.ClientInterface{}

	innerLog := p.Logger.With("account_id", accountID)
	for serviceName, newFunc := range globalServices {
		resourceClients[serviceName] = newFunc(awsCfg,
			p.db, innerLog, accountID, "us-east-1")
	}
	globalServicesFetched := map[string]bool{}
	for _, region := range p.regions {
		//Find a better way in AWS SDK V2 to decide if a region is disabled.
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

		innerLog := p.Logger.With("account_id", accountID, "region", region)
		for serviceName, newFunc := range regionalServices {
			resourceClients[serviceName] = newFunc(awsCfg,
				p.db, innerLog, accountID, region)
		}

		g := errgroup.Group{}
		for _, r := range p.config.Resources {
			resourcePath := strings.Split(r.Name, ".")
			serviceName := resourcePath[0]
			resourceName := resourcePath[1]
			resourceConfig := r.Other
			if globalServices[serviceName] != nil {
				if globalServicesFetched[serviceName] {
					continue
				}
				globalServicesFetched[serviceName] = true
			}
			g.Go(func() error {
				err := resourceClients[serviceName].CollectResource(resourceName, resourceConfig)
				if err != nil {
					var ae smithy.APIError
					if errors.As(err, &ae) {
						switch ae.ErrorCode() {
						case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation":
							p.Logger.Info("Skipping resource. Access denied", "account_id", accountID, "region", region, "resource", resourceName, "error", err)
							return nil
						case "OptInRequired", "SubscriptionRequiredException":
							p.Logger.Info("Skipping resource. Service disabled", "account_id", accountID, "region", region, "resource", resourceName, "error", err)

							return nil
						}
					}
					return err
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}

	}
	return nil
}

func (p *Provider) Fetch(data []byte) error {
	var awsCfg aws.Config
	ctx := context.Background()
	defaults.MustSet(&p.config)
	if err := yaml.Unmarshal(data, &p.config); err != nil {
		return err
	}

	if err := p.validateFetchConfig(); err != nil {
		return err
	}

	p.regions = p.config.Regions
	if len(p.regions) == 0 {
		p.regions = allRegions
		p.Logger.Info(fmt.Sprintf("No regions specified in config.yml. Assuming all %d regions", len(p.regions)))
	}

	if len(p.config.Accounts) == 0 {
		p.config.Accounts = append(p.config.Accounts, Account{
			ID:      "default",
			RoleARN: "default",
		})
	}
	p.Logger.Info("Configuring SDK retryer", "retry_attempts", p.config.MaxRetries, "max_backoff", p.config.MaxBackoff)
	p.retryOpt = config.WithRetryer(func() aws.Retryer {
		return retry.AddWithMaxBackoffDelay(retry.AddWithMaxAttempts(retry.NewStandard(), p.config.MaxRetries), time.Second*time.Duration(p.config.MaxBackoff))
	})

	g := errgroup.Group{}
	for _, account := range p.config.Accounts {
		var err error
		if account.ID != "default" && account.RoleARN != "" {
			// assume role if specified (SDK takes it from default or env var: AWS_PROFILE)
			awsCfg, err = config.LoadDefaultConfig(ctx, p.retryOpt)
			if err != nil {
				_ = g.Wait()
				return err
			}
			awsCfg.Credentials = stscreds.NewAssumeRoleProvider(sts.NewFromConfig(awsCfg), account.RoleARN)
		} else if account.ID != "default" {
			awsCfg, err = config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(account.ID), p.retryOpt)
		} else {
			awsCfg, err = config.LoadDefaultConfig(ctx, p.retryOpt)
		}
		if err != nil {
			_ = g.Wait()
			return err
		}
		svc := sts.NewFromConfig(awsCfg)
		output, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}, func(o *sts.Options) {
			o.Region = "us-east-1"
		})
		if err != nil {
			_ = g.Wait()
			return err
		}
		accountID := *output.Account
		g.Go(func() error {
			return p.fetchAccount(accountID, awsCfg, svc)
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
