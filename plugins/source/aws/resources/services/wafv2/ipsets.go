package wafv2

import (
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Ipsets() *schema.Table {
	return &schema.Table{
		Name:                "aws_wafv2_ipsets",
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_IPSet.html`,
		Resolver:            fetchWafv2Ipsets,
		Transform:           transformers.TransformWithStruct(&types.IPSet{}),
		PreResourceResolver: getIpset,
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "addresses",
				Type:     schema.TypeInetArray,
				Resolver: resolveIpsetAddresses,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIpsetTags,
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
