package glue

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DevEndpoints() *schema.Table {
	tableName := "aws_glue_dev_endpoints"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_DevEndpoint.html`,
		Resolver:    fetchGlueDevEndpoints,
		Transform:   transformers.TransformWithStruct(&types.DevEndpoint{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveGlueDevEndpointArn,
				PrimaryKeyComponent: true,
			},
			tagsCol(func(cl *client.Client, resource *schema.Resource) string {
				return devEndpointARN(cl, aws.ToString(resource.Item.(types.DevEndpoint).EndpointName))
			}),
		},
	}
}

func fetchGlueDevEndpoints(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlue).Glue
	paginator := glue.NewGetDevEndpointsPaginator(svc, &glue.GetDevEndpointsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- page.DevEndpoints
	}
	return nil
}
func resolveGlueDevEndpointArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, devEndpointARN(cl, aws.ToString(resource.Item.(types.DevEndpoint).EndpointName)))
}

func devEndpointARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "devEndpoint/" + name,
	}.String()
}
