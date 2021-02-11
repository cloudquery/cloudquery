package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/cq-provider-aws/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Bucket struct {
	_               interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID              uint        `gorm:"primarykey"`
	AccountID       string      `neo:"unique"`
	Region          string      `neo:"unique"`
	CreationDate    *time.Time
	Name            *string                 `neo:"unique"`
	Grants          []*BucketGrant          `gorm:"constraint:OnDelete:CASCADE;"`
	CORSRules       []*BucketCorsRule       `gorm:"constraint:OnDelete:CASCADE;"`
	EncryptionRules []*BucketEncryptionRule `gorm:"constraint:OnDelete:CASCADE;"`
	// The bucket policy as a JSON document.
	Policy *string

	// Specifies whether MFA delete is enabled in the bucket versioning configuration.
	// This element is only returned if the bucket has been configured with MFA
	// delete. If the bucket has never been so configured, this element is not returned.
	MFADelete *string

	// The versioning state of the bucket.
	Status *string

	// Logging
	LoggingTargetPrefix *string
	LoggingTargetBucket *string
}

func (Bucket) TableName() string {
	return "aws_s3_buckets"
}

type BucketEncryptionRule struct {
	ID        uint   `gorm:"primarykey"`
	BucketID  uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	KMSMasterKeyID *string
	SSEAlgorithm   *string
}

func (BucketEncryptionRule) TableName() string {
	return "aws_s3_bucket_encryption_rules"
}

type BucketGrant struct {
	ID        uint   `gorm:"primarykey"`
	BucketID  uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	// The person being granted permissions.
	GranteeDisplayName  *string
	GranteeEmailAddress *string
	GranteeID           *string
	GranteeType         *string
	GranteeURI          *string

	// Specifies the permission given to the grantee.
	Permission *string
}

func (BucketGrant) TableName() string {
	return "aws_s3_bucket_grants"
}

type BucketCorsRule struct {
	ID        uint   `gorm:"primarykey"`
	BucketID  uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

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

func (BucketCorsRule) TableName() string {
	return "aws_s3_bucket_cors_rules"
}

func (c *Client) transformGrants(values []*s3.Grant) []*BucketGrant {
	var tValues []*BucketGrant
	for _, v := range values {
		tValue := BucketGrant{
			AccountID:  c.accountID,
			Region:     c.region,
			Permission: v.Permission,
		}
		if v.Grantee != nil {
			tValue.GranteeDisplayName = v.Grantee.DisplayName
			tValue.GranteeType = v.Grantee.Type
			tValue.GranteeID = v.Grantee.ID
			tValue.GranteeEmailAddress = v.Grantee.EmailAddress
			tValue.GranteeURI = v.Grantee.URI
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformBucketCorsRules(values []*s3.CORSRule) []*BucketCorsRule {
	var tValues []*BucketCorsRule
	for _, v := range values {
		tValues = append(tValues, &BucketCorsRule{
			AccountID:      c.accountID,
			Region:         c.region,
			AllowedHeaders: common.StringListToString(v.AllowedHeaders),
			AllowedMethods: common.StringListToString(v.AllowedMethods),
			AllowedOrigins: common.StringListToString(v.AllowedOrigins),
			ExposeHeaders:  common.StringListToString(v.ExposeHeaders),
			MaxAgeSeconds:  v.MaxAgeSeconds,
		})
	}
	return tValues
}

func (c *Client) transformEncryptionRules(values []*s3.ServerSideEncryptionRule) []*BucketEncryptionRule {
	var tValues []*BucketEncryptionRule
	for _, v := range values {
		if v.ApplyServerSideEncryptionByDefault != nil {
			tValues = append(tValues, &BucketEncryptionRule{
				AccountID:      c.accountID,
				Region:         c.region,
				KMSMasterKeyID: v.ApplyServerSideEncryptionByDefault.KMSMasterKeyID,
				SSEAlgorithm:   v.ApplyServerSideEncryptionByDefault.SSEAlgorithm,
			})
		}
	}
	return tValues
}

func (c *Client) transformBucket(value *s3.Bucket) (*Bucket, error) {
	loggingOutput, err := c.svc.GetBucketLogging(&s3.GetBucketLoggingInput{Bucket: value.Name})
	if err != nil {
		return nil, err
	}

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

	encryptionOutput, err := c.svc.GetBucketEncryption(&s3.GetBucketEncryptionInput{
		Bucket: value.Name,
	})
	if err != nil {
		var aerr awserr.Error
		var ok bool
		if aerr, ok = err.(awserr.Error); !ok {
			return nil, err
		}
		if aerr.Code() != "ServerSideEncryptionConfigurationNotFoundError" {
			return nil, err
		}
	}
	var EncryptionRules []*BucketEncryptionRule
	if encryptionOutput.ServerSideEncryptionConfiguration != nil {
		EncryptionRules = c.transformEncryptionRules(encryptionOutput.ServerSideEncryptionConfiguration.Rules)
	}

	res := Bucket{
		Region:          c.region,
		AccountID:       c.accountID,
		CreationDate:    value.CreationDate,
		Name:            value.Name,
		Grants:          c.transformGrants(aclOutput.Grants),
		CORSRules:       c.transformBucketCorsRules(CORSOutput.CORSRules),
		EncryptionRules: EncryptionRules,
		Policy:          policyOutput.Policy,
		Status:          versioningOutput.Status,
		MFADelete:       versioningOutput.MFADelete,
	}

	if loggingOutput.LoggingEnabled != nil {
		res.LoggingTargetBucket = loggingOutput.LoggingEnabled.TargetBucket
		res.LoggingTargetPrefix = loggingOutput.LoggingEnabled.TargetPrefix
	}

	return &res, nil
}

func (c *Client) transformBuckets(values []*s3.Bucket) ([]*Bucket, error) {
	var tValues []*Bucket
	for _, v := range values {
		output, err := c.svc.GetBucketLocation(&s3.GetBucketLocationInput{
			Bucket: v.Name,
		})

		if err != nil {
			if err.(awserr.Error).Code() == "NoSuchBucket" {
				// https://aws.amazon.com/premiumsupport/knowledge-center/s3-listing-deleted-bucket/
				// deleted buckets may show up
				c.log.Debug("Skipping bucket (already deleted)", zap.String("bucket", *v.Name))
				continue
			}
			return nil, err
		}
		c.region = "us-east-1"
		if output.LocationConstraint != nil {
			// This is a weird corner case by AWS API https://github.com/aws/aws-sdk-net/issues/323#issuecomment-196584538
			c.region = aws.StringValue(output.LocationConstraint)
		}
		c.awsConfig.Region = aws.String(c.region)
		c.svc = s3.New(c.session, c.awsConfig)

		tBucket, err := c.transformBucket(v)
		if err != nil {
			return nil, err
		}
		tValues = append(tValues, tBucket)
	}
	return tValues, nil
}

var BucketTables = []interface{}{
	&Bucket{},
	&BucketGrant{},
	&BucketCorsRule{},
	&BucketEncryptionRule{},
}

func (c *Client) buckets(gConfig interface{}) error {
	var config s3.ListBucketsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	output, err := c.svc.ListBuckets(&config)
	if err != nil {
		return err
	}
	c.db.Where("account_id", c.accountID).Delete(BucketTables...)
	tBuckets, err := c.transformBuckets(output.Buckets)
	if err != nil {
		return err
	}
	c.db.ChunkedCreate(tBuckets)
	c.log.Info("Fetched resources", zap.String("resource", "s3.buckets"), zap.Int("count", len(tBuckets)))

	return nil
}
