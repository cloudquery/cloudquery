package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
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
	db              *gorm.DB
	config          Config
	accountID       string
	resourceClients map[string]resource.ClientInterface
	log             *zap.Logger
}

type Config struct {
	Region    string
	Resources []struct {
		Name  string
		Other map[string]interface{} `mapstructure:",remain"`
	}
}

type NewResourceFunc func(session *session.Session, db *gorm.DB, log *zap.Logger,
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
}

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

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(p.config.Region)},
	)
	if err != nil {
		return err
	}
	p.session = sess

	svc := sts.New(p.session)
	output, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return err
	}
	p.accountID = aws.StringValue(output.Account)

	for _, resource := range p.config.Resources {
		err := p.collectResource(resource.Name, resource.Other)
		if err != nil {
			return err
		}
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

	if p.resourceClients[service] == nil {
		p.resourceClients[service] = resourceFactory[service](p.session, p.db, p.log, p.accountID, p.config.Region)
	}
	p.db.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("aws_%s_", service),
	}
	return p.resourceClients[service].CollectResource(resourceName, config)
}
