package networkmanager

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func GlobalNetworks() *schema.Table {
	tableName := "aws_networkmanager_global_networks"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_GlobalNetwork.html
The  'request_region' column is added to show region of where the request was made from.`,
		Resolver:  fetchNetworks,
		Transform: transformers.TransformWithStruct(&types.GlobalNetwork{}),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "networkmanager"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			// Only including the request_region in the PK because the ARN doesn't include it as it is a global resource
			client.RequestRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("GlobalNetworkArn"),
				PrimaryKeyComponent: true,
			},

			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: schema.Tables{
			transitGatewayRegistration(),
			sites(),
			links(),
		},
	}
}

func fetchNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNetworkmanager).Networkmanager
	paginator := networkmanager.NewDescribeGlobalNetworksPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *networkmanager.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.GlobalNetworks
	}
	return nil
}
