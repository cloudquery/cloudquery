package bss

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/codegen/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	maxLimit = 100
)

// See: https://help.aliyun.com/document_detail/100400.html
func fetchBssBillOverview(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	spec := meta.(*client.Client).Spec
	cli, err := bssopenapi.NewClientWithAccessKey(spec.RegionID, spec.AccessKey, spec.SecretKey)
	if err != nil {
		return err
	}

	billingCycles := GetBillingCycles()
	if spec.BillHistory > 0 {
		billingCycles = append(GetHistoryBillingCycles(spec.BillHistory), billingCycles...)
	}
	for _, billingCycle := range billingCycles {
		request := bssopenapi.CreateQueryBillOverviewRequest()
		request.BillingCycle = billingCycle
		response, err := cli.QueryBillOverview(request)
		if err != nil {
			return err
		}
		if response.GetHttpStatus() != http.StatusOK {
			return fmt.Errorf("httpcode %d", response.GetHttpStatus())
		}
		for _, item := range response.Data.Items.Item {
			res <- ToBSSBillOverview(response.Data.BillingCycle, response.Data.AccountID, response.Data.AccountName, item)
		}
	}
	return nil
}

func ToBSSBillOverview(billingCycle, accountId, accountName string, item bssopenapi.Item) *models.BillOverview {
	return &models.BillOverview{
		BillingCycle:          billingCycle,
		AccountID:             accountId,
		AccountName:           accountName,
		DeductedByCoupons:     item.DeductedByCoupons,
		RoundDownDiscount:     item.RoundDownDiscount,
		ProductName:           item.ProductName,
		ProductDetail:         item.ProductDetail,
		ProductCode:           item.ProductCode,
		BillAccountID:         item.BillAccountID,
		ProductType:           item.ProductType,
		DeductedByCashCoupons: item.DeductedByCashCoupons,
		OutstandingAmount:     item.OutstandingAmount,
		BizType:               item.BizType,
		PaymentAmount:         item.PaymentAmount,
		PipCode:               item.PipCode,
		DeductedByPrepaidCard: item.DeductedByPrepaidCard,
		InvoiceDiscount:       item.InvoiceDiscount,
		Item:                  item.Item,
		SubscriptionType:      item.SubscriptionType,
		PretaxGrossAmount:     item.PretaxGrossAmount,
		PretaxAmount:          item.PretaxAmount,
		OwnerID:               item.OwnerID,
		Currency:              item.Currency,
		CommodityCode:         item.CommodityCode,
		BillAccountName:       item.BillAccountName,
		AdjustAmount:          item.AdjustAmount,
		CashAmount:            item.CashAmount,
	}
}

/*
* https://help.aliyun.com/document_detail/100392.html
 */
func fetchBssBill(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	spec := meta.(*client.Client).Spec
	cli, err := bssopenapi.NewClientWithAccessKey(spec.RegionID, spec.AccessKey, spec.SecretKey)
	if err != nil {
		return err
	}
	billingCycles := GetBillingCycles()
	if spec.BillHistory > 0 {
		billingCycles = append(GetHistoryBillingCycles(spec.BillHistory), billingCycles...)
	}
	for _, billingCycle := range billingCycles {
		request := bssopenapi.CreateQueryBillRequest()
		request.BillingCycle = billingCycle
		pageNum := 1
		total := 0
		request.PageNum = requests.NewInteger(pageNum)
		request.PageSize = requests.NewInteger(maxLimit)
		for {
			response, err := cli.QueryBill(request)
			if err != nil {
				return err
			}
			if response.GetHttpStatus() != http.StatusOK {
				return fmt.Errorf("httpcode %d", response.GetHttpStatus())
			}
			for _, item := range response.Data.Items.Item {
				res <- ToBSSBill(response.Data.BillingCycle, response.Data.AccountID, response.Data.AccountName, item)
			}
			total += len(response.Data.Items.Item)
			if len(response.Data.Items.Item) == 0 || total >= response.Data.TotalCount {
				break
			}
			pageNum++
			request.PageNum = requests.NewInteger(pageNum)
		}
	}
	return nil
}

func ToBSSBill(billingCycle, accountId, accountName string, item bssopenapi.Item) *models.Bill {
	return &models.Bill{
		BillingCycle:          billingCycle,
		AccountID:             accountId,
		AccountName:           accountName,
		ProductName:           item.ProductName,
		SubOrderId:            item.SubOrderId,
		DeductedByCashCoupons: item.DeductedByCashCoupons,
		PaymentTime:           item.PaymentTime,
		PaymentAmount:         item.PaymentAmount,
		DeductedByPrepaidCard: item.DeductedByPrepaidCard,
		InvoiceDiscount:       item.InvoiceDiscount,
		UsageEndTime:          item.UsageEndTime,
		Item:                  item.Item,
		SubscriptionType:      item.SubscriptionType,
		PretaxGrossAmount:     item.PretaxGrossAmount,
		Currency:              item.Currency,
		CommodityCode:         item.CommodityCode,
		UsageStartTime:        item.UsageStartTime,
		AdjustAmount:          item.AdjustAmount,
		Status:                item.Status,
		DeductedByCoupons:     item.DeductedByCoupons,
		RoundDownDiscount:     item.RoundDownDiscount,
		ProductDetail:         item.ProductDetail,
		ProductCode:           item.ProductCode,
		ProductType:           item.ProductType,
		OutstandingAmount:     item.OutstandingAmount,
		PipCode:               item.PipCode,
		PretaxAmount:          item.PretaxAmount,
		OwnerID:               item.OwnerID,
		RecordID:              item.RecordID,
		CashAmount:            item.CashAmount,
	}
}

func GetBillingCycles() []string {
	var months []string
	curMonth := time.Now().Format("2006-01")
	months = append(months, curMonth)
	lastMonth := time.Now().AddDate(0, 0, -5).Format("2006-01")
	if lastMonth != curMonth {
		months = append(months, lastMonth)
	}
	return months
}

func GetHistoryBillingCycles(history int) []string {
	var months []string
	for month := 1; month <= history; month++ {
		months = append(months, time.Now().AddDate(0, -month, 0).Format("2006-01"))
	}
	return months
}
