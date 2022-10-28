package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ELBv2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "load_balancers",
			Struct:      &types.LoadBalancer{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html",
			SkipFields:  []string{"LoadBalancerArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:          "web_acl_arn",
						Type:          schema.TypeString,
						Resolver:      `resolveElbv2loadBalancerWebACLArn`,
						IgnoreInTests: true,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveElbv2loadBalancerTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("LoadBalancerArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"Listeners()",
				"LoadBalancerAttributes()",
			},
		},
		{
			SubService:  "listeners",
			Struct:      &types.Listener{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Listener.html",
			SkipFields:  []string{"ListenerArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ListenerArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveElbv2listenerTags`,
					},
				}...),
			Relations: []string{
				"ListenerCertificates()",
			},
		},
		{
			SubService:  "listener_certificates",
			Struct:      &types.Certificate{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Certificate.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "listener_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "load_balancer_attributes",
			Struct:      &types.LoadBalancerAttribute{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancerAttribute.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "load_balancer_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "target_groups",
			Struct:      &types.TargetGroup{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html",
			SkipFields:  []string{"TargetGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveElbv2targetGroupTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("TargetGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"TargetGroupTargetHealthDescriptions()",
			},
		},
		{
			SubService:  "target_group_target_health_descriptions",
			Struct:      &types.TargetHealthDescription{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetHealthDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "target_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "elbv2"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("elasticloadbalancing")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
