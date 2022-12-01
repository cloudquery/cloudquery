package recipes

import (
	types "github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
)

func SSOAdminResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "instances",
			Struct:      &types.InstanceMetadata{},
			Description: "https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html",
			Relations: []string{
				"PermissionSets()",
			},
		},
		{
			SubService:          "permission_sets",
			Struct:              &types.PermissionSet{},
			Description:         "https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_PermissionSet.html",
			PreResourceResolver: "getSsoadminPermissionSet",
			Relations: []string{
				"AccountAssignments()",
			},
		},
		{
			SubService:  "account_assignments",
			Struct:      &types.AccountAssignment{},
			Description: "https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html",
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "ssoadmin"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("identitystore")`
	}
	return resources
}
