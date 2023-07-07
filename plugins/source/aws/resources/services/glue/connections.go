package glue

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Connections() *schema.Table {
	tableName := "aws_glue_connections"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Connection.html`,
		Resolver:    fetchGlueConnections,
		Transform:   transformers.TransformWithStruct(&types.Connection{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveGlueConnectionArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchGlueConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	paginator := glue.NewGetConnectionsPaginator(svc, &glue.GetConnectionsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ConnectionList
	}
	return nil
}
func resolveGlueConnectionArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.Connection)
	return resource.Set(c.Name, connectionARN(cl, &r))
}

func connectionARN(cl *client.Client, c *types.Connection) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("connection/%s", aws.ToString(c.Name)),
	}.String()
}
