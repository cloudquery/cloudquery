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
		Resolver:     fetchElasticbeanstalkEnvironments,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "abortable_operation_in_progress",
				Type: schema.TypeBool,
			},
			{
				Name: "application_name",
				Type: schema.TypeString,
			},
			{
				Name:     "cname",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CNAME"),
			},
			{
				Name: "date_created",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "date_updated",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "endpoint_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EndpointURL"),
			},
			{
				Name: "environment_arn",
				Type: schema.TypeString,
			},
			{
				Name: "environment_id",
				Type: schema.TypeString,
			},
			{
				Name: "environment_name",
				Type: schema.TypeString,
			},
			{
				Name: "health",
				Type: schema.TypeString,
			},
			{
				Name: "health_status",
				Type: schema.TypeString,
			},
			{
				Name: "operations_role",
				Type: schema.TypeString,
			},
			{
				Name: "platform_arn",
				Type: schema.TypeString,
			},
			{
				Name:     "load_balancer_domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Resources.LoadBalancer.Domain"),
			},
			{
				Name:     "load_balancer_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Resources.LoadBalancer.LoadBalancerName"),
			},
			{
				Name: "solution_stack_name",
				Type: schema.TypeString,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "template_name",
				Type: schema.TypeString,
			},
			{
				Name:     "tier_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier.Name"),
			},
			{
				Name:     "tier_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier.Type"),
			},
			{
				Name:     "tier_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier.Version"),
			},
			{
				Name: "version_label",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_elasticbeanstalk_environment_links",
				Resolver: fetchElasticbeanstalkEnvironmentLinks,
				Columns: []schema.Column{
					{
						Name:     "environment_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "environment_name",
						Type: schema.TypeString,
					},
					{
						Name: "link_name",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_elasticbeanstalk_env_resources_load_balancer_listeners",
				Resolver: fetchElasticbeanstalkEnvResourcesLoadBalancerListeners,
				Columns: []schema.Column{
					{
						Name:     "environment_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "port",
						Type: schema.TypeInt,
					},
					{
						Name: "protocol",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElasticbeanstalkEnvironments(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
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
func fetchElasticbeanstalkEnvironmentLinks(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(types.EnvironmentDescription)
	res <- p.EnvironmentLinks
	return nil
}
func fetchElasticbeanstalkEnvResourcesLoadBalancerListeners(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(types.EnvironmentDescription)
	if p.Resources != nil && p.Resources.LoadBalancer != nil {
		res <- p.Resources.LoadBalancer.Listeners
	}
	return nil
}
