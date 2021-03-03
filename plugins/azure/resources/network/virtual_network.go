package network

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
)

type VirtualNetwork struct {
	ID                         uint    `gorm:"primarykey"`
	SubscriptionID             string  `neo:"unique"`
	ResourceID                 *string `neo:"unique"`
	Name                       *string
	Type                       *string
	Location                   *string
	Etag                       *string
	Tags                       []*VirtualNetworkTag           `gorm:"constraint:OnDelete:CASCADE;"`
	AddressPrefixes            []*VirtualNetworkAddressPrefix `gorm:"constraint:OnDelete:CASCADE;"`
	BGPRegionalCommunity       *string
	BGPVirtualNetworkCommunity *string
	DdosProtectionPlanID       *string
	DNSServers                 []*VirtualNetworkDNSServer `gorm:"constraint:OnDelete:CASCADE;"`
	DdosProtectionEnabled      *bool
	VMProtectionEnabled        *bool
	ExtendedLocationName       *string
	ExtendedLocationType       *string
	IPAllocations              []*VirtualNetworkIPAllocation `gorm:"constraint:OnDelete:CASCADE;"`
	ProvisioningState          string
	ResourceGUID               *string
	Subnets                    []*Subnet                `gorm:"constraint:OnDelete:CASCADE;"`
	Peerings                   []*VirtualNetworkPeering `gorm:"constraint:OnDelete:CASCADE;"`
}

type VirtualNetworkDNSServer struct {
	ID               uint   `gorm:"primarykey"`
	VirtualNetworkID uint   `neo:"ignore"`
	SubscriptionID   string `gorm:"-"`
	IP               string
}

type VirtualNetworkAddressPrefix struct {
	ID               uint   `gorm:"primarykey"`
	VirtualNetworkID uint   `neo:"ignore"`
	SubscriptionID   string `gorm:"-"`

	AddressPrefix string
}

type VirtualNetworkIPAllocation struct {
	ID               uint   `gorm:"primarykey"`
	VirtualNetworkID uint   `neo:"ignore"`
	SubscriptionID   string `gorm:"-"`

	ResourceID *string `neo:"unique"`
}

type VirtualNetworkTag struct {
	ID               uint   `gorm:"primarykey"`
	VirtualNetworkID uint   `neo:"ignore"`
	SubscriptionID   string `gorm:"-"`

	Key   string
	Value *string
}

type VirtualNetworkPeering struct {
	ID               uint   `gorm:"primarykey"`
	VirtualNetworkID uint   `neo:"ignore"`
	SubscriptionID   string `gorm:"-"`

	RemoteVirtualNetworkID *string
}

type Subnet struct {
	ID               uint    `gorm:"primarykey"`
	SubscriptionID   string  `neo:"unique"`
	VirtualNetworkID uint    `neo:"ignore"`
	ResourceID       *string `neo:"unique"`
	Name             *string
	Etag             *string
	Purpose          *string

	AddressPrefix                     *string
	AddressPrefixes                   []*SubnetAddressPrefix `gorm:"constraint:OnDelete:CASCADE;"`
	ProvisioningState                 string
	PrivateEndpointNetworkPolicies    *string
	PrivateLinkServiceNetworkPolicies *string
	NatGatewayID                      *string
}

type SubnetAddressPrefix struct {
	ID             uint   `gorm:"primarykey"`
	SubnetID       uint   `neo:"ignore"`
	SubscriptionID string `gorm:"-"`

	AddressPrefix string
}

func (VirtualNetwork) TableName() string {
	return "azure_network_virtual_networks"
}

func (VirtualNetworkTag) TableName() string {
	return "azure_network_virtual_network_tags"
}

func (VirtualNetworkAddressPrefix) TableName() string {
	return "azure_network_virtual_network_address_prefixes"
}

func (VirtualNetworkDNSServer) TableName() string {
	return "azure_network_virtual_network_dns_servers"
}

func (VirtualNetworkIPAllocation) TableName() string {
	return "azure_network_virtual_network_ip_allocations"
}

func (VirtualNetworkPeering) TableName() string {
	return "azure_network_virtual_network_peerings"
}

func (Subnet) TableName() string {
	return "azure_network_subnets"
}

func (SubnetAddressPrefix) TableName() string {
	return "azure_network_subnet_address_prefixes"
}

