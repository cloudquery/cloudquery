package wafregional

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func WebAcls() *schema.Table {
	tableName := "aws_wafregional_web_acls"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_WebACL.html`,
		Resolver:    fetchWafregionalWebAcls,
		Transform:   transformers.TransformWithStruct(&types.WebACL{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("WebACLArn"),
				PrimaryKey: true,
			},
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    resolveWafregionalWebACLTags,
				Description: `Web ACL tags.`,
			},
			{
				Name:     "resources_for_web_acl",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: resolveWafregionalWebACLResourcesForWebACL,
			},
		},
	}
}

func fetchWafregionalWebAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafregional
	var params wafregional.ListWebACLsInput
	for {
		result, err := svc.ListWebACLs(ctx, &params, func(o *wafregional.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, item := range result.WebACLs {
			detail, err := svc.GetWebACL(
				ctx,
				&wafregional.GetWebACLInput{WebACLId: item.WebACLId},
				func(o *wafregional.Options) {
					o.Region = cl.Region
				},
			)
			if err != nil {
				return err
			}
			if detail.WebACL == nil {
				continue
			}
			res <- *detail.WebACL
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}
func resolveWafregionalWebACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafregional
	params := wafregional.ListTagsForResourceInput{ResourceARN: resource.Item.(types.WebACL).WebACLArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params, func(o *wafregional.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if result.TagInfoForResource != nil {
			client.TagsIntoMap(result.TagInfoForResource.TagList, tags)
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return resource.Set(c.Name, tags)
}

func resolveWafregionalWebACLResourcesForWebACL(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wafregional
	output, err := service.ListResourcesForWebACL(ctx, &wafregional.ListResourcesForWebACLInput{
		WebACLId: resource.Item.(types.WebACL).WebACLId,
	}, func(o *wafregional.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.ResourceArns)
}
