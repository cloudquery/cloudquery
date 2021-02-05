package compute

import (
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type VpnGateway struct {
	_                 interface{} `neo:"raw:MERGE (a:GCPProject {project_id: $project_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                uint        `gorm:"primarykey"`
	ProjectID         string      `neo:"unique"`
	CreationTimestamp string
	Description       string
	ResourceID        uint64 `neo:"unique"`
	Kind              string
	LabelFingerprint  string
	Labels            []*VpnGatewayLabel `gorm:"constraint:OnDelete:CASCADE;"`
	Name              string
	Network           string
	Region            string
	SelfLink          string
	VpnInterfaces     []*VpnGatewayVpnGatewayInterface `gorm:"constraint:OnDelete:CASCADE;"`
}

func (VpnGateway) TableName() string {
	return "gcp_compute_vpn_gateways"
}

type VpnGatewayLabel struct {
	ID           uint   `gorm:"primarykey"`
	VpnGatewayID uint   `neo:"ignore"`
	ProjectID    string `gorm:"-"`
	Key          string
	Value        string
}

func (VpnGatewayLabel) TableName() string {
	return "gcp_compute_vpn_gateway_label"
}

type VpnGatewayVpnGatewayInterface struct {
	ID           uint   `gorm:"primarykey"`
	VpnGatewayID uint   `neo:"ignore"`
	ProjectID    string `gorm:"-"`
	ResourceID   int64
	IpAddress    string
}

func (VpnGatewayVpnGatewayInterface) TableName() string {
	return "gcp_compute_vpn_gateway_interfaces"
}

func (c *Client) transformVpnGatewayVpnGatewayLabels(values map[string]string) []*VpnGatewayLabel {
	var tValues []*VpnGatewayLabel
	for k, v := range values {
		tValues = append(tValues, &VpnGatewayLabel{
			ProjectID: c.projectID,
			Key:       k,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformVpnGatewayVpnGatewayInterface(value *compute.VpnGatewayVpnGatewayInterface) *VpnGatewayVpnGatewayInterface {
	return &VpnGatewayVpnGatewayInterface{
		ProjectID:  c.projectID,
		ResourceID: value.Id,
		IpAddress:  value.IpAddress,
	}
}

func (c *Client) transformVpnGatewayVpnGatewayInterfaces(values []*compute.VpnGatewayVpnGatewayInterface) []*VpnGatewayVpnGatewayInterface {
	var tValues []*VpnGatewayVpnGatewayInterface
	for _, v := range values {
		tValues = append(tValues, c.transformVpnGatewayVpnGatewayInterface(v))
	}
	return tValues
}

func (c *Client) transformVpnGateway(value *compute.VpnGateway) *VpnGateway {
	return &VpnGateway{
		ProjectID:         c.projectID,
		CreationTimestamp: value.CreationTimestamp,
		Description:       value.Description,
		ResourceID:        value.Id,
		Kind:              value.Kind,
		LabelFingerprint:  value.LabelFingerprint,
		Labels:            c.transformVpnGatewayVpnGatewayLabels(value.Labels),
		Name:              value.Name,
		Network:           value.Network,
		Region:            value.Region,
		SelfLink:          value.SelfLink,
		VpnInterfaces:     c.transformVpnGatewayVpnGatewayInterfaces(value.VpnInterfaces),
	}
}

func (c *Client) transformVpnGateways(values []*compute.VpnGateway) []*VpnGateway {
	var tValues []*VpnGateway
	for _, v := range values {
		tValues = append(tValues, c.transformVpnGateway(v))
	}
	return tValues
}

type VpnGatewayConfig struct {
	Filter string
}

var VPNGatewayTables = []interface{}{
	&VpnGateway{},
	&VpnGatewayVpnGatewayInterface{},
	&VpnGatewayLabel{},
}

func (c *Client) vpnGateways(gConfig interface{}) error {
	var config VpnGatewayConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	c.db.Where("project_id", c.projectID).Delete(VPNGatewayTables...)
	nextPageToken := ""
	for {
		call := c.svc.VpnGateways.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var tValues []*VpnGateway
		for _, items := range output.Items {
			tValues = append(tValues, c.transformVpnGateways(items.VpnGateways)...)
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", zap.String("resource", "compute.vpn_gateways"), zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
