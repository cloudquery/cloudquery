//go:generate mockgen -destination=./mocks/logic.go -package=mocks . MonitorDiagnosticSettingsClient,WorkflowsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/monitor/mgmt/insights"
	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	"github.com/Azure/go-autorest/autorest"
)

type LogicClient struct {
	DiagnosticSettings MonitorDiagnosticSettingsClient
	Workflows          WorkflowsClient
}

type WorkflowsClient interface {
	ListBySubscription(ctx context.Context, top *int32, filter string) (result logic.WorkflowListResultPage, err error)
}

type MonitorDiagnosticSettingsClient interface {
	List(ctx context.Context, resourceURI string) (result insights.DiagnosticSettingsResourceCollection, err error)
}

func NewLogicClient(subscriptionId string, auth autorest.Authorizer) LogicClient {
	diagnosticSettings := insights.NewDiagnosticSettingsClient(subscriptionId)
	diagnosticSettings.Authorizer = auth
	workflowsSvc := logic.NewWorkflowsClient(subscriptionId)
	workflowsSvc.Authorizer = auth
	return LogicClient{
		DiagnosticSettings: diagnosticSettings,
		Workflows:          workflowsSvc,
	}
}
