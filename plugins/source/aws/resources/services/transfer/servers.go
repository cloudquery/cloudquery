package transfer

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Servers() *schema.Table {
	tableName := "aws_transfer_servers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/transfer/latest/userguide/API_DescribedServer.html`,
		Resolver:            fetchTransferServers,
		PreResourceResolver: getServer,
		Transform:           transformers.TransformWithStruct(&types.DescribedServer{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "transfer"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    resolveServersTags,
				Description: `Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described`,
			},
		},
	}
}
func fetchTransferServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Transfer
	input := transfer.ListServersInput{MaxResults: aws.Int32(1000)}
	paginator := transfer.NewListServersPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *transfer.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Servers
	}
	return nil
}

func getServer(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Transfer
	server := resource.Item.(types.ListedServer)

	desc, err := svc.DescribeServer(ctx, &transfer.DescribeServerInput{ServerId: server.ServerId}, func(o *transfer.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = desc.Server
	return nil
}

func resolveServersTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Transfer
	server := resource.Item.(*types.DescribedServer)
	input := transfer.ListTagsForResourceInput{Arn: server.Arn}
	var tags []types.Tag
	paginator := transfer.NewListTagsForResourcePaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *transfer.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				continue
			}
			return err
		}
		tags = append(tags, page.Tags...)
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
}
