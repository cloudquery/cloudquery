package route53

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func trafficPolicyVersions() *schema.Table {
	tableName := "aws_route53_traffic_policy_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicy.html`,
		Resolver:    fetchRoute53TrafficPolicyVersions,
		Transform:   transformers.TransformWithStruct(&types.TrafficPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "traffic_policy_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Id"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "version",
				Type:                arrow.PrimitiveTypes.Int64,
				Resolver:            schema.PathResolver("Version"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "document",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Document"),
			},
		},
	}
}
