package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"log"
	"time"
)

type AccessKey struct {
	ID                  uint `gorm:"primarykey"`
	AccountID           string
	AccessKeyId         *string
	CreateDate          *time.Time
	Status              *string
	UserName            *string
	LastUsed            *time.Time
	LastUsedServiceName *string
}

func (c *Client) transformAccessKey(value *iam.AccessKeyMetadata) *AccessKey {
	output, err := c.svc.GetAccessKeyLastUsed(&iam.GetAccessKeyLastUsedInput{AccessKeyId: value.AccessKeyId})
	if err != nil {
		log.Fatal(err)
	}

	res := AccessKey{
		AccountID:           c.accountID,
		AccessKeyId:         value.AccessKeyId,
		CreateDate:          value.CreateDate,
		Status:              value.Status,
		UserName:            value.UserName,
		LastUsed:            output.AccessKeyLastUsed.LastUsedDate,
		LastUsedServiceName: output.AccessKeyLastUsed.ServiceName,
	}
	if output.AccessKeyLastUsed != nil {
		res.LastUsed = output.AccessKeyLastUsed.LastUsedDate
		res.LastUsedServiceName = output.AccessKeyLastUsed.ServiceName
	}
	return &res
}

func (c *Client) transformAccessKeys(values []*iam.AccessKeyMetadata) []*AccessKey {
	var tValues []*AccessKey
	for _, v := range values {
		tValues = append(tValues, c.transformAccessKey(v))
	}
	return tValues
}

func (c *Client) accessKeys(gConfig interface{}) error {
	var config iam.ListAccessKeysInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["iamAccessKey"] {
		err := c.db.AutoMigrate(
			&AccessKey{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["iamAccessKey"] = true
	}
	for {
		output, err := c.svc.ListAccessKeys(&config)
		if err != nil {
			return err
		}
		c.db.Where("account_id = ?", c.accountID).Delete(&AccessKey{})
		common.ChunkedCreate(c.db, c.transformAccessKeys(output.AccessKeyMetadata))
		c.log.Info("Fetched resources", zap.Int("count", len(output.AccessKeyMetadata)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
