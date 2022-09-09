package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func InstanceSnapshots() *schema.Table {
	return &schema.Table{
		Name:          "aws_lightsail_instance_snapshots",
		Description:   "Describes an instance snapshot",
		Resolver:      fetchLightsailInstanceSnapshots,
		Multiplex:     client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the snapshot (eg, arn:aws:lightsail:us-east-2:123456789101:InstanceSnapshot/d23b5706-3322-4d83-81e5-12345EXAMPLE)",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the snapshot was created (eg, 1479907467024)",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "from_blueprint_id",
				Description: "The blueprint ID from which you created the snapshot (eg, os_debian_8_3)",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_bundle_id",
				Description: "The bundle ID from which you created the snapshot (eg, micro_1_0)",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_instance_arn",
				Description: "The Amazon Resource Name (ARN) of the instance from which the snapshot was created (eg, arn:aws:lightsail:us-east-2:123456789101:Instance/64b8404c-ccb1-430b-8daf-12345EXAMPLE)",
				Type:        schema.TypeString,
			},
			{
				Name:        "from_instance_name",
				Description: "The instance from which the snapshot was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_from_auto_snapshot",
				Description: "A Boolean value indicating whether the snapshot was created from an automatic snapshot",
				Type:        schema.TypeBool,
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:        "name",
				Description: "The name of the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "progress",
				Description: "The progress of the snapshot",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "The type of resource (usually InstanceSnapshot)",
				Type:        schema.TypeString,
			},
			{
				Name:        "size_in_gb",
				Description: "The size in GB of the SSD",
				Type:        schema.TypeInt,
			},
			{
				Name:        "state",
				Description: "The state the snapshot is in",
				Type:        schema.TypeString,
			},
			{
				Name:        "support_code",
				Description: "The support code",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name: "from_attached_disks",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailInstanceSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetInstanceSnapshotsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetInstanceSnapshots(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.InstanceSnapshots
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
