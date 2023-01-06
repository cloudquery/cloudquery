package models

type BillOverview struct {
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

type Bill struct {
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
