package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	s3controlTypes "github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func S3Resources() []*Resource {
	resources := []*Resource{

		{
			SubService: "accounts",
			Struct:     &s3controlTypes.PublicAccessBlockConfiguration{},
			SkipFields: []string{"ARN"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService: "buckets",
			Struct:     &s3.WrappedBucket{},
			SkipFields: []string{""},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveBucketARN()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"BucketEncryptionRules()",
				"BucketLifecycles()",
				"BucketGrants()",
				"BucketCorsRules()",
			},
		},
		{
			SubService: "bucket_encryption_rules",
			Struct:     &types.ServerSideEncryptionRule{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeUUID,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "bucket_lifecycles",
			Struct:     &types.LifecycleRule{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeUUID,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "bucket_grants",
			Struct:     &types.Grant{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeUUID,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "bucket_cors_rules",
			Struct:     &types.CORSRule{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeUUID,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "s3"
		r.Multiplex = "client.AccountMultiplex"
	}
	return resources
}
