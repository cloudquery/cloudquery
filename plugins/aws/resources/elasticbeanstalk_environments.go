package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ElasticbeanstalkEnvironments() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticbeanstalk_environments",
		Description:  "Describes the properties of an environment.",
		Resolver:     fetchElasticbeanstalkEnvironments,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:        "abortable_operation_in_progress",
				Description: "Indicates if there is an in-progress environment configuration update or application version deployment that you can cancel.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "application_name",
				Description: "The name of the application associated with this environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cname",
				Description: "The URL to the CNAME for this environment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CNAME"),
			},
			{
				Name:        "date_created",
				Description: "The creation date for this environment.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "date_updated",
				Description: "The last modified date for this environment.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "Describes this environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "endpoint_url",
				Description: "For load-balanced, autoscaling environments, the URL to the LoadBalancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EndpointURL"),
			},
			{
				Name:        "environment_arn",
				Description: "The environment's Amazon Resource Name (ARN), which can be used in other API requests that require an ARN.",
				Type:        schema.TypeString,
			},
			{
				Name:        "environment_id",
				Description: "The ID of this environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "environment_name",
				Description: "The name of this environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "health",
				Description: "Describes the health status of the environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "health_status",
				Description: "Returns the health status of the application running in your environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "operations_role",
				Description: "The Amazon Resource Name (ARN) of the environment's operations role.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform_arn",
				Description: "The ARN of the platform version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "load_balancer_domain",
				Description: "The domain name of the LoadBalancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Resources.LoadBalancer.Domain"),
			},
			{
				Name:        "listeners",
				Description: "A list of Listeners used by the LoadBalancer.",
				Type:        schema.TypeJSON,
				Resolver:    resolveElasticbeanstalkEnvironmentListeners,
			},
			{
				Name:        "load_balancer_name",
				Description: "The name of the LoadBalancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Resources.LoadBalancer.LoadBalancerName"),
			},
			{
				Name:        "solution_stack_name",
				Description: "The name of the SolutionStack deployed with this environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The current operational status of the environment:  * Launching: Environment is in the process of initial deployment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "template_name",
				Description: "The name of the configuration template used to originally launch this environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tier_name",
				Description: "The name of this environment tier.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Tier.Name"),
			},
			{
				Name:        "tier_type",
				Description: "The type of this environment tier.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Tier.Type"),
			},
			{
				Name:        "tier_version",
				Description: "The version of this environment tier.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Tier.Version"),
			},
			{
				Name:        "version_label",
				Description: "The application version deployed in this environment.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticbeanstalk_environment_links",
				Description: "A link to another environment, defined in the environment's manifest.",
				Resolver:    fetchElasticbeanstalkEnvironmentLinks,
				Columns: []schema.Column{
					{
						Name:        "environment_id",
						Description: "Unique ID of aws_elasticbeanstalk_environments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "environment_name",
						Description: "The name of the linked environment (the dependency).",
						Type:        schema.TypeString,
					},
					{
						Name:        "link_name",
						Description: "The name of the link.",
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
func fetchElasticbeanstalkEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config elasticbeanstalk.DescribeEnvironmentsInput
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	for {
		response, err := svc.DescribeEnvironments(ctx, &config, func(options *elasticbeanstalk.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Environments
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveElasticbeanstalkEnvironmentListeners(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.EnvironmentDescription)
	if p.Resources == nil || p.Resources.LoadBalancer == nil {
		return nil
	}
	listeners := make(map[int32]*string, len(p.Resources.LoadBalancer.Listeners))
	for _, l := range p.Resources.LoadBalancer.Listeners {
		listeners[l.Port] = l.Protocol
	}
	return resource.Set(c.Name, listeners)
}

func fetchElasticbeanstalkEnvironmentLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(types.EnvironmentDescription)
	res <- p.EnvironmentLinks
	return nil
}
