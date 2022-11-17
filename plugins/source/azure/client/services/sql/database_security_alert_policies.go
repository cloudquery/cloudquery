// Code generated by codegen; DO NOT EDIT.
package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

type (
	RuntimePagerArmsqlDatabaseSecurityAlertPoliciesClientListByDatabaseResponse = runtime.Pager[armsql.DatabaseSecurityAlertPoliciesClientListByDatabaseResponse]
)

//go:generate mockgen -package=mocks -destination=../../mocks/sql/database_security_alert_policies.go -source=database_security_alert_policies.go DatabaseSecurityAlertPoliciesClient
type DatabaseSecurityAlertPoliciesClient interface {
	Get(context.Context, string, string, string, armsql.SecurityAlertPolicyName, *armsql.DatabaseSecurityAlertPoliciesClientGetOptions) (armsql.DatabaseSecurityAlertPoliciesClientGetResponse, error)
	NewListByDatabasePager(string, string, string, *armsql.DatabaseSecurityAlertPoliciesClientListByDatabaseOptions) *RuntimePagerArmsqlDatabaseSecurityAlertPoliciesClientListByDatabaseResponse
}
