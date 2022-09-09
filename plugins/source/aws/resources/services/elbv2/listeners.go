package elbv2

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	notSupportedGatewayLB = regexp.MustCompile("This operation does not support Gateway Load Balancer Listeners")
)

func Elbv2Listeners() *schema.Table {
	return &schema.Table{
		Name:          "aws_elbv2_listeners",
		Description:   "Information about a listener.",
		Resolver:      fetchElbv2Listeners,
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
				Name:        "load_balancer_cq_id",
				Description: "Unique CloudQuery ID of aws_elbv2_load_balancers table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv2listenerTags,
			},
			{
				Name:        "alpn_policy",
				Description: "[TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the listener.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ListenerArn"),
			},
			{
				Name:        "load_balancer_arn",
				Description: "The Amazon Resource Name (ARN) of the load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "port",
				Description: "The port on which the load balancer is listening.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "protocol",
				Description: "The protocol for connections from clients to the load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ssl_policy",
				Description: "[HTTPS or TLS listener] The security policy that defines which protocols and ciphers are supported.",
				Type:        schema.TypeString,
			},
			{
				Name:        "default_actions",
				Description: "Information about default actions",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("DefaultActions"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_elbv2_listener_certificates",
				Description:   "Information about an SSL server certificate.",
				Resolver:      fetchElbv2ListenerCertificates,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "listener_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv2_listeners table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "certificate_arn",
						Description: "The Amazon Resource Name (ARN) of the certificate.",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_default",
						Description: "Indicates whether the certificate is the default certificate",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchElbv2Listeners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	lb := parent.Item.(types.LoadBalancer)
	config := elbv2.DescribeListenersInput{
		LoadBalancerArn: lb.LoadBalancerArn,
	}
	c := meta.(*client.Client)
	svc := c.Services().ELBv2
	for {
		response, err := svc.DescribeListeners(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Listeners
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
func resolveElbv2listenerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ELBv2
	listener := resource.Item.(types.Listener)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*listener.ListenerArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		return err
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, td := range tagsOutput.TagDescriptions {
		for _, s := range td.Tags {
			tags[*s.Key] = s.Value
		}
	}

	return resource.Set(c.Name, tags)
}
func fetchElbv2ListenerCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	region := c.Region
	svc := c.Services().ELBv2
	listener := parent.Item.(types.Listener)
	config := elbv2.DescribeListenerCertificatesInput{ListenerArn: listener.ListenerArn}
	for {
		response, err := svc.DescribeListenerCertificates(ctx, &config, func(options *elbv2.Options) {
			options.Region = region
		})
		if err != nil {
			if client.IsErrorRegex(err, "ValidationError", notSupportedGatewayLB) {
				c.Logger().Debug().Msg("ELBv2: DescribeListenerCertificates not supported for Gateway Load Balancers")
				return nil
			}
			return err
		}
		res <- response.Certificates
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
