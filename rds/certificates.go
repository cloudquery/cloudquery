package rds

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Certificate struct {
	_                         interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                        uint        `gorm:"primarykey"`
	AccountID                 string
	Region                    string
	CertificateArn            *string `neo:"unique"`
	CertificateIdentifier     *string
	CertificateType           *string
	CustomerOverride          *bool
	CustomerOverrideValidTill *time.Time
	Thumbprint                *string
	ValidFrom                 *time.Time
	ValidTill                 *time.Time
}

func (Certificate) TableName() string {
	return "aws_rds_certificates"
}

func (c *Client) transformCertificates(values *[]types.Certificate) []*Certificate {
	var tValues []*Certificate
	for _, value := range *values {
		tValues = append(tValues, &Certificate{
			Region:                    c.region,
			AccountID:                 c.accountID,
			CertificateArn:            value.CertificateArn,
			CertificateIdentifier:     value.CertificateIdentifier,
			CertificateType:           value.CertificateType,
			CustomerOverride:          value.CustomerOverride,
			CustomerOverrideValidTill: value.CustomerOverrideValidTill,
			Thumbprint:                value.Thumbprint,
			ValidFrom:                 value.ValidFrom,
			ValidTill:                 value.ValidTill,
		})
	}
	return tValues
}

var CertificateTables = []interface{}{
	&Certificate{},
}

func (c *Client) certificates(gConfig interface{}) error {
	var config rds.DescribeCertificatesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	ctx := context.Background()
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(CertificateTables...)

	for {
		output, err := c.svc.DescribeCertificates(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformCertificates(&output.Certificates))
		c.log.Info("Fetched resources", zap.String("resource", "rds.certificates"), zap.Int("count", len(output.Certificates)))
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
