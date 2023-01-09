package eventbridge

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Endpoints() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_endpoints",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Endpoint.html`,
		Resolver:    fetchEventbridgeEndpoints,
		Multiplex:   client.AccountMultiplex,
		Transform:   transformers.TransformWithStruct(&types.Endpoint{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
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
