// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"

func Armeducation() []Table {
	tables := []Table{
		{
      Name: "join_request_details",
      Struct: &armeducation.JoinRequestDetails{},
      ResponseStruct: &armeducation.JoinRequestsClientListResponse{},
      Client: &armeducation.JoinRequestsClient{},
      ListFunc: (&armeducation.JoinRequestsClient{}).NewListPager,
			NewFunc: armeducation.NewJoinRequestsClient,
		},
		{
      Name: "lab_details",
      Struct: &armeducation.LabDetails{},
      ResponseStruct: &armeducation.LabsClientListResponse{},
      Client: &armeducation.LabsClient{},
      ListFunc: (&armeducation.LabsClient{}).NewListPager,
			NewFunc: armeducation.NewLabsClient,
		},
		{
      Name: "student_details",
      Struct: &armeducation.StudentDetails{},
      ResponseStruct: &armeducation.StudentsClientListResponse{},
      Client: &armeducation.StudentsClient{},
      ListFunc: (&armeducation.StudentsClient{}).NewListPager,
			NewFunc: armeducation.NewStudentsClient,
		},
		{
      Name: "grant_details",
      Struct: &armeducation.GrantDetails{},
      ResponseStruct: &armeducation.GrantsClientListResponse{},
      Client: &armeducation.GrantsClient{},
      ListFunc: (&armeducation.GrantsClient{}).NewListPager,
			NewFunc: armeducation.NewGrantsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armeducation"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armeducation()...)
}