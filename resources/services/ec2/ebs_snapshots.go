package ec2

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2EbsSnapshots() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_ebs_snapshots",
		Description:   "Describes a snapshot.",
		Resolver:      fetchEc2EbsSnapshots,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "snapshot_id"}},
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
				Name:     "create_volume_permissions",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2ebsSnapshotCreateVolumePermissions,
			},
			{
				Name:        "data_encryption_key_id",
				Description: "The data encryption key identifier for the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "The description for the snapshot.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encrypted",
				Description: "Indicates whether the snapshot is encrypted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the volume encryption key for the parent volume.",
				Type:        schema.TypeString,
			},
			{
				Name:        "outpost_arn",
				Description: "The ARN of the AWS Outpost on which the snapshot is stored",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner_alias",
				Description: "The AWS owner alias, from an Amazon-maintained list (amazon)",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner_id",
				Description: "The AWS account ID of the EBS snapshot owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "progress",
				Description: "The progress of the snapshot, as a percentage.",
				Type:        schema.TypeString,
			},
			{
				Name:        "snapshot_id",
				Description: "The ID of the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "start_time",
				Description: "The time stamp when the snapshot was initiated.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "state",
				Description: "The snapshot state.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_message",
				Description: "Encrypted Amazon EBS snapshots are copied asynchronously",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the snapshot.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2ebsSnapshotTags,
			},
			{
				Name:        "volume_id",
				Description: "The ID of the volume that was used to create the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "volume_size",
				Description: "The size of the volume, in GiB.",
				Type:        schema.TypeInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2EbsSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeSnapshotsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	config.OwnerIds = []string{c.AccountID}
	for {
		output, err := svc.DescribeSnapshots(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Snapshots
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveEc2ebsSnapshotCreateVolumePermissions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Snapshot)
	cl := meta.(*client.Client)
	svc := cl.Services().EC2
	output, err := svc.DescribeSnapshotAttribute(ctx, &ec2.DescribeSnapshotAttributeInput{
		Attribute:  types.SnapshotAttributeNameCreateVolumePermission,
		SnapshotId: r.SnapshotId,
	}, func(options *ec2.Options) {
		options.Region = cl.Region
	})

	if err != nil {
		if client.IsAWSError(err, "InvalidSnapshot.NotFound") {
			meta.Logger().Debug("Failed extracting snapshot volume permissions", "SnapshotId", r.SnapshotId, "error", err)
			return nil
		}
		return diag.WrapError(err)
	}

	createVolumePermissions := make([]map[string]string, len(output.CreateVolumePermissions))
	for i, p := range output.CreateVolumePermissions {
		createVolumePermissions[i] = map[string]string{}
		createVolumePermissions[i]["group"] = string(p.Group)
		if p.UserId != nil {
			createVolumePermissions[i]["user_id"] = *p.UserId
		} else {
			createVolumePermissions[i]["user_id"] = ""
		}
	}
	b, err := json.Marshal(createVolumePermissions)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set("create_volume_permissions", b))
}
func resolveEc2ebsSnapshotTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Snapshot)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return diag.WrapError(resource.Set("tags", tags))
}
