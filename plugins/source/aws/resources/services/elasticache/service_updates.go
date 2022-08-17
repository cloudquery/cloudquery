package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource service_updates --config ./gen.hcl --output .
func ServiceUpdates() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_service_updates",
		Description:  "An update that you can apply to your Redis clusters.",
		Resolver:     fetchElasticacheServiceUpdates,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticache"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "name"}},
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
				Name:        "auto_update_after_recommended_apply_by_date",
				Description: "Indicates whether the service update will be automatically applied once the recommended apply-by date has expired.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "engine",
				Description: "The Elasticache engine to which the update applies",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The Elasticache engine version to which the update applies",
				Type:        schema.TypeString,
			},
			{
				Name:        "estimated_update_time",
				Description: "The estimated length of time the service update will take",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "Provides details of the service update",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceUpdateDescription"),
			},
			{
				Name:        "end_date",
				Description: "The date after which the service update is no longer available",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ServiceUpdateEndDate"),
			},
			{
				Name:        "name",
				Description: "The unique ID of the service update",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceUpdateName"),
			},
			{
				Name:        "recommended_apply_by_date",
				Description: "The recommendend date to apply the service update in order to ensure compliance. For information on compliance, see Self-Service Security Updates for Compliance (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/elasticache-compliance.html#elasticache-compliance-self-service).",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ServiceUpdateRecommendedApplyByDate"),
			},
			{
				Name:        "release_date",
				Description: "The date when the service update is initially available",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ServiceUpdateReleaseDate"),
			},
			{
				Name:        "severity",
				Description: "The severity of the service update",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceUpdateSeverity"),
			},
			{
				Name:        "status",
				Description: "The status of the service update",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceUpdateStatus"),
			},
			{
				Name:        "type",
				Description: "Reflects the nature of the service update",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceUpdateType"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElasticacheServiceUpdates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeServiceUpdatesPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.ServiceUpdates
	}
	return nil
}
