package appmesh

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Meshes() *schema.Table {
	tableName := "aws_appmesh_meshes"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_MeshData.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:            fetchMeshes,
		PreResourceResolver: getMesh,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "appmesh"),
		Transform:           transformers.TransformWithStruct(&types.MeshData{}),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Metadata.Arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveMeshTags,
			},
		},
		Relations: []*schema.Table{
			virtualServices(),
			virtualNodes(),
			virtualRouters(),
			virtualGateways(),
		},
	}
}

func fetchMeshes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppmesh).Appmesh

	paginator := appmesh.NewListMeshesPaginator(svc, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *appmesh.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Meshes
	}
	return nil
}

func getMesh(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppmesh).Appmesh
	mesh := resource.Item.(types.MeshRef)
	input := appmesh.DescribeMeshInput{
		MeshName:  mesh.MeshName,
		MeshOwner: mesh.MeshOwner,
	}
	output, err := svc.DescribeMesh(ctx, &input, func(o *appmesh.Options) { o.Region = cl.Region })
	if err != nil {
		return err
	}
	resource.Item = output.Mesh
	return nil
}

func resolveMeshTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	mesh := resource.Item.(*types.MeshData)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppmesh).Appmesh

	paginator := appmesh.NewListTagsForResourcePaginator(svc, &appmesh.ListTagsForResourceInput{
		ResourceArn: mesh.Metadata.Arn,
	})
	tags := make(map[string]string)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *appmesh.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		client.TagsIntoMap(page.Tags, tags)
	}
	return resource.Set(c.Name, tags)
}
