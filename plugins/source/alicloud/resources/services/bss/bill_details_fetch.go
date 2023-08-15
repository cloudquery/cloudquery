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

/*
* https://help.aliyun.com/document_detail/100392.html
 */
func fetchBillDetails(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	billingCycles := getDetailsBillingCycles()
	if c.Spec.BillHistoryMonths > 0 {
		billingCycles = append(getHistoryBillingCycles(c.Spec.BillHistoryMonths), billingCycles...)
	}
	for _, billingCycle := range billingCycles {
		billingDates := getDetailsBillingDates(billingCycle)
		for _, billingDate := range billingDates {
			request := bssopenapi.CreateQueryInstanceBillRequest()
			request.BillingCycle = billingCycle
			pageNum := 1
			total := 0
			detailsMaxLimit := 100
			request.PageNum = requests.NewInteger(pageNum)
			request.BillingDate = billingDate
			request.PageSize = requests.NewInteger(detailsMaxLimit)
			request.Granularity = "DAILY"
			for {
				response, err := c.Services().BSS.QueryInstanceBill(request)
				if err != nil {
					return err
				}
				if !response.Success {
					code := response.GetHttpStatus()
					return fmt.Errorf("got response status code %d (%v)", code, http.StatusText(code))
				}
				for _, item := range response.Data.Items.Item {
					res <- &BillDetailsModel{
						BillingCycle:          response.Data.BillingCycle,
						BillingDate:           item.BillingDate,
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
						ResourceGroup:         item.ResourceGroup,
						InstanceID:            item.InstanceID,
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
	}
	return nil
}

func getDetailsBillingCycles() []string {
	var months []string
	curMonth := time.Now().Format("2006-01")
	months = append(months, curMonth)
	lastMonth := time.Now().AddDate(0, 0, -5).Format("2006-01")
	if lastMonth != curMonth {
		months = append(months, lastMonth)
	}
	return months
}

func getDetailsBillingDates(month string) []string {
	var days []string
	firstDate := month + "-01"
	first, _ := time.ParseInLocation("2006-01-02", firstDate, time.Local)
	for curr := first; curr.Month() == first.Month(); curr = curr.AddDate(0, 0, 1) {
		days = append(days, curr.Format("2006-01-02"))
	}
	return days
}
