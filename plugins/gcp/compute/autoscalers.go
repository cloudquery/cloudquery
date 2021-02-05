package compute

import (
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type Autoscaler struct {
	_         interface{} `neo:"raw:MERGE (a:GCPProject {project_id: $project_id}) MERGE (a) - [:Resource] -> (n)"`
	ID        uint        `gorm:"primarykey"`
	ProjectID string      `neo:"unique"`

	CoolDownPeriodSec                         int64
	CpuUtilizationUtilizationTarget           float64
	CustomMetris                              []*AutoscalerPolicyCustomMetric `gorm:"constraint:OnDelete:CASCADE;"`
	LoadBalancingUtilizationUtilizationTarget float64
	MaxNumReplicas                            int64
	MinNumReplicas                            int64
	Mode                                      string
	MaxScaledInReplicasCalculated             int64
	MaxScaledInReplicasFixed                  int64
	MaxScaledInReplicasPercent                int64
	TimeWindowSec                             int64

	CreationTimestamp string
	Description       string
	ResourceID        uint64 `neo:"unique"`
	Kind              string
	Name              string
	RecommendedSize   int64
	Region            string
	SelfLink          string
	Status            string
	StatusDetails     []*AutoscalerStatusDetails `gorm:"constraint:OnDelete:CASCADE;"`
	Target            string
	Zone              string
}

func (Autoscaler) TableName() string {
	return "gcp_compute_autoscalers"
}

type AutoscalerPolicyCustomMetric struct {
	ID           uint   `gorm:"primarykey"`
	AutoscalerID uint   `neo:"ignore"`
	ProjectID    string `gorm:"-"`

	Filter                   string
	Metric                   string
	SingleInstanceAssignment float64
	UtilizationTarget        float64
	UtilizationTargetType    string
}

func (AutoscalerPolicyCustomMetric) TableName() string {
	return "gcp_compute_autoscaler_policy_custom_metrics"
}

type AutoscalerStatusDetails struct {
	ID           uint   `gorm:"primarykey"`
	AutoscalerID uint   `neo:"ignore"`
	ProjectID    string `gorm:"-"`

	Message string
	Type    string
}

func (AutoscalerStatusDetails) TableName() string {
	return "gcp_compute_autoscaler_status_details"
}

func (c *Client) transformAutoscalerAutoscalingPolicyCustomMetricUtilization(value *compute.AutoscalingPolicyCustomMetricUtilization) *AutoscalerPolicyCustomMetric {
	return &AutoscalerPolicyCustomMetric{
		ProjectID:                c.projectID,
		Filter:                   value.Filter,
		Metric:                   value.Metric,
		SingleInstanceAssignment: value.SingleInstanceAssignment,
		UtilizationTarget:        value.UtilizationTarget,
		UtilizationTargetType:    value.UtilizationTargetType,
	}
}

func (c *Client) transformAutoscalerAutoscalingPolicyCustomMetricUtilizations(values []*compute.AutoscalingPolicyCustomMetricUtilization) []*AutoscalerPolicyCustomMetric {
	var tValues []*AutoscalerPolicyCustomMetric
	for _, v := range values {
		tValues = append(tValues, c.transformAutoscalerAutoscalingPolicyCustomMetricUtilization(v))
	}
	return tValues
}

func (c *Client) transformAutoscalerStatusDetails(value *compute.AutoscalerStatusDetails) *AutoscalerStatusDetails {
	return &AutoscalerStatusDetails{
		ProjectID: c.projectID,
		Message:   value.Message,
		Type:      value.Type,
	}
}

func (c *Client) transformAutoscalerStatusDetailss(values []*compute.AutoscalerStatusDetails) []*AutoscalerStatusDetails {
	var tValues []*AutoscalerStatusDetails
	for _, v := range values {
		tValues = append(tValues, c.transformAutoscalerStatusDetails(v))
	}
	return tValues
}

func (c *Client) transformAutoscaler(value *compute.Autoscaler) *Autoscaler {
	res := Autoscaler{
		ProjectID: c.projectID,

		CreationTimestamp: value.CreationTimestamp,
		Description:       value.Description,
		ResourceID:        value.Id,
		Kind:              value.Kind,
		Name:              value.Name,
		RecommendedSize:   value.RecommendedSize,
		Region:            value.Region,
		SelfLink:          value.SelfLink,
		Status:            value.Status,
		StatusDetails:     c.transformAutoscalerStatusDetailss(value.StatusDetails),
		Target:            value.Target,
		Zone:              value.Zone,
	}

	if value.AutoscalingPolicy != nil {
		res.CoolDownPeriodSec = value.AutoscalingPolicy.CoolDownPeriodSec

		if value.AutoscalingPolicy.CpuUtilization != nil {
			res.CpuUtilizationUtilizationTarget = value.AutoscalingPolicy.CpuUtilization.UtilizationTarget
		}

		res.CustomMetris = c.transformAutoscalerAutoscalingPolicyCustomMetricUtilizations(value.AutoscalingPolicy.CustomMetricUtilizations)
		if value.AutoscalingPolicy.LoadBalancingUtilization != nil {
			res.LoadBalancingUtilizationUtilizationTarget = value.AutoscalingPolicy.LoadBalancingUtilization.UtilizationTarget
		}

		res.MaxNumReplicas = value.AutoscalingPolicy.MaxNumReplicas
		res.MinNumReplicas = value.AutoscalingPolicy.MinNumReplicas
		res.Mode = value.AutoscalingPolicy.Mode

		if value.AutoscalingPolicy.ScaleInControl != nil {
			if value.AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas != nil {
				res.MaxScaledInReplicasCalculated = value.AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Calculated
				res.MaxScaledInReplicasFixed = value.AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Fixed
				res.MaxScaledInReplicasPercent = value.AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Percent
			}
			res.TimeWindowSec = value.AutoscalingPolicy.ScaleInControl.TimeWindowSec
		}
	}

	return &res
}

func (c *Client) transformAutoscalers(values []*compute.Autoscaler) []*Autoscaler {
	var tValues []*Autoscaler
	for _, v := range values {
		tValues = append(tValues, c.transformAutoscaler(v))
	}
	return tValues
}

type AutoscalerConfig struct {
	Filter string
}

var AutoscalerTables = []interface{}{
	&Autoscaler{},
	&AutoscalerPolicyCustomMetric{},
	&AutoscalerStatusDetails{},
}

func (c *Client) autoscalers(gConfig interface{}) error {
	var config AutoscalerConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("project_id", c.projectID).Delete(AutoscalerTables...)
	nextPageToken := ""
	for {
		call := c.svc.Autoscalers.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var tValues []*Autoscaler
		for _, items := range output.Items {
			tValues = append(tValues, c.transformAutoscalers(items.Autoscalers)...)
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", zap.String("resource", "compute.autoscalers"), zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
