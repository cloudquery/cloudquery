package ssm

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Parameters() *schema.Table {
	tableName := "aws_ssm_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ParameterMetadata.html`,
		Resolver:    fetchSsmParameters,
		Transform:   transformers.TransformWithStruct(&types.ParameterMetadata{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "name",
				Type:                arrow.BinaryTypes.String,
				Description:         `The parameter name`,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveParameterTags,
			},
		},
	}
}

func fetchSsmParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsm).Ssm
	paginator := ssm.NewDescribeParametersPaginator(svc, &ssm.DescribeParametersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Parameters
	}
	return nil
}
func resolveParameterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsm).Ssm
	pm := resource.Item.(types.ParameterMetadata)
	resp, err := svc.ListTagsForResource(ctx, &ssm.ListTagsForResourceInput{
		ResourceId:   pm.Name,
		ResourceType: types.ResourceTypeForTaggingParameter,
	}, func(options *ssm.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	tags := make(map[string]string)
	client.TagsIntoMap(resp.TagList, tags)
	return resource.Set(c.Name, tags)
}
