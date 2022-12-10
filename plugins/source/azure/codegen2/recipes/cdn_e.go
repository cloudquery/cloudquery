package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func ArmcdnE() []Table {
	tables := []Table{
		{
			Name:           "profiles",
			Service: 				"armcdn",
			Struct:         &armcdn.Profile{},
			ResponseStruct: &armcdn.ProfilesClientListResponse{},
			Client:         &armcdn.ProfilesClient{},
			ListFunc:       (&armcdn.ProfilesClient{}).NewListPager,
			NewFunc:        armcdn.NewProfilesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Cdn)`,
			Relations: []*Table{
				{
					Name:           "endpoints",
					Service: 				"armcdn",
					Struct:         &armcdn.Endpoint{},
					ResponseStruct: &armcdn.EndpointsClientListByProfileResponse{},
					Client:         &armcdn.EndpointsClient{},
					ListFunc:       (&armcdn.EndpointsClient{}).NewListByProfilePager,
					NewFunc:        armcdn.NewEndpointsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints",
					SkipFetch: 		true,
				},
			},
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, ArmcdnE()...)
}
