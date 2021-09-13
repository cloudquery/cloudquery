package integration_tests

//todo fix resource
//func TestIntegrationComputeVpnGateways(t *testing.T) {
//	testIntegrationHelper(t, resources.ComputeVpnGateways(), []string{
//		"network.tf",
//	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: resources.ComputeVpnGateways().Name,
//			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("ssl-proxy-%s%s", res.Prefix, res.Suffix)})
//			},
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					Data: map[string]interface{}{
//						"name":         fmt.Sprintf("ssl-proxy-%s%s", res.Prefix, res.Suffix),
//						"kind":         "compute#targetSslProxy",
//						"proxy_header": "NONE",
//					},
//				},
//			},
//		}
//	})
//}
