package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func DirectconnectLags() *schema.Table {
	return &schema.Table{
		Name:         "aws_directconnect_lags",
		Description:  "Information about Direct Connect Link Aggregation Group (LAG)",
		Resolver:     fetchDirectconnectLags,
		Multiplex:    client.ServiceAccountRegionMultiplexer("directconnect"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
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
				Resolver:    resolveDirectconnectLagConnectionIds,
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
				Name:        "id",
				Description: "The ID of the LAG.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LagId"),
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
				Resolver:    resolveDirectconnectLagTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_directconnect_lag_mac_sec_keys",
				Description: "The MAC Security (MACsec) security keys associated with the LAG.",
				Resolver:    fetchDirectconnectLagMacSecKeys,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"lag_cq_id", "secret_arn"}},
				Columns: []schema.Column{
					{
						Name:        "lag_cq_id",
						Description: "Unique CloudQuery ID of aws_directconnect_lags table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "lag_id",
						Description: "The ID of the LAG.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "ckn",
						Description: "The Connection Key Name (CKN) for the MAC Security secret key.",
						Type:        schema.TypeString,
					},
					{
						Name:        "secret_arn",
						Description: "The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecretARN"),
					},
					{
						Name:        "start_on",
						Description: "The date that the MAC Security (MACsec) secret key takes effect. The value is displayed in UTC format.",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "The state of the MAC Security secret key. The possible values are: associating, associated, disassociating, disassociated",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDirectconnectLags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config directconnect.DescribeLagsInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeLags(ctx, &config, func(options *directconnect.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- output.Lags
	return nil
}

func resolveDirectconnectLagTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Lag)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}

func fetchDirectconnectLagMacSecKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	connection, ok := parent.Item.(types.Lag)
	if !ok {
		return fmt.Errorf("not a direct connect LAG")
	}
	res <- connection.MacSecKeys
	return nil
}

func resolveDirectconnectLagConnectionIds(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Lag)
	connectionIds := make([]*string, len(r.Connections))
	for i, connection := range r.Connections {
		connectionIds[i] = connection.ConnectionId
	}
	return resource.Set("connection_ids", connectionIds)
}
