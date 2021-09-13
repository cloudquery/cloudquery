package integration_tests

//todo fix resource
//func TestIntegrationComputeInterconnects(t *testing.T) {
//	testIntegrationHelper(t, resources.ComputeInterconnects(), []string{"gcp_compute_interconnects.tf", "network.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: resources.ComputeInterconnects().Name,
//			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq.Where(squirrel.Like{"name": fmt.Sprintf("ssl-policies-policy-%s%s", res.Prefix, res.Suffix)})
//			},
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					Data: map[string]interface{}{
//						"name":            fmt.Sprintf("ssl-policies-policy-%s%s", res.Prefix, res.Suffix),
//						"description":     "",
//						"kind":            "compute#sslPolicy",
//						"min_tls_version": "TLS_1_2",
//						"profile":         "CUSTOM",
//						"custom_features": []interface{}{
//							"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
//							"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
//						},
//					},
//				},
//			},
//		}
//	})
//}
