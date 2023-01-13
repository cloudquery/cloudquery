package bss

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Bill() *schema.Table {
	return &schema.Table{
		Name:        "alicloud_bss_bill",
		Resolver:    fetchBssBill,
		Description: "https://help.aliyun.com/document_detail/100400.html",
		Transform: transformers.TransformWithStruct(
			&BillModel{},
			transformers.WithPrimaryKeys(
				"billing_cycle", "account_id", "product_code",
				"bill_account_id", "product_type", "pip_code",
				"subscription_type", "commodity_code",
			),
		),
		Columns: []schema.Column{},
	}
}

type BillModel struct {
	BillingCycle          string  `json:"BillingCycle"`
	AccountID             string  `json:"AccountID"`
	AccountName           string  `json:"AccountName"`
	ProductName           string  `json:"ProductName"`
	SubOrderId            string  `json:"SubOrderId"`
	DeductedByCashCoupons float64 `json:"DeductedByCashCoupons"`
	PaymentTime           string  `json:"PaymentTime"`
	PaymentAmount         float64 `json:"PaymentAmount"`
	DeductedByPrepaidCard float64 `json:"DeductedByPrepaidCard"`
	InvoiceDiscount       float64 `json:"InvoiceDiscount"`
	UsageEndTime          string  `json:"UsageEndTime"`
	Item                  string  `json:"Item"`
	SubscriptionType      string  `json:"SubscriptionType"`
	PretaxGrossAmount     float64 `json:"PretaxGrossAmount"`
	Currency              string  `json:"Currency"`
	CommodityCode         string  `json:"CommodityCode"`
	UsageStartTime        string  `json:"UsageStartTime"`
	AdjustAmount          float64 `json:"AdjustAmount"`
	Status                string  `json:"Status"`
	DeductedByCoupons     float64 `json:"DeductedByCoupons"`
	RoundDownDiscount     string  `json:"RoundDownDiscount"`
	ProductDetail         string  `json:"ProductDetail"`
	ProductCode           string  `json:"ProductCode"`
	ProductType           string  `json:"ProductType"`
	OutstandingAmount     float64 `json:"OutstandingAmount"`
	PipCode               string  `json:"PipCode"`
	PretaxAmount          float64 `json:"PretaxAmount"`
	OwnerID               string  `json:"OwnerID"`
	RecordID              string  `json:"RecordID"`
	CashAmount            float64 `json:"CashAmount"`
}
