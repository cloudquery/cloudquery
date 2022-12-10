// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"

func Armreservations() []Table {
	tables := []Table{
		{
			Service:        "armreservations",
			Name:           "reservation_order",
			Struct:         &armreservations.ReservationOrderResponse{},
			ResponseStruct: &armreservations.ReservationOrderClientListResponse{},
			Client:         &armreservations.ReservationOrderClient{},
			ListFunc:       (&armreservations.ReservationOrderClient{}).NewListPager,
			NewFunc:        armreservations.NewReservationOrderClient,
			URL:            "/providers/Microsoft.Capacity/reservationOrders",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Capacity)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armreservations()...)
}
