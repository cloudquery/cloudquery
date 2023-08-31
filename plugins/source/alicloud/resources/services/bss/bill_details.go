package bss

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillDetails() *schema.Table {
	return &schema.Table{
		Name:        "alicloud_bss_bill_details",
		Title:       "Alibaba Cloud BSS Bill Details",
		Description: "https://help.aliyun.com/document_detail/100392.html",
		Resolver:    fetchBillDetails,
		Multiplex:   client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(
			&BillDetailsModel{},
			transformers.WithPrimaryKeys(
				"BillingCycle", "BillingDate", "AccountID", "ProductCode",
				"ProductType", "PipCode", "RecordID", "SubscriptionType", "CommodityCode", "InstanceID",
			),
		),
	}
}

type BillDetailsModel struct {
	BillingCycle          string  `json:"BillingCycle"`
	BillingDate           string  `json:"BillingDate"`
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
	ResourceGroup         string  `json:"ResourceGroup"`
	InstanceID            string  `json:"InstanceID"`
	CashAmount            float64 `json:"CashAmount"`
}
