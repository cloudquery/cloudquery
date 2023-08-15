package detective

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Graphs() *schema.Table {
	tableName := "aws_detective_graphs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/detective/latest/APIReference/API_ListGraphs.html`,
		Resolver:    fetchGraphs,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "api.detective"),
		Transform:   transformers.TransformWithStruct(&types.Graph{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveGraphTags,
			},
		},
		Relations: schema.Tables{
			members(),
		},
	}
}

func fetchGraphs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Detective
	config := detective.ListGraphsInput{}
	paginator := detective.NewListGraphsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *detective.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- page.GraphList
	}

	return nil
}
func resolveGraphTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	graph := resource.Item.(types.Graph)
	cl := meta.(*client.Client)
	svc := cl.Services().Detective
	input := &detective.ListTagsForResourceInput{
		ResourceArn: graph.Arn,
	}
	response, err := svc.ListTagsForResource(ctx, input, func(options *detective.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
