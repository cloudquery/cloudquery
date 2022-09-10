package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EksClusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_eks_clusters",
		Description: "An object representing an Amazon EKS cluster.",
		Resolver:    fetchEksClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("eks"),
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the cluster.",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "certificate_authority",
				Type: schema.TypeJSON,
			},
			{
				Name:          "client_request_token",
				Description:   "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.",
				Type:          schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The Unix epoch timestamp in seconds for when the cluster was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "endpoint",
				Description: "The endpoint for your Kubernetes API server.",
				Type:        schema.TypeString,
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "kubernetes_network_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KubernetesNetworkConfig"),
			},
			{
				Name:        "name",
				Description: "The name of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform_version",
				Description: "The platform version of your Amazon EKS cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:     "resources_vpc_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourcesVpcConfig"),
			},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The current status of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The metadata that you apply to the cluster to assist with categorization and organization.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "version",
				Description: "The Kubernetes server version for the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encryption_configs",
				Description: "The encryption configuration for the cluster.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("EncryptionConfig"),
			},
			{
				Name:        "logging",
				Description: "Cluster logging",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Logging"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEksClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config eks.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Eks
	for {
		listClustersOutput, err := svc.ListClusters(ctx, &config)
		if err != nil {
			return err
		}
		for _, name := range listClustersOutput.Clusters {
			describeClusterOutput, err := svc.DescribeCluster(ctx, &eks.DescribeClusterInput{Name: &name}, func(options *eks.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			res <- describeClusterOutput.Cluster
		}
		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}
