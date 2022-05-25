package route53

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type Route53HostedZoneWrapper struct {
	types.HostedZone
	Tags            map[string]string
	DelegationSetId *string
	VPCs            []types.VPC
}

func Route53HostedZones() *schema.Table {
	return &schema.Table{
		Name:         "aws_route53_hosted_zones",
		Description:  "A complex type that contains general information about the hosted zone.",
		Resolver:     fetchRoute53HostedZones,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the hosted zone.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "arn",
				Description: "Amazon Resource Name (ARN) of the route53 hosted zone.",
				Type:        schema.TypeString,
				Resolver:    resolveRoute53HostedZoneArn,
			},
			{
				Name:          "delegation_set_id",
				Description:   "A complex type that lists the Amazon Route 53 name servers for the specified hosted zone.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "caller_reference",
				Description: "The value that you specified for CallerReference when you created the hosted zone.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID that Amazon Route 53 assigned to the hosted zone when you created it.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "name",
				Description: "The name of the domain.",
				Type:        schema.TypeString,
			},
			{
				Name:        "config_comment",
				Description: "Any comments that you want to include about the hosted zone.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Comment"),
			},
			{
				Name:        "config_private_zone",
				Description: "A value that indicates whether this is a private hosted zone.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Config.PrivateZone"),
			},
			{
				Name:          "linked_service_description",
				Description:   "If the health check or hosted zone was created by another service, an optional description that can be provided by the other service.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LinkedService.Description"),
				IgnoreInTests: true,
			},
			{
				Name:          "linked_service_principal",
				Description:   "If the health check or hosted zone was created by another service, the service that created the resource.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LinkedService.ServicePrincipal"),
				IgnoreInTests: true,
			},
			{
				Name:        "resource_record_set_count",
				Description: "The number of resource record sets in the hosted zone.",
				Type:        schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_route53_hosted_zone_query_logging_configs",
				Description:   "A complex type that contains information about a configuration for DNS query logging.",
				Resolver:      fetchRoute53HostedZoneQueryLoggingConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "hosted_zone_cq_id",
						Description: "Unique CloudQuery ID of aws_route53_hosted_zones table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cloud_watch_logs_log_group_arn",
						Description: "The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The ID for a configuration for DNS query logging.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "Amazon Resource Name (ARN) of the route53 hosted zone query logging config.",
						Type:        schema.TypeString,
						Resolver:    resolveRoute53HostedZoneQueryLoggingConfigsArn,
					},
				},
			},
			{
				Name:        "aws_route53_hosted_zone_resource_record_sets",
				Description: "Information about the resource record set to create or delete.",
				Resolver:    fetchRoute53HostedZoneResourceRecordSets,
				Columns: []schema.Column{
					{
						Name:        "hosted_zone_cq_id",
						Description: "Unique CloudQuery ID of aws_route53_hosted_zones table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "resource_records",
						Type:     schema.TypeStringArray,
						Resolver: resolveRoute53hostedZoneResourceRecordSetResourceRecords,
					},
					{
						Name:        "name",
						Description: "For ChangeResourceRecordSets requests, the name of the record that you want to create, update, or delete.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The DNS record type.",
						Type:        schema.TypeString,
					},
					{
						Name:          "dns_name",
						Description:   "Alias resource record sets only: The value that you specify depends on where you want to route queries: Amazon API Gateway custom regional APIs and edge-optimized APIs Specify the applicable domain name for your API.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("AliasTarget.DNSName"),
						IgnoreInTests: true,
					},
					{
						Name:        "evaluate_target_health",
						Description: "Applies only to alias, failover alias, geolocation alias, latency alias, and weighted alias resource record sets: When EvaluateTargetHealth is true, an alias resource record set inherits the health of the referenced AWS resource, such as an ELB load balancer or another resource record set in the hosted zone.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("AliasTarget.EvaluateTargetHealth"),
					},
					{
						Name:        "failover",
						Description: "Failover resource record sets only: To configure failover, you add the Failover element to two resource record sets.",
						Type:        schema.TypeString,
					},
					{
						Name:          "geo_location_continent_code",
						Description:   "The two-letter code for the continent.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("GeoLocation.ContinentCode"),
						IgnoreInTests: true,
					},
					{
						Name:          "geo_location_country_code",
						Description:   "For geolocation resource record sets, the two-letter code for a country.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("GeoLocation.CountryCode"),
						IgnoreInTests: true,
					},
					{
						Name:          "geo_location_subdivision_code",
						Description:   "For geolocation resource record sets, the two-letter code for a state of the United States.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("GeoLocation.SubdivisionCode"),
						IgnoreInTests: true,
					},
					{
						Name:          "health_check_id",
						Description:   "If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the HealthCheckId element and specify the ID of the applicable health check.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "multi_value_answer",
						Description:   "Multivalue answer resource record sets only: To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer.",
						Type:          schema.TypeBool,
						IgnoreInTests: true,
					},
					{
						Name:        "region",
						Description: "Latency-based resource record sets only: The Amazon EC2 Region where you created the resource that this resource record set refers to.",
						Type:        schema.TypeString,
					},
					{
						Name:          "set_identifier",
						Description:   "Resource record sets that have a routing policy other than simple: An identifier that differentiates among multiple resource record sets that have the same combination of name and type, such as multiple weighted resource record sets named acme.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "ttl",
						Description: "The resource record cache time to live (TTL), in seconds.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("TTL"),
					},
					{
						Name:          "traffic_policy_instance_id",
						Description:   "When you create a traffic policy instance, Amazon Route 53 automatically creates a resource record set.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "weight",
						Description:   "Weighted resource record sets only: Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set.",
						Type:          schema.TypeBigInt,
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:          "aws_route53_hosted_zone_traffic_policy_instances",
				Description:   "A complex type that contains settings for the new traffic policy instance.",
				Resolver:      fetchRoute53HostedZoneTrafficPolicyInstances,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "hosted_zone_cq_id",
						Description: "Unique CloudQuery ID of aws_route53_hosted_zones table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The ID that Amazon Route 53 assigned to the new traffic policy instance.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "message",
						Description: "If State is Failed, an explanation of the reason for the failure.",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The DNS name, such as www.",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "The value of State is one of the following values: Applied Amazon Route 53 has finished creating resource record sets, and changes have propagated to all Route 53 edge locations.",
						Type:        schema.TypeString,
					},
					{
						Name:        "ttl",
						Description: "The TTL that Amazon Route 53 assigned to all of the resource record sets that it created in the specified hosted zone.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("TTL"),
					},
					{
						Name:        "traffic_policy_id",
						Description: "The ID of the traffic policy that Amazon Route 53 used to create resource record sets in the specified hosted zone.",
						Type:        schema.TypeString,
					},
					{
						Name:        "traffic_policy_type",
						Description: "The DNS type that Amazon Route 53 assigned to all of the resource record sets that it created for this traffic policy instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "traffic_policy_version",
						Description: "The version of the traffic policy that Amazon Route 53 used to create resource record sets in the specified hosted zone.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "arn",
						Description: "Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.",
						Type:        schema.TypeString,
						Resolver:    resolveRoute53HostedZoneTrafficPolicyInstancesArn,
					},
				},
			},
			{
				Name:          "aws_route53_hosted_zone_vpc_association_authorizations",
				Description:   "(Private hosted zones only) A complex type that contains information about an Amazon VPC.",
				Resolver:      fetchRoute53HostedZoneVpcAssociationAuthorizations,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "hosted_zone_cq_id",
						Description: "Unique CloudQuery ID of aws_route53_hosted_zones table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "vpc_id",
						Description: "(Private hosted zones only) The ID of an Amazon VPC.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VPCId"),
					},
					{
						Name:        "vpc_region",
						Description: "(Private hosted zones only) The region that an Amazon VPC was created in.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VPCRegion"),
					},
					{
						Name:        "vpc_arn",
						Description: "Amazon Resource Name (ARN) of the ec2 vpc.",
						Type:        schema.TypeString,
						Resolver:    resolveRoute53HostedZoneVpcArn,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
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
			return diag.WrapError(err)
		}
		for _, h := range hostedZones {
			gotHostedZone, err := svc.GetHostedZone(ctx, &route53.GetHostedZoneInput{Id: h.Id})
			if err != nil {
				return diag.WrapError(err)
			}
			var delegationSetId *string
			if gotHostedZone.DelegationSet != nil {
				delegationSetId = gotHostedZone.DelegationSet.Id
			}
			wrapper := Route53HostedZoneWrapper{
				HostedZone:      h,
				Tags:            client.TagsToMap(getRoute53tagsByResourceID(*h.Id, tagsResponse.ResourceTagSets)),
				DelegationSetId: delegationSetId,
				VPCs:            gotHostedZone.VPCs,
			}
			res <- wrapper
		}
		return nil
	}

	for {
		response, err := svc.ListHostedZones(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}

		for i := 0; i < len(response.HostedZones); i += 10 {
			end := i + 10

			if end > len(response.HostedZones) {
				end = len(response.HostedZones)
			}
			zones := response.HostedZones[i:end]
			err := processHostedZonesBundle(zones)
			if err != nil {
				return diag.WrapError(err)
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
	r := parent.Item.(Route53HostedZoneWrapper)
	svc := meta.(*client.Client).Services().Route53
	config := route53.ListQueryLoggingConfigsInput{HostedZoneId: r.Id}
	for {
		response, err := svc.ListQueryLoggingConfigs(ctx, &config, func(options *route53.Options) {})
		if err != nil {
			return diag.WrapError(err)
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
	r := parent.Item.(Route53HostedZoneWrapper)
	svc := meta.(*client.Client).Services().Route53
	config := route53.ListResourceRecordSetsInput{HostedZoneId: r.Id}
	for {
		response, err := svc.ListResourceRecordSets(ctx, &config, func(options *route53.Options) {})
		if err != nil {
			return diag.WrapError(err)
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
func resolveRoute53hostedZoneResourceRecordSetResourceRecords(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ResourceRecordSet)
	recordSets := make([]string, 0, len(r.ResourceRecords))
	for _, t := range r.ResourceRecords {
		recordSets = append(recordSets, *t.Value)
	}
	return resource.Set(c.Name, recordSets)
}
func fetchRoute53HostedZoneTrafficPolicyInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(Route53HostedZoneWrapper)
	config := route53.ListTrafficPolicyInstancesByHostedZoneInput{HostedZoneId: r.Id}
	svc := meta.(*client.Client).Services().Route53
	for {
		response, err := svc.ListTrafficPolicyInstancesByHostedZone(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.TrafficPolicyInstances
		if aws.ToString(response.TrafficPolicyInstanceNameMarker) == "" {
			break
		}
		config.TrafficPolicyInstanceNameMarker = response.TrafficPolicyInstanceNameMarker
	}
	return nil
}
func fetchRoute53HostedZoneVpcAssociationAuthorizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(Route53HostedZoneWrapper)
	res <- r.VPCs
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
	hz := resource.Item.(Route53HostedZoneWrapper)
	return resource.Set(c.Name, cl.PartitionGlobalARN(client.Route53Service, "hostedzone", *hz.Id))
}
func resolveRoute53HostedZoneQueryLoggingConfigsArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ql := resource.Item.(types.QueryLoggingConfig)
	return resource.Set(c.Name, cl.PartitionGlobalARN(client.Route53Service, "queryloggingconfig", *ql.Id))
}
func resolveRoute53HostedZoneTrafficPolicyInstancesArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	tp := resource.Item.(types.TrafficPolicyInstance)
	return resource.Set(c.Name, cl.PartitionGlobalARN(client.Route53Service, "trafficpolicyinstance", *tp.Id))
}
func resolveRoute53HostedZoneVpcArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	vpc := resource.Item.(types.VPC)
	return resource.Set(c.Name, cl.ARN(client.EC2Service, "vpc", *vpc.VPCId))
}
