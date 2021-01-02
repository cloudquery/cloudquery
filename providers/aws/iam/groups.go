package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Group struct {
	ID         uint `gorm:"primarykey"`
	AccountID  string
	Arn        *string `neo:"unique"`
	CreateDate *time.Time
	GroupId    *string
	GroupName  *string
	Path       *string
}

func (Group) TableName() string {
	return "aws_iam_groups"
}

func (c *Client) transformGroup(value *iam.Group) *Group {
	return &Group{
		AccountID:  c.accountID,
		Arn:        value.Arn,
		CreateDate: value.CreateDate,
		GroupId:    value.GroupId,
		GroupName:  value.GroupName,
		Path:       value.Path,
	}
}

func (c *Client) transformGroups(values []*iam.Group) []*Group {
	var tValues []*Group
	for _, v := range values {
		tValues = append(tValues, c.transformGroup(v))
	}
	return tValues
}

var GroupTables = []interface{}{
	&Group{},
}

func (c *Client) groups(gConfig interface{}) error {
	var config iam.ListGroupsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("account_id", c.accountID).Delete(GroupTables...)

	for {
		output, err := c.svc.ListGroups(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformGroups(output.Groups))
		c.log.Info("Fetched resources", zap.String("resource", "iam.groups"), zap.Int("count", len(output.Groups)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
