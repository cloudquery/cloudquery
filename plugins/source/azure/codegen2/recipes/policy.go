// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"

func Armpolicy() []Table {
	tables := []Table{
		{
			Name:           "assignments",
			Struct:         &armpolicy.Assignment{},
			ResponseStruct: &armpolicy.AssignmentsClientListResponse{},
			Client:         &armpolicy.AssignmentsClient{},
			ListFunc:       (&armpolicy.AssignmentsClient{}).NewListPager,
			NewFunc:        armpolicy.NewAssignmentsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyAssignments",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Authorization")`,
		},
		{
			Name:           "data_policy_manifests",
			Struct:         &armpolicy.DataPolicyManifest{},
			ResponseStruct: &armpolicy.DataPolicyManifestsClientListResponse{},
			Client:         &armpolicy.DataPolicyManifestsClient{},
			ListFunc:       (&armpolicy.DataPolicyManifestsClient{}).NewListPager,
			NewFunc:        armpolicy.NewDataPolicyManifestsClient,
			URL:            "/providers/Microsoft.Authorization/dataPolicyManifests",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Authorization")`,
		},
		{
			Name:           "definitions",
			Struct:         &armpolicy.Definition{},
			ResponseStruct: &armpolicy.DefinitionsClientListResponse{},
			Client:         &armpolicy.DefinitionsClient{},
			ListFunc:       (&armpolicy.DefinitionsClient{}).NewListPager,
			NewFunc:        armpolicy.NewDefinitionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Authorization")`,
		},
		{
			Name:           "exemptions",
			Struct:         &armpolicy.Exemption{},
			ResponseStruct: &armpolicy.ExemptionsClientListResponse{},
			Client:         &armpolicy.ExemptionsClient{},
			ListFunc:       (&armpolicy.ExemptionsClient{}).NewListPager,
			NewFunc:        armpolicy.NewExemptionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyExemptions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Authorization")`,
		},
		{
			Name:           "set_definitions",
			Struct:         &armpolicy.SetDefinition{},
			ResponseStruct: &armpolicy.SetDefinitionsClientListResponse{},
			Client:         &armpolicy.SetDefinitionsClient{},
			ListFunc:       (&armpolicy.SetDefinitionsClient{}).NewListPager,
			NewFunc:        armpolicy.NewSetDefinitionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Authorization")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armpolicy"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armpolicy()...)
}
