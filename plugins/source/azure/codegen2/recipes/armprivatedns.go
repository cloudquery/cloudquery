// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"

func Armprivatedns() []Table {
	tables := []Table{
		{
      Name: "virtual_network_link",
      Struct: &armprivatedns.VirtualNetworkLink{},
      ResponseStruct: &armprivatedns.VirtualNetworkLinksClientListResponse{},
      Client: &armprivatedns.VirtualNetworkLinksClient{},
      ListFunc: (&armprivatedns.VirtualNetworkLinksClient{}).NewListPager,
			NewFunc: armprivatedns.NewVirtualNetworkLinksClient,
		},
		{
      Name: "private_zone",
      Struct: &armprivatedns.PrivateZone{},
      ResponseStruct: &armprivatedns.PrivateZonesClientListResponse{},
      Client: &armprivatedns.PrivateZonesClient{},
      ListFunc: (&armprivatedns.PrivateZonesClient{}).NewListPager,
			NewFunc: armprivatedns.NewPrivateZonesClient,
		},
		{
      Name: "record_set",
      Struct: &armprivatedns.RecordSet{},
      ResponseStruct: &armprivatedns.RecordSetsClientListResponse{},
      Client: &armprivatedns.RecordSetsClient{},
      ListFunc: (&armprivatedns.RecordSetsClient{}).NewListPager,
			NewFunc: armprivatedns.NewRecordSetsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armprivatedns"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armprivatedns()...)
}