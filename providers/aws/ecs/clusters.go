package ecs

import (
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Cluster struct {
	ID                                uint `gorm:"primarykey"`
	AccountID                         string
	Region                            string
	ActiveServicesCount               *int64
	AttachmentsStatus                 *string
	CapacityProviders                 *string
	ClusterArn                        *string
	ClusterName                       *string
	DefaultCapacityProviderStrategy   []*ClusterCapacityProviderStrategyItem `gorm:"constraint:OnDelete:CASCADE;"`
	PendingTasksCount                 *int64
	RegisteredContainerInstancesCount *int64
	RunningTasksCount                 *int64
	Settings                          []*ClusterSetting      `gorm:"constraint:OnDelete:CASCADE;"`
	Statistics                        []*ClusterKeyValuePair `gorm:"constraint:OnDelete:CASCADE;"`
	Status                            *string
	Tags                              []*ClusterTag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Cluster) TableName() string {
	return "aws_ecs_clusters"
}

type ClusterKeyValuePair struct {
	ID        uint `gorm:"primarykey"`
	ClusterID uint
	Name      *string
	Value     *string
}

func (ClusterKeyValuePair) TableName() string {
	return "aws_ecs_cluster_key_value_pairs"
}

type ClusterCapacityProviderStrategyItem struct {
	ID               uint `gorm:"primarykey"`
	ClusterID        uint
	Base             *int64
	CapacityProvider *string
	Weight           *int64
}

func (ClusterCapacityProviderStrategyItem) TableName() string {
	return "aws_ecs_cluster_capacity_provider_strategy_items"
}

type ClusterSetting struct {
	ID        uint `gorm:"primarykey"`
	ClusterID uint
	Name      *string
	Value     *string
}

func (ClusterSetting) TableName() string {
	return "aws_ecs_cluster_settings"
}

type ClusterTag struct {
	ID        uint `gorm:"primarykey"`
	ClusterID uint
	Key       *string
	Value     *string
}

func (ClusterTag) TableName() string {
	return "aws_ecs_cluster_tags"
}

func (c *Client) transformClusterKeyValuePair(value *ecs.KeyValuePair) *ClusterKeyValuePair {
	return &ClusterKeyValuePair{
		Name:  value.Name,
		Value: value.Value,
	}
}

func (c *Client) transformClusterKeyValuePairs(values []*ecs.KeyValuePair) []*ClusterKeyValuePair {
	var tValues []*ClusterKeyValuePair
	for _, v := range values {
		tValues = append(tValues, c.transformClusterKeyValuePair(v))
	}
	return tValues
}

func (c *Client) transformClusterCapacityProviderStrategyItem(value *ecs.CapacityProviderStrategyItem) *ClusterCapacityProviderStrategyItem {
	return &ClusterCapacityProviderStrategyItem{
		Base:             value.Base,
		CapacityProvider: value.CapacityProvider,
		Weight:           value.Weight,
	}
}

func (c *Client) transformClusterCapacityProviderStrategyItems(values []*ecs.CapacityProviderStrategyItem) []*ClusterCapacityProviderStrategyItem {
	var tValues []*ClusterCapacityProviderStrategyItem
	for _, v := range values {
		tValues = append(tValues, c.transformClusterCapacityProviderStrategyItem(v))
	}
	return tValues
}

func (c *Client) transformClusterSetting(value *ecs.ClusterSetting) *ClusterSetting {
	return &ClusterSetting{
		Name:  value.Name,
		Value: value.Value,
	}
}

func (c *Client) transformClusterSettings(values []*ecs.ClusterSetting) []*ClusterSetting {
	var tValues []*ClusterSetting
	for _, v := range values {
		tValues = append(tValues, c.transformClusterSetting(v))
	}
	return tValues
}

func (c *Client) transformClusterTag(value *ecs.Tag) *ClusterTag {
	return &ClusterTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformClusterTags(values []*ecs.Tag) []*ClusterTag {
	var tValues []*ClusterTag
	for _, v := range values {
		tValues = append(tValues, c.transformClusterTag(v))
	}
	return tValues
}

func (c *Client) transformCluster(value *ecs.Cluster) *Cluster {
	return &Cluster{
		Region:                            c.region,
		AccountID:                         c.accountID,
		ActiveServicesCount:               value.ActiveServicesCount,
		AttachmentsStatus:                 value.AttachmentsStatus,
		CapacityProviders:                 common.StringListToString(value.CapacityProviders),
		ClusterArn:                        value.ClusterArn,
		ClusterName:                       value.ClusterName,
		DefaultCapacityProviderStrategy:   c.transformClusterCapacityProviderStrategyItems(value.DefaultCapacityProviderStrategy),
		PendingTasksCount:                 value.PendingTasksCount,
		RegisteredContainerInstancesCount: value.RegisteredContainerInstancesCount,
		RunningTasksCount:                 value.RunningTasksCount,
		Settings:                          c.transformClusterSettings(value.Settings),
		Statistics:                        c.transformClusterKeyValuePairs(value.Statistics),
		Status:                            value.Status,
		Tags:                              c.transformClusterTags(value.Tags),
	}
}

func (c *Client) transformClusters(values []*ecs.Cluster) []*Cluster {
	var tValues []*Cluster
	for _, v := range values {
		tValues = append(tValues, c.transformCluster(v))
	}
	return tValues
}

func (c *Client) clusters(gConfig interface{}) error {
	var config ecs.DescribeClustersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ecsCluster"] {
		err := c.db.AutoMigrate(
			&Cluster{},
			&ClusterKeyValuePair{},
			&ClusterCapacityProviderStrategyItem{},
			&ClusterSetting{},
			&ClusterTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ecsCluster"] = true
	}

	output, err := c.svc.DescribeClusters(&config)
	if err != nil {
		return err
	}
	c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Cluster{})
	common.ChunkedCreate(c.db, c.transformClusters(output.Clusters))
	c.log.Info("Fetched resources", zap.String("resource", "ecs.cluster"), zap.Int("count", len(output.Clusters)))
	return nil
}
