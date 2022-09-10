package redshift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RedshiftClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_redshift_clusters",
		Description:  "Describes a cluster.",
		Resolver:     fetchRedshiftClusters,
		Multiplex:    client.ServiceAccountRegionMultiplexer("redshift"),
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
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.RedshiftService, func(resource *schema.Resource) ([]string, error) {
					return []string{fmt.Sprintf("cluster:%s", *resource.Item.(types.Cluster).ClusterIdentifier)}, nil
				}),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Name:          "cluster_snapshot_copy_status_destination_region",
				Description:   "The destination region that snapshots are automatically copied to when cross-region snapshot copy is enabled.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ClusterSnapshotCopyStatus.DestinationRegion"),
				IgnoreInTests: true,
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
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ClusterSnapshotCopyStatus.RetentionPeriod"),
			},
			{
				Name:          "cluster_snapshot_copy_status_snapshot_copy_grant_name",
				Description:   "The name of the snapshot copy grant.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ClusterSnapshotCopyStatus.SnapshotCopyGrantName"),
				IgnoreInTests: true,
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
				Name: "data_transfer_progress",
				Type: schema.TypeJSON,
			},
			{
				Name:          "elastic_ip_status_elastic_ip",
				Description:   "The elastic IP (EIP) address for the cluster.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ElasticIpStatus.ElasticIp"),
				IgnoreInTests: true,
			},
			{
				Name:          "elastic_ip_status",
				Description:   "The status of the elastic IP (EIP) address.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ElasticIpStatus.Status"),
				IgnoreInTests: true,
			},
			{
				Name:          "elastic_resize_number_of_node_options",
				Description:   "The number of nodes that you can resize the cluster to with the elastic resize method.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
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
				Name:          "expected_next_snapshot_schedule_time",
				Description:   "The date and time when the next snapshot is expected to be taken for clusters with a valid snapshot schedule and backups enabled.",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:          "expected_next_snapshot_schedule_time_status",
				Description:   "The status of next expected snapshot for clusters having a valid snapshot schedule and backups enabled.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "hsm_status_hsm_client_certificate_identifier",
				Description:   "Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("HsmStatus.HsmClientCertificateIdentifier"),
				IgnoreInTests: true,
			},
			{
				Name:          "hsm_status_hsm_configuration_identifier",
				Description:   "Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("HsmStatus.HsmConfigurationIdentifier"),
				IgnoreInTests: true,
			},
			{
				Name:          "hsm_status",
				Description:   "Reports whether the Amazon Redshift cluster has finished applying any HSM settings changes specified in a modify cluster command.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("HsmStatus.Status"),
				IgnoreInTests: true,
			},
			{
				Name:          "kms_key_id",
				Description:   "The AWS Key Management Service (AWS KMS) key ID of the encryption key used to encrypt data in the cluster.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
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
				Name:          "modify_status",
				Description:   "The status of a modify operation, if any, initiated for the cluster.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
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
				Name:          "pending_actions",
				Description:   "Cluster operations that are waiting to be started.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_automated_snapshot_retention_period",
				Description:   "The pending or in-progress change of the automated snapshot retention period.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("PendingModifiedValues.AutomatedSnapshotRetentionPeriod"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_cluster_identifier",
				Description:   "The pending or in-progress change of the new identifier for the cluster.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PendingModifiedValues.ClusterIdentifier"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_cluster_type",
				Description:   "The pending or in-progress change of the cluster type.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PendingModifiedValues.ClusterType"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_cluster_version",
				Description:   "The pending or in-progress change of the service version.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PendingModifiedValues.ClusterVersion"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_encryption_type",
				Description:   "The encryption type for a cluster.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PendingModifiedValues.EncryptionType"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_enhanced_vpc_routing",
				Description:   "An option that specifies whether to create the cluster with enhanced VPC routing enabled.",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("PendingModifiedValues.EnhancedVpcRouting"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_maintenance_track_name",
				Description:   "The name of the maintenance track that the cluster will change to during the next maintenance window.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PendingModifiedValues.MaintenanceTrackName"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_master_user_password",
				Description:   "The pending or in-progress change of the master user password for the cluster.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PendingModifiedValues.MasterUserPassword"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_node_type",
				Description:   "The pending or in-progress change of the cluster's node type.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("PendingModifiedValues.NodeType"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_number_of_nodes",
				Description:   "The pending or in-progress change of the number of nodes in the cluster.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("PendingModifiedValues.NumberOfNodes"),
				IgnoreInTests: true,
			},
			{
				Name:          "pending_modified_values_publicly_accessible",
				Description:   "The pending or in-progress change of the ability to connect to the cluster from the public network.",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("PendingModifiedValues.PubliclyAccessible"),
				IgnoreInTests: true,
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
				Name:          "resize_info_resize_type",
				Description:   "Returns the value ClassicResize.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ResizeInfo.ResizeType"),
				IgnoreInTests: true,
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
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("RestoreStatus.ElapsedTimeInSeconds"),
			},
			{
				Name:        "restore_status_estimated_time_to_completion_in_seconds",
				Description: "The estimate of the time remaining before the restore will complete.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("RestoreStatus.EstimatedTimeToCompletionInSeconds"),
			},
			{
				Name:        "restore_status_progress_in_mega_bytes",
				Description: "The number of megabytes that have been transferred from snapshot storage.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("RestoreStatus.ProgressInMegaBytes"),
			},
			{
				Name:        "restore_status_snapshot_size_in_mega_bytes",
				Description: "The size of the set of snapshot data used to restore the cluster.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("RestoreStatus.SnapshotSizeInMegaBytes"),
			},
			{
				Name:          "restore_status",
				Description:   "The status of the restore action.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("RestoreStatus.Status"),
				IgnoreInTests: true,
			},
			{
				Name:          "snapshot_schedule_identifier",
				Description:   "A unique identifier for the cluster snapshot schedule.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
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
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "total_storage_capacity_in_mega_bytes",
				Description: "The total storage capacity of the cluster in megabytes.",
				Type:        schema.TypeInt,
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
			{
				Name:        "cluster_nodes",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "cluster_parameter_groups",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "cluster_security_groups",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "deferred_maintenance_windows",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "endpoint",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "iam_roles",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "vpc_security_groups",
				Type: 			schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			Snapshots(),
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchRedshiftClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config redshift.DescribeClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Redshift
	for {
		response, err := svc.DescribeClusters(ctx, &config)
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

func resolveRedshiftClusterLoggingStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Cluster)

	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	cfg := redshift.DescribeLoggingStatusInput{
		ClusterIdentifier: r.ClusterIdentifier,
	}
	response, err := svc.DescribeLoggingStatus(ctx, &cfg)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, response)
}

