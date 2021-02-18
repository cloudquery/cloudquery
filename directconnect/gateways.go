package directconnect

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
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

func (c *Client) transformGateways(values *[]types.DirectConnectGateway) []*Gateway {
	var tValues []*Gateway
	for _, v := range *values {
		tValues = append(tValues, &Gateway{
			Region:                    c.region,
			AccountID:                 c.accountID,
			AmazonSideAsn:             v.AmazonSideAsn,
			DirectConnectGatewayId:    v.DirectConnectGatewayId,
			DirectConnectGatewayName:  v.DirectConnectGatewayName,
			DirectConnectGatewayState: aws.String(string(v.DirectConnectGatewayState)),
			OwnerAccount:              v.OwnerAccount,
			StateChangeError:          v.StateChangeError,
		})
	}
	return tValues
}

var GatewayTables = []interface{}{
	&Gateway{},
}

func (c *Client) gateways(gConfig interface{}) error {
	ctx := context.Background()
	var config directconnect.DescribeDirectConnectGatewaysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(GatewayTables...)
	for {
		output, err := c.svc.DescribeDirectConnectGateways(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformGateways(&output.DirectConnectGateways))
		c.log.Info("Fetched resources", zap.String("resource", "directconnect.gateways"), zap.Int("count", len(output.DirectConnectGateways)))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
