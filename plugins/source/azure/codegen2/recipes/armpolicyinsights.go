// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"

func Armpolicyinsights() []Table {
	tables := []Table{
		{
      Name: "slim_policy_metadata",
      Struct: &armpolicyinsights.SlimPolicyMetadata{},
      ResponseStruct: &armpolicyinsights.PolicyMetadataClientListResponse{},
      Client: &armpolicyinsights.PolicyMetadataClient{},
      ListFunc: (&armpolicyinsights.PolicyMetadataClient{}).NewListPager,
			NewFunc: armpolicyinsights.NewPolicyMetadataClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armpolicyinsights"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armpolicyinsights()...)
}