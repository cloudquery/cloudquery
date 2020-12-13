package ecr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Image struct {
	ID          uint `gorm:"primarykey"`
	AccountID   string
	Region      string
	ImageDigest *string
	ImageTag    *string
}

func (Image) TableName() string {
	return "aws_ecr_images"
}

func (c *Client) transformImageIdentifier(value *ecr.ImageIdentifier) *Image {
	return &Image{
		Region:      c.region,
		AccountID:   c.accountID,
		ImageDigest: value.ImageDigest,
		ImageTag:    value.ImageTag,
	}
}

func (c *Client) transformImageIdentifiers(values []*ecr.ImageIdentifier) []*Image {
	var tValues []*Image
	for _, v := range values {
		tValues = append(tValues, c.transformImageIdentifier(v))
	}
	return tValues
}

func (c *Client) imageIdentifiers(gConfig interface{}) error {
	var config ecr.ListImagesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ecrImageIdentifier"] {
		err := c.db.AutoMigrate(
			&Image{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ecrImageIdentifier"] = true
	}
	for {
		output, err := c.svc.ListImages(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Image{})
		common.ChunkedCreate(c.db, c.transformImageIdentifiers(output.ImageIds))
		c.log.Info("Fetched resources", zap.String("resource", "ecr.images"), zap.Int("count", len(output.ImageIds)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
