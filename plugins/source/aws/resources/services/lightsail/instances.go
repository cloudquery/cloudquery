package lightsail

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_instances",
		Description: "Describes an instance (a virtual private server)",
		Resolver:    fetchLightsailInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:     "access_details",
				Type:     schema.TypeJSON,
				Resolver: resolveLightsailInstanceAccessDetails,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the instance (eg, arn:aws:lightsail:us-east-2:123456789101:Instance/244ad76f-8aad-4741-809f-12345EXAMPLE)",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "blueprint_id",
				Description: "The blueprint ID (eg, os_amlinux_2016_03)",
				Type:        schema.TypeString,
			},
			{
				Name:        "blueprint_name",
				Description: "The friendly name of the blueprint (eg, Amazon Linux)",
				Type:        schema.TypeString,
			},
			{
				Name:        "bundle_id",
				Description: "The bundle for the instance (eg, micro_1_0)",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the instance was created (eg, 147973490917) in Unix time format",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "hardware",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Hardware"),
			},
			{
				Name:        "ip_address_type",
				Description: "The IP address type of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "ipv6_addresses",
				Description: "The IPv6 addresses of the instance",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "is_static_ip",
				Description: "A Boolean value indicating whether this instance has a static IP assigned to it",
				Type:        schema.TypeBool,
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:        "name",
				Description: "The name the user gave the instance (eg, Amazon_Linux-1GB-Ohio-1)",
				Type:        schema.TypeString,
			},
			{
				Name:        "networking_monthly_transfer_gb_per_month_allocated",
				Description: "The amount allocated per month (in GB)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Networking.MonthlyTransfer.GbPerMonthAllocated"),
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP address of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_ip_address",
				Description: "The public IP address of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "The type of resource (usually Instance)",
				Type:        schema.TypeString,
			},
			{
				Name:        "ssh_key_name",
				Description: "The name of the SSH key being used to connect to the instance (eg, LightsailDefaultKeyPair)",
				Type:        schema.TypeString,
			},
			{
				Name:     "state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:        "support_code",
				Description: "The support code",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "username",
				Description: "The user name for connecting to the instance (eg, ec2-user)",
				Type:        schema.TypeString,
			},
			{
				Name:        "add_ons",
				Description: "Describes an add-on that is enabled for an Amazon Lightsail resource",
				Type:        schema.TypeJSON,
			},
			{
				Name: "networking",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_lightsail_instance_port_states",
				Description: "Describes open ports on an instance, the IP addresses allowed to connect to the instance through the ports, and the protocol",
				Resolver:    fetchLightsailInstancePortStates,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cidr_list_aliases",
						Description: "An alias that defines access for a preconfigured range of IP addresses",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "cidrs",
						Description: "The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "from_port",
						Description: "The first port in a range of open ports on an instance",
						Type:        schema.TypeInt,
					},
					{
						Name:        "ipv6_cidrs",
						Description: "The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "protocol",
						Description: "The IP protocol name",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "Specifies whether the instance port is open or closed",
						Type:        schema.TypeString,
					},
					{
						Name:        "to_port",
						Description: "The last port in a range of open ports on an instance",
						Type:        schema.TypeInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	input := lightsail.GetInstancesInput{}
	for {
		output, err := svc.GetInstances(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.Instances

		if aws.ToString(output.NextPageToken) == "" {
			break
		}
		input.PageToken = output.NextPageToken
	}
	return nil
}
func resolveLightsailInstanceAccessDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services().Lightsail
	input := lightsail.GetInstanceAccessDetailsInput{InstanceName: r.Name}
	output, err := svc.GetInstanceAccessDetails(ctx, &input)
	if err != nil {
		return err
	}
	j, err := json.Marshal(output.AccessDetails)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func fetchLightsailInstancePortStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services().Lightsail
	input := lightsail.GetInstancePortStatesInput{InstanceName: r.Name}
	output, err := svc.GetInstancePortStates(ctx, &input)
	if err != nil {
		return err
	}

	res <- output.PortStates
	return nil
}
