package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DirectconnectLags() *schema.Table {
	return &schema.Table{
		Name:          "aws_directconnect_lags",
		Description:   "Information about Direct Connect Link Aggregation Group (LAG)",
		Resolver:      fetchDirectconnectLags,
		Multiplex:     client.ServiceAccountRegionMultiplexer("directconnect"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.DirectConnectService, func(resource *schema.Resource) ([]string, error) {
					return []string{"dxlag", *resource.Item.(types.Lag).LagId}, nil
				}),
			},
			{
				Name:        "allows_hosted_connections",
				Description: "Indicates whether the LAG can host other connections.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "aws_device_v2",
				Description: "The AWS Direct Connect endpoint that hosts the LAG.",
				Type:        schema.TypeString,
			},
			{
				Name:        "connection_ids",
				Description: "The list of IDs of Direct Connect Connections bundled by the LAG",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Connections.ConnectionId"),
			},
			{
				Name:        "connections_bandwidth",
				Description: "The individual bandwidth of the physical connections bundled by the LAG.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encryption_mode",
				Description: "The LAG MAC Security (MACsec) encryption mode.",
				Type:        schema.TypeString,
			},
			{
				Name:        "has_logical_redundancy",
				Description: "Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).",
				Type:        schema.TypeString,
			},
			{
				Name:        "jumbo_frame_capable",
				Description: "Indicates whether jumbo frames (9001 MTU) are supported.",
				Type:        schema.TypeBool,
			},
			{
				Name:            "id",
				Description:     "The ID of the LAG.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("LagId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "The name of the LAG.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LagName"),
			},
			{
				Name:        "state",
				Description: "The state of the LAG. Possible values are: requested, pending, available, down, deleting, deleted, unknown",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LagState"),
			},
			{
				Name:        "location",
				Description: "The location of the LAG.",
				Type:        schema.TypeString,
			},
			{
				Name:        "mac_sec_capable",
				Description: "Indicates whether the LAG supports MAC Security (MACsec).",
				Type:        schema.TypeBool,
			},
			{
				Name:        "minimum_links",
				Description: "The minimum number of physical dedicated connections that must be operational for the LAG itself to be operational.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "number_of_connections",
				Description: "The number of physical dedicated connections bundled by the LAG, up to a maximum of 10.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "owner_account",
				Description: "The ID of the AWS account that owns the LAG.",
				Type:        schema.TypeString,
			},
			{
				Name:        "provider_name",
				Description: "The name of the service provider associated with the LAG.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the LAG.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "mac_sec_keys",
				Description: "The MAC Security (MACsec) security keys associated with the connection.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("MacSecKeys"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchDirectconnectLags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config directconnect.DescribeLagsInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeLags(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.Lags
	return nil
}
