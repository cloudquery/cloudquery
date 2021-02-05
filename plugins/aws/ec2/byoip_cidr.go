package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.uber.org/zap"
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

var ByoipCidrTables = []interface{}{
	&ByoipCidr{},
}

func (c *Client) byoipCidrs(_ interface{}) error {
	MaxResults := int64(100)
	config := ec2.DescribeByoipCidrsInput{
		MaxResults: &MaxResults,
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ByoipCidrTables...)
	for {
		output, err := c.svc.DescribeByoipCidrs(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformByoipCidrs(output.ByoipCidrs))
		c.log.Info("Fetched resources", zap.String("resource", "ec2.byoip_cidrs"), zap.Int("count", len(output.ByoipCidrs)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
