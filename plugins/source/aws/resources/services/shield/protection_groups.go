package shield

import (
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ProtectionGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_shield_protection_groups",
		Description: `https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_ProtectionGroup.html`,
		Resolver:    fetchShieldProtectionGroups,
		Transform:   transformers.TransformWithStruct(&types.ProtectionGroup{}),
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProtectionGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveShieldProtectionGroupTags,
			},
		},
	}
}
