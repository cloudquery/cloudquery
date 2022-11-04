// Code generated by codegen; DO NOT EDIT.

package shield

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Protections() *schema.Table {
	return &schema.Table{
		Name:        "aws_shield_protections",
		Description: `https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Protection.html`,
		Resolver:    fetchShieldProtections,
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
			{
				Name:     "application_layer_automatic_response_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ApplicationLayerAutomaticResponseConfiguration"),
			},
			{
				Name:     "health_check_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("HealthCheckIds"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "resource_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceArn"),
			},
		},
	}
}
