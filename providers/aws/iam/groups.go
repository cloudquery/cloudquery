package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Group struct {
	ID         uint `gorm:"primarykey"`
	AccountID  string
	Region     string
	Arn        *string
	CreateDate *time.Time
	GroupId    *string
	GroupName  *string
	Path       *string
}

func (c *Client) transformGroup(value *iam.Group) *Group {
	return &Group{
		Region:     c.region,
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

func (c *Client) groups(gConfig interface{}) error {
	var config iam.ListGroupsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["iamGroup"] {
		err := c.db.AutoMigrate(
			&Group{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["iamGroup"] = true
	}
	for {
		output, err := c.svc.ListGroups(&config)
		if err != nil {
			return err
		}
		c.log.Debug("deleting previous Groups", zap.String("region", c.region), zap.String("account_id", c.accountID))
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Group{})
		common.ChunkedCreate(c.db, c.transformGroups(output.Groups))
		c.log.Info("populating Groups", zap.Int("count", len(output.Groups)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
