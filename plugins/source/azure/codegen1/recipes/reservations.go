// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"

func Armreservations() []*Table {
	tables := []*Table{
		{
			NewFunc: armreservations.NewReservationOrderClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations",
			URL: "/providers/Microsoft.Capacity/reservationOrders",
		},
		{
			NewFunc: armreservations.NewOperationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations",
			URL: "/providers/Microsoft.Capacity/operations",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armreservations())
}