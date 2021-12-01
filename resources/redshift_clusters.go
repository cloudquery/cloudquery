package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RedshiftClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_redshift_clusters",
		Description:  "Describes a cluster.",
		Resolver:     fetchRedshiftClusters,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "allow_version_upgrade",
				Description: "A boolean value that, if true, indicates that major version upgrades will be applied automatically to the cluster during the maintenance window.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "automated_snapshot_retention_period",
				Description: "The number of days that automatic cluster snapshots are retained.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "availability_zone",
				Description: "The name of the Availability Zone in which the cluster is located.",
				Type:        schema.TypeString,
			},
			{
				Name:        "availability_zone_relocation_status",
				Description: "Describes the status of the Availability Zone relocation operation.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_availability_status",
				Description: "The availability status of the cluster for queries.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_create_time",
				Description: "The date and time that the cluster was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "id",
				Description: "The unique identifier of the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterIdentifier"),
			},
			{
				Name:        "cluster_namespace_arn",
				Description: "The namespace Amazon Resource Name (ARN) of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_public_key",
				Description: "The public key for the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_revision_number",
				Description: "The specific revision number of the database in the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_snapshot_copy_status_destination_region",
				Description: "The destination region that snapshots are automatically copied to when cross-region snapshot copy is enabled.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterSnapshotCopyStatus.DestinationRegion"),
			},
			{
				Name:        "cluster_snapshot_copy_status_manual_snapshot_retention_period",
				Description: "The number of days that automated snapshots are retained in the destination region after they are copied from a source region.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ClusterSnapshotCopyStatus.ManualSnapshotRetentionPeriod"),
			},
			{
				Name:        "cluster_snapshot_copy_status_retention_period",
				Description: "The number of days that automated snapshots are retained in the destination region after they are copied from a source region.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ClusterSnapshotCopyStatus.RetentionPeriod"),
			},
			{
				Name:        "cluster_snapshot_copy_status_snapshot_copy_grant_name",
				Description: "The name of the snapshot copy grant.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterSnapshotCopyStatus.SnapshotCopyGrantName"),
			},
			{
				Name:        "cluster_status",
				Description: "The current state of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_subnet_group_name",
				Description: "The name of the subnet group that is associated with the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_version",
				Description: "The version ID of the Amazon Redshift engine that is running on the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "db_name",
				Description: "The name of the initial database that was created when the cluster was created.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBName"),
			},
			{
				Name:        "data_transfer_progress_current_rate_in_mega_bytes_per_second",
				Description: "Describes the data transfer rate in MB's per second.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("DataTransferProgress.CurrentRateInMegaBytesPerSecond"),
			},
			{
				Name:        "data_transfer_progress_data_transferred_in_mega_bytes",
				Description: "Describes the total amount of data that has been transferred in MB's.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DataTransferProgress.DataTransferredInMegaBytes"),
			},
			{
				Name:        "data_transfer_progress_elapsed_time_in_seconds",
				Description: "Describes the number of seconds that have elapsed during the data transfer.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DataTransferProgress.ElapsedTimeInSeconds"),
			},
			{
				Name:        "data_transfer_progress_estimated_time_to_completion_in_seconds",
				Description: "Describes the estimated number of seconds remaining to complete the transfer.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DataTransferProgress.EstimatedTimeToCompletionInSeconds"),
			},
			{
				Name:        "data_transfer_progress_status",
				Description: "Describes the status of the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataTransferProgress.Status"),
			},
			{
				Name:        "data_transfer_progress_total_data_in_mega_bytes",
				Description: "Describes the total amount of data to be transferred in megabytes.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DataTransferProgress.TotalDataInMegaBytes"),
			},
			{
				Name:        "elastic_ip_status_elastic_ip",
				Description: "The elastic IP (EIP) address for the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ElasticIpStatus.ElasticIp"),
			},
			{
				Name:        "elastic_ip_status",
				Description: "The status of the elastic IP (EIP) address.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ElasticIpStatus.Status"),
			},
			{
				Name:        "elastic_resize_number_of_node_options",
				Description: "The number of nodes that you can resize the cluster to with the elastic resize method.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encrypted",
				Description: "A boolean value that, if true, indicates that data in the cluster is encrypted at rest.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "endpoint_address",
				Description: "The DNS address of the Cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Endpoint.Address"),
			},
			{
				Name:        "endpoint_port",
				Description: "The port that the database engine is listening on.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Endpoint.Port"),
			},
			{
				Name:        "enhanced_vpc_routing",
				Description: "An option that specifies whether to create the cluster with enhanced VPC routing enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "expected_next_snapshot_schedule_time",
				Description: "The date and time when the next snapshot is expected to be taken for clusters with a valid snapshot schedule and backups enabled.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "expected_next_snapshot_schedule_time_status",
				Description: "The status of next expected snapshot for clusters having a valid snapshot schedule and backups enabled.",
				Type:        schema.TypeString,
			},
			{
				Name:        "hsm_status_hsm_client_certificate_identifier",
				Description: "Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HsmStatus.HsmClientCertificateIdentifier"),
			},
			{
				Name:        "hsm_status_hsm_configuration_identifier",
				Description: "Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HsmStatus.HsmConfigurationIdentifier"),
			},
			{
				Name:        "hsm_status",
				Description: "Reports whether the Amazon Redshift cluster has finished applying any HSM settings changes specified in a modify cluster command.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HsmStatus.Status"),
			},
			{
				Name:        "kms_key_id",
				Description: "The AWS Key Management Service (AWS KMS) key ID of the encryption key used to encrypt data in the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "maintenance_track_name",
				Description: "The name of the maintenance track for the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "manual_snapshot_retention_period",
				Description: "The default number of days to retain a manual snapshot.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "master_username",
				Description: "The master user name for the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "modify_status",
				Description: "The status of a modify operation, if any, initiated for the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "next_maintenance_window_start_time",
				Description: "The date and time in UTC when system maintenance can begin.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "node_type",
				Description: "The node type for the nodes in the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "number_of_nodes",
				Description: "The number of compute nodes in the cluster.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "pending_actions",
				Description: "Cluster operations that are waiting to be started.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "pending_modified_values_automated_snapshot_retention_period",
				Description: "The pending or in-progress change of the automated snapshot retention period.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PendingModifiedValues.AutomatedSnapshotRetentionPeriod"),
			},
			{
				Name:        "pending_modified_values_cluster_identifier",
				Description: "The pending or in-progress change of the new identifier for the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.ClusterIdentifier"),
			},
			{
				Name:        "pending_modified_values_cluster_type",
				Description: "The pending or in-progress change of the cluster type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.ClusterType"),
			},
			{
				Name:        "pending_modified_values_cluster_version",
				Description: "The pending or in-progress change of the service version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.ClusterVersion"),
			},
			{
				Name:        "pending_modified_values_encryption_type",
				Description: "The encryption type for a cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.EncryptionType"),
			},
			{
				Name:        "pending_modified_values_enhanced_vpc_routing",
				Description: "An option that specifies whether to create the cluster with enhanced VPC routing enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PendingModifiedValues.EnhancedVpcRouting"),
			},
			{
				Name:        "pending_modified_values_maintenance_track_name",
				Description: "The name of the maintenance track that the cluster will change to during the next maintenance window.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.MaintenanceTrackName"),
			},
			{
				Name:        "pending_modified_values_master_user_password",
				Description: "The pending or in-progress change of the master user password for the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.MasterUserPassword"),
			},
			{
				Name:        "pending_modified_values_node_type",
				Description: "The pending or in-progress change of the cluster's node type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.NodeType"),
			},
			{
				Name:        "pending_modified_values_number_of_nodes",
				Description: "The pending or in-progress change of the number of nodes in the cluster.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PendingModifiedValues.NumberOfNodes"),
			},
			{
				Name:        "pending_modified_values_publicly_accessible",
				Description: "The pending or in-progress change of the ability to connect to the cluster from the public network.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PendingModifiedValues.PubliclyAccessible"),
			},
			{
				Name:        "preferred_maintenance_window",
				Description: "The weekly time range, in Universal Coordinated Time (UTC), during which system maintenance can occur.",
				Type:        schema.TypeString,
			},
			{
				Name:        "publicly_accessible",
				Description: "A boolean value that, if true, indicates that the cluster can be accessed from a public network.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "resize_info_allow_cancel_resize",
				Description: "A boolean value indicating if the resize operation can be cancelled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResizeInfo.AllowCancelResize"),
			},
			{
				Name:        "resize_info_resize_type",
				Description: "Returns the value ClassicResize.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResizeInfo.ResizeType"),
			},
			{
				Name:        "restore_status_current_restore_rate_in_mega_bytes_per_second",
				Description: "The number of megabytes per second being transferred from the backup storage.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("RestoreStatus.CurrentRestoreRateInMegaBytesPerSecond"),
			},
			{
				Name:        "restore_status_elapsed_time_in_seconds",
				Description: "The amount of time an in-progress restore has been running, or the amount of time it took a completed restore to finish.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RestoreStatus.ElapsedTimeInSeconds"),
			},
			{
				Name:        "restore_status_estimated_time_to_completion_in_seconds",
				Description: "The estimate of the time remaining before the restore will complete.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RestoreStatus.EstimatedTimeToCompletionInSeconds"),
			},
			{
				Name:        "restore_status_progress_in_mega_bytes",
				Description: "The number of megabytes that have been transferred from snapshot storage.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RestoreStatus.ProgressInMegaBytes"),
			},
			{
				Name:        "restore_status_snapshot_size_in_mega_bytes",
				Description: "The size of the set of snapshot data used to restore the cluster.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RestoreStatus.SnapshotSizeInMegaBytes"),
			},
			{
				Name:        "restore_status",
				Description: "The status of the restore action.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RestoreStatus.Status"),
			},
			{
				Name:        "snapshot_schedule_identifier",
				Description: "A unique identifier for the cluster snapshot schedule.",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshot_schedule_state",
				Description: "The current state of the cluster snapshot schedule.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The list of tags for the cluster.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRedshiftClusterTags,
			},
			{
				Name:        "total_storage_capacity_in_mega_bytes",
				Description: "The total storage capacity of the cluster in megabytes.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "vpc_id",
				Description: "The identifier of the VPC the cluster is in, if the cluster is in a VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "logging_status",
				Description: "Describes the status of logging for a cluster.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRedshiftClusterLoggingStatus,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_redshift_cluster_nodes",
				Description: "The identifier of a node in a cluster.",
				Resolver:    fetchRedshiftClusterNodes,
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "node_role",
						Description: "Whether the node is a leader node or a compute node.",
						Type:        schema.TypeString,
					},
					{
						Name:        "private_ip_address",
						Description: "The private IP address of a node within a cluster.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateIPAddress"),
					},
					{
						Name:        "public_ip_address",
						Description: "The public IP address of a node within a cluster.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PublicIPAddress"),
					},
				},
			},
			{
				Name:        "aws_redshift_cluster_parameter_groups",
				Description: "Describes the status of a parameter group.",
				Resolver:    fetchRedshiftClusterParameterGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_cq_id", "parameter_group_name"}},
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "parameter_apply_status",
						Description: "The status of parameter updates.",
						Type:        schema.TypeString,
					},
					{
						Name:        "parameter_group_name",
						Description: "The name of the cluster parameter group.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_redshift_cluster_parameters",
						Description: "Describes a parameter in a cluster parameter group.",
						Resolver:    fetchRedshiftClusterParameter,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_parameter_group_cq_id", "parameter_name"}},
						Columns: []schema.Column{
							{
								Name:        "cluster_parameter_group_cq_id",
								Description: "Unique CloudQuery ID of aws_redshift_cluster_parameter_groups table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "allowed_values",
								Description: "The valid range of values for the parameter.",
								Type:        schema.TypeString,
							},
							{
								Name:        "apply_type",
								Description: "Specifies how to apply the WLM configuration parameter",
								Type:        schema.TypeString,
							},
							{
								Name:        "data_type",
								Description: "The data type of the parameter.",
								Type:        schema.TypeString,
							},
							{
								Name:        "description",
								Description: "A description of the parameter.",
								Type:        schema.TypeString,
							},
							{
								Name:        "is_modifiable",
								Description: "If true, the parameter can be modified",
								Type:        schema.TypeBool,
							},
							{
								Name:        "minimum_engine_version",
								Description: "The earliest engine version to which the parameter can apply.",
								Type:        schema.TypeString,
							},
							{
								Name:        "parameter_name",
								Description: "The name of the parameter.",
								Type:        schema.TypeString,
							},
							{
								Name:        "parameter_value",
								Description: "The value of the parameter",
								Type:        schema.TypeString,
							},
							{
								Name:        "source",
								Description: "The source of the parameter value, such as \"engine-default\" or \"user\".",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_redshift_cluster_parameter_group_status_lists",
						Description: "Describes the status of a parameter group.",
						Resolver:    fetchRedshiftClusterParameterGroupStatusLists,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_parameter_group_cq_id", "parameter_name"}},
						Columns: []schema.Column{
							{
								Name:        "cluster_parameter_group_cq_id",
								Description: "Unique CloudQuery ID of aws_redshift_cluster_parameter_groups table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "parameter_apply_error_description",
								Description: "The error that prevented the parameter from being applied to the database.",
								Type:        schema.TypeString,
							},
							{
								Name:        "parameter_apply_status",
								Description: "The status of the parameter that indicates whether the parameter is in sync with the database, waiting for a cluster reboot, or encountered an error when being applied.",
								Type:        schema.TypeString,
							},
							{
								Name:        "parameter_name",
								Description: "The name of the parameter.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_redshift_cluster_security_groups",
				Description: "Describes a cluster security group.",
				Resolver:    fetchRedshiftClusterSecurityGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_cq_id", "cluster_security_group_name"}},
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cluster_security_group_name",
						Description: "The name of the cluster security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the cluster security group.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_redshift_cluster_deferred_maintenance_windows",
				Description: "Describes a deferred maintenance window .",
				Resolver:    fetchRedshiftClusterDeferredMaintenanceWindows,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_cq_id", "defer_maintenance_identifier"}},
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "defer_maintenance_end_time",
						Description: "A timestamp for the end of the time period when we defer maintenance.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "defer_maintenance_identifier",
						Description: "A unique identifier for the maintenance window.",
						Type:        schema.TypeString,
					},
					{
						Name:        "defer_maintenance_start_time",
						Description: "A timestamp for the beginning of the time period when we defer maintenance.",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "aws_redshift_cluster_endpoint_vpc_endpoints",
				Description: "The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.",
				Resolver:    fetchRedshiftClusterEndpointVpcEndpoints,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_cq_id", "vpc_endpoint_id"}},
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "vpc_endpoint_id",
						Description: "The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vpc_id",
						Description: "The VPC identifier that the endpoint is associated.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces",
						Description: "Describes a network interface.",
						Resolver:    fetchRedshiftClusterEndpointVpcEndpointNetworkInterfaces,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_endpoint_vpc_endpoint_cq_id", "network_interface_id"}},
						Columns: []schema.Column{
							{
								Name:        "cluster_endpoint_vpc_endpoint_cq_id",
								Description: "Unique CloudQuery ID of aws_redshift_cluster_endpoint_vpc_endpoints table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "availability_zone",
								Description: "The Availability Zone.",
								Type:        schema.TypeString,
							},
							{
								Name:        "network_interface_id",
								Description: "The network interface identifier.",
								Type:        schema.TypeString,
							},
							{
								Name:        "private_ip_address",
								Description: "The IPv4 address of the network interface within the subnet.",
								Type:        schema.TypeString,
							},
							{
								Name:        "subnet_id",
								Description: "The subnet identifier.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_redshift_cluster_iam_roles",
				Description: "An AWS Identity and Access Management (IAM) role that can be used by the associated Amazon Redshift cluster to access other AWS services.",
				Resolver:    fetchRedshiftClusterIamRoles,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_cq_id", "iam_role_arn"}},
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "apply_status",
						Description: "A value that describes the status of the IAM role's association with an Amazon Redshift cluster.",
						Type:        schema.TypeString,
					},
					{
						Name:        "iam_role_arn",
						Description: "The Amazon Resource Name (ARN) of the IAM role, for example, arn:aws:iam::123456789012:role/RedshiftCopyUnload.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_redshift_cluster_vpc_security_groups",
				Description: "Describes the members of a VPC security group.",
				Resolver:    fetchRedshiftClusterVpcSecurityGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cluster_cq_id", "vpc_security_group_id"}},
				Columns: []schema.Column{
					{
						Name:        "cluster_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_clusters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "status",
						Description: "The status of the VPC security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vpc_security_group_id",
						Description: "The identifier of the VPC security group.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRedshiftClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func resolveRedshiftClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected Cluster but got %T", r)
	}
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}
func resolveRedshiftClusterLoggingStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("expected Cluster but got %T", r)
	}

	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	cfg := redshift.DescribeLoggingStatusInput{
		ClusterIdentifier: r.ClusterIdentifier,
	}
	response, err := svc.DescribeLoggingStatus(ctx, &cfg, func(o *redshift.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	return resource.Set(c.Name, response)
}
func fetchRedshiftClusterNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.ClusterNodes
	return nil
}
func fetchRedshiftClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.ClusterParameterGroups
	return nil
}
func fetchRedshiftClusterParameter(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	parameterGroup, ok := parent.Item.(types.ClusterParameterGroupStatus)
	if !ok {
		return fmt.Errorf("not redshift cluster parameter group")
	}
	config := redshift.DescribeClusterParametersInput{
		ParameterGroupName: parameterGroup.ParameterGroupName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Redshift
	for {
		response, err := svc.DescribeClusterParameters(ctx, &config, func(o *redshift.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Parameters
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}

	return nil
}
func fetchRedshiftClusterParameterGroupStatusLists(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	parameterGroup, ok := parent.Item.(types.ClusterParameterGroupStatus)
	if !ok {
		return fmt.Errorf("not redshift cluster parameter group")
	}
	res <- parameterGroup.ClusterParameterStatusList
	return nil
}
func fetchRedshiftClusterSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.ClusterSecurityGroups
	return nil
}
func fetchRedshiftClusterDeferredMaintenanceWindows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.DeferredMaintenanceWindows
	return nil
}
func fetchRedshiftClusterEndpointVpcEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.Endpoint.VpcEndpoints
	return nil
}
func fetchRedshiftClusterEndpointVpcEndpointNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vpcEndpoint, ok := parent.Item.(types.VpcEndpoint)
	if !ok {
		return fmt.Errorf("not vpc endpoint")
	}
	res <- vpcEndpoint.NetworkInterfaces
	return nil
}
func fetchRedshiftClusterIamRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.IamRoles
	return nil
}
func fetchRedshiftClusterVpcSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cluster, ok := parent.Item.(types.Cluster)
	if !ok {
		return fmt.Errorf("not redshift cluster")
	}
	res <- cluster.VpcSecurityGroups
	return nil
}
