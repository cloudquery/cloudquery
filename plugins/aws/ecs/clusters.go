package ecs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cloudquery/cq-provider-aws/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Cluster struct {
	_                                 interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                                uint        `gorm:"primarykey"`
	AccountID                         string
	Region                            string
	ActiveServicesCount               *int32
	AttachmentsStatus                 *string
	CapacityProviders                 *string
	ClusterArn                        *string `neo:"unique"`
	ClusterName                       *string
	DefaultCapacityProviderStrategy   []*ClusterCapacityProviderStrategyItem `gorm:"constraint:OnDelete:CASCADE;"`
	PendingTasksCount                 *int32
	RegisteredContainerInstancesCount *int32
	RunningTasksCount                 *int32
	Settings                          []*ClusterSetting      `gorm:"constraint:OnDelete:CASCADE;"`
	Statistics                        []*ClusterKeyValuePair `gorm:"constraint:OnDelete:CASCADE;"`
	Status                            *string
	Tags                              []*ClusterTag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Cluster) TableName() string {
	return "aws_ecs_clusters"
}

type ClusterKeyValuePair struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Name  *string
	Value *string
}

func (ClusterKeyValuePair) TableName() string {
	return "aws_ecs_cluster_key_value_pairs"
}

type ClusterCapacityProviderStrategyItem struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Base             *int32
	CapacityProvider *string
	Weight           *int32
}

func (ClusterCapacityProviderStrategyItem) TableName() string {
	return "aws_ecs_cluster_capacity_provider_strategy_items"
}

type ClusterSetting struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Name  *string
	Value *string
}

func (ClusterSetting) TableName() string {
	return "aws_ecs_cluster_settings"
}

type ClusterTag struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (ClusterTag) TableName() string {
	return "aws_ecs_cluster_tags"
}

func (c *Client) transformClusterKeyValuePairs(values *[]types.KeyValuePair) []*ClusterKeyValuePair {
	var tValues []*ClusterKeyValuePair
	for _, value := range *values {
		tValues = append(tValues, &ClusterKeyValuePair{
			AccountID: c.accountID,
			Region:    c.region,
			Name:      value.Name,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformClusterCapacityProviderStrategyItems(values *[]types.CapacityProviderStrategyItem) []*ClusterCapacityProviderStrategyItem {
	var tValues []*ClusterCapacityProviderStrategyItem
	for _, value := range *values {
		tValues = append(tValues, &ClusterCapacityProviderStrategyItem{
			AccountID:        c.accountID,
			Region:           c.region,
			Base:             &value.Base,
			CapacityProvider: value.CapacityProvider,
			Weight:           &value.Weight,
		})
	}
	return tValues
}

func (c *Client) transformClusterSettings(values *[]types.ClusterSetting) []*ClusterSetting {
	var tValues []*ClusterSetting
	for _, value := range *values {
		tValues = append(tValues, &ClusterSetting{
			AccountID: c.accountID,
			Region:    c.region,
			Name:      aws.String(string(value.Name)),
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformClusterTags(values *[]types.Tag) []*ClusterTag {
	var tValues []*ClusterTag
	for _, value := range *values {
		tValues = append(tValues, &ClusterTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformClusters(values *[]types.Cluster) []*Cluster {
	var tValues []*Cluster
	for _, value := range *values {
		tValues = append(tValues, &Cluster{
			Region:                            c.region,
			AccountID:                         c.accountID,
			ActiveServicesCount:               &value.ActiveServicesCount,
			AttachmentsStatus:                 value.AttachmentsStatus,
			CapacityProviders:                 common.StringListToString(&value.CapacityProviders),
			ClusterArn:                        value.ClusterArn,
			ClusterName:                       value.ClusterName,
			DefaultCapacityProviderStrategy:   c.transformClusterCapacityProviderStrategyItems(&value.DefaultCapacityProviderStrategy),
			PendingTasksCount:                 &value.PendingTasksCount,
			RegisteredContainerInstancesCount: &value.RegisteredContainerInstancesCount,
			RunningTasksCount:                 &value.RunningTasksCount,
			Settings:                          c.transformClusterSettings(&value.Settings),
			Statistics:                        c.transformClusterKeyValuePairs(&value.Statistics),
			Status:                            value.Status,
			Tags:                              c.transformClusterTags(&value.Tags),
		})
	}
	return tValues
}

var ClusterTables = []interface{}{
	&Cluster{},
	&ClusterKeyValuePair{},
	&ClusterCapacityProviderStrategyItem{},
	&ClusterSetting{},
	&ClusterTag{},

	&Service{},
	&ServiceSecurityGroups{},
	&ServiceSubnets{},
	&ServiceCapProviderStrategy{},
	&ServiceLoadBalancer{},
	&ServicePlacementConstraint{},
	&ServicePlacementStrategy{},
	&ServiceRegistry{},
	&ServiceTag{},
}

func (c *Client) clusters(gConfig interface{}) error {
	ctx := context.Background()
	var config ecs.DescribeClustersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ClusterTables...)
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(&Service{})
	var listConfig ecs.ListClustersInput
	for {
		listOutput, err := c.svc.ListClusters(ctx, &listConfig)
		if err != nil {
			return err
		}
		err = c.services(&listOutput.ClusterArns)
		if err != nil {
			return err
		}
		config.Clusters = listOutput.ClusterArns
		output, err := c.svc.DescribeClusters(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformClusters(&output.Clusters))
		c.log.Info("Fetched resources", zap.String("resource", "ecs.cluster"), zap.Int("count", len(output.Clusters)))

		if listOutput.NextToken == nil {
			break
		}
		listConfig.NextToken = listOutput.NextToken
	}
	return nil
}
