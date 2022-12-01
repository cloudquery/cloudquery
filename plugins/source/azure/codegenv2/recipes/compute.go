package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"

func ComputeResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "disks",
			Struct: &armcompute.Disk{},
			ResponseStruct: &armcompute.DisksClientListResponse{},
			Client: &armcompute.DisksClient{},
			ListFunc: (&armcompute.DisksClient{}).NewListPager,
			NewFunc: armcompute.NewDisksClient,
			OutputField: "Value",
		},
		{
			SubService: "virtual_machines",
			Struct: &armcompute.VirtualMachine{},
			ResponseStruct: &armcompute.VirtualMachinesClientListAllResponse{},
			Client: &armcompute.VirtualMachinesClient{},
			ListFunc: (&armcompute.VirtualMachinesClient{}).NewListAllPager,
			NewFunc: armcompute.NewVirtualMachinesClient,
			OutputField: "Value",
			ListAll: true,
		},
		{
			SubService: "virtual_machine_scale_sets",
			Struct: &armcompute.VirtualMachineScaleSet{},
			ResponseStruct: &armcompute.VirtualMachineScaleSetsClientListAllResponse{},
			Client: &armcompute.VirtualMachineScaleSetsClient{},
			ListFunc: (&armcompute.VirtualMachineScaleSetsClient{}).NewListAllPager,
			NewFunc: armcompute.NewVirtualMachineScaleSetsClient,
			OutputField: "Value",
			ListAll: true,
		},
		// {
		// 	SubService: "availability_sets",
		// 	Struct: &armcompute.AvailabilitySet{},
		// 	ResponseStruct: &armcompute.AvailabilitySetsClientListResponse{},
		// 	Client: &armcompute.AvailabilitySetsClient{},
		// 	ListFunc: (&armcompute.AvailabilitySetsClient{}).NewListPager,
		// 	NewFunc: armcompute.NewAvailabilitySetsClient,
		// 	OutputField: "Value",
		// 	ListAll: true,
		// },
		// {
		// 	SubService: "capacity_reservation_groups",
		// 	Struct: &armcompute.CapacityReservationGroup{},
		// 	ResponseStruct: &armcompute.CapacityReservationGroupsClientListBySubscriptionResponse{},
		// 	Client: &armcompute.CapacityReservationGroupsClient{},
		// 	ListFunc: (&armcompute.CapacityReservationGroupsClient{}).NewListBySubscriptionPager,
		// 	NewFunc: armcompute.NewCapacityReservationGroupsClient,
		// 	OutputField: "Value",
		// },
		// {
		// 	SubService: "capacity_reservations",
		// 	Struct: &armcompute.CapacityReservation{},
		// 	ResponseStruct: &armcompute.CapacityReservationsClientListByCapacityReservationGroupResponse{},
		// 	Client: &armcompute.CapacityReservationsClient{},
		// 	ListFunc: (&armcompute.CapacityReservationsClient{}).NewListByCapacityReservationGroupPager,
		// 	NewFunc: armcompute.NewCapacityReservationsClient,
		// 	OutputField: "Value",
		// },
		// {
		// 	SubService: "cloud_services",
		// 	Struct: &armcompute.CloudService{},
		// 	ResponseStruct: &armcompute.CloudServicesClientListAllOptions{},
		// 	Client: &armcompute.CloudServicesClient{},
		// 	ListFunc: (&armcompute.CloudServicesClient{}).NewListAllPager,
		// 	NewFunc: armcompute.NewCloudServicesClient,
		// 	OutputField: "Value",
		// },
	}

	for _, r := range resources {
		r.ImportPath = "compute/armcompute"
		r.Service = "armcompute"
		r.Template = "list"
	}

	return resources
}