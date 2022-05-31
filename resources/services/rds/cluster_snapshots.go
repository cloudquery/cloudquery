package rds

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsClusterSnapshots() *schema.Table {
	return &schema.Table{
		Name:          "aws_rds_cluster_snapshots",
		Description:   "Contains the details for an Amazon RDS DB cluster snapshot This data type is used as a response element in the DescribeDBClusterSnapshots action.",
		Resolver:      fetchRdsClusterSnapshots,
		Multiplex:     client.ServiceAccountRegionMultiplexer("rds"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size in gibibytes (GiB).",
				Type:        schema.TypeInt,
			},
			{
				Name:        "availability_zones",
				Description: "Provides the list of Availability Zones (AZs) where instances in the DB cluster snapshot can be restored.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "cluster_create_time",
				Description: "Specifies the time when the DB cluster was created, in Universal Coordinated Time (UTC).",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "db_cluster_identifier",
				Description: "Specifies the DB cluster identifier of the DB cluster that this DB cluster snapshot was created from.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB cluster snapshot.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterSnapshotArn"),
			},
			{
				Name:        "db_cluster_snapshot_identifier",
				Description: "Specifies the identifier for the DB cluster snapshot.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterSnapshotIdentifier"),
			},
			{
				Name:        "engine",
				Description: "Specifies the name of the database engine for this DB cluster snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_mode",
				Description: "Provides the engine mode of the database engine for this DB cluster snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "Provides the version of the database engine for this DB cluster snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "True if mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:        "kms_key_id",
				Description: "If StorageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "license_model",
				Description: "Provides the license model information for this DB cluster snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "master_username",
				Description: "Provides the master username for this DB cluster snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "percent_progress",
				Description: "Specifies the percentage of the estimated data that has been transferred.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "port",
				Description: "Specifies the port that the DB cluster was listening on at the time of the snapshot.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "snapshot_create_time",
				Description: "Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "snapshot_type",
				Description: "Provides the type of the DB cluster snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_db_cluster_snapshot_arn",
				Description: "If the DB cluster snapshot was copied from a source DB cluster snapshot, the Amazon Resource Name (ARN) for the source DB cluster snapshot, otherwise, a null value.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceDBClusterSnapshotArn"),
			},
			{
				Name:        "status",
				Description: "Specifies the status of this DB cluster snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "storage_encrypted",
				Description: "Specifies whether the DB cluster snapshot is encrypted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "vpc_id",
				Description: "Provides the VPC ID associated with the DB cluster snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRDSClusterSnapshotTags,
			},
			{
				Name:        "attributes",
				Description: "Snapshot attribute names and values",
				Type:        schema.TypeJSON,
				Resolver:    resolveRDSClusterSnapshotAttributes,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsClusterSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().RDS
	var input rds.DescribeDBClusterSnapshotsInput
	for {
		output, err := svc.DescribeDBClusterSnapshots(ctx, &input, func(o *rds.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return nil
		}
		res <- output.DBClusterSnapshots
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveRDSClusterSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	tags := map[string]*string{}
	for _, t := range s.TagList {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}

func resolveRDSClusterSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	c := meta.(*client.Client)
	svc := c.Services().RDS
	out, err := svc.DescribeDBClusterSnapshotAttributes(
		ctx,
		&rds.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		func(o *rds.Options) {
			o.Region = c.Region
		},
	)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	if out.DBClusterSnapshotAttributesResult == nil {
		return nil
	}

	b, err := json.Marshal(out.DBClusterSnapshotAttributesResult.DBClusterSnapshotAttributes)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(column.Name, b)
}
