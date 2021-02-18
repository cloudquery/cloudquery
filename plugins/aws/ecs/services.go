package ecs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/aws-sdk-go/aws"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	_         interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID        uint        `gorm:"primarykey"`
	AccountID string
	Region    string

	CapProviderStrategy []*ServiceCapProviderStrategy `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterArn          *string
	CreatedAt           *time.Time
	CreatedBy           *string

	DeploymentConfigurationMaximumPercent        *int32
	DeploymentConfigurationMinimumHealthyPercent *int32

	DeploymentControllerType *string
	DesiredCount             int32
	EnableECSManagedTags     *bool

	HealthCheckGracePeriodSeconds *int32
	LaunchType                    *string
	LoadBalancers                 []*ServiceLoadBalancer `gorm:"constraint:OnDelete:CASCADE;"`

	AssignPublicIp       *string
	SecurityGroups       []*ServiceSecurityGroups `gorm:"constraint:OnDelete:CASCADE;"`
	Subnets              []*ServiceSubnets        `gorm:"constraint:OnDelete:CASCADE;"`
	PendingCount         *int32
	PlacementConstraints []*ServicePlacementConstraint `gorm:"constraint:OnDelete:CASCADE;"`
	PlacementStrategy    []*ServicePlacementStrategy   `gorm:"constraint:OnDelete:CASCADE;"`
	PlatformVersion      *string
	PropagateTags        *string
	RoleArn              *string
	RunningCount         *int32
	SchedulingStrategy   *string
	ServiceArn           *string `neo:"unique"`
	ServiceName          *string
	ServiceRegistries    []*ServiceRegistry `gorm:"constraint:OnDelete:CASCADE;"`
	Status               *string
	Tags                 []*ServiceTag `gorm:"constraint:OnDelete:CASCADE;"`
	TaskDefinition       *string
}

func (Service) TableName() string {
	return "aws_ecs_services"
}

type ServiceCapProviderStrategy struct {
	ID        uint   `gorm:"primarykey"`
	ServiceID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Base             *int32
	CapacityProvider *string
	Weight           *int32
}

func (ServiceCapProviderStrategy) TableName() string {
	return "aws_ecs_service_cap_provider_strategies"
}

type ServiceSecurityGroups struct {
	ID        uint   `gorm:"primarykey"`
	ServiceID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Value string
}

func (ServiceSecurityGroups) TableName() string {
	return "aws_ecs_service_security_groups"
}

type ServiceSubnets struct {
	ID        uint   `gorm:"primarykey"`
	ServiceID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Value string
}

func (ServiceSubnets) TableName() string {
	return "aws_ecs_service_subnets"
}

type ServiceLoadBalancer struct {
	ID        uint   `gorm:"primarykey"`
	ServiceID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ContainerName    *string
	ContainerPort    *int32
	LoadBalancerName *string
	TargetGroupArn   *string
}

func (ServiceLoadBalancer) TableName() string {
	return "aws_ecs_service_load_balancers"
}

type ServicePlacementConstraint struct {
	ID        uint   `gorm:"primarykey"`
	ServiceID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Expression *string
	Type       *string
}

func (ServicePlacementConstraint) TableName() string {
	return "aws_ecs_service_placement_constraints"
}

type ServicePlacementStrategy struct {
	ID        uint   `gorm:"primarykey"`
	ServiceID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Field *string
	Type  *string
}

func (ServicePlacementStrategy) TableName() string {
	return "aws_ecs_service_placement_strategies"
}

type ServiceRegistry struct {
	ID        uint   `gorm:"primarykey"`
	ServiceID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ContainerName *string
	ContainerPort *int32
	Port          *int32
	RegistryArn   *string
}

func (ServiceRegistry) TableName() string {
	return "aws_ecs_service_registries"
}

type ServiceTag struct {
	ID        uint   `gorm:"primarykey"`
	ServiceID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (ServiceTag) TableName() string {
	return "aws_ecs_service_tags"
}

func (c *Client) transformServices(values *[]types.Service) []*Service {
	var tValues []*Service
	for _, value := range *values {
		tValue := Service{
			AccountID:                     c.accountID,
			Region:                        c.region,
			CapProviderStrategy:           c.transformServiceCapacityProviderStrategyItems(&value.CapacityProviderStrategy),
			ClusterArn:                    value.ClusterArn,
			CreatedAt:                     value.CreatedAt,
			CreatedBy:                     value.CreatedBy,
			DesiredCount:                  value.DesiredCount,
			EnableECSManagedTags:          &value.EnableECSManagedTags,
			HealthCheckGracePeriodSeconds: value.HealthCheckGracePeriodSeconds,
			LaunchType:                    aws.String(string(value.LaunchType)),
			LoadBalancers:                 c.transformServiceLoadBalancers(&value.LoadBalancers),
			PendingCount:                  &value.PendingCount,
			PlacementConstraints:          c.transformServicePlacementConstraints(&value.PlacementConstraints),
			PlacementStrategy:             c.transformServicePlacementStrategies(&value.PlacementStrategy),
			PlatformVersion:               value.PlatformVersion,
			PropagateTags:                 aws.String(string(value.PropagateTags)),
			RoleArn:                       value.RoleArn,
			RunningCount:                  &value.RunningCount,
			SchedulingStrategy:            aws.String(string(value.SchedulingStrategy)),
			ServiceArn:                    value.ServiceArn,
			ServiceName:                   value.ServiceName,
			ServiceRegistries:             c.transformServiceRegistries(&value.ServiceRegistries),
			Status:                        value.Status,
			Tags:                          c.transformServiceTags(&value.Tags),
			TaskDefinition:                value.TaskDefinition,
		}
		if value.DeploymentConfiguration != nil {
			tValue.DeploymentConfigurationMaximumPercent = value.DeploymentConfiguration.MaximumPercent
			tValue.DeploymentConfigurationMinimumHealthyPercent = value.DeploymentConfiguration.MinimumHealthyPercent
		}
		if value.DeploymentController != nil {
			tValue.DeploymentControllerType = aws.String(string(value.DeploymentController.Type))
		}
		if value.NetworkConfiguration != nil && value.NetworkConfiguration.AwsvpcConfiguration != nil {
			tValue.SecurityGroups = c.transformServiceAwsVpcConfigurationSecurityGroups(&value.NetworkConfiguration.AwsvpcConfiguration.SecurityGroups)
			tValue.Subnets = c.transformServiceAwsVpcConfigurationSubnets(&value.NetworkConfiguration.AwsvpcConfiguration.Subnets)
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServiceCapacityProviderStrategyItems(values *[]types.CapacityProviderStrategyItem) []*ServiceCapProviderStrategy {
	var tValues []*ServiceCapProviderStrategy
	for _, value := range *values {
		tValue := ServiceCapProviderStrategy{
			AccountID:        c.accountID,
			Region:           c.region,
			Base:             &value.Base,
			CapacityProvider: value.CapacityProvider,
			Weight:           &value.Weight,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformServiceAwsVpcConfigurationSecurityGroups(values *[]string) []*ServiceSecurityGroups {
	var tValues []*ServiceSecurityGroups
	for _, v := range *values {
		tValues = append(tValues, &ServiceSecurityGroups{
			AccountID: c.accountID,
			Region:    c.region,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformServiceAwsVpcConfigurationSubnets(values *[]string) []*ServiceSubnets {
	var tValues []*ServiceSubnets
	for _, v := range *values {
		tValues = append(tValues, &ServiceSubnets{
			AccountID: c.accountID,
			Region:    c.region,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformServiceLoadBalancers(values *[]types.LoadBalancer) []*ServiceLoadBalancer {
	var tValues []*ServiceLoadBalancer
	for _, value := range *values {
		tValue := ServiceLoadBalancer{
			AccountID:        c.accountID,
			Region:           c.region,
			ContainerName:    value.ContainerName,
			ContainerPort:    value.ContainerPort,
			LoadBalancerName: value.LoadBalancerName,
			TargetGroupArn:   value.TargetGroupArn,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServicePlacementConstraints(values *[]types.PlacementConstraint) []*ServicePlacementConstraint {
	var tValues []*ServicePlacementConstraint
	for _, value := range *values {
		tValue := ServicePlacementConstraint{
			AccountID:  c.accountID,
			Region:     c.region,
			Expression: value.Expression,
			Type:       aws.String(string(value.Type)),
		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServicePlacementStrategies(values *[]types.PlacementStrategy) []*ServicePlacementStrategy {
	var tValues []*ServicePlacementStrategy
	for _, value := range *values {
		tValue := ServicePlacementStrategy{
			AccountID: c.accountID,
			Region:    c.region,
			Field:     value.Field,
			Type:      aws.String(string(value.Type)),
		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServiceRegistries(values *[]types.ServiceRegistry) []*ServiceRegistry {
	var tValues []*ServiceRegistry
	for _, value := range *values {
		tValue := ServiceRegistry{
			AccountID:     c.accountID,
			Region:        c.region,
			ContainerName: value.ContainerName,
			ContainerPort: value.ContainerPort,
			Port:          value.Port,
			RegistryArn:   value.RegistryArn,
		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServiceTags(values *[]types.Tag) []*ServiceTag {
	var tValues []*ServiceTag
	for _, value := range *values {
		tValue := ServiceTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) services(clusters *[]string) error {
	ctx := context.Background()
	var listInput ecs.ListServicesInput
	for _, cluster := range *clusters {
		listInput.Cluster = &cluster
		listInput.NextToken = nil
		for {
			listOutput, err := c.svc.ListServices(ctx, &listInput)
			if err != nil {
				return err
			}
			if len(listOutput.ServiceArns) == 0 {
				c.log.Info("Fetched resources", zap.String("resource", "ecs.services"), zap.Int("count", 0))
				break
			}
			output, err := c.svc.DescribeServices(ctx, &ecs.DescribeServicesInput{
				Cluster:  &cluster,
				Services: listOutput.ServiceArns,
			})
			if err != nil {
				return err
			}
			c.db.ChunkedCreate(c.transformServices(&output.Services))
			c.log.Info("Fetched resources", zap.String("resource", "ecs.services"), zap.Int("count", len(output.Services)))

			if listOutput.NextToken == nil {
				break
			}
			listInput.NextToken = listOutput.NextToken
		}
	}
	return nil
}
