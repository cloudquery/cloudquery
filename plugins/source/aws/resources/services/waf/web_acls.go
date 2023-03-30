package waf

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WebAcls() *schema.Table {
	tableName := "aws_waf_web_acls"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_WebACLSummary.html`,
		Resolver:    fetchWafWebAcls,
		Transform:   transformers.TransformWithStruct(&WebACLWrapper{}, transformers.WithUnwrapStructFields("WebACL")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "waf"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebACLArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafWebACLTags,
			},
		},
	}
}
