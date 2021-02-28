package compute

import (
	"google.golang.org/api/compute/v1"
)


type Network struct {
	ID uint `gorm:"primarykey"`
	ProjectID string
	IPv4Range string
	AutoCreateSubnetworks bool
	CreationTimestamp string
	Description string
	GatewayIPv4 string
	Id uint64
	Kind string
	Mtu int64
	Name string
	Peerings []*NetworkPeering `gorm:"constraint:OnDelete:CASCADE;"`

	RoutingConfigRoutingMode string

	SelfLink string
	Subnetworks []*NetworkSubnetwork `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Network) TableName() string {
	return "gcp_compute_networks"
}

type NetworkPeering struct {
	ID uint `gorm:"primarykey"`
	ProjectID string `gorm:"-"`
	NetworkID uint `neo:"ignore"`
	AutoCreateRoutes bool
	ExchangeSubnetRoutes bool
	ExportCustomRoutes bool
	ExportSubnetRoutesWithPublicIp bool
	ImportCustomRoutes bool
	ImportSubnetRoutesWithPublicIp bool
	Name string
	Network string
	PeerMtu int64
	State string
	StateDetails string
}

func (NetworkPeering) TableName() string {
	return "gcp_compute_network_peerings"
}

type NetworkSubnetwork struct {
	ID uint`gorm:"primarykey"`
	NetworkID uint
	Value string
}
func (NetworkSubnetwork) TableName() string {
	return "gcp_compute_network_subnetworks"
}

func (c *Client) transformNetworks(values []*compute.Network) []*Network {
	var tValues []*Network
	for _, value := range values {
		tValue := Network {
			ProjectID: c.projectID,
			IPv4Range: value.IPv4Range,
			AutoCreateSubnetworks: value.AutoCreateSubnetworks,
			CreationTimestamp: value.CreationTimestamp,
			Description: value.Description,
			GatewayIPv4: value.GatewayIPv4,
			Id: value.Id,
			Kind: value.Kind,
			Mtu: value.Mtu,
			Name: value.Name,
			Peerings: c.transformNetworkPeerings(value.Peerings),
			SelfLink: value.SelfLink,
			Subnetworks: c.transformNetworkSubnetworks(value.Subnetworks),
		}
		if value.RoutingConfig != nil {

			tValue.RoutingConfigRoutingMode = value.RoutingConfig.RoutingMode

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformNetworkPeerings(values []*compute.NetworkPeering) []*NetworkPeering {
	var tValues []*NetworkPeering
	for _, value := range values {
		tValue := NetworkPeering {
			ProjectID: c.projectID,
			AutoCreateRoutes: value.AutoCreateRoutes,
			ExchangeSubnetRoutes: value.ExchangeSubnetRoutes,
			ExportCustomRoutes: value.ExportCustomRoutes,
			ExportSubnetRoutesWithPublicIp: value.ExportSubnetRoutesWithPublicIp,
			ImportCustomRoutes: value.ImportCustomRoutes,
			ImportSubnetRoutesWithPublicIp: value.ImportSubnetRoutesWithPublicIp,
			Name: value.Name,
			Network: value.Network,
			PeerMtu: value.PeerMtu,
			State: value.State,
			StateDetails: value.StateDetails,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformNetworkSubnetworks(values []string) []*NetworkSubnetwork {
	var tValues []*NetworkSubnetwork
	for _, v := range values {
		tValues = append(tValues, &NetworkSubnetwork{
			Value: v,
		})
	}
	return tValues
}



var NetworkTables = []interface{} {
	&Network{},
	&NetworkPeering{},
	&NetworkSubnetwork{},
}

func (c *Client)networks(_ interface{}) error {

	nextPageToken := ""
	c.db.Where("project_id", c.projectID).Delete(NetworkTables...)
	for {
		call := c.svc.Networks.List(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformNetworks(output.Items))
		c.log.Info("populating Networks", "resource", "compute.networks", "count", len(output.Items))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

