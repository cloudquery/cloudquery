package directconnect

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func DirectconnectConnections() *schema.Table {
	return &schema.Table{
		Name:          "aws_directconnect_connections",
		Description:   "Information about a Direct Connect Connection",
		Resolver:      fetchDirectconnectConnections,
		Multiplex:     client.ServiceAccountRegionMultiplexer("directconnect"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		IgnoreInTests: true,
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.DirectConnectService, func(resource *schema.Resource) ([]string, error) {
					return []string{"dxcon", *resource.Item.(types.Connection).ConnectionId}, nil
				}),
			},
			{
				Name:        "aws_device_v2",
				Description: "The Direct Connect endpoint on which the physical connection terminates.",
				Type:        schema.TypeString,
			},
			{
				Name:        "bandwidth",
				Description: "The bandwidth of the connection.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the connection.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConnectionId"),
			},
			{
				Name:        "name",
				Description: "The name of the connection.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConnectionName"),
			},
			{
				Name:        "connection_state",
				Description: "The state of the connection. Possible values are: ordering, requested, pending, available, down, deleting, deleted, rejected, unknown",
				Type:        schema.TypeString,
			},
			{
				Name:        "encryption_mode",
				Description: "The MAC Security (MACsec) connection encryption mode. The valid values are: no_encrypt, should_encrypt, and must_encrypt.",
				Type:        schema.TypeString,
			},
			{
				Name:        "has_logical_redundancy",
				Description: "Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6). Valid values are: yes, no, unknown",
				Type:        schema.TypeString,
			},
			{
				Name:        "jumbo_frame_capable",
				Description: "Indicates whether jumbo frames (9001 MTU) are supported.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "lag_id",
				Description: "The ID of the LAG.",
				Type:        schema.TypeString,
			},
			{
				Name:        "loa_issue_time",
				Description: "The time of the most recent call to DescribeLoa for this connection.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "location",
				Description: "The location of the connection.",
				Type:        schema.TypeString,
			},
			{
				Name:        "mac_sec_capable",
				Description: "Indicates whether the connection supports MAC Security (MACsec).",
				Type:        schema.TypeBool,
			},
			{
				Name:        "owner_account",
				Description: "The ID of the AWS account that owns the connection.",
				Type:        schema.TypeString,
			},
			{
				Name:        "partner_name",
				Description: "The name of the AWS Direct Connect service provider associated with the connection.",
				Type:        schema.TypeString,
			},
			{
				Name:        "port_encryption_status",
				Description: "The MAC Security (MACsec) port link status of the connection. The valid values are Encryption Up, which means that there is an active Connection Key Name, or Encryption Down.",
				Type:        schema.TypeString,
			},
			{
				Name:        "provider_name",
				Description: "The name of the service provider associated with the connection.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the connection.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDirectconnectConnectionTags,
			},
			{
				Name:        "vlan",
				Description: "The ID of the VLAN.",
				Type:        schema.TypeInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_directconnect_connection_mac_sec_keys",
				Description:   "The MAC Security (MACsec) security keys associated with the connection.",
				Resolver:      fetchDirectconnectConnectionMacSecKeys,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"connection_cq_id", "secret_arn"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "connection_cq_id",
						Description: "Unique CloudQuery ID of aws_directconnect_connections table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "connection_id",
						Description: "The ID of the connection.",
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
func fetchDirectconnectConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config directconnect.DescribeConnectionsInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeConnections(ctx, &config, func(options *directconnect.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- output.Connections
	return nil
}

func resolveDirectconnectConnectionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Connection)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchDirectconnectConnectionMacSecKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	connection, ok := parent.Item.(types.Connection)
	if !ok {
		return fmt.Errorf("not a direct connect connection")
	}
	res <- connection.MacSecKeys
	return nil
}
