package route53

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRoute53HostedZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config route53.ListHostedZonesInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	processHostedZonesBundle := func(hostedZones []types.HostedZone) error {
		tagsCfg := &route53.ListTagsForResourcesInput{ResourceType: types.TagResourceTypeHostedzone, ResourceIds: make([]string, 0, len(hostedZones))}
		for i := range hostedZones {
			parsedId := strings.Replace(*hostedZones[i].Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
			hostedZones[i].Id = &parsedId
			tagsCfg.ResourceIds = append(tagsCfg.ResourceIds, parsedId)
		}
		tagsResponse, err := svc.ListTagsForResources(ctx, tagsCfg)
		if err != nil {
			return err
		}
		for _, h := range hostedZones {
			gotHostedZone, err := svc.GetHostedZone(ctx, &route53.GetHostedZoneInput{Id: h.Id})
			if err != nil {
				return err
			}
			var delegationSetId *string
			if gotHostedZone.DelegationSet != nil {
				delegationSetId = gotHostedZone.DelegationSet.Id
			}
			res <- &models.Route53HostedZoneWrapper{
				HostedZone:      h,
				Tags:            client.TagsToMap(getRoute53tagsByResourceID(*h.Id, tagsResponse.ResourceTagSets)),
				DelegationSetId: delegationSetId,
				VPCs:            gotHostedZone.VPCs,
			}
		}
		return nil
	}

	for {
		response, err := svc.ListHostedZones(ctx, &config)
		if err != nil {
			return err
		}

		for i := 0; i < len(response.HostedZones); i += 10 {
			end := i + 10

			if end > len(response.HostedZones) {
				end = len(response.HostedZones)
			}
			zones := response.HostedZones[i:end]
			err := processHostedZonesBundle(zones)
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
func fetchRoute53HostedZoneQueryLoggingConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*models.Route53HostedZoneWrapper)
	svc := meta.(*client.Client).Services().Route53
	config := route53.ListQueryLoggingConfigsInput{HostedZoneId: r.Id}
	for {
		response, err := svc.ListQueryLoggingConfigs(ctx, &config, func(options *route53.Options) {})
		if err != nil {
			return err
		}
		res <- response.QueryLoggingConfigs
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchRoute53HostedZoneResourceRecordSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*models.Route53HostedZoneWrapper)
	svc := meta.(*client.Client).Services().Route53
	config := route53.ListResourceRecordSetsInput{HostedZoneId: r.Id}
	for {
		response, err := svc.ListResourceRecordSets(ctx, &config, func(options *route53.Options) {})
		if err != nil {
			return err
		}

		res <- response.ResourceRecordSets
		if !response.IsTruncated {
			break
		}

		config.StartRecordIdentifier = response.NextRecordIdentifier
		config.StartRecordType = response.NextRecordType
		config.StartRecordName = response.NextRecordName
	}

	return nil
}
func fetchRoute53HostedZoneTrafficPolicyInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*models.Route53HostedZoneWrapper)
	config := route53.ListTrafficPolicyInstancesByHostedZoneInput{HostedZoneId: r.Id}
	svc := meta.(*client.Client).Services().Route53
	for {
		response, err := svc.ListTrafficPolicyInstancesByHostedZone(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.TrafficPolicyInstances
		if aws.ToString(response.TrafficPolicyInstanceNameMarker) == "" {
			break
		}
		config.TrafficPolicyInstanceNameMarker = response.TrafficPolicyInstanceNameMarker
	}
	return nil
}

func getRoute53tagsByResourceID(id string, set []types.ResourceTagSet) []types.Tag {
	for _, s := range set {
		if *s.ResourceId == id {
			return s.Tags
		}
	}
	return nil
}
func resolveRoute53HostedZoneArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	hz := resource.Item.(*models.Route53HostedZoneWrapper)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Route53Service),
		Region:    "",
		AccountID: "",
		Resource:  fmt.Sprintf("hostedzone/%s", aws.ToString(hz.Id)),
	}.String())
}
func resolveRoute53HostedZoneQueryLoggingConfigsArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ql := resource.Item.(types.QueryLoggingConfig)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Route53Service),
		Region:    "",
		AccountID: "",
		Resource:  fmt.Sprintf("queryloggingconfig/%s", aws.ToString(ql.Id)),
	}.String())
}
func resolveRoute53HostedZoneTrafficPolicyInstancesArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	tp := resource.Item.(types.TrafficPolicyInstance)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Route53Service),
		Region:    "",
		AccountID: "",
		Resource:  fmt.Sprintf("trafficpolicyinstance/%s", aws.ToString(tp.Id)),
	}.String())
}
