// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Disks() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_disks",
		Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Disk.html",
		Resolver:    fetchLightsailDisks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "add_ons",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AddOns"),
			},
			{
				Name:     "attached_to",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AttachedTo"),
			},
			{
				Name:     "attachment_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AttachmentState"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "gb_in_use",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("GbInUse"),
			},
			{
				Name:     "iops",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Iops"),
			},
			{
				Name:     "is_attached",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsAttached"),
			},
			{
				Name:     "is_system_disk",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsSystemDisk"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Path"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "size_in_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeInGb"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
		},

		Relations: []*schema.Table{
			DiskSnapshot(),
		},
	}
}
