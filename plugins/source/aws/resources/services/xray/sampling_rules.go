package xray

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SamplingRules() *schema.Table {
	tableName := "aws_xray_sampling_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_SamplingRuleRecord.html`,
		Resolver:    fetchXraySamplingRules,
		Transform:   transformers.TransformWithStruct(&types.SamplingRuleRecord{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "xray"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("SamplingRule.RuleARN"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveXraySamplingRuleTags,
			},
		},
	}
}

func fetchXraySamplingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := xray.NewGetSamplingRulesPaginator(cl.Services().Xray, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(o *xray.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.SamplingRuleRecords
	}
	return nil
}
func resolveXraySamplingRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	sr := resource.Item.(types.SamplingRuleRecord)
	cl := meta.(*client.Client)
	svc := cl.Services().Xray
	params := xray.ListTagsForResourceInput{ResourceARN: sr.SamplingRule.RuleARN}

	output, err := svc.ListTagsForResource(ctx, &params, func(o *xray.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
