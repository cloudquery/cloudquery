package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DaxClusters() *schema.Table {
	return &schema.Table{
		Name:          "aws_dax_clusters",
		Description:   "Information about a DAX cluster.",
		Resolver:      fetchDaxClusters,
		Multiplex:     client.ServiceAccountRegionMultiplexer("dax"),
		IgnoreInTests: true,
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
				Name:        "tags",
				Description: "The tags associated with the cluster.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDaxClusterTags,
			},
			{
				Name:        "active_nodes",
				Description: "The number of nodes in the cluster that are active (i.e., capable of serving requests).",
				Type:        schema.TypeInt,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) that uniquely identifies the cluster.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "cluster_discovery_endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ClusterDiscoveryEndpoint"),
			},
			{
				Name:        "cluster_endpoint_encryption_type",
				Description: "The type of encryption supported by the cluster's endpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the DAX cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterName"),
			},
			{
				Name:        "description",
				Description: "The description of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "iam_role_arn",
				Description: "A valid Amazon Resource Name (ARN) that identifies an IAM role",
				Type:        schema.TypeString,
			},
			{
				Name:        "node_ids_to_remove",
				Description: "A list of nodes to be removed from the cluster.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "node_type",
				Description: "The node type for the nodes in the cluster",
				Type:        schema.TypeString,
			},
			{
				Name:     "notification_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NotificationConfiguration"),
			},
			{
				Name:     "parameter_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ParameterGroup"),
			},
			{
				Name:        "preferred_maintenance_window",
				Description: "A range of time when maintenance of DAX cluster software will be performed",
				Type:        schema.TypeString,
			},
			{
				Name:        "sse_description",
				Type: 			schema.TypeJSON,
				Resolver: schema.PathResolver("SSEDescription"),
			},
			{
				Name:        "security_groups",
				Description: "A list of security groups, and the status of each, for the nodes in the cluster.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDaxClusterSecurityGroups,
			},
			{
				Name:        "status",
				Description: "The current status of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_group",
				Description: "The subnet group where the DAX cluster is running.",
				Type:        schema.TypeString,
			},
			{
				Name:        "total_nodes",
				Description: "The total number of nodes in the cluster.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "nodes",
				Description: "Represents an individual node within a DAX cluster.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Nodes"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchDaxClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().DAX

	config := dax.DescribeClustersInput{}
	for {
		output, err := svc.DescribeClusters(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.Clusters

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
func resolveDaxClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.Cluster)

	cl := meta.(*client.Client)
	svc := cl.Services().DAX
	response, err := svc.ListTags(ctx, &dax.ListTagsInput{
		ResourceName: cluster.ClusterArn,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.Tags))
}
func resolveDaxClusterSecurityGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Cluster)
	val := make(map[string]interface{}, len(r.SecurityGroups))
	for i := range r.SecurityGroups {
		val[aws.ToString(r.SecurityGroups[i].SecurityGroupIdentifier)] = map[string]interface{}{
			"identifier": r.SecurityGroups[i].SecurityGroupIdentifier,
			"status":     r.SecurityGroups[i].Status,
		}
	}
	return resource.Set(c.Name, val)
}
