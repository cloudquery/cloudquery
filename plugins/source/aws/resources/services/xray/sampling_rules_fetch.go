package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchXraySamplingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := xray.NewGetSamplingRulesPaginator(meta.(*client.Client).Services().Xray, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
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

	output, err := svc.ListTagsForResource(ctx, &params)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
