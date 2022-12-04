// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"

func Armdevtestlabs() []Table {
	tables := []Table{
		{
      Name: "artifact",
      Struct: &armdevtestlabs.Artifact{},
      ResponseStruct: &armdevtestlabs.ArtifactsClientListResponse{},
      Client: &armdevtestlabs.ArtifactsClient{},
      ListFunc: (&armdevtestlabs.ArtifactsClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewArtifactsClient,
		},
		{
      Name: "service_fabric",
      Struct: &armdevtestlabs.ServiceFabric{},
      ResponseStruct: &armdevtestlabs.ServiceFabricsClientListResponse{},
      Client: &armdevtestlabs.ServiceFabricsClient{},
      ListFunc: (&armdevtestlabs.ServiceFabricsClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewServiceFabricsClient,
		},
		{
      Name: "schedule",
      Struct: &armdevtestlabs.Schedule{},
      ResponseStruct: &armdevtestlabs.VirtualMachineSchedulesClientListResponse{},
      Client: &armdevtestlabs.VirtualMachineSchedulesClient{},
      ListFunc: (&armdevtestlabs.VirtualMachineSchedulesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewVirtualMachineSchedulesClient,
		},
		{
      Name: "artifact_source",
      Struct: &armdevtestlabs.ArtifactSource{},
      ResponseStruct: &armdevtestlabs.ArtifactSourcesClientListResponse{},
      Client: &armdevtestlabs.ArtifactSourcesClient{},
      ListFunc: (&armdevtestlabs.ArtifactSourcesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewArtifactSourcesClient,
		},
		{
      Name: "operation_metadata",
      Struct: &armdevtestlabs.OperationMetadata{},
      ResponseStruct: &armdevtestlabs.ProviderOperationsClientListResponse{},
      Client: &armdevtestlabs.ProviderOperationsClient{},
      ListFunc: (&armdevtestlabs.ProviderOperationsClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewProviderOperationsClient,
		},
		{
      Name: "secret",
      Struct: &armdevtestlabs.Secret{},
      ResponseStruct: &armdevtestlabs.SecretsClientListResponse{},
      Client: &armdevtestlabs.SecretsClient{},
      ListFunc: (&armdevtestlabs.SecretsClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewSecretsClient,
		},
		{
      Name: "lab_virtual_machine",
      Struct: &armdevtestlabs.LabVirtualMachine{},
      ResponseStruct: &armdevtestlabs.VirtualMachinesClientListResponse{},
      Client: &armdevtestlabs.VirtualMachinesClient{},
      ListFunc: (&armdevtestlabs.VirtualMachinesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewVirtualMachinesClient,
		},
		{
      Name: "gallery_image",
      Struct: &armdevtestlabs.GalleryImage{},
      ResponseStruct: &armdevtestlabs.GalleryImagesClientListResponse{},
      Client: &armdevtestlabs.GalleryImagesClient{},
      ListFunc: (&armdevtestlabs.GalleryImagesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewGalleryImagesClient,
		},
		{
      Name: "schedule",
      Struct: &armdevtestlabs.Schedule{},
      ResponseStruct: &armdevtestlabs.ServiceFabricSchedulesClientListResponse{},
      Client: &armdevtestlabs.ServiceFabricSchedulesClient{},
      ListFunc: (&armdevtestlabs.ServiceFabricSchedulesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewServiceFabricSchedulesClient,
		},
		{
      Name: "user",
      Struct: &armdevtestlabs.User{},
      ResponseStruct: &armdevtestlabs.UsersClientListResponse{},
      Client: &armdevtestlabs.UsersClient{},
      ListFunc: (&armdevtestlabs.UsersClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewUsersClient,
		},
		{
      Name: "custom_image",
      Struct: &armdevtestlabs.CustomImage{},
      ResponseStruct: &armdevtestlabs.CustomImagesClientListResponse{},
      Client: &armdevtestlabs.CustomImagesClient{},
      ListFunc: (&armdevtestlabs.CustomImagesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewCustomImagesClient,
		},
		{
      Name: "disk",
      Struct: &armdevtestlabs.Disk{},
      ResponseStruct: &armdevtestlabs.DisksClientListResponse{},
      Client: &armdevtestlabs.DisksClient{},
      ListFunc: (&armdevtestlabs.DisksClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewDisksClient,
		},
		{
      Name: "dtl_environment",
      Struct: &armdevtestlabs.DtlEnvironment{},
      ResponseStruct: &armdevtestlabs.EnvironmentsClientListResponse{},
      Client: &armdevtestlabs.EnvironmentsClient{},
      ListFunc: (&armdevtestlabs.EnvironmentsClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewEnvironmentsClient,
		},
		{
      Name: "formula",
      Struct: &armdevtestlabs.Formula{},
      ResponseStruct: &armdevtestlabs.FormulasClientListResponse{},
      Client: &armdevtestlabs.FormulasClient{},
      ListFunc: (&armdevtestlabs.FormulasClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewFormulasClient,
		},
		{
      Name: "virtual_network",
      Struct: &armdevtestlabs.VirtualNetwork{},
      ResponseStruct: &armdevtestlabs.VirtualNetworksClientListResponse{},
      Client: &armdevtestlabs.VirtualNetworksClient{},
      ListFunc: (&armdevtestlabs.VirtualNetworksClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewVirtualNetworksClient,
		},
		{
      Name: "arm_template",
      Struct: &armdevtestlabs.ArmTemplate{},
      ResponseStruct: &armdevtestlabs.ArmTemplatesClientListResponse{},
      Client: &armdevtestlabs.ArmTemplatesClient{},
      ListFunc: (&armdevtestlabs.ArmTemplatesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewArmTemplatesClient,
		},
		{
      Name: "schedule",
      Struct: &armdevtestlabs.Schedule{},
      ResponseStruct: &armdevtestlabs.SchedulesClientListResponse{},
      Client: &armdevtestlabs.SchedulesClient{},
      ListFunc: (&armdevtestlabs.SchedulesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewSchedulesClient,
		},
		{
      Name: "notification_channel",
      Struct: &armdevtestlabs.NotificationChannel{},
      ResponseStruct: &armdevtestlabs.NotificationChannelsClientListResponse{},
      Client: &armdevtestlabs.NotificationChannelsClient{},
      ListFunc: (&armdevtestlabs.NotificationChannelsClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewNotificationChannelsClient,
		},
		{
      Name: "policy",
      Struct: &armdevtestlabs.Policy{},
      ResponseStruct: &armdevtestlabs.PoliciesClientListResponse{},
      Client: &armdevtestlabs.PoliciesClient{},
      ListFunc: (&armdevtestlabs.PoliciesClient{}).NewListPager,
			NewFunc: armdevtestlabs.NewPoliciesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armdevtestlabs"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdevtestlabs()...)
}