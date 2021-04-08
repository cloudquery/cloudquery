package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2Images() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_images",
		Resolver:     fetchEc2Images,
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
				Name: "architecture",
				Type: schema.TypeString,
			},
			{
				Name: "creation_date",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "ena_support",
				Type: schema.TypeBool,
			},
			{
				Name: "hypervisor",
				Type: schema.TypeString,
			},
			{
				Name: "image_id",
				Type: schema.TypeString,
			},
			{
				Name: "image_location",
				Type: schema.TypeString,
			},
			{
				Name: "image_owner_alias",
				Type: schema.TypeString,
			},
			{
				Name: "image_type",
				Type: schema.TypeString,
			},
			{
				Name: "kernel_id",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name: "platform",
				Type: schema.TypeString,
			},
			{
				Name: "platform_details",
				Type: schema.TypeString,
			},
			{
				Name:     "product_codes",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2imageProductCodes,
			},
			{
				Name: "public",
				Type: schema.TypeBool,
			},
			{
				Name: "ramdisk_id",
				Type: schema.TypeString,
			},
			{
				Name: "root_device_name",
				Type: schema.TypeString,
			},
			{
				Name: "root_device_type",
				Type: schema.TypeString,
			},
			{
				Name: "sriov_net_support",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name:     "state_reason_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateReason.Code"),
			},
			{
				Name:     "state_reason_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateReason.Message"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2imageTags,
			},
			{
				Name: "usage_operation",
				Type: schema.TypeString,
			},
			{
				Name: "virtualization_type",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_image_block_device_mappings",
				Resolver: fetchEc2ImageBlockDeviceMappings,
				Columns: []schema.Column{
					{
						Name:     "image_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "device_name",
						Type: schema.TypeString,
					},
					{
						Name:     "ebs_delete_on_termination",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Ebs.DeleteOnTermination"),
					},
					{
						Name:     "ebs_encrypted",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Ebs.Encrypted"),
					},
					{
						Name:     "ebs_iops",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Ebs.Iops"),
					},
					{
						Name:     "ebs_kms_key_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ebs.KmsKeyId"),
					},
					{
						Name:     "ebs_outpost_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ebs.OutpostArn"),
					},
					{
						Name:     "ebs_snapshot_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ebs.SnapshotId"),
					},
					{
						Name:     "ebs_throughput",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Ebs.Throughput"),
					},
					{
						Name:     "ebs_volume_size",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Ebs.VolumeSize"),
					},
					{
						Name:     "ebs_volume_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ebs.VolumeType"),
					},
					{
						Name: "no_device",
						Type: schema.TypeString,
					},
					{
						Name: "virtual_name",
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
func fetchEc2Images(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	svc := c.Services().EC2
	response, err := svc.DescribeImages(ctx, &ec2.DescribeImagesInput{Owners: []string{"self"}}, func(options *ec2.Options) {
		options.Region = c.Region
		options.EndpointResolver = ec2.EndpointResolverFromURL(fmt.Sprintf("https://ec2.%s.amazonaws.com", c.Region))
	})
	if err != nil {
		return err
	}
	res <- response.Images
	return nil
}
func resolveEc2imageProductCodes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Image)
	productCodes := map[string]string{}
	for _, t := range r.ProductCodes {
		productCodes[*t.ProductCodeId] = string(t.ProductCodeType)
	}
	resource.Set("product_codes", productCodes)
	return nil
}
func resolveEc2imageTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Image)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
func fetchEc2ImageBlockDeviceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.Image)
	res <- r.BlockDeviceMappings
	return nil
}
