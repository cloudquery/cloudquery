package s3

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/mapstructure"
	"github.com/cloudquery/cloudquery/providers/common"
	"go.uber.org/zap"
	"time"
)

type Bucket struct {
	ID           uint `gorm:"primarykey"`
	AccountID    string
	Region       string
	CreationDate *time.Time
	Name         *string
	Grants       []*BucketGrant    `gorm:"constraint:OnDelete:CASCADE;"`
	CORSRules    []*BucketCorsRule `gorm:"constraint:OnDelete:CASCADE;"`

	// The bucket policy as a JSON document.
	Policy *string

	// Specifies whether MFA delete is enabled in the bucket versioning configuration.
	// This element is only returned if the bucket has been configured with MFA
	// delete. If the bucket has never been so configured, this element is not returned.
	MFADelete *string

	// The versioning state of the bucket.
	Status *string
}

type BucketGrant struct {
	ID       uint `gorm:"primarykey"`
	BucketID uint

	// The person being granted permissions.
	Grantee *s3.Grantee `gorm:"embedded;embeddedPrefix:s3_grantee_"`

	// Specifies the permission given to the grantee.
	Permission *string
}

type BucketCorsRule struct {
	ID       uint `gorm:"primarykey"`
	BucketID uint

	// Headers that are specified in the Access-Control-Request-Headers header.
	// These headers are allowed in a preflight OPTIONS request. In response to
	// any preflight OPTIONS request, Amazon S3 returns any requested headers that
	// are allowed.
	AllowedHeaders *string

	// An HTTP method that you allow the origin to execute. Valid values are GET,
	// PUT, HEAD, POST, and DELETE.
	//
	// AllowedMethods is a required field
	AllowedMethods *string

	// One or more origins you want customers to be able to access the bucket from.
	//
	// AllowedOrigins is a required field
	AllowedOrigins *string

	// One or more headers in the response that you want customers to be able to
	// access from their applications (for example, from a JavaScript XMLHttpRequest
	// object).
	ExposeHeaders *string

	// The time in seconds that your browser is to cache the preflight response
	// for the specified resource.
	MaxAgeSeconds *int64
}

func (c *Client) transformGrants(values []*s3.Grant) []*BucketGrant {
	var tValues []*BucketGrant
	for _, v := range values {
		tValues = append(tValues, &BucketGrant{
			Grantee:    v.Grantee,
			Permission: v.Permission,
		})
	}
	return tValues
}

func (c *Client) transformBucketCorsRules(values []*s3.CORSRule) []*BucketCorsRule {
	var tValues []*BucketCorsRule
	for _, v := range values {
		tValues = append(tValues, &BucketCorsRule{
			AllowedHeaders: common.StringListToString(v.AllowedHeaders),
			AllowedMethods: common.StringListToString(v.AllowedMethods),
			AllowedOrigins: common.StringListToString(v.AllowedOrigins),
			ExposeHeaders:  common.StringListToString(v.ExposeHeaders),
			MaxAgeSeconds:  v.MaxAgeSeconds,
		})
	}
	return tValues
}

func (c *Client) transformBucket(value *s3.Bucket) (*Bucket, error) {
	aclOutput, err := c.svc.GetBucketAcl(&s3.GetBucketAclInput{Bucket: value.Name})
	if err != nil {
		return nil, err
	}

	CORSOutput, err := c.svc.GetBucketCors(&s3.GetBucketCorsInput{
		Bucket: value.Name,
	})
	if err != nil && err.(awserr.Error).Code() != "NoSuchCORSConfiguration" {
		return nil, err
	}

	policyOutput, err := c.svc.GetBucketPolicy(&s3.GetBucketPolicyInput{
		Bucket: value.Name,
	})
	if err != nil && err.(awserr.Error).Code() != "NoSuchBucketPolicy" {
		return nil, err
	}

	versioningOutput, err := c.svc.GetBucketVersioning(&s3.GetBucketVersioningInput{
		Bucket: value.Name,
	})
	if err != nil {
		return nil, err
	}

	return &Bucket{
		Region:       c.region,
		AccountID:    c.accountID,
		CreationDate: value.CreationDate,
		Name:         value.Name,
		Grants:       c.transformGrants(aclOutput.Grants),
		CORSRules:    c.transformBucketCorsRules(CORSOutput.CORSRules),
		Policy:       policyOutput.Policy,
		Status:       versioningOutput.Status,
		MFADelete:    versioningOutput.MFADelete,
	}, nil
}

func (c *Client) transformBuckets(values []*s3.Bucket) ([]*Bucket, error) {
	var tValues []*Bucket
	for _, v := range values {
		tBucket, err := c.transformBucket(v)
		if err != nil {
			return nil, err
		}
		tValues = append(tValues, tBucket)
	}
	return tValues, nil
}

func (c *Client) Buckets(gConfig interface{}) error {
	var config s3.ListBucketsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["s3Bucket"] {
		err := c.db.AutoMigrate(
			&Bucket{},
			&BucketGrant{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["s3Bucket"] = true
	}
	output, err := c.svc.ListBuckets(&config)
	if err != nil {
		return err
	}
	c.log.Debug("deleting previous Buckets", zap.String("region", c.region), zap.String("account_id", c.accountID))
	c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Bucket{})
	tBuckets, err := c.transformBuckets(output.Buckets)
	if err != nil {
		return err
	}
	common.ChunkedCreate(c.db, tBuckets)
	c.log.Info("populating Buckets", zap.Int("count", len(output.Buckets)))

	return nil
}
