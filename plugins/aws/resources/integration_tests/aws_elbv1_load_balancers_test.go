package integration_tests

//
//import (
//	"testing"
//
//	"github.com/cloudquery/cq-provider-aws/resources"
//
//	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
//)
//
//func TestIntegrationElbv1LoadBalancers(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.Elbv1LoadBalancers(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_elbv1_load_balancers",
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					Data: map[string]interface{}{
//						"attributes_access_log_enabled":                false,
//						"attributes_connection_settings_idle_timeout":  float64(400),
//						"attributes_connection_draining_timeout":       float64(400),
//						"attributes_cross_zone_load_balancing_enabled": true,
//						"attributes_connection_draining_enabled":       true,
//						"health_check_healthy_threshold":               float64(2),
//						"health_check_interval":                        float64(30),
//						"scheme":                                       "internet-facing",
//					},
//				},
//			},
//			Relations: []*providertest.ResourceIntegrationVerification{
//				{
//					Name:           "aws_elbv1_load_balancer_listeners",
//					ForeignKeyName: "load_balancer_id",
//					ExpectedValues: []providertest.ExpectedValue{
//						{
//							Count: 1,
//							Data: map[string]interface{}{
//								"listener_instance_port":      float64(8000),
//								"listener_load_balancer_port": float64(80),
//								"listener_protocol":           "HTTP",
//								"listener_instance_protocol":  "HTTP",
//							},
//						},
//					},
//				},
//				{
//					Name:           "aws_elbv1_load_balancer_policies",
//					ForeignKeyName: "load_balancer_id",
//					ExpectedValues: []providertest.ExpectedValue{
//						{
//							Count: 1,
//							Data: map[string]interface{}{
//								"policy_type_name": "SSLNegotiationPolicyType",
//							},
//						},
//					},
//				},
//			},
//		}
//	})
//}
