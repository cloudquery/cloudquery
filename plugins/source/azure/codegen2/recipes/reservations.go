// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"

func Armreservations() []Table {
	tables := []Table{
		{
			Name:           "reservation_order",
			Struct:         &armreservations.ReservationOrderResponse{},
			ResponseStruct: &armreservations.ReservationOrderClientListResponse{},
			Client:         &armreservations.ReservationOrderClient{},
			ListFunc:       (&armreservations.ReservationOrderClient{}).NewListPager,
			NewFunc:        armreservations.NewReservationOrderClient,
			URL:            "/providers/Microsoft.Capacity/reservationOrders",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Capacity")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armreservations"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armreservations()...)
}
