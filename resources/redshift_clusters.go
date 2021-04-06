package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func RedshiftClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_redshift_clusters",
		Resolver:     fetchRedshiftClusters,
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
				Name: "allow_version_upgrade",
				Type: schema.TypeBool,
			},
			{
				Name: "automated_snapshot_retention_period",
				Type: schema.TypeInt,
			},
			{
				Name: "availability_zone",
				Type: schema.TypeString,
			},
			{
				Name: "availability_zone_relocation_status",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_availability_status",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_create_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "cluster_identifier",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_namespace_arn",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_public_key",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_revision_number",
				Type: schema.TypeString,
			},
			{
				Name:     "cluster_snapshot_copy_status_destination_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterSnapshotCopyStatus.DestinationRegion"),
			},
			{
				Name:     "cluster_snapshot_copy_status_manual_snapshot_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ClusterSnapshotCopyStatus.ManualSnapshotRetentionPeriod"),
			},
			{
				Name:     "cluster_snapshot_copy_status_retention_period",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClusterSnapshotCopyStatus.RetentionPeriod"),
			},
			{
				Name:     "cluster_snapshot_copy_status_snapshot_copy_grant_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterSnapshotCopyStatus.SnapshotCopyGrantName"),
			},
			{
				Name: "cluster_status",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_subnet_group_name",
				Type: schema.TypeString,
			},
			{
				Name: "cluster_version",
				Type: schema.TypeString,
			},
			{
				Name:     "db_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBName"),
			},
			{
				Name:     "data_transfer_progress_current_rate_in_mega_bytes_per_second",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("DataTransferProgress.CurrentRateInMegaBytesPerSecond"),
			},
			{
				Name:     "data_transfer_progress_data_transferred_in_mega_bytes",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DataTransferProgress.DataTransferredInMegaBytes"),
			},
			{
				Name:     "data_transfer_progress_elapsed_time_in_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DataTransferProgress.ElapsedTimeInSeconds"),
			},
			{
				Name:     "data_transfer_progress_estimated_time_to_completion_in_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DataTransferProgress.EstimatedTimeToCompletionInSeconds"),
			},
			{
				Name:     "data_transfer_progress_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataTransferProgress.Status"),
			},
			{
				Name:     "data_transfer_progress_total_data_in_mega_bytes",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DataTransferProgress.TotalDataInMegaBytes"),
			},
			{
				Name:     "elastic_ip_status_elastic_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ElasticIpStatus.ElasticIp"),
			},
			{
				Name:     "elastic_ip_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ElasticIpStatus.Status"),
			},
			{
				Name: "elastic_resize_number_of_node_options",
				Type: schema.TypeString,
			},
			{
				Name: "encrypted",
				Type: schema.TypeBool,
			},
			{
				Name:     "endpoint_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint.Address"),
			},
			{
				Name:     "endpoint_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Endpoint.Port"),
			},
			{
				Name: "enhanced_vpc_routing",
				Type: schema.TypeBool,
			},
			{
				Name: "expected_next_snapshot_schedule_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "expected_next_snapshot_schedule_time_status",
				Type: schema.TypeString,
			},
			{
				Name:     "hsm_status_hsm_client_certificate_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HsmStatus.HsmClientCertificateIdentifier"),
			},
			{
				Name:     "hsm_status_hsm_configuration_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HsmStatus.HsmConfigurationIdentifier"),
			},
			{
				Name:     "hsm_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HsmStatus.Status"),
			},
			{
				Name: "kms_key_id",
				Type: schema.TypeString,
			},
			{
				Name: "maintenance_track_name",
				Type: schema.TypeString,
			},
			{
				Name: "manual_snapshot_retention_period",
				Type: schema.TypeInt,
			},
			{
				Name: "master_username",
				Type: schema.TypeString,
			},
			{
				Name: "modify_status",
				Type: schema.TypeString,
			},
			{
				Name: "next_maintenance_window_start_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "node_type",
				Type: schema.TypeString,
			},
			{
				Name: "number_of_nodes",
				Type: schema.TypeInt,
			},
			{
				Name: "pending_actions",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "pending_modified_values_automated_snapshot_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingModifiedValues.AutomatedSnapshotRetentionPeriod"),
			},
			{
				Name:     "pending_modified_values_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.ClusterIdentifier"),
			},
			{
				Name:     "pending_modified_values_cluster_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.ClusterType"),
			},
			{
				Name:     "pending_modified_values_cluster_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.ClusterVersion"),
			},
			{
				Name:     "pending_modified_values_encryption_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.EncryptionType"),
			},
			{
				Name:     "pending_modified_values_enhanced_vpc_routing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PendingModifiedValues.EnhancedVpcRouting"),
			},
			{
				Name:     "pending_modified_values_maintenance_track_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.MaintenanceTrackName"),
			},
			{
				Name:     "pending_modified_values_master_user_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.MasterUserPassword"),
			},
			{
				Name:     "pending_modified_values_node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.NodeType"),
			},
			{
				Name:     "pending_modified_values_number_of_nodes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingModifiedValues.NumberOfNodes"),
			},
			{
				Name:     "pending_modified_values_publicly_accessible",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PendingModifiedValues.PubliclyAccessible"),
			},
			{
				Name: "preferred_maintenance_window",
				Type: schema.TypeString,
			},
			{
				Name: "publicly_accessible",
				Type: schema.TypeBool,
			},
			{
				Name:     "resize_info_allow_cancel_resize",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ResizeInfo.AllowCancelResize"),
			},
			{
				Name:     "resize_info_resize_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResizeInfo.ResizeType"),
			},
			{
				Name:     "restore_status_current_restore_rate_in_mega_bytes_per_second",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("RestoreStatus.CurrentRestoreRateInMegaBytesPerSecond"),
			},
			{
				Name:     "restore_status_elapsed_time_in_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("RestoreStatus.ElapsedTimeInSeconds"),
			},
			{
				Name:     "restore_status_estimated_time_to_completion_in_seconds",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("RestoreStatus.EstimatedTimeToCompletionInSeconds"),
			},
			{
				Name:     "restore_status_progress_in_mega_bytes",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("RestoreStatus.ProgressInMegaBytes"),
			},
			{
				Name:     "restore_status_snapshot_size_in_mega_bytes",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("RestoreStatus.SnapshotSizeInMegaBytes"),
			},
			{
				Name:     "restore_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RestoreStatus.Status"),
			},
			{
				Name: "snapshot_schedule_identifier",
				Type: schema.TypeString,
			},
			{
				Name: "snapshot_schedule_state",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRedshiftClusterTags,
			},
			{
				Name: "total_storage_capacity_in_mega_bytes",
				Type: schema.TypeBigInt,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_redshift_cluster_nodes",
				Resolver: fetchRedshiftClusterNodes,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "node_role",
						Type: schema.TypeString,
					},
					{
						Name:     "private_ip_address",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PrivateIPAddress"),
					},
					{
						Name:     "public_ip_address",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PublicIPAddress"),
					},
				},
			},
			{
				Name:     "aws_redshift_cluster_parameter_groups",
				Resolver: fetchRedshiftClusterParameterGroups,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "parameter_apply_status",
						Type: schema.TypeString,
					},
					{
						Name: "parameter_group_name",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_redshift_cluster_parameter_group_cluster_parameter_status_lists",
						Resolver: fetchRedshiftClusterParameterGroupClusterParameterStatusLists,
						Columns: []schema.Column{
							{
								Name:     "cluster_parameter_group_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "parameter_apply_error_description",
								Type: schema.TypeString,
							},
							{
								Name: "parameter_apply_status",
								Type: schema.TypeString,
							},
							{
								Name: "parameter_name",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_redshift_cluster_security_groups",
				Resolver: fetchRedshiftClusterSecurityGroups,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "cluster_security_group_name",
						Type: schema.TypeString,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_redshift_cluster_deferred_maintenance_windows",
				Resolver: fetchRedshiftClusterDeferredMaintenanceWindows,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "defer_maintenance_end_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "defer_maintenance_identifier",
						Type: schema.TypeString,
					},
					{
						Name: "defer_maintenance_start_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "aws_redshift_cluster_endpoint_vpc_endpoints",
				Resolver: fetchRedshiftClusterEndpointVpcEndpoints,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "vpc_endpoint_id",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_id",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces",
						Resolver: fetchRedshiftClusterEndpointVpcEndpointNetworkInterfaces,
						Columns: []schema.Column{
							{
								Name:     "clusterendpoint_vpc_endpoint_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "availability_zone",
								Type: schema.TypeString,
							},
							{
								Name: "network_interface_id",
								Type: schema.TypeString,
							},
							{
								Name: "private_ip_address",
								Type: schema.TypeString,
							},
							{
								Name: "subnet_id",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_redshift_cluster_iam_roles",
				Resolver: fetchRedshiftClusterIamRoles,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "apply_status",
						Type: schema.TypeString,
					},
					{
						Name: "iam_role_arn",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_redshift_cluster_vpc_security_groups",
				Resolver: fetchRedshiftClusterVpcSecurityGroups,
				Columns: []schema.Column{
					{
						Name:     "cluster_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_security_group_id",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRedshiftClusters(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var config redshift.DescribeClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Redshift
	for {
		response, err := svc.DescribeClusters(ctx, &config, func(o *redshift.Options) {
			o.Region = c.Region
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
func resolveRedshiftClusterTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Cluster)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
func fetchRedshiftClusterNodes(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.ClusterNodes
	return nil
}
func fetchRedshiftClusterParameterGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.ClusterParameterGroups
	return nil
}
func fetchRedshiftClusterParameterGroupClusterParameterStatusLists(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	parameterGroup, ok := parent.Item.(types.ClusterParameterGroupStatus)
	if !ok {
		return fmt.Errorf("not redshift cluster parameter group")
	}
	res <- parameterGroup.ClusterParameterStatusList
	return nil
}
func fetchRedshiftClusterSecurityGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.ClusterSecurityGroups
	return nil
}
func fetchRedshiftClusterDeferredMaintenanceWindows(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.DeferredMaintenanceWindows
	return nil
}
func fetchRedshiftClusterEndpointVpcEndpoints(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.Endpoint.VpcEndpoints
	return nil
}
func fetchRedshiftClusterEndpointVpcEndpointNetworkInterfaces(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vpcEndpoint, ok := parent.Item.(types.VpcEndpoint)
	if !ok {
		return fmt.Errorf("not vpc endpoint")
	}
	res <- vpcEndpoint.NetworkInterfaces
	return nil
}
func fetchRedshiftClusterIamRoles(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.IamRoles
	return nil
}
func fetchRedshiftClusterVpcSecurityGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.VpcSecurityGroups
	return nil
}
