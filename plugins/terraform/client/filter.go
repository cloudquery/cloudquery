package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

// DeleteLineageSerialFilter will delete duplicate fetches
func DeleteLineageSerialFilter(meta schema.ClientMeta, parent *schema.Resource) []interface{} {
	client := meta.(*Client)
	backend := client.Backend()
	return []interface{}{"lineage", backend.Data.State.Lineage, "serial", backend.Data.State.Serial}
}
