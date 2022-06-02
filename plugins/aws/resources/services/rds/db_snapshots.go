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

func RdsDbSnapshots() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_db_snapshots",
		Description:  "Contains the details of an Amazon RDS DB snapshot",
		Resolver:     fetchRdsDbSnapshots,
		Multiplex:    client.ServiceAccountRegionMultiplexer("rds"),
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
				Name:        "allocated_storage",
				Description: "Specifies the allocated storage size in gibibytes (GiB).",
				Type:        schema.TypeInt,
			},
			{
				Name:        "availability_zone",
				Description: "Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "db_instance_identifier",
				Description: "Specifies the DB instance identifier of the DB instance this DB snapshot was created from.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBInstanceIdentifier"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the DB snapshot.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSnapshotArn"),
			},
			{
				Name:        "db_snapshot_identifier",
				Description: "Specifies the identifier for the DB snapshot.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSnapshotIdentifier"),
			},
			{
				Name:        "dbi_resource_id",
				Description: "The identifier for the source DB instance, which can't be changed and which is unique to an AWS Region.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encrypted",
				Description: "Specifies whether the DB snapshot is encrypted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "engine",
				Description: "Specifies the name of the database engine.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "Specifies the version of the database engine.",
				Type:        schema.TypeString,
			},
			{
				Name:        "iam_database_authentication_enabled",
				Description: "True if mapping of AWS Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:        "instance_create_time",
				Description: "Specifies the time in Coordinated Universal Time (UTC) when the DB instance, from which the snapshot was taken, was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "iops",
				Description:   "Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:          "kms_key_id",
				Description:   "If Encrypted is true, the AWS KMS key identifier for the encrypted DB snapshot. The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the AWS KMS customer master key (CMK).",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "license_model",
				Description: "License model information for the restored DB instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "master_username",
				Description: "Provides the master username for the DB snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "option_group_name",
				Description: "Provides the option group name for the DB snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "percent_progress",
				Description: "The percentage of the estimated data that has been transferred.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "port",
				Description: "Specifies the port that the database engine was listening on at the time of the snapshot.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "processor_features",
				Description: "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance when the DB snapshot was created.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRDSDBSnapshotProcessorFeatures,
			},
			{
				Name:        "snapshot_create_time",
				Description: "Specifies when the snapshot was taken in Coordinated Universal Time (UTC).",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "snapshot_type",
				Description: "Provides the type of the DB snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:          "source_db_snapshot_identifier",
				Description:   "The DB snapshot Amazon Resource Name (ARN) that the DB snapshot was copied from. It only has value in case of cross-customer or cross-region copy.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("SourceDBSnapshotIdentifier"),
				IgnoreInTests: true,
			},
			{
				Name:          "source_region",
				Description:   "The AWS Region that the DB snapshot was created in or copied from.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "status",
				Description: "Specifies the status of this DB snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "storage_type",
				Description: "Specifies the storage type associated with DB snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:          "tde_credential_arn",
				Description:   "The ARN from the key store with which to associate the instance for TDE encryption.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "timezone",
				Description:   "The time zone of the DB snapshot",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "vpc_id",
				Description: "Provides the VPC ID associated with the DB snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRDSDBSnapshotTags,
			},
			{
				Name:        "attributes",
				Description: "Snapshot attribute names and values",
				Type:        schema.TypeJSON,
				Resolver:    resolveRDSDBSnapshotAttributes,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsDbSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().RDS
	var input rds.DescribeDBSnapshotsInput
	for {
		output, err := svc.DescribeDBSnapshots(ctx, &input, func(o *rds.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return nil
		}
		res <- output.DBSnapshots
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveRDSDBSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.DBSnapshot)
	tags := map[string]*string{}
	for _, t := range s.TagList {
		tags[*t.Key] = t.Value
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}

func resolveRDSDBSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBSnapshot)
	c := meta.(*client.Client)
	svc := c.Services().RDS
	out, err := svc.DescribeDBSnapshotAttributes(
		ctx,
		&rds.DescribeDBSnapshotAttributesInput{DBSnapshotIdentifier: s.DBSnapshotIdentifier},
		func(o *rds.Options) {
			o.Region = c.Region
		},
	)
	if err != nil {
		return diag.WrapError(err)
	}
	if out.DBSnapshotAttributesResult == nil {
		return nil
	}

	b, err := json.Marshal(out.DBSnapshotAttributesResult.DBSnapshotAttributes)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(column.Name, b))
}

func resolveRDSDBSnapshotProcessorFeatures(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBSnapshot)

	b, err := json.Marshal(s.ProcessorFeatures)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(column.Name, b))
}
