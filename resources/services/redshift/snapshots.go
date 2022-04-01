package redshift

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:          "aws_redshift_snapshots",
		Description:   "Describes a snapshot.",
		Resolver:      fetchRedshiftSnapshots,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"cluster_identifier", "cluster_create_time"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "actual_incremental_backup_size",
				Description: "The size of the incremental backup in megabytes.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("ActualIncrementalBackupSizeInMegaBytes"),
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone in which the cluster was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "backup_progress",
				Description: "The number of megabytes that have been transferred to the snapshot backup.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("BackupProgressInMegaBytes"),
			},
			{
				Name:        "cluster_create_time",
				Description: "The time (UTC) when the cluster was originally created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "cluster_identifier",
				Description: "The identifier of the cluster for which the snapshot was taken.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cluster_version",
				Description: "The version ID of the Amazon Redshift engine that is running on the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "current_backup_rate",
				Description: "The number of megabytes per second being transferred to the snapshot backup. Returns 0 for a completed backup.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("CurrentBackupRateInMegaBytesPerSecond"),
			},
			{
				Name:        "db_name",
				Description: "The name of the database that was created when the cluster was created.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBName"),
			},
			{
				Name:        "elapsed_time",
				Description: "The amount of time an in-progress snapshot backup has been running, or the amount of time it took a completed backup to finish, in seconds.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ElapsedTimeInSeconds"),
			},
			{
				Name:        "encrypted",
				Description: "If true, the data in the snapshot is encrypted at rest.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "encrypted_with_hsm",
				Description: "A boolean that indicates whether the snapshot data is encrypted using the HSM keys of the source cluster.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EncryptedWithHSM"),
			},
			{
				Name:        "engine_full_version",
				Description: "The cluster version of the cluster used to create the snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "enhanced_vpc_routing",
				Description: "An option that specifies whether to create the cluster with enhanced VPC routing enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "estimated_seconds_to_completion",
				Description: "The estimate of the time remaining before the snapshot backup will complete. Returns 0 for a completed backup.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "kms_key_id",
				Description: "The AWS Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.",
				Type:        schema.TypeString,
			},
			{
				Name:        "maintenance_track_name",
				Description: "The name of the maintenance track for the snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "manual_snapshot_remaining_days",
				Description: "The number of days until a manual snapshot will pass its retention period.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "manual_snapshot_retention_period",
				Description: "The number of days that a manual snapshot is retained",
				Type:        schema.TypeInt,
			},
			{
				Name:        "master_username",
				Description: "The master user name for the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "node_type",
				Description: "The node type of the nodes in the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "number_of_nodes",
				Description: "The number of nodes in the cluster.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "owner_account",
				Description: "For manual snapshots, the AWS customer account used to create or copy the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "port",
				Description: "The port that the cluster is listening on.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "restorable_node_types",
				Description: "The list of node types that this cluster snapshot is able to restore into.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "snapshot_create_time",
				Description: "The time (in UTC format) when Amazon Redshift began the snapshot",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "snapshot_identifier",
				Description: "The snapshot identifier that is provided in the request.",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshot_retention_start_time",
				Description: "A timestamp representing the start of the retention period for the snapshot.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "snapshot_type",
				Description: "The snapshot type",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_region",
				Description: "The source region from which the snapshot was copied.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The snapshot status.",
				Type:        schema.TypeString,
			},
			{
				Name:        "total_backup_size_in_mega_bytes",
				Description: "The size of the complete set of backup data that would be used to restore the cluster.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "vpc_id",
				Description: "The VPC identifier of the cluster if the snapshot is from a cluster in a VPC. Otherwise, this field is not in the output.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Tags consisting of a name/value pair for a resource.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSnapshotTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_redshift_snapshot_accounts_with_restore_access",
				Description: "Describes an AWS customer account authorized to restore a snapshot.",
				Resolver:    fetchRedshiftSnapshotAccountsWithRestoreAccesses,
				Columns: []schema.Column{
					{
						Name:        "snapshot_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_snapshots table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_alias",
						Description: "The identifier of an AWS support account authorized to restore a snapshot",
						Type:        schema.TypeString,
					},
					{
						Name:        "account_id",
						Description: "The identifier of an AWS customer account authorized to restore a snapshot.",
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

func fetchRedshiftSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift
	cluster := parent.Item.(types.Cluster)
	params := redshift.DescribeClusterSnapshotsInput{
		ClusterExists:     aws.Bool(true),
		ClusterIdentifier: cluster.ClusterIdentifier,
		MaxRecords:        aws.Int32(100),
	}
	for {
		result, err := svc.DescribeClusterSnapshots(ctx, &params, func(o *redshift.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.Snapshots
		if aws.ToString(result.Marker) == "" {
			break
		}
		params.Marker = result.Marker
	}
	return nil
}

func fetchRedshiftSnapshotAccountsWithRestoreAccesses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	s := parent.Item.(types.Snapshot)
	res <- s.AccountsWithRestoreAccess
	return nil
}

func resolveSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.Snapshot)
	tags := make(map[string]string, len(s.Tags))
	for _, v := range s.Tags {
		tags[aws.ToString(v.Key)] = aws.ToString(v.Value)
	}
	b, err := json.Marshal(tags)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
