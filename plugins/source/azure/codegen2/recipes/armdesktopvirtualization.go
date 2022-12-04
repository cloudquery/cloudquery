// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"

func Armdesktopvirtualization() []Table {
	tables := []Table{
		{
      Name: "msix_package",
      Struct: &armdesktopvirtualization.MSIXPackage{},
      ResponseStruct: &armdesktopvirtualization.MSIXPackagesClientListResponse{},
      Client: &armdesktopvirtualization.MSIXPackagesClient{},
      ListFunc: (&armdesktopvirtualization.MSIXPackagesClient{}).NewListPager,
			NewFunc: armdesktopvirtualization.NewMSIXPackagesClient,
		},
		{
      Name: "session_host",
      Struct: &armdesktopvirtualization.SessionHost{},
      ResponseStruct: &armdesktopvirtualization.SessionHostsClientListResponse{},
      Client: &armdesktopvirtualization.SessionHostsClient{},
      ListFunc: (&armdesktopvirtualization.SessionHostsClient{}).NewListPager,
			NewFunc: armdesktopvirtualization.NewSessionHostsClient,
		},
		{
      Name: "application",
      Struct: &armdesktopvirtualization.Application{},
      ResponseStruct: &armdesktopvirtualization.ApplicationsClientListResponse{},
      Client: &armdesktopvirtualization.ApplicationsClient{},
      ListFunc: (&armdesktopvirtualization.ApplicationsClient{}).NewListPager,
			NewFunc: armdesktopvirtualization.NewApplicationsClient,
		},
		{
      Name: "start_menu_item",
      Struct: &armdesktopvirtualization.StartMenuItem{},
      ResponseStruct: &armdesktopvirtualization.StartMenuItemsClientListResponse{},
      Client: &armdesktopvirtualization.StartMenuItemsClient{},
      ListFunc: (&armdesktopvirtualization.StartMenuItemsClient{}).NewListPager,
			NewFunc: armdesktopvirtualization.NewStartMenuItemsClient,
		},
		{
      Name: "user_session",
      Struct: &armdesktopvirtualization.UserSession{},
      ResponseStruct: &armdesktopvirtualization.UserSessionsClientListResponse{},
      Client: &armdesktopvirtualization.UserSessionsClient{},
      ListFunc: (&armdesktopvirtualization.UserSessionsClient{}).NewListPager,
			NewFunc: armdesktopvirtualization.NewUserSessionsClient,
		},
		{
      Name: "host_pool",
      Struct: &armdesktopvirtualization.HostPool{},
      ResponseStruct: &armdesktopvirtualization.HostPoolsClientListResponse{},
      Client: &armdesktopvirtualization.HostPoolsClient{},
      ListFunc: (&armdesktopvirtualization.HostPoolsClient{}).NewListPager,
			NewFunc: armdesktopvirtualization.NewHostPoolsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armdesktopvirtualization"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdesktopvirtualization()...)
}