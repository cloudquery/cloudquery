package elasticbeanstalk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ElasticbeanstalkEnvironments() *schema.Table {
	return &schema.Table{
		Name:          "aws_elasticbeanstalk_environments",
		Description:   "Describes the properties of an environment.",
		Resolver:      fetchElasticbeanstalkEnvironments,
		Multiplex:     client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
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
				Name:        "tags",
				Type:        schema.TypeJSON,
				Description: "Any tags assigned to the resource",
				Resolver:    resolveElasticbeanstalkEnvironmentTags,
			},
			{
				Name:        "abortable_operation_in_progress",
				Description: "Indicates if there is an in-progress environment configuration update or application version deployment that you can cancel",
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
				Description: "For load-balanced, autoscaling environments, the URL to the LoadBalancer",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EndpointURL"),
			},
			{
				Name:        "arn",
				Description: "The environment's Amazon Resource Name (ARN), which can be used in other API requests that require an ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EnvironmentArn"),
			},
			{
				Name:        "id",
				Description: "The ID of this environment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EnvironmentId"),
			},
			{
				Name:        "name",
				Description: "The name of this environment.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EnvironmentName"),
			},
			{
				Name:        "health",
				Description: "Describes the health status of the environment",
				Type:        schema.TypeString,
			},
			{
				Name:        "health_status",
				Description: "Returns the health status of the application running in your environment",
				Type:        schema.TypeString,
			},
			{
				Name:        "operations_role",
				Description: "The Amazon Resource Name (ARN) of the environment's operations role",
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
				Description: "The current operational status of the environment:  * Launching: Environment is in the process of initial deployment.  * Updating: Environment is in the process of updating its configuration settings or application version.  * Ready: Environment is available to have an action performed on it, such as update or terminate.  * Terminating: Environment is in the shut-down process.  * Terminated: Environment is not running.",
				Type:        schema.TypeString,
			},
			{
				Name:        "template_name",
				Description: "The name of the configuration template used to originally launch this environment.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tier_name",
				Description: "The name of this environment tier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Tier.Name"),
			},
			{
				Name:        "tier_type",
				Description: "The type of this environment tier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Tier.Type"),
			},
			{
				Name:        "tier_version",
				Description: "The version of this environment tier",
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
				Name:          "aws_elasticbeanstalk_configuration_settings",
				Description:   "Describes the settings for a configuration set.",
				Resolver:      fetchElasticbeanstalkConfigurationSettings,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"environment_cq_id", "application_arn"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "environment_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticbeanstalk_environments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "application_name",
						Description: "The name of the application associated with this configuration set.",
						Type:        schema.TypeString,
					},
					{
						Name:        "application_arn",
						Description: "The arn of the associated application.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ApplicationArn"),
					},
					{
						Name:        "date_created",
						Description: "The date (in UTC time) when this configuration set was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "date_updated",
						Description: "The date (in UTC time) when this configuration set was last modified.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "deployment_status",
						Description: "If this configuration set is associated with an environment, the DeploymentStatus parameter indicates the deployment status of this configuration set:  * null: This configuration is not associated with a running environment.  * pending: This is a draft configuration that is not deployed to the associated environment but is in the process of deploying.  * deployed: This is the configuration that is currently deployed to the associated running environment.  * failed: This is a draft configuration that failed to successfully deploy.",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "Describes this configuration set.",
						Type:        schema.TypeString,
					},
					{
						Name:        "environment_name",
						Description: "If not null, the name of the environment for this configuration set.",
						Type:        schema.TypeString,
					},
					{
						Name:        "platform_arn",
						Description: "The ARN of the platform version.",
						Type:        schema.TypeString,
					},
					{
						Name:        "solution_stack_name",
						Description: "The name of the solution stack this configuration set uses.",
						Type:        schema.TypeString,
					},
					{
						Name:        "template_name",
						Description: "If not null, the name of the configuration template for this configuration set.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_elasticbeanstalk_configuration_setting_options",
						Description:   "A specification identifying an individual configuration option along with its current value",
						Resolver:      fetchElasticbeanstalkConfigurationSettingOptionSettings,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "configuration_setting_cq_id",
								Description: "Unique CloudQuery ID of aws_elasticbeanstalk_configuration_setting_options table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "namespace",
								Description: "A unique namespace that identifies the option's associated AWS resource.",
								Type:        schema.TypeString,
							},
							{
								Name:        "option_name",
								Description: "The name of the configuration option.",
								Type:        schema.TypeString,
							},
							{
								Name:        "resource_name",
								Description: "A unique resource name for the option setting",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The current value for the configuration option.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:          "aws_elasticbeanstalk_configuration_options",
				Description:   "Describes the possible values for a configuration option.",
				Resolver:      fetchElasticbeanstalkConfigurationOptions,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"environment_cq_id", "application_arn", "name", "namespace"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "environment_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticbeanstalk_environments table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "application_arn",
						Description: "The arn of the associated application.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ApplicationArn"),
					}, {
						Name:        "name",
						Description: "The name of the configuration option.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Name"),
					},
					{
						Name:        "namespace",
						Description: "A unique namespace identifying the option's associated AWS resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Namespace"),
					},
					{
						Name: "change_severity",
						Description: `An indication of which action is required if the value for this configuration option changes:
						* NoInterruption : There is no interruption to the environment or application availability.
						* RestartEnvironment : The environment is entirely restarted, all AWS resources are deleted and recreated, and the environment is unavailable during the process.
						* RestartApplicationServer : The environment is available the entire time`,
						Type: schema.TypeString,
					},
					{
						Name:        "default_value",
						Description: "The default value for this configuration option.",
						Type:        schema.TypeString,
					},
					{
						Name:        "max_length",
						Description: "If specified, the configuration option must be a string value no longer than this value.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "max_value",
						Description: "If specified, the configuration option must be a numeric value less than this value.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "min_value",
						Description: "If specified, the configuration option must be a numeric value greater than this value.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "regex_label",
						Description: "A unique name representing this regular expression.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Regex.Label"),
					},
					{
						Name:        "regex_pattern",
						Description: "The regular expression pattern that a string configuration option value with this restriction must match.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Regex.Pattern"),
					},
					{
						Name:        "user_defined",
						Description: "An indication of whether the user defined this configuration option:  * true : This configuration option was defined by the user",
						Type:        schema.TypeBool,
					},
					{
						Name:        "value_options",
						Description: "If specified, values for the configuration option are selected from this list.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "value_type",
						Description: "An indication of which type of values this option has and whether it is allowable to select one or more than one of the possible values:  * Scalar : Values for this option are a single selection from the possible values, or an unformatted string, or numeric value governed by the MIN/MAX/Regex constraints.  * List : Values for this option are multiple selections from the possible values.  * Boolean : Values for this option are either true or false .  * Json : Values for this option are a JSON representation of a ConfigDocument.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_elasticbeanstalk_environment_links",
				Description:   "A link to another environment, defined in the environment's manifest",
				Resolver:      fetchElasticbeanstalkEnvironmentLinks,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"environment_cq_id", "link_name"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "environment_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticbeanstalk_environments table (FK)",
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
func fetchElasticbeanstalkEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elasticbeanstalk.DescribeEnvironmentsInput
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	for {
		response, err := svc.DescribeEnvironments(ctx, &config, func(options *elasticbeanstalk.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Environments
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveElasticbeanstalkEnvironmentTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.EnvironmentDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", resource.Item)
	}
	if p.Resources == nil || p.Resources.LoadBalancer == nil {
		return nil
	}
	listeners := make(map[int32]*string, len(p.Resources.LoadBalancer.Listeners))
	for _, l := range p.Resources.LoadBalancer.Listeners {
		listeners[l.Port] = l.Protocol
	}
	return resource.Set(c.Name, listeners)
}
func resolveElasticbeanstalkEnvironmentListeners(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.EnvironmentDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", resource.Item)
	}
	svc := meta.(*client.Client).Services().ElasticBeanstalk
	tagsOutput, err := svc.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: p.EnvironmentArn,
	}, func(o *elasticbeanstalk.Options) {})
	if err != nil {
		return diag.WrapError(err)
	}
	if len(tagsOutput.ResourceTags) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, s := range tagsOutput.ResourceTags {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}
func fetchElasticbeanstalkEnvironmentLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(types.EnvironmentDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", parent.Item)
	}
	res <- p.EnvironmentLinks
	return nil
}

func fetchElasticbeanstalkConfigurationOptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(types.EnvironmentDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", parent.Item)
	}
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	configOptionsIn := elasticbeanstalk.DescribeConfigurationOptionsInput{
		ApplicationName: p.ApplicationName,
		EnvironmentName: p.EnvironmentName,
	}
	output, err := svc.DescribeConfigurationOptions(ctx, &configOptionsIn, func(options *elasticbeanstalk.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}

	for _, option := range output.Options {
		res <- ConfigOptions{
			option, client.GenerateResourceARN("elasticbeanstalk", "application", *p.ApplicationName, c.Region, c.AccountID),
		}
	}

	return nil
}

type ConfigOptions struct {
	types.ConfigurationOptionDescription
	ApplicationArn string
}

func fetchElasticbeanstalkConfigurationSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(types.EnvironmentDescription)
	if !ok {
		return fmt.Errorf("expected types.EnvironmentDescription but got %T", parent.Item)
	}
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk

	configOptionsIn := elasticbeanstalk.DescribeConfigurationSettingsInput{
		ApplicationName: p.ApplicationName,
		EnvironmentName: p.EnvironmentName,
	}
	output, err := svc.DescribeConfigurationSettings(ctx, &configOptionsIn, func(options *elasticbeanstalk.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}

	for _, option := range output.ConfigurationSettings {
		res <- ConfigSettings{
			option, client.GenerateResourceARN("elasticbeanstalk", "application", *p.ApplicationName, c.Region, c.AccountID),
		}
	}

	return nil
}

type ConfigSettings struct {
	types.ConfigurationSettingsDescription
	ApplicationArn string
}

func fetchElasticbeanstalkConfigurationSettingOptionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	option, ok := parent.Item.(ConfigSettings)
	if !ok {
		meta.Logger().Error("parent.Item", "Item", parent.Item)
		return fmt.Errorf("not %T", option)
	}
	for _, t := range option.OptionSettings {
		res <- t
	}

	return nil
}
