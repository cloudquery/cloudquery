package emr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Cluster struct {
	_                       interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                      uint        `gorm:"primarykey"`
	AccountID               string
	Region                  string
	ClusterArn              *string `neo:"unique"`
	ClusterID               *string
	Name                    *string
	NormalizedInstanceHours *int64
	OutpostArn              *string

	StatusState               *string
	StatusChangeReasonCode    *string
	StatusChangeReasonMessage *string

	StatusCreationDateTime *time.Time
	StatusEndDateTime      *time.Time
	StatusReadyDateTime    *time.Time
}

func (Cluster) TableName() string {
	return "aws_emr_clusters"
}

func (c *Client) transformCluster(value *emr.ClusterSummary) *Cluster {
	res := Cluster{
		Region:                  c.region,
		AccountID:               c.accountID,
		ClusterArn:              value.ClusterArn,
		ClusterID:               value.Id,
		Name:                    value.Name,
		NormalizedInstanceHours: value.NormalizedInstanceHours,
		OutpostArn:              value.OutpostArn,
	}

	if value.Status != nil {
		res.StatusState = value.Status.State
		if value.Status.StateChangeReason != nil {
			res.StatusChangeReasonCode = value.Status.StateChangeReason.Code
			res.StatusChangeReasonMessage = value.Status.StateChangeReason.Message
		}
		if value.Status.Timeline != nil {
			res.StatusCreationDateTime = value.Status.Timeline.CreationDateTime
			res.StatusEndDateTime = value.Status.Timeline.EndDateTime
			res.StatusReadyDateTime = value.Status.Timeline.ReadyDateTime
		}
	}

	return &res
}

func (c *Client) transformClusters(values []*emr.ClusterSummary) []*Cluster {
	var tValues []*Cluster
	for _, v := range values {
		tValues = append(tValues, c.transformCluster(v))
	}
	return tValues
}

var ClusterTables = []interface{}{
	&Cluster{},
}

func (c *Client) clusters(gConfig interface{}) error {
	var config emr.ListClustersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ClusterTables...)
	for {
		output, err := c.svc.ListClusters(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformClusters(output.Clusters))
		c.log.Info("Fetched resources", zap.String("resource", "emr.clusters"), zap.Int("count", len(output.Clusters)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
