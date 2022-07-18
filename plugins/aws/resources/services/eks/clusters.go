package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EksClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_eks_clusters",
		Description:  "An object representing an Amazon EKS cluster.",
		Resolver:     fetchEksClusters,
		Multiplex:    client.ServiceAccountRegionMultiplexer("eks"),
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
				Description: "The Amazon Resource Name (ARN) of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_authority_data",
				Description: "The Base64-encoded certificate data required to communicate with your cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificateAuthority.Data"),
			},
			{
				Name:          "client_request_token",
				Description:   "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
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
				Name:        "identity_oidc_issuer",
				Description: "The issuer URL for the OIDC identity provider.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Oidc.Issuer"),
			},
			{
				Name:        "kubernetes_network_config_service_ipv4_cidr",
				Description: "The CIDR block that Kubernetes service IP addresses are assigned from.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KubernetesNetworkConfig.ServiceIpv4Cidr"),
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
				Name:        "resources_vpc_config_cluster_security_group_id",
				Description: "The cluster security group that was created by Amazon EKS for the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourcesVpcConfig.ClusterSecurityGroupId"),
			},
			{
				Name:        "resources_vpc_config_endpoint_private_access",
				Description: "This parameter indicates whether the Amazon EKS private API server endpoint is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourcesVpcConfig.EndpointPrivateAccess"),
			},
			{
				Name:        "resources_vpc_config_endpoint_public_access",
				Description: "This parameter indicates whether the Amazon EKS public API server endpoint is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourcesVpcConfig.EndpointPublicAccess"),
			},
			{
				Name:        "resources_vpc_config_public_access_cidrs",
				Description: "The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ResourcesVpcConfig.PublicAccessCidrs"),
			},
			{
				Name:        "resources_vpc_config_security_group_ids",
				Description: "The security groups associated with the cross-account elastic network interfaces that are used to allow communication between your nodes and the Kubernetes control plane.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ResourcesVpcConfig.SecurityGroupIds"),
			},
			{
				Name:        "resources_vpc_config_subnet_ids",
				Description: "The subnets associated with your cluster.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ResourcesVpcConfig.SubnetIds"),
			},
			{
				Name:        "resources_vpc_config_vpc_id",
				Description: "The VPC associated with your cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourcesVpcConfig.VpcId"),
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
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_eks_cluster_encryption_configs",
				Description:   "The encryption configuration for the cluster.",
				Resolver:      fetchEksClusterEncryptionConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_eks_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "provider_key_arn",
						Description: "Amazon Resource Name (ARN) or alias of the customer master key (CMK).",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Provider.KeyArn"),
					},
					{
						Name:        "resources",
						Description: "Specifies the resources to be encrypted.",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "aws_eks_cluster_loggings",
				Description: "An object representing the enabled or disabled Kubernetes control plane logs for your cluster.",
				Resolver:    fetchEksClusterLoggings,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_eks_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "enabled",
						Description: "If a log type is enabled, that log type exports its control plane logs to CloudWatch Logs.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "types",
						Description: "The available cluster control plane log types.",
						Type:        schema.TypeStringArray,
						Resolver:    resolveEksClusterLoggingTypes,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEksClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config eks.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Eks
	for {
		listClustersOutput, err := svc.ListClusters(ctx, &config, func(options *eks.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, name := range listClustersOutput.Clusters {
			describeClusterOutput, err := svc.DescribeCluster(ctx, &eks.DescribeClusterInput{Name: &name}, func(options *eks.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
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
func fetchEksClusterEncryptionConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*types.Cluster)
	res <- p.EncryptionConfig
	return nil
}
func fetchEksClusterLoggings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*types.Cluster)
	if p.Logging == nil {
		return nil
	}
	res <- p.Logging.ClusterLogging
	return nil
}
func resolveEksClusterLoggingTypes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	logSetup := resource.Item.(types.LogSetup)
	logTypes := make([]string, len(logSetup.Types))
	for i, l := range logSetup.Types {
		logTypes[i] = string(l)
	}
	return diag.WrapError(resource.Set("types", logTypes))
}
