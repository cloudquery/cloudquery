package bss

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchBssBillOverview(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	spec := meta.(*client.Client).Spec
	c := meta.(*client.Client)
	billingCycles := getBillingCycles()
	if spec.BillHistoryMonths > 0 {
		billingCycles = append(getHistoryBillingCycles(spec.BillHistoryMonths), billingCycles...)
	}
	for _, billingCycle := range billingCycles {
		request := bssopenapi.CreateQueryBillOverviewRequest()
		request.BillingCycle = billingCycle
		response, err := c.Services().BSS.QueryBillOverview(request)
		if err != nil {
			return err
		}
		if !response.Success {
			code := response.GetHttpStatus()
			return fmt.Errorf("got response status code %d (%v)", code, http.StatusText(code))
		}
		for _, item := range response.Data.Items.Item {
			res <- &BillOverviewModel{
				BillingCycle:          response.Data.BillingCycle,
				AccountID:             response.Data.AccountID,
				AccountName:           response.Data.AccountName,
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
	}
	return nil
}
