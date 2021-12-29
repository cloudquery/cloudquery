package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2EbsVolumes() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_ebs_volumes",
		Resolver:     fetchEc2EbsVolumes,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VolumeId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the ebs volume",
				Type:        schema.TypeString,
				Resolver:    resolveEbsVolumeArn,
			},
			{
				Name: "availability_zone",
				Type: schema.TypeString,
			},
			{
				Name: "create_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "encrypted",
				Type: schema.TypeBool,
			},
			{
				Name: "fast_restored",
				Type: schema.TypeBool,
			},
			{
				Name: "iops",
				Type: schema.TypeInt,
			},
			{
				Name: "kms_key_id",
				Type: schema.TypeString,
			},
			{
				Name: "multi_attach_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "outpost_arn",
				Type: schema.TypeString,
			},
			{
				Name: "size",
				Type: schema.TypeInt,
			},
			{
				Name: "snapshot_id",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2EbsVolumeTags,
			},
			{
				Name: "throughput",
				Type: schema.TypeInt,
			},
			{
				Name: "volume_type",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_ebs_volume_attachments",
				Resolver: fetchEc2EbsVolumeAttachments,
				Options:  schema.TableCreationOptions{PrimaryKeys: []string{"ebs_volume_cq_id", "device"}},
				Columns: []schema.Column{
					{
						Name:     "ebs_volume_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "attach_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "delete_on_termination",
						Type: schema.TypeBool,
					},
					{
						Name: "device",
						Type: schema.TypeString,
					},
					{
						Name: "instance_id",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "volume_id",
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
func fetchEc2EbsVolumes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2
	config := ec2.DescribeVolumesInput{}
	for {
		response, err := svc.DescribeVolumes(ctx, &config, func(o *ec2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Volumes
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveEc2EbsVolumeTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Volume)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2EbsVolumeAttachments(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	volume, ok := parent.Item.(types.Volume)
	if !ok {
		return fmt.Errorf("not ec2 ebs volume")
	}
	res <- volume.Attachments
	return nil
}

func resolveEbsVolumeArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ebs, ok := resource.Item.(types.Volume)
	if !ok {
		return fmt.Errorf("not ec2 ebs volume")
	}
	return resource.Set(c.Name, client.GenerateResourceARN("ec2", "volume", *ebs.VolumeId, cl.Region, cl.AccountID))
}
