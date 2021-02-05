package compute

import (
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type Address struct {
	_                 interface{} `neo:"raw:MERGE (a:GCPProject {project_id: $project_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                uint        `gorm:"primarykey"`
	ProjectID         string      `neo:"unique"`
	Address           string
	AddressType       string
	CreationTimestamp string
	Description       string
	ResourceID        uint64 `neo:"unique"`
	IpVersion         string
	Kind              string
	Name              string
	Network           string
	NetworkTier       string
	PrefixLength      int64
	Purpose           string
	Region            string
	SelfLink          string
	Status            string
	Subnetwork        string
	Users             []*AddressUser `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Address) TableName() string {
	return "gcp_compute_addresses"
}

type AddressUser struct {
	ID        uint   `gorm:"primarykey"`
	AddressID uint   `neo:"ignore"`
	ProjectID string `gorm:"-"`
	Value     string
}

func (AddressUser) TableName() string {
	return "gcp_compute_address_users"
}

func (c *Client) transformAddressUsers(values []string) []*AddressUser {
	var tValues []*AddressUser
	for _, v := range values {
		tValues = append(tValues, &AddressUser{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformAddress(value *compute.Address) *Address {
	return &Address{
		ProjectID:         c.projectID,
		Address:           value.Address,
		AddressType:       value.AddressType,
		CreationTimestamp: value.CreationTimestamp,
		Description:       value.Description,
		ResourceID:        value.Id,
		IpVersion:         value.IpVersion,
		Kind:              value.Kind,
		Name:              value.Name,
		Network:           value.Network,
		NetworkTier:       value.NetworkTier,
		PrefixLength:      value.PrefixLength,
		Purpose:           value.Purpose,
		Region:            value.Region,
		SelfLink:          value.SelfLink,
		Status:            value.Status,
		Subnetwork:        value.Subnetwork,
		Users:             c.transformAddressUsers(value.Users),
	}
}

func (c *Client) transformAddresses(values []*compute.Address) []*Address {
	var tValues []*Address
	for _, v := range values {
		tValues = append(tValues, c.transformAddress(v))
	}
	return tValues
}

type AddressConfig struct {
	Filter string
}

var AddressTables = []interface{}{
	&Address{},
	&AddressUser{},
}

func (c *Client) addresses(gConfig interface{}) error {
	var config AddressConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("project_id", c.projectID).Delete(AddressTables...)
	nextPageToken := ""
	for {
		call := c.svc.Addresses.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var tValues []*Address
		for _, items := range output.Items {
			tValues = append(tValues, c.transformAddresses(items.Addresses)...)
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", zap.String("resource", "compute.addresses"), zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
