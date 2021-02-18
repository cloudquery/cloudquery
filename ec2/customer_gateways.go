package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type CustomerGateway struct {
	_                 interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                uint        `gorm:"primarykey"`
	AccountID         string      `neo:"unique"`
	Region            string      `neo:"unique"`
	BgpAsn            *string
	CertificateArn    *string
	CustomerGatewayId *string `neo:"unique"`
	DeviceName        *string
	IpAddress         *string
	State             *string
	Tags              []*CustomerGatewayTag `gorm:"constraint:OnDelete:CASCADE;"`
	Type              *string
}

func (CustomerGateway) TableName() string {
	return "aws_ec2_customer_gateways"
}

type CustomerGatewayTag struct {
	ID                uint   `gorm:"primarykey"`
	CustomerGatewayID uint   `neo:"ignore"`
	AccountID         string `gorm:"-"`
	Region            string `gorm:"-"`

	Key   *string
	Value *string
}

func (CustomerGatewayTag) TableName() string {
	return "aws_ec2_customer_gateway_tags"
}

func (c *Client) transformCustomerGatewayTags(values *[]types.Tag) []*CustomerGatewayTag {
	var tValues []*CustomerGatewayTag
	for _, v := range *values {
		tValues = append(tValues, &CustomerGatewayTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       v.Key,
			Value:     v.Value,
		})
	}
	return tValues
}

func (c *Client) transformCustomerGateways(values *[]types.CustomerGateway) []*CustomerGateway {
	var tValues []*CustomerGateway
	for _, v := range *values {
		tValues = append(tValues, &CustomerGateway{
			Region:            c.region,
			AccountID:         c.accountID,
			BgpAsn:            v.BgpAsn,
			CertificateArn:    v.CertificateArn,
			CustomerGatewayId: v.CustomerGatewayId,
			DeviceName:        v.DeviceName,
			IpAddress:         v.IpAddress,
			State:             v.State,
			Tags:              c.transformCustomerGatewayTags(&v.Tags),
			Type:              v.Type,
		})
	}
	return tValues
}

var CustomerGatewayTables = []interface{}{
	&CustomerGateway{},
	&CustomerGatewayTag{},
}

func (c *Client) customerGateways(gConfig interface{}) error {
	ctx := context.Background()
	var config ec2.DescribeCustomerGatewaysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	output, err := c.svc.DescribeCustomerGateways(ctx, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(CustomerGatewayTables...)
	c.db.ChunkedCreate(c.transformCustomerGateways(&output.CustomerGateways))
	c.log.Info("Fetched resources", zap.String("resource", "ec2.customer_gateways"), zap.Int("count", len(output.CustomerGateways)))
	return nil
}
