package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_load_balancers",
		Description: "Describes a load balancer",
		Resolver:    fetchLightsailLoadBalancers,
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the load balancer",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "configuration_options",
				Description: "A string to string map of the configuration options for your load balancer Valid values are listed below",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "created_at",
				Description: "The date when your load balancer was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "dns_name",
				Description: "The DNS name of your Lightsail load balancer",
				Type:        schema.TypeString,
			},
			{
				Name:        "health_check_path",
				Description: "The path you specified to perform your health checks",
				Type:        schema.TypeString,
			},
			{
				Name:        "https_redirection_enabled",
				Description: "A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer",
				Type:        schema.TypeBool,
			},
			{
				Name:        "instance_port",
				Description: "The port where the load balancer will direct traffic to your Lightsail instances",
				Type:        schema.TypeInt,
			},
			{
				Name:        "ip_address_type",
				Description: "The IP address type of the load balancer",
				Type:        schema.TypeString,
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:        "name",
				Description: "The name of the load balancer (eg, my-load-balancer)",
				Type:        schema.TypeString,
			},
			{
				Name:        "protocol",
				Description: "The protocol you have enabled for your load balancer",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_ports",
				Description: "An array of public port settings for your load balancer",
				Type:        schema.TypeIntArray,
				Resolver:    resolveLoadBalancersPublicPorts,
			},
			{
				Name:        "resource_type",
				Description: "Type of the lightsail resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The status of your load balancer",
				Type:        schema.TypeString,
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
				Name:        "tls_policy_name",
				Description: "The name of the TLS security policy for the load balancer",
				Type:        schema.TypeString,
			},
			{
				Name: "instance_health_summary",
				Type: schema.TypeJSON,
			},
			{
				Name: "tsl_certificate_summaries",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_lightsail_load_balancer_tls_certificates",
				Description:   "Describes a load balancer SSL/TLS certificate",
				Resolver:      fetchLightsailLoadBalancerTlsCertificates,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_load_balancers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the SSL/TLS certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "The time when you created your SSL/TLS certificate",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "domain_name",
						Description: "The domain name for your SSL/TLS certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "domain_validation_records",
						Description: "An array of LoadBalancerTlsCertificateDomainValidationRecord objects describing the records",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "failure_reason",
						Description: "The validation failure reason, if any, of the certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_attached",
						Description: "When true, the SSL/TLS certificate is attached to the Lightsail load balancer",
						Type:        schema.TypeBool,
					},
					{
						Name:        "issued_at",
						Description: "The time when the SSL/TLS certificate was issued",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "issuer",
						Description: "The issuer of the certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_algorithm",
						Description: "The algorithm used to generate the key pair (the public and private key)",
						Type:        schema.TypeString,
					},
					{
						Name:        "load_balancer_name",
						Description: "The load balancer name where your SSL/TLS certificate is attached",
						Type:        schema.TypeString,
					},
					{
						Name:     "location",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Location"),
					},
					{
						Name:        "name",
						Description: "The name of the SSL/TLS certificate (eg, my-certificate)",
						Type:        schema.TypeString,
					},
					{
						Name:        "not_after",
						Description: "The timestamp when the SSL/TLS certificate expires",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "not_before",
						Description: "The timestamp when the SSL/TLS certificate is first valid",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "renewal_summary_domain_validation_options",
						Description: "Contains information about the validation of each domain name in the certificate, as it pertains to Lightsail's managed renewal",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("RenewalSummary.DomainValidationOptions"),
					},
					{
						Name:     "renewal_summary",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("RenewalSummary"),
					},
					{
						Name:        "resource_type",
						Description: "The resource type (eg, LoadBalancerTlsCertificate)  * Instance - A Lightsail instance (a virtual private server)  * StaticIp - A static IP address  * KeyPair - The key pair used to connect to a Lightsail instance  * InstanceSnapshot - A Lightsail instance snapshot  * Domain - A DNS zone  * PeeredVpc - A peered VPC  * LoadBalancer - A Lightsail load balancer  * LoadBalancerTlsCertificate - An SSL/TLS certificate associated with a Lightsail load balancer  * Disk - A Lightsail block storage disk  * DiskSnapshot - A block storage disk snapshot",
						Type:        schema.TypeString,
					},
					{
						Name:        "revocation_reason",
						Description: "The reason the certificate was revoked",
						Type:        schema.TypeString,
					},
					{
						Name:        "revoked_at",
						Description: "The timestamp when the certificate was revoked",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "serial",
						Description: "The serial number of the certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "signature_algorithm",
						Description: "The algorithm that was used to sign the certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The validation status of the SSL/TLS certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "subject",
						Description: "The name of the entity that is associated with the public key contained in the certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "subject_alternative_names",
						Description: "An array of strings that specify the alternate domains (eg, example2com) and subdomains (eg, blogexamplecom) for the certificate",
						Type:        schema.TypeStringArray,
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
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetLoadBalancersInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetLoadBalancers(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.LoadBalancers
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func resolveLoadBalancersPublicPorts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.LoadBalancer)
	ports := make([]int, 0, len(r.PublicPorts))
	for _, p := range r.PublicPorts {
		ports = append(ports, int(p))
	}
	return resource.Set(c.Name, ports)
}
func fetchLightsailLoadBalancerTlsCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.LoadBalancer)
	input := lightsail.GetLoadBalancerTlsCertificatesInput{
		LoadBalancerName: r.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetLoadBalancerTlsCertificates(ctx, &input)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.TlsCertificates
	return nil
}
