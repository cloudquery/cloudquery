package emr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Cluster struct {
	ID                      uint `gorm:"primarykey"`
	AccountID               string
	Region                  string
	ClusterArn              *string
	Id                      *string
	Name                    *string
	NormalizedInstanceHours *int64
	OutpostArn              *string
	Status                  *ClusterStatus `gorm:"constraint:OnDelete:CASCADE;"`
}

type ClusterStatus struct {
	ID                uint `gorm:"primarykey"`
	ClusterID         uint
	State             *string
	StateChangeReason *emr.ClusterStateChangeReason `gorm:"embedded;embeddedPrefix:state_change_reason_"`
	Timeline          *emr.ClusterTimeline          `gorm:"embedded;embeddedPrefix:timeline_"`
}

func (c *Client) transformClusterStatus(value *emr.ClusterStatus) *ClusterStatus {
	return &ClusterStatus{
		State:             value.State,
		StateChangeReason: value.StateChangeReason,
		Timeline:          value.Timeline,
	}
}

func (c *Client) transformCluster(value *emr.ClusterSummary) *Cluster {
	return &Cluster{
		Region:                  c.region,
		AccountID:               c.accountID,
		ClusterArn:              value.ClusterArn,
		Id:                      value.Id,
		Name:                    value.Name,
		NormalizedInstanceHours: value.NormalizedInstanceHours,
		OutpostArn:              value.OutpostArn,
		Status:                  c.transformClusterStatus(value.Status),
	}
}

func (c *Client) transformClusters(values []*emr.ClusterSummary) []*Cluster {
	var tValues []*Cluster
	for _, v := range values {
		tValues = append(tValues, c.transformCluster(v))
	}
	return tValues
}

func (c *Client) clusters(gConfig interface{}) error {
	var config emr.ListClustersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["EMRClusterSummary"] {
		err := c.db.AutoMigrate(
			&Cluster{},
			&ClusterStatus{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["EMRClusterSummary"] = true
	}
	for {
		output, err := c.svc.ListClusters(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Cluster{})
		common.ChunkedCreate(c.db, c.transformClusters(output.Clusters))
		c.log.Info("Fetched resources", zap.Int("count", len(output.Clusters)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
