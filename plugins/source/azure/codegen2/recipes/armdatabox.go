// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"

func Armdatabox() []Table {
	tables := []Table{
		{
      Name: "job_resource",
      Struct: &armdatabox.JobResource{},
      ResponseStruct: &armdatabox.JobsClientListResponse{},
      Client: &armdatabox.JobsClient{},
      ListFunc: (&armdatabox.JobsClient{}).NewListPager,
			NewFunc: armdatabox.NewJobsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/jobs",
		},
	}

	for i := range tables {
		tables[i].Service = "armdatabox"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdatabox()...)
}