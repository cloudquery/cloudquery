package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func DirectconnectGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_directconnect_gateways",
		Description:  "Information about a Direct Connect gateway, which enables you to connect virtual interfaces and virtual private gateway or transit gateways.",
		Resolver:     fetchDirectconnectGateways,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "id"}},
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
				Name:        "amazon_side_asn",
				Description: "The autonomous system number (ASN) for the Amazon side of the connection.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "id",
				Description: "The ID of the Direct Connect gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectConnectGatewayId"),
			},
			{
				Name:        "name",
				Description: "The name of the Direct Connect gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectConnectGatewayName"),
			},
			{
				Name:        "state",
				Description: "The state of the Direct Connect gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectConnectGatewayState"),
			},
			{
				Name:        "owner_account",
				Description: "The ID of the AWS account that owns the Direct Connect gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_change_error",
				Description: "The error message if the state of an object failed to advance.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_directconnect_gateway_associations",
				Description: "Information about the association between an Direct Connect Gateway and either a Virtual Private Gateway, or Transit Gateway",
				Resolver:    fetchDirectconnectGatewayAssociations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"gateway_cq_id", "association_id"}},
				Columns: []schema.Column{
					{
						Name:        "gateway_cq_id",
						Description: "Unique CloudQuery ID of aws_directconnect_gateways table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "gateway_id",
						Description: "The ID of the Direct Connect gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "allowed_prefixes_to_direct_connect_gateway",
						Description: "The Amazon VPC prefixes to advertise to the Direct Connect gateway.",
						Type:        schema.TypeStringArray,
						Resolver:    resolveDirectconnectGatewayAssociationAllowedPrefixes,
					},
					{
						Name:        "associated_gateway_id",
						Description: "The ID of the associated gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AssociatedGateway.Id"),
					},
					{
						Name:        "associated_gateway_owner_account",
						Description: "The ID of the AWS account that owns the associated virtual private gateway or transit gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AssociatedGateway.OwnerAccount"),
					},
					{
						Name:        "associated_gateway_region",
						Description: "The Region where the associated gateway is located.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AssociatedGateway.Region"),
					},
					{
						Name:        "associated_gateway_type",
						Description: "The type of associated gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AssociatedGateway.Type"),
					},
					{
						Name:        "association_id",
						Description: "The ID of the Direct Connect gateway association",
						Type:        schema.TypeString,
					},
					{
						Name:        "association_state",
						Description: "The state of the association.",
						Type:        schema.TypeString,
					},
					{
						Name:        "direct_connect_gateway_owner_account",
						Description: "The ID of the AWS account that owns the associated gateway.",
						Type:        schema.TypeString,
					},
					{
						Name:        "state_change_error",
						Description: "The error message if the state of an object failed to advance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_gateway_id",
						Description: "The ID of the virtual private gateway. Applies only to private virtual interfaces.",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_gateway_owner_account",
						Description: "The ID of the AWS account that owns the virtual private gateway.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_id",
						Description: "The ID of the Direct Connect gateway association",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AssociationId"),
					},
				},
			},
			{
				Name:        "aws_directconnect_gateway_attachments",
				Description: "Information about the attachment between a Direct Connect gateway and virtual interfaces.",
				Resolver:    fetchDirectconnectGatewayAttachments,
				Columns: []schema.Column{
					{
						Name:        "gateway_cq_id",
						Description: "Unique CloudQuery ID of aws_directconnect_gateways table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "gateway_id",
						Description: "The ID of the Direct Connect gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "attachment_state",
						Description: "The state of the attachment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "attachment_type",
						Description: "The type of attachment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "state_change_error",
						Description: "The error message if the state of an object failed to advance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_interface_id",
						Description: "The ID of the virtual interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_interface_owner_account",
						Description: "The ID of the AWS account that owns the virtual interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_interface_region",
						Description: "The AWS Region where the virtual interface is located.",
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
func fetchDirectconnectGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config directconnect.DescribeDirectConnectGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	for {
		output, err := svc.DescribeDirectConnectGateways(ctx, &config, func(options *directconnect.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.DirectConnectGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchDirectconnectGatewayAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	gateway, ok := parent.Item.(types.DirectConnectGateway)
	if !ok {
		return fmt.Errorf("not direct connect gateway")
	}

	var config directconnect.DescribeDirectConnectGatewayAssociationsInput
	config.DirectConnectGatewayId = gateway.DirectConnectGatewayId

	c := meta.(*client.Client)
	svc := c.Services().Directconnect

	for {
		output, err := svc.DescribeDirectConnectGatewayAssociations(ctx, &config, func(options *directconnect.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.DirectConnectGatewayAssociations
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchDirectconnectGatewayAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	gateway, ok := parent.Item.(types.DirectConnectGateway)
	if !ok {
		return fmt.Errorf("not direct connect gateway")
	}

	var config directconnect.DescribeDirectConnectGatewayAttachmentsInput
	config.DirectConnectGatewayId = gateway.DirectConnectGatewayId

	c := meta.(*client.Client)
	svc := c.Services().Directconnect

	for {
		output, err := svc.DescribeDirectConnectGatewayAttachments(ctx, &config, func(options *directconnect.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.DirectConnectGatewayAttachments
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveDirectconnectGatewayAssociationAllowedPrefixes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DirectConnectGatewayAssociation)
	allowedPrefixes := make([]*string, len(r.AllowedPrefixesToDirectConnectGateway))
	for i, prefix := range r.AllowedPrefixesToDirectConnectGateway {
		allowedPrefixes[i] = prefix.Cidr
	}
	return resource.Set(c.Name, allowedPrefixes)
}
