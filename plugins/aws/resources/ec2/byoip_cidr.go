package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ByoipCidr struct {
	_             interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID            uint        `gorm:"primarykey"`
	AccountID     string
	Region        string
	Cidr          *string
	Description   *string
	State         *string
	StatusMessage *string
}

func (ByoipCidr) TableName() string {
	return "aws_ec2_byoip_cidrs"
}

func (c *Client) transformByoipCidrs(values *[]types.ByoipCidr) []*ByoipCidr {
	var tValues []*ByoipCidr
	for _, v := range *values {
		tValues = append(tValues, &ByoipCidr{
			Region:        c.region,
			AccountID:     c.accountID,
			Cidr:          v.Cidr,
			Description:   v.Description,
			State:         aws.String(string(v.State)),
			StatusMessage: v.StatusMessage,
		})
	}
	return tValues
}

var ByoipCidrTables = []interface{}{
	&ByoipCidr{},
}

func (c *Client) byoipCidrs(_ interface{}) error {
	ctx := context.Background()
	config := ec2.DescribeByoipCidrsInput{
		MaxResults: 100,
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ByoipCidrTables...)
	for {
		output, err := c.svc.DescribeByoipCidrs(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformByoipCidrs(&output.ByoipCidrs))
		c.log.Info("Fetched resources", "resource", "ec2.byoip_cidrs", "count", len(output.ByoipCidrs))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
