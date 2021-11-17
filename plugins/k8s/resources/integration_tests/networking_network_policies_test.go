package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationNetworkPolicies(t *testing.T) {
	schema := resources.NetworkingNetworkPolicies()
	k8sTestIntegrationHelper(t, schema, nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: schema.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("network-policy-%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{

				Count: 1,
				Data: map[string]interface{}{
					"name":        fmt.Sprintf("network-policy-%s%s", res.Prefix, res.Suffix),
					"labels":      nil,
					"annotations": nil,
					"policy_types": []interface{}{
						"Ingress",
						"Egress",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "k8s_networking_network_policy_ingress",
					ForeignKeyName: "network_policy_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data:  map[string]interface{}{},
						},
					},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "k8s_networking_network_policy_ingress_ports",
							ForeignKeyName: "network_policy_ingress_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"protocol":     "UDP",
										"port_type":    float64(0),
										"port_int_val": float64(8125),
									},
								},
							},
						},
						{
							Name:           "k8s_networking_network_policy_ingress_ports",
							ForeignKeyName: "network_policy_ingress_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"protocol":     "TCP",
										"port_type":    float64(1),
										"port_str_val": "http",
									},
								},
							},
						},
						{
							Name:           "k8s_networking_network_policy_ingress_from",
							ForeignKeyName: "network_policy_ingress_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"pod_selector_match_labels":            nil,
										"pod_selector_match_expressions":       nil,
										"namespace_selector_match_labels":      map[string]interface{}{"name": "default"},
										"namespace_selector_match_expressions": nil,
										"ip_block_cidr":                        "",
										"ip_block_except":                      nil,
									},
								},
							},
						},
						{
							Name:           "k8s_networking_network_policy_ingress_from",
							ForeignKeyName: "network_policy_ingress_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"pod_selector_match_labels":            nil,
										"pod_selector_match_expressions":       nil,
										"namespace_selector_match_labels":      nil,
										"namespace_selector_match_expressions": nil,
										"ip_block_cidr":                        "10.0.0.0/8",
										"ip_block_except": []interface{}{
											"10.0.0.0/24",
											"10.0.1.0/24",
										},
									},
								},
							},
						},
					},
				},
			},
		}
	})
}
