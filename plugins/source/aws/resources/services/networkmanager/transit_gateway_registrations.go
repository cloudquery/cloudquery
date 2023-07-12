package networkmanager

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func transitGatewayRegistration() *schema.Table {
	tableName := "aws_networkmanager_transit_gateway_registrations"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_TransitGatewayRegistration.html
The  'request_region' column is added to show region of where the request was made from.`,
		Resolver:  fetchTransitGatewayRegistration,
		Transform: transformers.TransformWithStruct(&types.TransitGatewayRegistration{}, transformers.WithPrimaryKeys("GlobalNetworkId", "TransitGatewayArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:       "request_region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSRegion,
				PrimaryKey: true,
			}},
		Relations: schema.Tables{},
	}
}

func fetchTransitGatewayRegistration(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Networkmanager
	globalNetwork := parent.Item.(types.GlobalNetwork)
	input := &networkmanager.GetTransitGatewayRegistrationsInput{
		GlobalNetworkId: globalNetwork.GlobalNetworkId,
	}
	paginator := networkmanager.NewGetTransitGatewayRegistrationsPaginator(svc, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *networkmanager.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TransitGatewayRegistrations
	}
	return nil
}
