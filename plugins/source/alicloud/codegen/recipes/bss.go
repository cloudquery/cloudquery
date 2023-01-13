package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/codegen/models"
)

func BSS() []*Resource {
	return []*Resource{
		{
			Service:    "bss",
			SubService: "bill_overview",
			Struct:     new(models.BillOverview),
			TableName:  "bss_bill_overview",
			PKColumns:  []string{"billing_cycle", "account_id", "product_code", "bill_account_id", "product_type", "pip_code", "subscription_type", "commodity_code"},
		},
		{
			Service:    "bss",
			SubService: "bill",
			Struct:     new(models.Bill),
			TableName:  "bss_bill",
			PKColumns:  []string{"billing_cycle", "account_id", "product_code", "product_type", "pip_code", "subscription_type", "commodity_code", "record_id"},
		},
	}
}
