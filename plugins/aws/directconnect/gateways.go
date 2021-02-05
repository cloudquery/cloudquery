package directconnect

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Gateway struct {
	_                         interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                        uint        `gorm:"primarykey"`
	AccountID                 string      `neo:"unique"`
	Region                    string      `neo:"unique"`
	AmazonSideAsn             *int64
	DirectConnectGatewayId    *string `neo:"unique"`
	DirectConnectGatewayName  *string
	DirectConnectGatewayState *string
	OwnerAccount              *string
	StateChangeError          *string
}

func (Gateway) TableName() string {
	return "aws_directconnect_gateways"
}

func (c *Client) transformGateway(value *directconnect.Gateway) *Gateway {
	return &Gateway{
		Region:                    c.region,
		AccountID:                 c.accountID,
		AmazonSideAsn:             value.AmazonSideAsn,
		DirectConnectGatewayId:    value.DirectConnectGatewayId,
		DirectConnectGatewayName:  value.DirectConnectGatewayName,
		DirectConnectGatewayState: value.DirectConnectGatewayState,
		OwnerAccount:              value.OwnerAccount,
		StateChangeError:          value.StateChangeError,
	}
}

func (c *Client) transformGateways(values []*directconnect.Gateway) []*Gateway {
	var tValues []*Gateway
	for _, v := range values {
		tValues = append(tValues, c.transformGateway(v))
	}
	return tValues
}

var GatewayTables = []interface{}{
	&Gateway{},
}

func (c *Client) gateways(gConfig interface{}) error {
	var config directconnect.DescribeDirectConnectGatewaysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(GatewayTables...)
	for {
		output, err := c.svc.DescribeDirectConnectGateways(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformGateways(output.DirectConnectGateways))
		c.log.Info("Fetched resources", zap.String("resource", "directconnect.gateways"), zap.Int("count", len(output.DirectConnectGateways)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
