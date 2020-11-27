package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type ByoipCidr struct {
	ID            uint `gorm:"primarykey"`
	AccountID     string
	Region        string
	Cidr          *string
	Description   *string
	State         *string
	StatusMessage *string
}

func (c *Client) transformByoipCidr(value *ec2.ByoipCidr) *ByoipCidr {
	return &ByoipCidr{
		Region:        c.region,
		AccountID:     c.accountID,
		Cidr:          value.Cidr,
		Description:   value.Description,
		State:         value.State,
		StatusMessage: value.StatusMessage,
	}
}

func (c *Client) transformByoipCidrs(values []*ec2.ByoipCidr) []*ByoipCidr {
	var tValues []*ByoipCidr
	for _, v := range values {
		tValues = append(tValues, c.transformByoipCidr(v))
	}
	return tValues
}

func (c *Client) ByoipCidrs(gConfig interface{}) error {
	var config ec2.DescribeByoipCidrsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ec2ByoipCidr"] {
		err := c.db.AutoMigrate(
			&ByoipCidr{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ec2ByoipCidr"] = true
	}
	for {
		output, err := c.svc.DescribeByoipCidrs(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous ByoipCidrs", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&ByoipCidr{})
		common.ChunkedCreate(c.db, c.transformByoipCidrs(output.ByoipCidrs))
		c.log.Info("populating ByoipCidrs", zap.Int("count", len(output.ByoipCidrs)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
