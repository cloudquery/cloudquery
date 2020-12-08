package directconnect

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Gateway struct {
	ID                        uint `gorm:"primarykey"`
	AccountID                 string
	Region                    string
	AmazonSideAsn             *int64
	DirectConnectGatewayId    *string
	DirectConnectGatewayName  *string
	DirectConnectGatewayState *string
	OwnerAccount              *string
	StateChangeError          *string
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

func (c *Client) gateways(gConfig interface{}) error {
	var config directconnect.DescribeDirectConnectGatewaysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["directconnectGateway"] {
		err := c.db.AutoMigrate(
			&Gateway{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["directconnectGateway"] = true
	}
	for {
		output, err := c.svc.DescribeDirectConnectGateways(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous gateways", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Gateway{})
		common.ChunkedCreate(c.db, c.transformGateways(output.DirectConnectGateways))
		c.log.Info("populating gateways", zap.Int("count", len(output.DirectConnectGateways)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
