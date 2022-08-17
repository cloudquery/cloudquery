package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource users --config ./gen.hcl --output .
func Users() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_users",
		Description:  "Describes Elasticache users",
		Resolver:     fetchElasticacheUsers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticache"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "access_string",
				Description: "Access permissions string used for this user.",
				Type:        schema.TypeString,
			},
			{
				Name:        "authentication_password_count",
				Description: "The number of passwords belonging to the user",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Authentication.PasswordCount"),
			},
			{
				Name:        "authentication_type",
				Description: "Indicates whether the user requires a password to authenticate.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Authentication.Type"),
			},
			{
				Name:        "engine",
				Description: "The current supported value is Redis.",
				Type:        schema.TypeString,
			},
			{
				Name:        "minimum_engine_version",
				Description: "The minimum engine version required, which is Redis 6.0",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "Indicates the user status",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_group_ids",
				Description: "Returns a list of the user group IDs the user belongs to.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "user_id",
				Description: "The ID of the user.",
				Type:        schema.TypeString,
			},
			{
				Name:        "user_name",
				Description: "The username of the user.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticacheUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeUsersPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.Users
	}
	return nil
}
