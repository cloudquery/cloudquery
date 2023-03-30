package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Endpoints() *schema.Table {
	tableName := "aws_eventbridge_endpoints"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Endpoint.html`,
		Resolver:    fetchEndpoints,
		Multiplex:   client.AccountMultiplex(tableName),
		Transform:   transformers.TransformWithStruct(&types.Endpoint{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input eventbridge.ListEndpointsInput
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	for {
		response, err := svc.ListEndpoints(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Endpoints
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
