package compute

import (
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type Address struct {
	ID                uint `gorm:"primarykey"`
	ProjectID         string
	Address           string
	AddressType       string
	CreationTimestamp string
	Description       string
	ResourceID        uint64
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

type AddressUser struct {
	ID        uint `gorm:"primarykey"`
	AddressID uint
	Value     string
}

func (c *Client) transformAddressUsers(values []string) []*AddressUser {
	var tValues []*AddressUser
	for _, v := range values {
		tValues = append(tValues, &AddressUser{
			Value: v,
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

func (c *Client) addresses(gConfig interface{}) error {
	var config AddressConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["computeAddress"] {
		err := c.db.AutoMigrate(
			&Address{},
			&AddressUser{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["computeAddress"] = true
	}
	nextPageToken := ""
	for {
		call := c.svc.Addresses.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id = ?", c.projectID).Delete(&Address{})
		var tValues []*Address
		for _, items := range output.Items {
			tValues = append(tValues, c.transformAddresses(items.Addresses)...)
		}
		common.ChunkedCreate(c.db, tValues)
		c.log.Info("populating Addresss", zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
