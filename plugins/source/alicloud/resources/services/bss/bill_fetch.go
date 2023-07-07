package bss

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	maxLimit = 100
)

/*
* https://help.aliyun.com/document_detail/100392.html
 */
func fetchBssBill(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	billingCycles := getBillingCycles()
	if c.Spec.BillHistoryMonths > 0 {
		billingCycles = append(getHistoryBillingCycles(c.Spec.BillHistoryMonths), billingCycles...)
	}
	for _, billingCycle := range billingCycles {
		request := bssopenapi.CreateQueryBillRequest()
		request.BillingCycle = billingCycle
		pageNum := 1
		total := 0
		request.PageNum = requests.NewInteger(pageNum)
		request.PageSize = requests.NewInteger(maxLimit)
		for {
			response, err := c.Services().BSS.QueryBill(request)
			if err != nil {
				return err
			}
			if !response.Success {
				code := response.GetHttpStatus()
				return fmt.Errorf("got response status code %d (%v)", code, http.StatusText(code))
			}
			for _, item := range response.Data.Items.Item {
				res <- &BillModel{
					BillingCycle:          response.Data.BillingCycle,
					AccountID:             response.Data.AccountID,
					AccountName:           response.Data.AccountName,
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

func getBillingCycles() []string {
	var months []string
	curMonth := time.Now().Format("2006-01")
	months = append(months, curMonth)
	lastMonth := time.Now().AddDate(0, 0, -5).Format("2006-01")
	if lastMonth != curMonth {
		months = append(months, lastMonth)
	}
	return months
}

func getHistoryBillingCycles(history int) []string {
	var months []string
	for month := 1; month <= history; month++ {
		months = append(months, time.Now().AddDate(0, -month, 0).Format("2006-01"))
	}
	return months
}
