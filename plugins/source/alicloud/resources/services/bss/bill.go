package bss

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Bill() *schema.Table {
	return &schema.Table{
		Name:        "alicloud_bss_bill",
		Resolver:    fetchBssBill,
		Description: "https://help.aliyun.com/document_detail/100392.html",
		Transform: transformers.TransformWithStruct(
			&bssopenapi.Item{},
			transformers.WithPrimaryKeys(
				"billing_cycle", "account_id", "product_code",
				"bill_account_id", "product_type", "pip_code",
				"subscription_type", "commodity_code",
			),
		),
		Columns: []schema.Column{},
	}
}
