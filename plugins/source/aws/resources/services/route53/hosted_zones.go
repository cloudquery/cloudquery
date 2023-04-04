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
	"github.com/cloudquery/plugin-sdk/transformers"
)

func HostedZones() *schema.Table {
	tableName := "aws_route53_hosted_zones"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_HostedZone.html`,
		Resolver:    fetchRoute53HostedZones,
		Transform: transformers.TransformWithStruct(
			&models.Route53HostedZoneWrapper{},
			transformers.WithUnwrapStructFields("HostedZone"),
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"vp_cs": "vpcs"})),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveRoute53HostedZoneArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			hostedZoneQueryLoggingConfigs(),
			hostedZoneResourceRecordSets(),
			hostedZoneTrafficPolicyInstances(),
		},
	}
}

func fetchRoute53HostedZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
				Tags:            client.TagsToMap(getTags(*h.Id, tagsResponse.ResourceTagSets)),
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
