package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"golang.org/x/sync/errgroup"
)

func Ec2Images() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_images",
		Description:   "Describes an image.",
		Resolver:      fetchEc2Images,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNWithRegion(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"image", *resource.Item.(types.Image).ImageId}, nil
				}),
			},
			{
				Name:        "id",
				Description: "The ID of the AMI.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ImageId"),
			},
			{
				Name:        "architecture",
				Description: "The architecture of the image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_date",
				Description: "The date and time the image was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    client.ResolveTimestampField("CreationDate"),
			},
			{
				Name:        "description",
				Description: "The description of the AMI that was provided during image creation.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ena_support",
				Description: "Specifies whether enhanced networking with ENA is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "hypervisor",
				Description: "The hypervisor type of the image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "image_location",
				Description: "The location of the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "image_owner_alias",
				Description: "The AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "image_type",
				Description: "The type of image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "kernel_id",
				Description: "The kernel associated with the image, if any.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the AMI that was provided during image creation.",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner_id",
				Description: "The AWS account ID of the image owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform",
				Description: "This value is set to windows for Windows AMIs; otherwise, it is blank.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform_details",
				Description: "The platform details associated with the billing code of the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "product_codes",
				Description: "Any product codes associated with the AMI.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2imageProductCodes,
			},
			{
				Name:        "public",
				Description: "Indicates whether the image has public launch permissions.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "ramdisk_id",
				Description: "The RAM disk associated with the image, if any.",
				Type:        schema.TypeString,
			},
			{
				Name:        "root_device_name",
				Description: "The device name of the root device volume (for example, /dev/sda1).",
				Type:        schema.TypeString,
			},
			{
				Name:        "root_device_type",
				Description: "The type of root device used by the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "sriov_net_support",
				Description: "Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_reason_code",
				Description: "The reason code for the state change.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StateReason.Code"),
			},
			{
				Name:        "state_reason_message",
				Description: "The message for the state change.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StateReason.Message"),
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the image.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2imageTags,
			},
			{
				Name:        "usage_operation",
				Description: "The operation of the Amazon EC2 instance and the billing code that is associated with the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "virtualization_type",
				Description: "The type of virtualization of the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_launched_time",
				Description: "The timestamp of the last time the AMI was used for an EC2 instance launch.",
				Type:        schema.TypeTimestamp,
				Resolver:    resolveEc2ImageLastLaunchedTime,
			},
			{
				Name:        "deprecation_time",
				Description: "The date and time to deprecate the AMI.",
				Type:        schema.TypeTimestamp,
				Resolver:    client.ResolveTimestampField("DeprecationTime"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_ec2_image_block_device_mappings",
				Description:   "Describes a block device mapping.",
				Resolver:      fetchEc2ImageBlockDeviceMappings,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "image_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_images table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "device_name",
						Description: "The device name (for example, /dev/sdh or xvdh).",
						Type:        schema.TypeString,
					},
					{
						Name:        "ebs_delete_on_termination",
						Description: "Indicates whether the EBS volume is deleted on instance termination.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Ebs.DeleteOnTermination"),
					},
					{
						Name:        "ebs_encrypted",
						Description: "Indicates whether the encryption state of an EBS volume is changed while being restored from a backing snapshot.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Ebs.Encrypted"),
					},
					{
						Name:        "ebs_iops",
						Description: "The number of I/O operations per second (IOPS).",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Ebs.Iops"),
					},
					{
						Name:        "ebs_kms_key_id",
						Description: "Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed CMK under which the EBS volume is encrypted.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.KmsKeyId"),
					},
					{
						Name:        "ebs_outpost_arn",
						Description: "The ARN of the Outpost on which the snapshot is stored.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.OutpostArn"),
					},
					{
						Name:        "ebs_snapshot_id",
						Description: "The ID of the snapshot.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.SnapshotId"),
					},
					{
						Name:        "ebs_throughput",
						Description: "The throughput that the volume supports, in MiB/s.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Ebs.Throughput"),
					},
					{
						Name:        "ebs_volume_size",
						Description: "The size of the volume, in GiBs.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Ebs.VolumeSize"),
					},
					{
						Name:        "ebs_volume_type",
						Description: "The volume type.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.VolumeType"),
					},
					{
						Name:        "no_device",
						Description: "To omit the device from the block device mapping, specify an empty string.",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_name",
						Description: "The virtual device name (ephemeralN).",
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
func fetchEc2Images(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	svc := c.Services().EC2
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// fetch ec2.Images owned by this account
		response, err := svc.DescribeImages(ctx, &ec2.DescribeImagesInput{Owners: []string{"self"}}, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Images
		return nil
	})

	g.Go(func() error {
		// fetch ec2.Images that are shared with this account
		response, err := svc.DescribeImages(ctx, &ec2.DescribeImagesInput{ExecutableUsers: []string{"self"}}, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Images
		return nil
	})

	if err := g.Wait(); err != nil {
		return diag.WrapError(err)
	}
	return nil
}

func resolveEc2imageProductCodes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Image)
	productCodes := map[string]string{}
	for _, t := range r.ProductCodes {
		productCodes[*t.ProductCodeId] = string(t.ProductCodeType)
	}
	return diag.WrapError(resource.Set("product_codes", productCodes))
}
func resolveEc2imageTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Image)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return diag.WrapError(resource.Set("tags", tags))
}
func fetchEc2ImageBlockDeviceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Image)
	res <- r.BlockDeviceMappings
	return nil
}

func resolveEc2ImageLastLaunchedTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	image := resource.Item.(types.Image)
	svc := cl.Services().EC2
	if *image.OwnerId != cl.AccountID {
		return nil
	}
	opts := ec2.DescribeImageAttributeInput{
		Attribute: types.ImageAttributeNameLastLaunchedTime,
		ImageId:   image.ImageId,
	}
	result, err := svc.DescribeImageAttribute(ctx, &opts, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	if result.LastLaunchedTime == nil || result.LastLaunchedTime.Value == nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339, *result.LastLaunchedTime.Value)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, t))
}
