package wafv2

import (
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ipsets() *schema.Table {
	tableName := "aws_wafv2_ipsets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_IPSet.html`,
		Resolver:            fetchWafv2Ipsets,
		Transform:           client.TransformWithStruct(&types.IPSet{}),
		PreResourceResolver: getIpset,
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer(tableName, "waf-regional"),
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
