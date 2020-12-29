package ecs

import (
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/cloudquery/cloudquery/providers/common"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	ID        uint `gorm:"primarykey"`
	AccountID string
	Region    string

	CapProviderStrategy []*ServiceCapProviderStrategy `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterArn          *string
	CreatedAt           *time.Time
	CreatedBy           *string

	DeploymentConfigurationMaximumPercent        *int64
	DeploymentConfigurationMinimumHealthyPercent *int64

	DeploymentControllerType *string
	DesiredCount             *int64
	EnableECSManagedTags     *bool

	HealthCheckGracePeriodSeconds *int64
	LaunchType                    *string
	LoadBalancers                 []*ServiceLoadBalancer `gorm:"constraint:OnDelete:CASCADE;"`

	AssignPublicIp       *string
	SecurityGroups       []*ServiceSecurityGroups `gorm:"constraint:OnDelete:CASCADE;"`
	Subnets              []*ServiceSubnets        `gorm:"constraint:OnDelete:CASCADE;"`
	PendingCount         *int64
	PlacementConstraints []*ServicePlacementConstraint `gorm:"constraint:OnDelete:CASCADE;"`
	PlacementStrategy    []*ServicePlacementStrategy   `gorm:"constraint:OnDelete:CASCADE;"`
	PlatformVersion      *string
	PropagateTags        *string
	RoleArn              *string
	RunningCount         *int64
	SchedulingStrategy   *string
	ServiceArn           *string
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
	ID        uint `gorm:"primarykey"`
	ServiceID uint

	Base             *int64
	CapacityProvider *string
	Weight           *int64
}

func (ServiceCapProviderStrategy) TableName() string {
	return "aws_ecs_service_cap_provider_strategies"
}

type ServiceSecurityGroups struct {
	ID        uint `gorm:"primarykey"`
	ServiceID uint
	Value     *string
}

func (ServiceSecurityGroups) TableName() string {
	return "aws_ecs_service_security_groups"
}

type ServiceSubnets struct {
	ID        uint `gorm:"primarykey"`
	ServiceID uint
	Value     *string
}

func (ServiceSubnets) TableName() string {
	return "aws_ecs_service_subnets"
}

type ServiceLoadBalancer struct {
	ID        uint `gorm:"primarykey"`
	ServiceID uint

	ContainerName    *string
	ContainerPort    *int64
	LoadBalancerName *string
	TargetGroupArn   *string
}

func (ServiceLoadBalancer) TableName() string {
	return "aws_ecs_service_load_balancers"
}

type ServicePlacementConstraint struct {
	ID        uint `gorm:"primarykey"`
	ServiceID uint

	Expression *string
	Type       *string
}

func (ServicePlacementConstraint) TableName() string {
	return "aws_ecs_service_placement_constraints"
}

type ServicePlacementStrategy struct {
	ID        uint `gorm:"primarykey"`
	ServiceID uint

	Field *string
	Type  *string
}

func (ServicePlacementStrategy) TableName() string {
	return "aws_ecs_service_placement_strategies"
}

type ServiceRegistry struct {
	ID        uint `gorm:"primarykey"`
	ServiceID uint

	ContainerName *string
	ContainerPort *int64
	Port          *int64
	RegistryArn   *string
}

func (ServiceRegistry) TableName() string {
	return "aws_ecs_service_registries"
}

type ServiceTag struct {
	ID        uint `gorm:"primarykey"`
	ServiceID uint

	Key   *string
	Value *string
}

func (ServiceTag) TableName() string {
	return "aws_ecs_service_tags"
}

func (c *Client) transformServices(values []*ecs.Service) []*Service {
	var tValues []*Service
	for _, value := range values {
		tValue := Service{
			AccountID:                     c.accountID,
			Region:                        c.region,
			CapProviderStrategy:           c.transformServiceCapacityProviderStrategyItems(value.CapacityProviderStrategy),
			ClusterArn:                    value.ClusterArn,
			CreatedAt:                     value.CreatedAt,
			CreatedBy:                     value.CreatedBy,
			DesiredCount:                  value.DesiredCount,
			EnableECSManagedTags:          value.EnableECSManagedTags,
			HealthCheckGracePeriodSeconds: value.HealthCheckGracePeriodSeconds,
			LaunchType:                    value.LaunchType,
			LoadBalancers:                 c.transformServiceLoadBalancers(value.LoadBalancers),
			PendingCount:                  value.PendingCount,
			PlacementConstraints:          c.transformServicePlacementConstraints(value.PlacementConstraints),
			PlacementStrategy:             c.transformServicePlacementStrategies(value.PlacementStrategy),
			PlatformVersion:               value.PlatformVersion,
			PropagateTags:                 value.PropagateTags,
			RoleArn:                       value.RoleArn,
			RunningCount:                  value.RunningCount,
			SchedulingStrategy:            value.SchedulingStrategy,
			ServiceArn:                    value.ServiceArn,
			ServiceName:                   value.ServiceName,
			ServiceRegistries:             c.transformServiceRegistries(value.ServiceRegistries),
			Status:                        value.Status,
			Tags:                          c.transformServiceTags(value.Tags),
			TaskDefinition:                value.TaskDefinition,
		}
		if value.DeploymentConfiguration != nil {
			tValue.DeploymentConfigurationMaximumPercent = value.DeploymentConfiguration.MaximumPercent
			tValue.DeploymentConfigurationMinimumHealthyPercent = value.DeploymentConfiguration.MinimumHealthyPercent
		}
		if value.DeploymentController != nil {
			tValue.DeploymentControllerType = value.DeploymentController.Type
		}
		if value.NetworkConfiguration != nil && value.NetworkConfiguration.AwsvpcConfiguration != nil {
			tValue.SecurityGroups = c.transformServiceAwsVpcConfigurationSecurityGroups(value.NetworkConfiguration.AwsvpcConfiguration.SecurityGroups)
			tValue.Subnets = c.transformServiceAwsVpcConfigurationSubnets(value.NetworkConfiguration.AwsvpcConfiguration.Subnets)
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServiceCapacityProviderStrategyItems(values []*ecs.CapacityProviderStrategyItem) []*ServiceCapProviderStrategy {
	var tValues []*ServiceCapProviderStrategy
	for _, value := range values {
		tValue := ServiceCapProviderStrategy{
			Base:             value.Base,
			CapacityProvider: value.CapacityProvider,
			Weight:           value.Weight,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformServiceAwsVpcConfigurationSecurityGroups(values []*string) []*ServiceSecurityGroups {
	var tValues []*ServiceSecurityGroups
	for _, v := range values {
		tValues = append(tValues, &ServiceSecurityGroups{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformServiceAwsVpcConfigurationSubnets(values []*string) []*ServiceSubnets {
	var tValues []*ServiceSubnets
	for _, v := range values {
		tValues = append(tValues, &ServiceSubnets{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformServiceLoadBalancers(values []*ecs.LoadBalancer) []*ServiceLoadBalancer {
	var tValues []*ServiceLoadBalancer
	for _, value := range values {
		tValue := ServiceLoadBalancer{
			ContainerName:    value.ContainerName,
			ContainerPort:    value.ContainerPort,
			LoadBalancerName: value.LoadBalancerName,
			TargetGroupArn:   value.TargetGroupArn,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServicePlacementConstraints(values []*ecs.PlacementConstraint) []*ServicePlacementConstraint {
	var tValues []*ServicePlacementConstraint
	for _, value := range values {
		tValue := ServicePlacementConstraint{
			Expression: value.Expression,
			Type:       value.Type,
		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServicePlacementStrategies(values []*ecs.PlacementStrategy) []*ServicePlacementStrategy {
	var tValues []*ServicePlacementStrategy
	for _, value := range values {
		tValue := ServicePlacementStrategy{
			Field: value.Field,
			Type:  value.Type,
		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServiceRegistries(values []*ecs.ServiceRegistry) []*ServiceRegistry {
	var tValues []*ServiceRegistry
	for _, value := range values {
		tValue := ServiceRegistry{
			ContainerName: value.ContainerName,
			ContainerPort: value.ContainerPort,
			Port:          value.Port,
			RegistryArn:   value.RegistryArn,
		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformServiceTags(values []*ecs.Tag) []*ServiceTag {
	var tValues []*ServiceTag
	for _, value := range values {
		tValue := ServiceTag{
			Key:   value.Key,
			Value: value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) services(clusters []*string) error {

	var listInput ecs.ListServicesInput
	for _, cluster := range clusters {
		listInput.Cluster = cluster
		for {
			listOutput, err := c.svc.ListServices(&listInput)
			if err != nil {
				return err
			}
			output, err := c.svc.DescribeServices(&ecs.DescribeServicesInput{
				Cluster:  cluster,
				Services: listOutput.ServiceArns,
			})
			if err != nil {
				return err
			}
			common.ChunkedCreate(c.db, c.transformServices(output.Services))
			c.log.Info("Fetched resources", zap.String("resource", "ecs.services"), zap.Int("count", len(output.Services)))

			if listOutput.NextToken == nil {
				break
			}
			listInput.NextToken = listOutput.NextToken
		}
	}
	return nil
}
