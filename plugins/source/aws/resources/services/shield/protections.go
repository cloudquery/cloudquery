package shield

import (
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Protections() *schema.Table {
	return &schema.Table{
		Name:        "aws_shield_protections",
		Description: `https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Protection.html`,
		Resolver:    fetchShieldProtections,
		Transform:   transformers.TransformWithStruct(&types.Protection{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("shield"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProtectionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveShieldProtectionTags,
			},
		},
	}
}
