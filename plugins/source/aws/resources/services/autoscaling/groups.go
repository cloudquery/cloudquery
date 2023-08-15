package autoscaling

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/autoscaling/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Groups() *schema.Table {
	tableName := "aws_autoscaling_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AutoScalingGroup.html`,
		Resolver:    fetchAutoscalingGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "autoscaling"),
		Transform: transformers.TransformWithStruct(&models.AutoScalingGroupWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"ar_ns": "arns"})),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "load_balancers",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveAutoscalingGroupLoadBalancers,
			},
			{
				Name:     "load_balancer_target_groups",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveAutoscalingGroupLoadBalancerTargetGroups,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("AutoScalingGroupARN"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "tags_raw",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
		Relations: []*schema.Table{
			groupScalingPolicies(),
			groupLifecycleHooks(),
		},
	}
}

func fetchAutoscalingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	processGroupsBundle := func(groups []types.AutoScalingGroup) error {
		input := autoscaling.DescribeNotificationConfigurationsInput{
			MaxRecords: aws.Int32(100),
		}
		for _, h := range groups {
			input.AutoScalingGroupNames = append(input.AutoScalingGroupNames, *h.AutoScalingGroupName)
		}
		var configurations []types.NotificationConfiguration
		paginator := autoscaling.NewDescribeNotificationConfigurationsPaginator(svc, &input)
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *autoscaling.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			configurations = append(configurations, page.NotificationConfigurations...)
		}
		for _, gr := range groups {
			wrapper := models.AutoScalingGroupWrapper{
				AutoScalingGroup:           gr,
				NotificationConfigurations: getNotificationConfigurationByGroupName(*gr.AutoScalingGroupName, configurations),
			}
			res <- wrapper
		}
		return nil
	}

	config := autoscaling.DescribeAutoScalingGroupsInput{}
	paginator := autoscaling.NewDescribeAutoScalingGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *autoscaling.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		groups := page.AutoScalingGroups
		for i := 0; i < len(groups); i += 255 {
			end := i + 255

			if end > len(groups) {
				end = len(groups)
			}
			t := groups[i:end]
			err := processGroupsBundle(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func resolveAutoscalingGroupLoadBalancers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLoadBalancersInput{AutoScalingGroupName: p.AutoScalingGroupName}
	j := map[string]any{}
	// No paginator available
	for {
		output, err := svc.DescribeLoadBalancers(ctx, &config, func(options *autoscaling.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return err
		}
		for _, lb := range output.LoadBalancers {
			j[*lb.LoadBalancerName] = *lb.State
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return resource.Set(c.Name, j)
}
func resolveAutoscalingGroupLoadBalancerTargetGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLoadBalancerTargetGroupsInput{AutoScalingGroupName: p.AutoScalingGroupName}
	j := map[string]any{}
	// No paginator available
	for {
		output, err := svc.DescribeLoadBalancerTargetGroups(ctx, &config, func(options *autoscaling.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return err
		}
		for _, lb := range output.LoadBalancerTargetGroups {
			j[*lb.LoadBalancerTargetGroupARN] = *lb.State
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return resource.Set(c.Name, j)
}

func getNotificationConfigurationByGroupName(name string, set []types.NotificationConfiguration) []types.NotificationConfiguration {
	var response []types.NotificationConfiguration
	for _, s := range set {
		if *s.AutoScalingGroupName == name {
			response = append(response, s)
		}
	}
	return response
}
