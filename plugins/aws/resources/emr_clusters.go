package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EmrClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_emr_clusters",
		Description:  "The summary description of the cluster.",
		Resolver:     fetchEmrClusters,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:        "cluster_arn",
				Description: "The Amazon Resource Name of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_id",
				Description: "The unique identifier for the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "name",
				Description: "The name of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "normalized_instance_hours",
				Description: "An approximation of the cost of the cluster, represented in m1.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost where the cluster is launched.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_state",
				Description: "The current state of the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.State"),
			},
			{
				Name:        "status_state_change_reason_code",
				Description: "The programmatic code for the state change reason.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.StateChangeReason.Code"),
			},
			{
				Name:        "status_state_change_reason_message",
				Description: "The descriptive message for the state change reason.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.StateChangeReason.Message"),
			},
			{
				Name:        "status_timeline_creation_date_time",
				Description: "The creation date and time of the cluster.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Status.Timeline.CreationDateTime"),
			},
			{
				Name:        "status_timeline_end_date_time",
				Description: "The date and time when the cluster was terminated.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Status.Timeline.EndDateTime"),
			},
			{
				Name:        "status_timeline_ready_date_time",
				Description: "The date and time when the cluster was ready to run steps.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Status.Timeline.ReadyDateTime"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEmrClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config emr.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().EMR
	for {
		response, err := svc.ListClusters(ctx, &config, func(options *emr.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Clusters
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
