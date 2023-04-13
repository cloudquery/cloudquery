package client

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func MultiplexByCustomer(meta schema.ClientMeta) []schema.ClientMeta {
	c := meta.(*Client)
	var res []schema.ClientMeta

	for managerID, customerIDs := range c.customers {
		m := c.withManagerID(managerID)
		for _, id := range customerIDs {
			res = append(res, m.withCustomerID(id))
		}
	}

	return res
}
