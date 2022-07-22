package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource load_balancers --config gen.hcl --output .
func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:         "aws_lightsail_load_balancers",
		Description:  "Describes a load balancer",
		Resolver:     fetchLightsailLoadBalancers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Description: "The Amazon Resource Name (ARN) of the load balancer",
				Type:        schema.TypeString,
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
				Name:        "availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.AvailabilityZone"),
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
				Resolver:    resolveLoadBalancersTags,
			},
			{
				Name:        "tls_policy_name",
				Description: "The name of the TLS security policy for the load balancer",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_lightsail_load_balancer_instance_health_summary",
				Description: "Describes information about the health of the instance",
				Resolver:    fetchLightsailLoadBalancerInstanceHealthSummaries,
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_load_balancers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_health",
						Description: "Describes the overall instance health",
						Type:        schema.TypeString,
					},
					{
						Name:        "instance_health_reason",
						Description: "More information about the instance health",
						Type:        schema.TypeString,
					},
					{
						Name:        "instance_name",
						Description: "The name of the Lightsail instance for which you are requesting health check data",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_lightsail_load_balancer_tls_certificate_summaries",
				Description:   "Provides a summary of SSL/TLS certificate metadata",
				Resolver:      fetchLightsailLoadBalancerTlsCertificateSummaries,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_load_balancers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "is_attached",
						Description: "When true, the SSL/TLS certificate is attached to the Lightsail load balancer",
						Type:        schema.TypeBool,
					},
					{
						Name:        "name",
						Description: "The name of the SSL/TLS certificate",
						Type:        schema.TypeString,
					},
				},
			},
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
						Name:        "availability_zone",
						Description: "The Availability Zone",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.AvailabilityZone"),
					},
					{
						Name:        "region_name",
						Description: "The AWS Region name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.RegionName"),
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
						Name:        "renewal_summary_renewal_status",
						Description: "The renewal status of the certificate",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RenewalSummary.RenewalStatus"),
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
						Resolver:    resolveLoadBalancerTLSCertificatesTags,
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
		response, err := svc.GetLoadBalancers(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
	return diag.WrapError(resource.Set(c.Name, ports))
}
func resolveLoadBalancersTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.LoadBalancer)
	tags := make(map[string]string)
	client.TagsIntoMap(r.Tags, tags)
	return diag.WrapError(resource.Set(c.Name, tags))
}
func fetchLightsailLoadBalancerInstanceHealthSummaries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.LoadBalancer)
	res <- r.InstanceHealthSummary
	return nil
}
func fetchLightsailLoadBalancerTlsCertificateSummaries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.LoadBalancer)
	res <- r.TlsCertificateSummaries
	return nil
}
func fetchLightsailLoadBalancerTlsCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.LoadBalancer)
	input := lightsail.GetLoadBalancerTlsCertificatesInput{
		LoadBalancerName: r.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetLoadBalancerTlsCertificates(ctx, &input, func(options *lightsail.Options) {
		options.Region = c.Region
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	res <- response.TlsCertificates
	return nil
}
func resolveLoadBalancerTLSCertificatesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.LoadBalancerTlsCertificate)
	tags := make(map[string]string)
	client.TagsIntoMap(r.Tags, tags)
	return diag.WrapError(resource.Set(c.Name, tags))
}
