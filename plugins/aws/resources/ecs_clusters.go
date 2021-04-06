package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func EcsClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_ecs_clusters",
		Resolver:     fetchEcsClusters,
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
				Name: "active_services_count",
				Type: schema.TypeInt,
			},
			{
				Name: "attachments_status",
				Type: schema.TypeString,
			},
			{
				Name: "capacity_providers",
				Type: schema.TypeStringArray,
			},
			{
				Name: "cluster_arn",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_name",
				Type: schema.TypeString,
			},
			{
				Name:     "configuration_execute_command_configuration_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.ExecuteCommandConfiguration.KmsKeyId"),
			},
			{
				Name:     "configuration_execute_command_configuration_log_configuration_cloud_watch_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled"),
			},
			{
				Name:     "configuration_execute_command_configuration_log_configuration_cloud_watch_log_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName"),
			},
			{
				Name:     "configuration_execute_command_configuration_log_configuration_s_3_bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName"),
			},
			{
				Name:     "configuration_execute_command_configuration_log_configuration_s_3_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled"),
			},
			{
				Name:     "configuration_execute_command_configuration_log_configuration_s_3_key_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix"),
			},
			{
				Name:     "configuration_execute_command_configuration_logging",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.ExecuteCommandConfiguration.Logging"),
			},
			{
				Name: "pending_tasks_count",
				Type: schema.TypeInt,
			},
			{
				Name: "registered_container_instances_count",
				Type: schema.TypeInt,
			},
			{
				Name: "running_tasks_count",
				Type: schema.TypeInt,
			},
			{
				Name:     "settings",
				Type:     schema.TypeJSON,
				Resolver: resolveEcsClusterSettings,
			},
			{
				Name:     "statistics",
				Type:     schema.TypeJSON,
				Resolver: resolveEcsClusterStatistics,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEcsClusterTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ecs_cluster_attachments",
				Resolver: fetchEcsClusterAttachments,
				Columns: []schema.Column{
					{
						Name:     "clusters_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "id",
						Type: schema.TypeString,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_ecs_cluster_attachment_details",
						Resolver: fetchEcsClusterAttachmentDetails,
						Columns: []schema.Column{
							{
								Name:     "clusterattachments_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "name",
								Type: schema.TypeString,
							},
							{
								Name: "value",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_ecs_cluster_default_capacity_provider_strategies",
				Resolver: fetchEcsClusterDefaultCapacityProviderStrategies,
				Columns: []schema.Column{
					{
						Name:     "clusters_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "capacity_provider",
						Type: schema.TypeString,
					},
					{
						Name: "base",
						Type: schema.TypeInt,
					},
					{
						Name: "weight",
						Type: schema.TypeInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEcsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func resolveEcsClusterSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveEcsClusterStatistics(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func resolveEcsClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchEcsClusterAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchEcsClusterAttachmentDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchEcsClusterDefaultCapacityProviderStrategies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
