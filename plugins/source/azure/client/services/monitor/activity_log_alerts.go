// Code generated by codegen; DO NOT EDIT.
package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

type (
	RuntimePagerArmmonitorActivityLogAlertsClientListByResourceGroupResponse  = runtime.Pager[armmonitor.ActivityLogAlertsClientListByResourceGroupResponse]
	RuntimePagerArmmonitorActivityLogAlertsClientListBySubscriptionIDResponse = runtime.Pager[armmonitor.ActivityLogAlertsClientListBySubscriptionIDResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/monitor/activity_log_alerts.go -source=activity_log_alerts.go ActivityLogAlertsClient
type ActivityLogAlertsClient interface {
	Get(context.Context, string, string, *armmonitor.ActivityLogAlertsClientGetOptions) (armmonitor.ActivityLogAlertsClientGetResponse, error)
	NewListByResourceGroupPager(string, *armmonitor.ActivityLogAlertsClientListByResourceGroupOptions) *RuntimePagerArmmonitorActivityLogAlertsClientListByResourceGroupResponse
	NewListBySubscriptionIDPager(*armmonitor.ActivityLogAlertsClientListBySubscriptionIDOptions) *RuntimePagerArmmonitorActivityLogAlertsClientListBySubscriptionIDResponse
}
