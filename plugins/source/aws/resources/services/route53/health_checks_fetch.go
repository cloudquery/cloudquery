package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Route53HealthCheckWrapper struct {
	types.HealthCheck
	Tags map[string]string
}

func fetchRoute53HealthChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config route53.ListHealthChecksInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	processHealthChecksBundle := func(healthChecks []types.HealthCheck) error {
		tagsCfg := &route53.ListTagsForResourcesInput{ResourceType: types.TagResourceTypeHealthcheck, ResourceIds: make([]string, 0, len(healthChecks))}
		for _, h := range healthChecks {
			tagsCfg.ResourceIds = append(tagsCfg.ResourceIds, *h.Id)
		}
		tagsResponse, err := svc.ListTagsForResources(ctx, tagsCfg)
		if err != nil {
			return err
		}
		for _, h := range healthChecks {
			wrapper := Route53HealthCheckWrapper{
				HealthCheck: h,
				Tags:        client.TagsToMap(getRoute53tagsByResourceID(*h.Id, tagsResponse.ResourceTagSets)),
			}
			res <- wrapper
		}
		return nil
	}

	for {
		response, err := svc.ListHealthChecks(ctx, &config)
		if err != nil {
			return err
		}

		for i := 0; i < len(response.HealthChecks); i += 10 {
			end := i + 10

			if end > len(response.HealthChecks) {
				end = len(response.HealthChecks)
			}
			zones := response.HealthChecks[i:end]
			err := processHealthChecksBundle(zones)
			if err != nil {
				return err
			}
		}

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
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
