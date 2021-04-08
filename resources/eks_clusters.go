package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EksClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_eks_clusters",
		Resolver:     fetchEksClusters,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
			},
			{
				Name:     "certificate_authority_data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateAuthority.Data"),
			},
			{
				Name: "client_request_token",
				Type: schema.TypeString,
			},
			{
				Name: "created_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "endpoint",
				Type: schema.TypeString,
			},
			{
				Name:     "identity_oidc_issuer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Identity.Oidc.Issuer"),
			},
			{
				Name:     "kubernetes_network_config_service_ipv4_cidr",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KubernetesNetworkConfig.ServiceIpv4Cidr"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "platform_version",
				Type: schema.TypeString,
			},
			{
				Name:     "resources_vpc_config_cluster_security_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourcesVpcConfig.ClusterSecurityGroupId"),
			},
			{
				Name:     "resources_vpc_config_endpoint_private_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ResourcesVpcConfig.EndpointPrivateAccess"),
			},
			{
				Name:     "resources_vpc_config_endpoint_public_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ResourcesVpcConfig.EndpointPublicAccess"),
			},
			{
				Name:     "resources_vpc_config_public_access_cidrs",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ResourcesVpcConfig.PublicAccessCidrs"),
			},
			{
				Name:     "resources_vpc_config_security_group_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ResourcesVpcConfig.SecurityGroupIds"),
			},
			{
				Name:     "resources_vpc_config_subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ResourcesVpcConfig.SubnetIds"),
			},
			{
				Name:     "resources_vpc_config_vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourcesVpcConfig.VpcId"),
			},
			{
				Name: "role_arn",
				Type: schema.TypeString,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "version",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_eks_cluster_encryption_configs",
				Resolver: fetchEksClusterEncryptionConfigs,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "provider_key_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Provider.KeyArn"),
					},
					{
						Name: "resources",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "aws_eks_cluster_logging_cluster_loggings",
				Resolver: fetchEksClusterLoggingClusterLoggings,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "enabled",
						Type: schema.TypeBool,
					},
					{
						Name:     "types",
						Type:     schema.TypeStringArray,
						Resolver: resolveEksClusterLoggingClusterLoggingsLogType,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEksClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config eks.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Eks
	for {
		listClustersOutput, err := svc.ListClusters(ctx, &config, func(options *eks.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, name := range listClustersOutput.Clusters {
			describeClusterOutput, err := svc.DescribeCluster(ctx, &eks.DescribeClusterInput{Name: &name}, func(options *eks.Options) {
				options.Region = c.Region
			})
			if err != nil {
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
func fetchEksClusterEncryptionConfigs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(*types.Cluster)
	res <- p.EncryptionConfig
	return nil
}
func fetchEksClusterLoggingClusterLoggings(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(*types.Cluster)
	res <- p.Logging.ClusterLogging
	return nil
}

func resolveEksClusterLoggingClusterLoggingsLogType(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	logSetup := resource.Item.(types.LogSetup)
	logTypes := make([]string, len(logSetup.Types))
	for i, l := range logSetup.Types {
		logTypes[i] = string(l)
	}
	resource.Set("types", logTypes)
	return nil
}
