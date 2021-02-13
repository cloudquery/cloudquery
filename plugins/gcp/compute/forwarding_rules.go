package compute

import (
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type ForwardingRule struct {
	ID                   uint `gorm:"primarykey"`
	ProjectID            string
	IPAddress            string
	IPProtocol           string
	AllPorts             bool
	AllowGlobalAccess    bool
	BackendService       string
	CreationTimestamp    string
	Description          string
	Fingerprint          string
	Id                   uint64
	IpVersion            string
	IsMirroringCollector bool
	Kind                 string
	LoadBalancingScheme  string
	MetadataFilters      []*ForwardingRuleMetadataFilter `gorm:"constraint:OnDelete:CASCADE;"`
	Name                 string
	Network              string
	NetworkTier          string
	PortRange            string
	Ports                []*ForwardingRulePort `gorm:"constraint:OnDelete:CASCADE;"`
	Region               string
	SelfLink             string
	ServiceLabel         string
	ServiceName          string
	Subnetwork           string
	Target               string
}

func (ForwardingRule) TableName() string {
	return "gcp_compute_forwarding_rules"
}

type ForwardingRuleMetadataFilterLabelMatch struct {
	ID                             uint   `gorm:"primarykey"`
	ProjectID                      string `gorm:"-"`
	ForwardingRuleMetadataFilterID uint   `neo:"ignore"`
	Name                           string
	Value                          string
}

func (ForwardingRuleMetadataFilterLabelMatch) TableName() string {
	return "gcp_compute_forwarding_rule_metadata_filter_label_matches"
}

type ForwardingRuleMetadataFilter struct {
	ID               uint                                      `gorm:"primarykey"`
	ProjectID        string                                    `gorm:"-"`
	ForwardingRuleID uint                                      `neo:"ignore"`
	Labels           []*ForwardingRuleMetadataFilterLabelMatch `gorm:"constraint:OnDelete:CASCADE;"`
	MatchCriteria    string
}

func (ForwardingRuleMetadataFilter) TableName() string {
	return "gcp_compute_forwarding_rule_metadata_filters"
}

type ForwardingRulePort struct {
	ID               uint `gorm:"primarykey"`
	ForwardingRuleID uint
	Value            string
}

func (ForwardingRulePort) TableName() string {
	return "gcp_compute_forwarding_rule_ports"
}

func (c *Client) transformForwardingRules(values []*compute.ForwardingRule) []*ForwardingRule {
	var tValues []*ForwardingRule
	for _, value := range values {
		tValue := ForwardingRule{
			ProjectID:            c.projectID,
			IPAddress:            value.IPAddress,
			IPProtocol:           value.IPProtocol,
			AllPorts:             value.AllPorts,
			AllowGlobalAccess:    value.AllowGlobalAccess,
			BackendService:       value.BackendService,
			CreationTimestamp:    value.CreationTimestamp,
			Description:          value.Description,
			Fingerprint:          value.Fingerprint,
			Id:                   value.Id,
			IpVersion:            value.IpVersion,
			IsMirroringCollector: value.IsMirroringCollector,
			Kind:                 value.Kind,
			LoadBalancingScheme:  value.LoadBalancingScheme,
			MetadataFilters:      c.transformForwardingRuleMetadataFilters(value.MetadataFilters),
			Name:                 value.Name,
			Network:              value.Network,
			NetworkTier:          value.NetworkTier,
			PortRange:            value.PortRange,
			Ports:                c.transformForwardingRulePorts(value.Ports),
			Region:               value.Region,
			SelfLink:             value.SelfLink,
			ServiceLabel:         value.ServiceLabel,
			ServiceName:          value.ServiceName,
			Subnetwork:           value.Subnetwork,
			Target:               value.Target,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformForwardingRuleMetadataFilterLabelMatchs(values []*compute.MetadataFilterLabelMatch) []*ForwardingRuleMetadataFilterLabelMatch {
	var tValues []*ForwardingRuleMetadataFilterLabelMatch
	for _, value := range values {
		tValue := ForwardingRuleMetadataFilterLabelMatch{
			ProjectID: c.projectID,
			Name:      value.Name,
			Value:     value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformForwardingRuleMetadataFilters(values []*compute.MetadataFilter) []*ForwardingRuleMetadataFilter {
	var tValues []*ForwardingRuleMetadataFilter
	for _, value := range values {
		tValue := ForwardingRuleMetadataFilter{
			ProjectID:     c.projectID,
			Labels:        c.transformForwardingRuleMetadataFilterLabelMatchs(value.FilterLabels),
			MatchCriteria: value.FilterMatchCriteria,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformForwardingRulePorts(values []string) []*ForwardingRulePort {
	var tValues []*ForwardingRulePort
	for _, v := range values {
		tValues = append(tValues, &ForwardingRulePort{
			Value: v,
		})
	}
	return tValues
}

var ForwardingRuleTables = []interface{}{
	&ForwardingRule{},
	&ForwardingRuleMetadataFilter{},
	&ForwardingRuleMetadataFilterLabelMatch{},
	&ForwardingRulePort{},
}

func (c *Client) forwardingRules(_ interface{}) error {

	nextPageToken := ""
	c.db.Where("project_id", c.projectID).Delete(ForwardingRuleTables...)
	for {
		call := c.svc.ForwardingRules.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		for _, item := range output.Items {
			c.db.ChunkedCreate(c.transformForwardingRules(item.ForwardingRules))
			c.log.Info("populating ForwardingRules", zap.Int("count", len(item.ForwardingRules)))
		}
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
