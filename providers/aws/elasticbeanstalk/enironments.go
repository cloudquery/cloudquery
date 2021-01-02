package elasticbeanstalk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Environment struct {
	_                            interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                           uint        `gorm:"primarykey"`
	AccountID                    string
	Region                       string
	AbortableOperationInProgress *bool
	ApplicationName              *string
	CNAME                        *string
	DateCreated                  *time.Time
	DateUpdated                  *time.Time
	Description                  *string
	EndpointURL                  *string
	EnvironmentArn               *string `neo:"unique"`
	EnvironmentId                *string
	EnvironmentLinks             []*EnvironmentLink `gorm:"constraint:OnDelete:CASCADE;"`
	EnvironmentName              *string
	Health                       *string
	HealthStatus                 *string
	OperationsRole               *string
	PlatformArn                  *string

	LoadBalancerDomain    *string
	LoadBalancerListeners []*EnvironmentListener `gorm:"constraint:OnDelete:CASCADE;"`
	LoadBalancerName      *string

	SolutionStackName *string
	Status            *string
	TemplateName      *string

	TierName    *string
	TierType    *string
	TierVersion *string

	VersionLabel *string
}

func (Environment) TableName() string {
	return "aws_elasticbeanstalk_environments"
}

type EnvironmentLink struct {
	ID            uint   `gorm:"primarykey"`
	EnvironmentID uint   `neo:"ignore"`
	AccountID     string `gorm:"-"`
	Region        string `gorm:"-"`

	EnvironmentName *string
	LinkName        *string
}

func (EnvironmentLink) TableName() string {
	return "aws_elasticbeanstalk_environment_links"
}

type EnvironmentListener struct {
	ID            uint   `gorm:"primarykey"`
	EnvironmentID uint   `neo:"ignore"`
	AccountID     string `gorm:"-"`
	Region        string `gorm:"-"`

	Port     *int64
	Protocol *string
}

func (EnvironmentListener) TableName() string {
	return "aws_elasticbeanstalk_environment_listeners"
}

func (c *Client) transformEnvironmentLink(value *elasticbeanstalk.EnvironmentLink) *EnvironmentLink {
	return &EnvironmentLink{
		AccountID:       c.accountID,
		Region:          c.region,
		EnvironmentName: value.EnvironmentName,
		LinkName:        value.LinkName,
	}
}

func (c *Client) transformEnvironmentDescriptionEnvironmentLinks(values []*elasticbeanstalk.EnvironmentLink) []*EnvironmentLink {
	var tValues []*EnvironmentLink
	for _, v := range values {
		tValues = append(tValues, c.transformEnvironmentLink(v))
	}
	return tValues
}

func (c *Client) transformEnvironmentListener(value *elasticbeanstalk.Listener) *EnvironmentListener {
	return &EnvironmentListener{
		AccountID: c.accountID,
		Region:    c.region,
		Port:      value.Port,
		Protocol:  value.Protocol,
	}
}

func (c *Client) transformEnvironmentListeners(values []*elasticbeanstalk.Listener) []*EnvironmentListener {
	var tValues []*EnvironmentListener
	for _, v := range values {
		tValues = append(tValues, c.transformEnvironmentListener(v))
	}
	return tValues
}

func (c *Client) transformEnvironment(value *elasticbeanstalk.EnvironmentDescription) *Environment {
	res := Environment{
		Region:                       c.region,
		AccountID:                    c.accountID,
		AbortableOperationInProgress: value.AbortableOperationInProgress,
		ApplicationName:              value.ApplicationName,
		CNAME:                        value.CNAME,
		DateCreated:                  value.DateCreated,
		DateUpdated:                  value.DateUpdated,
		Description:                  value.Description,
		EndpointURL:                  value.EndpointURL,
		EnvironmentArn:               value.EnvironmentArn,
		EnvironmentId:                value.EnvironmentId,
		EnvironmentName:              value.EnvironmentName,
		Health:                       value.Health,
		HealthStatus:                 value.HealthStatus,
		OperationsRole:               value.OperationsRole,
		PlatformArn:                  value.PlatformArn,
		SolutionStackName:            value.SolutionStackName,
		Status:                       value.Status,
		TemplateName:                 value.TemplateName,
		VersionLabel:                 value.VersionLabel,
	}

	if value.Tier != nil {
		res.TierName = value.Tier.Name
		res.TierType = value.Tier.Type
		res.TierVersion = value.Tier.Version
	}

	if value.EnvironmentLinks != nil {
		res.EnvironmentLinks = c.transformEnvironmentDescriptionEnvironmentLinks(value.EnvironmentLinks)
	}

	if value.Resources != nil && value.Resources.LoadBalancer != nil {
		res.LoadBalancerDomain = value.Resources.LoadBalancer.Domain
		res.LoadBalancerListeners = c.transformEnvironmentListeners(value.Resources.LoadBalancer.Listeners)
		res.LoadBalancerName = value.Resources.LoadBalancer.LoadBalancerName
	}

	return &res
}

func (c *Client) transformEnvironments(values []*elasticbeanstalk.EnvironmentDescription) []*Environment {
	var tValues []*Environment
	for _, v := range values {
		tValues = append(tValues, c.transformEnvironment(v))
	}
	return tValues
}

var EnvironmentTables = []interface{}{
	&Environment{},
	&EnvironmentLink{},
	&EnvironmentListener{},
}

func (c *Client) environments(gConfig interface{}) error {
	var config elasticbeanstalk.DescribeEnvironmentsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(EnvironmentTables...)
	for {
		output, err := c.svc.DescribeEnvironments(&config)
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformEnvironments(output.Environments))
		c.log.Info("Fetched resources", zap.String("resource", "elasticbeanstalk.environments"), zap.Int("count", len(output.Environments)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
