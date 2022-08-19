package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource servers --config servers.hcl --output .
func Servers() *schema.Table {
	return &schema.Table{
		Name:         "aws_transfer_servers",
		Description:  "Describes the properties of a file transfer protocol-enabled server that was specified",
		Resolver:     fetchTransferServers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Description: "Specifies the unique Amazon Resource Name (ARN) of the server",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate",
				Description: "Specifies the ARN of the Amazon Web ServicesCertificate Manager (ACM) certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain",
				Description: "Specifies the domain of the storage system that is used for file transfers",
				Type:        schema.TypeString,
			},
			{
				Name:        "endpoint_details_address_allocation_ids",
				Description: "A list of address allocation IDs that are required to attach an Elastic IP address to your server's endpoint",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointDetails.AddressAllocationIds"),
			},
			{
				Name:        "endpoint_details_security_group_ids",
				Description: "A list of security groups IDs that are available to attach to your server's endpoint",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointDetails.SecurityGroupIds"),
			},
			{
				Name:        "endpoint_details_subnet_ids",
				Description: "A list of subnet IDs that are required to host your server endpoint in your VPC This property can only be set when EndpointType is set to VPC",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointDetails.SubnetIds"),
			},
			{
				Name:        "endpoint_details_vpc_endpoint_id",
				Description: "The ID of the VPC endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EndpointDetails.VpcEndpointId"),
			},
			{
				Name:        "endpoint_details_vpc_id",
				Description: "The VPC ID of the VPC in which a server's endpoint will be hosted",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EndpointDetails.VpcId"),
			},
			{
				Name:        "endpoint_type",
				Description: "Defines the type of endpoint that your server is connected to",
				Type:        schema.TypeString,
			},
			{
				Name:        "host_key_fingerprint",
				Description: "Specifies the Base64-encoded SHA256 fingerprint of the server's host key",
				Type:        schema.TypeString,
			},
			{
				Name:        "identity_provider_details_directory_id",
				Description: "The identifier of the Directory Service directory that you want to stop sharing",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IdentityProviderDetails.DirectoryId"),
			},
			{
				Name:        "identity_provider_details_function",
				Description: "The ARN for a lambda function to use for the Identity provider",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IdentityProviderDetails.Function"),
			},
			{
				Name:        "identity_provider_details_invocation_role",
				Description: "Provides the type of InvocationRole used to authenticate the user account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IdentityProviderDetails.InvocationRole"),
			},
			{
				Name:        "identity_provider_details_url",
				Description: "Provides the location of the service endpoint used to authenticate users",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IdentityProviderDetails.Url"),
			},
			{
				Name:        "identity_provider_type",
				Description: "The mode of authentication for a server",
				Type:        schema.TypeString,
			},
			{
				Name:        "logging_role",
				Description: "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFSevents",
				Type:        schema.TypeString,
			},
			{
				Name:        "post_authentication_login_banner",
				Description: "Specifies a string to display when users connect to a server",
				Type:        schema.TypeString,
			},
			{
				Name:        "pre_authentication_login_banner",
				Description: "Specifies a string to display when users connect to a server",
				Type:        schema.TypeString,
			},
			{
				Name:        "protocol_details_as2_transports",
				Description: "Indicates the transport method for the AS2 messages",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ProtocolDetails.As2Transports"),
			},
			{
				Name:        "protocol_details_passive_ip",
				Description: "Indicates passive mode, for FTP and FTPS protocols",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProtocolDetails.PassiveIp"),
			},
			{
				Name:        "protocol_details_set_stat_option",
				Description: "Use the SetStatOption to ignore the error that is generated when the client attempts to use SETSTAT on a file you are uploading to an S3 bucket",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProtocolDetails.SetStatOption"),
			},
			{
				Name:        "protocol_details_tls_session_resumption_mode",
				Description: "A property used with Transfer Family servers that use the FTPS protocol",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProtocolDetails.TlsSessionResumptionMode"),
			},
			{
				Name:        "protocols",
				Description: "Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "security_policy_name",
				Description: "Specifies the name of the security policy that is attached to the server",
				Type:        schema.TypeString,
			},
			{
				Name:        "server_id",
				Description: "Specifies the unique system-assigned identifier for a server that you instantiate",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The condition of the server that was described",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described",
				Type:        schema.TypeJSON,
				Resolver:    resolveServersTags,
			},
			{
				Name:        "user_count",
				Description: "Specifies the number of users that are assigned to a server you specified with the ServerId",
				Type:        schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_transfer_server_workflow_details_on_upload",
				Description: "Specifies the workflow ID for the workflow to assign and the execution role that's used for executing the workflow",
				Resolver:    schema.PathTableResolver("WorkflowDetails.OnUpload"),
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique CloudQuery ID of aws_transfer_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "execution_role",
						Description: "Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources",
						Type:        schema.TypeString,
					},
					{
						Name:        "workflow_id",
						Description: "A unique identifier for the workflow",
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

func fetchTransferServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Transfer
	input := transfer.ListServersInput{MaxResults: aws.Int32(1000)}
	for {
		result, err := svc.ListServers(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, server := range result.Servers {
			desc, err := svc.DescribeServer(ctx, &transfer.DescribeServerInput{ServerId: server.ServerId})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			if desc.Server != nil {
				res <- desc.Server
			}
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveServersTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Transfer
	server := resource.Item.(*types.DescribedServer)
	input := transfer.ListTagsForResourceInput{Arn: server.Arn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &input)
		if err != nil {
			if cl.IsNotFoundError(err) {
				continue
			}
			return diag.WrapError(err)
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
