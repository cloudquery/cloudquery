package bss

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillOverview() *schema.Table {
	return &schema.Table{
		Name:        "alicloud_bss_bill_overview",
		Title:       "Alibaba Cloud BSS Bill Overview",
		Description: "https://help.aliyun.com/document_detail/100400.html",
		Multiplex:   client.AccountMultiplex,
		Resolver:    fetchBssBillOverview,
		Transform: transformers.TransformWithStruct(
			&BillOverviewModel{},
			transformers.WithPrimaryKeys(
				"BillingCycle", "AccountID", "ProductCode",
				"BillAccountID", "ProductType", "PipCode",
				"SubscriptionType", "CommodityCode",
			),
		),
		Columns: []schema.Column{},
	}
}

type BillOverviewModel struct {
	BillingCycle          string  `json:"BillingCycle"`
	AccountID             string  `json:"AccountID"`
	AccountName           string  `json:"AccountName"`
	DeductedByCoupons     float64 `json:"DeductedByCoupons"`
	RoundDownDiscount     string  `json:"RoundDownDiscount"`
	ProductName           string  `json:"ProductName"`
	ProductDetail         string  `json:"ProductDetail"`
	ProductCode           string  `json:"ProductCode"`
	BillAccountID         string  `json:"BillAccountID"`
	ProductType           string  `json:"ProductType"`
	DeductedByCashCoupons float64 `json:"DeductedByCashCoupons"`
	OutstandingAmount     float64 `json:"OutstandingAmount"`
	BizType               string  `json:"BizType"`
	PaymentAmount         float64 `json:"PaymentAmount"`
	PipCode               string  `json:"PipCode"`
	DeductedByPrepaidCard float64 `json:"DeductedByPrepaidCard"`
	InvoiceDiscount       float64 `json:"InvoiceDiscount"`
	Item                  string  `json:"Item"`
	SubscriptionType      string  `json:"SubscriptionType"`
	PretaxGrossAmount     float64 `json:"PretaxGrossAmount"`
	PretaxAmount          float64 `json:"PretaxAmount"`
	OwnerID               string  `json:"OwnerID"`
	Currency              string  `json:"Currency"`
	CommodityCode         string  `json:"CommodityCode"`
	BillAccountName       string  `json:"BillAccountName"`
	AdjustAmount          float64 `json:"AdjustAmount"`
	CashAmount            float64 `json:"CashAmount"`
}
