package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elbv1/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ELBv1Resources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "load_balancers",
			Struct:      &models.ELBv1LoadBalancerWrapper{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_LoadBalancerDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveLoadBalancerARN()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...,
			),
			Relations: []string{
				"LoadBalancerPolicies()",
			},
		},
		{
			SubService:  "load_balancer_policies",
			Struct:      &types.PolicyDescription{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_PolicyDescription.html",
			SkipFields:  []string{"PolicyAttributeDescriptions"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "load_balancer_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "load_balancer_name",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("load_balancer_name")`,
					},
					{
						Name:     "policy_attribute_descriptions",
						Type:     schema.TypeJSON,
						Resolver: `resolveElbv1loadBalancerPolicyAttributeDescriptions`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "elbv1"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("elasticloadbalancing")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