func transformVirtualNetworks(subscriptionID string, auth autorest.Authorizer, values *[]network.VirtualNetwork) ([]*VirtualNetwork, error) {
	var tValues []*VirtualNetwork

	resourceGroupRe := regexp.MustCompile("resourceGroups/([a-zA-Z0-9-_]+)/")

	for _, value := range *values {
		tValue := VirtualNetwork{
			SubscriptionID:        subscriptionID,
			Location:              value.Location,
			ResourceID:            value.ID,
			Name:                  value.Name,
			Type:                  value.Type,
			Etag:                  value.Etag,
			Tags:                  transformVirtualNetworkTags(subscriptionID, value.Tags),
			AddressPrefixes:       transformVirtualNetworkAddressPrefix(subscriptionID, value.AddressSpace),
			DdosProtectionEnabled: value.EnableDdosProtection,
			VMProtectionEnabled:   value.EnableVMProtection,
			ProvisioningState:     string(value.ProvisioningState),
			ResourceGUID:          value.ResourceGUID,
			Subnets:               transformVirtualNetworkSubnets(subscriptionID, value.Subnets),
		}
		if value.VirtualNetworkPeerings != nil {
			tValue.Peerings = transformVirtualNetworkPeerings(subscriptionID, value.VirtualNetworkPeerings)
		}
		if value.IPAllocations != nil {
			tValue.IPAllocations = transformVirtualNetworkIPAllocations(subscriptionID, value.IPAllocations)
		}
		if value.DdosProtectionPlan != nil {
			tValue.DdosProtectionPlanID = value.DdosProtectionPlan.ID
		}
		if value.BgpCommunities != nil {
			tValue.BGPRegionalCommunity = value.BgpCommunities.RegionalCommunity
			tValue.BGPVirtualNetworkCommunity = value.BgpCommunities.VirtualNetworkCommunity
		}
		if value.DhcpOptions != nil && value.DhcpOptions.DNSServers != nil {
			tValue.DNSServers = transformVirtualNetworkDNSServers(subscriptionID, value.DhcpOptions.DNSServers)
		}
		if value.ExtendedLocation != nil {
			tValue.ExtendedLocationName = value.ExtendedLocation.Name
			tValue.ExtendedLocationType = value.ExtendedLocation.Type
		}
		match := resourceGroupRe.FindStringSubmatch(*value.ID)
		if len(match) < 2 {
			return nil, fmt.Errorf("couldn't extract resource group from %s", *value.ID)
		}
		tValues = append(tValues, &tValue)
	}
	return tValues, nil
}
func transformVirtualNetworkTags(subscriptionID string, values map[string]*string) []*VirtualNetworkTag {
	var tValues []*VirtualNetworkTag
	for k, v := range values {
		tValue := VirtualNetworkTag{
			SubscriptionID: subscriptionID,
			Key:            k,
			Value:          v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVirtualNetworkAddressPrefix(subscriptionID string, value *network.AddressSpace) []*VirtualNetworkAddressPrefix {
	var tValues []*VirtualNetworkAddressPrefix
	for _, value := range *value.AddressPrefixes {
		tValue := VirtualNetworkAddressPrefix{
			SubscriptionID: subscriptionID,
			AddressPrefix:  value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVirtualNetworkDNSServers(subscriptionID string, values *[]string) []*VirtualNetworkDNSServer {
	var tValues []*VirtualNetworkDNSServer
	for _, value := range *values {
		tValue := VirtualNetworkDNSServer{
			SubscriptionID: subscriptionID,
			IP:             value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVirtualNetworkIPAllocations(subscriptionID string, values *[]network.SubResource) []*VirtualNetworkIPAllocation {
	var tValues []*VirtualNetworkIPAllocation
	for _, value := range *values {
		tValue := VirtualNetworkIPAllocation{
			SubscriptionID: subscriptionID,
			ResourceID:     value.ID,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVirtualNetworkPeerings(subscriptionID string, values *[]network.VirtualNetworkPeering) []*VirtualNetworkPeering {
	var tValues []*VirtualNetworkPeering
	for _, value := range *values {
		tValue := VirtualNetworkPeering{
			SubscriptionID:         subscriptionID,
			RemoteVirtualNetworkID: value.VirtualNetworkPeeringPropertiesFormat.RemoteVirtualNetwork.ID,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformVirtualNetworkSubnets(subscriptionID string, values *[]network.Subnet) []*Subnet {
	var tValues []*Subnet
	for _, value := range *values {
		tValue := Subnet{
			SubscriptionID:                    subscriptionID,
			ResourceID:                        value.ID,
			Name:                              value.Name,
			Etag:                              value.Etag,
			Purpose:                           value.Purpose,
			AddressPrefix:                     value.AddressPrefix,
			AddressPrefixes:                   transformSubnetAddressPrefixes(subscriptionID, value.AddressPrefixes),
			ProvisioningState:                 string(value.ProvisioningState),
			PrivateEndpointNetworkPolicies:    value.PrivateEndpointNetworkPolicies,
			PrivateLinkServiceNetworkPolicies: value.PrivateLinkServiceNetworkPolicies,
			NatGatewayID:                      value.NatGateway.ID,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func transformSubnetAddressPrefixes(subscriptionID string, values *[]string) []*SubnetAddressPrefix {
	var tValues []*SubnetAddressPrefix
	for _, value := range *values {
		tValue := SubnetAddressPrefix{
			SubscriptionID: subscriptionID,
			AddressPrefix:  value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type VirtualNetworkConfig struct {
	Filter string
}

var VirtualNetworkTables = []interface{}{
	&VirtualNetwork{},
	&VirtualNetworkDNSServer{},
	&VirtualNetworkAddressPrefix{},
	&VirtualNetworkIPAllocation{},
	&VirtualNetworkTag{},
	&VirtualNetworkPeering{},
	&Subnet{},
	&SubnetAddressPrefix{},
}

func VirtualNetworks(subscriptionID string, auth autorest.Authorizer, db *database.Database, log hclog.Logger, gConfig interface{}) error {
	var config VirtualNetworkConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	svc := network.NewVirtualNetworksClient(subscriptionID)
	svc.Authorizer = auth
	output, err := svc.ListAll(ctx)
	if err != nil {
		return err
	}

	db.Where("subscription_id", subscriptionID).Delete(VirtualNetworkTables...)
	resourceCount := 0
	for output.NotDone() {
		values := output.Values()
		tValues, err := transformVirtualNetworks(subscriptionID, auth, &values)
		if err != nil {
			return err
		}
		db.ChunkedCreate(tValues)
		resourceCount += len(tValues)
		output.Next()
	}

	log.Info("Fetched resources", "count", resourceCount)
	return nil
}
