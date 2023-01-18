package wafv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/wafv2/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WebAcls() *schema.Table {
	return &schema.Table{
		Name:                "aws_wafv2_web_acls",
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_WebACL.html`,
		Resolver:            fetchWafv2WebAcls,
		PreResourceResolver: getWebAcl,
		Transform:           transformers.TransformWithStruct(&models.WebACLWrapper{}, transformers.WithUnwrapStructFields("WebACL")),
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWebACLTags,
			},
			{
				Name:     "resources_for_web_acl",
				Type:     schema.TypeStringArray,
				Resolver: resolveWafv2webACLResourcesForWebACL,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
