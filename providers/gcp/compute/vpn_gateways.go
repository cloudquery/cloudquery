package compute

import (
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type VpnGateway struct {
	ID                uint `gorm:"primarykey"`
	ProjectID         string
	CreationTimestamp string
	Description       string
	ResourceID        uint64
	Kind              string
	LabelFingerprint  string
	Labels            []*VpnGatewayLabel `gorm:"constraint:OnDelete:CASCADE;"`
	Name              string
	Network           string
	Region            string
	SelfLink          string
	VpnInterfaces     []*VpnGatewayVpnGatewayInterface `gorm:"constraint:OnDelete:CASCADE;"`
}

type VpnGatewayLabel struct {
	ID           uint `gorm:"primarykey"`
	VpnGatewayID uint
	Key          string
	Value        string
}

type VpnGatewayVpnGatewayInterface struct {
	ID           uint `gorm:"primarykey"`
	VpnGatewayID uint
	ResourceID   int64
	IpAddress    string
}

func (c *Client) transformVpnGatewayVpnGatewayLabels(values map[string]string) []*VpnGatewayLabel {
	var tValues []*VpnGatewayLabel
	for k, v := range values {
		tValues = append(tValues, &VpnGatewayLabel{
			Key:   k,
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformVpnGatewayVpnGatewayInterface(value *compute.VpnGatewayVpnGatewayInterface) *VpnGatewayVpnGatewayInterface {
	return &VpnGatewayVpnGatewayInterface{
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

func (c *Client) vpnGateways(gConfig interface{}) error {
	var config VpnGatewayConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["computeVpnGateway"] {
		err := c.db.AutoMigrate(
			&VpnGateway{},
			&VpnGatewayVpnGatewayInterface{},
			&VpnGatewayLabel{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["computeVpnGateway"] = true
	}
	nextPageToken := ""
	for {
		call := c.svc.VpnGateways.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id = ?", c.projectID).Delete(&VpnGateway{})
		var tValues []*VpnGateway
		for _, items := range output.Items {
			tValues = append(tValues, c.transformVpnGateways(items.VpnGateways)...)
		}
		common.ChunkedCreate(c.db, tValues)
		c.log.Info("Fetched resources", zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
