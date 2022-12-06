// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital"

func Armorbital() []Table {
	tables := []Table{
		{
			Name:           "contact_profile",
			Struct:         &armorbital.ContactProfile{},
			ResponseStruct: &armorbital.ContactProfilesClientListResponse{},
			Client:         &armorbital.ContactProfilesClient{},
			ListFunc:       (&armorbital.ContactProfilesClient{}).NewListPager,
			NewFunc:        armorbital.NewContactProfilesClient,
			URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/contactProfiles",
		},
		{
			Name:           "spacecraft",
			Struct:         &armorbital.Spacecraft{},
			ResponseStruct: &armorbital.SpacecraftsClientListResponse{},
			Client:         &armorbital.SpacecraftsClient{},
			ListFunc:       (&armorbital.SpacecraftsClient{}).NewListPager,
			NewFunc:        armorbital.NewSpacecraftsClient,
			URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts",
		},
	}

	for i := range tables {
		tables[i].Service = "armorbital"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armorbital()...)
}
