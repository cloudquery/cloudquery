package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type InternetGateway struct {
	_                 interface{}                  `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                uint                         `gorm:"primarykey"`
	AccountID         string                       `neo:"unique"`
	Region            string                       `neo:"unique"`
	Attachments       []*InternetGatewayAttachment `gorm:"constraint:OnDelete:CASCADE;"`
	InternetGatewayId *string                      `neo:"unique"`
	OwnerId           *string
	Tags              []*InternetGatewayTag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (InternetGateway) TableName() string {
	return "aws_ec2_internet_gateways"
}

type InternetGatewayAttachment struct {
	ID                uint `gorm:"primarykey"`
	InternetGatewayID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	State *string
	VpcId *string
}

func (InternetGatewayAttachment) TableName() string {
	return "aws_ec2_internet_gateway_attachments"
}

type InternetGatewayTag struct {
	ID                uint `gorm:"primarykey"`
	InternetGatewayID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (InternetGatewayTag) TableName() string {
	return "aws_ec2_internet_gateway_tags"
}

func (c *Client) transformInternetGatewayAttachments(values *[]types.InternetGatewayAttachment) []*InternetGatewayAttachment {
	var tValues []*InternetGatewayAttachment
	for _, value := range *values {
		tValues = append(tValues, &InternetGatewayAttachment{
			AccountID: c.accountID,
			Region:    c.region,
			State:     aws.String(string(value.State)),
			VpcId:     value.VpcId,
		})
	}
	return tValues
}

func (c *Client) transformInternetGatewayTags(values *[]types.Tag) []*InternetGatewayTag {
	var tValues []*InternetGatewayTag
	for _, value := range *values {
		tValues = append(tValues, &InternetGatewayTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		})
	}
	return tValues
}

func (c *Client) transformInternetGateways(values *[]types.InternetGateway) []*InternetGateway {
	var tValues []*InternetGateway
	for _, value := range *values {
		tValues = append(tValues, &InternetGateway{
			Region:            c.region,
			AccountID:         c.accountID,
			Attachments:       c.transformInternetGatewayAttachments(&value.Attachments),
			InternetGatewayId: value.InternetGatewayId,
			OwnerId:           value.OwnerId,
			Tags:              c.transformInternetGatewayTags(&value.Tags),
		})
	}
	return tValues
}

var InternetGatewayTables = []interface{}{
	&InternetGateway{},
	&InternetGatewayAttachment{},
	&InternetGatewayTag{},
}

func (c *Client) internetGateways(gConfig interface{}) error {
	ctx := context.Background()
	var config ec2.DescribeInternetGatewaysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(InternetGatewayTables...)
	for {
		output, err := c.svc.DescribeInternetGateways(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformInternetGateways(&output.InternetGateways))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.internet_gateways"), zap.Int("count", len(output.InternetGateways)))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
