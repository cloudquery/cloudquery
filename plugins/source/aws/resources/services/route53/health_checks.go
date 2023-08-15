package route53

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func HealthChecks() *schema.Table {
	tableName := "aws_route53_health_checks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheck.html`,
		Resolver:    fetchRoute53HealthChecks,
		Transform:   transformers.TransformWithStruct(&Route53HealthCheckWrapper{}, transformers.WithUnwrapStructFields("HealthCheck")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveHealthCheckArn(),
				PrimaryKey: true,
			},
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Description: `The tags associated with the health check.`,
			},
			{
				Name:     "cloud_watch_alarm_configuration_dimensions",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions,
			},
		},
	}
}

type Route53HealthCheckWrapper struct {
	types.HealthCheck
	Tags map[string]string
}

func fetchRoute53HealthChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config route53.ListHealthChecksInput
	cl := meta.(*client.Client)
	svc := cl.Services().Route53

	processHealthChecksBundle := func(healthChecks []types.HealthCheck) error {
		tagsCfg := &route53.ListTagsForResourcesInput{ResourceType: types.TagResourceTypeHealthcheck, ResourceIds: make([]string, 0, len(healthChecks))}
		for _, h := range healthChecks {
			tagsCfg.ResourceIds = append(tagsCfg.ResourceIds, *h.Id)
		}
		tagsResponse, err := svc.ListTagsForResources(ctx, tagsCfg, func(options *route53.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, h := range healthChecks {
			wrapper := Route53HealthCheckWrapper{
				HealthCheck: h,
				Tags:        client.TagsToMap(getTags(*h.Id, tagsResponse.ResourceTagSets)),
			}
			res <- wrapper
		}
		return nil
	}

	paginator := route53.NewListHealthChecksPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		for i := 0; i < len(page.HealthChecks); i += 10 {
			end := i + 10

			if end > len(page.HealthChecks) {
				end = len(page.HealthChecks)
			}
			zones := page.HealthChecks[i:end]
			err := processHealthChecksBundle(zones)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(Route53HealthCheckWrapper)

	if r.CloudWatchAlarmConfiguration == nil {
		return nil
	}
	tags := map[string]*string{}
	for _, t := range r.CloudWatchAlarmConfiguration.Dimensions {
		tags[*t.Name] = t.Value
	}
	return resource.Set(c.Name, tags)
}

func resolveHealthCheckArn() schema.ColumnResolver {
	return client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
		return []string{"healthcheck", *resource.Item.(Route53HealthCheckWrapper).Id}, nil
	})
}
