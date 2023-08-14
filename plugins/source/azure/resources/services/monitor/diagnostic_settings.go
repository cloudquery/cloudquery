package monitor

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type diagnosticSettingsWrapper struct {
	*armmonitor.DiagnosticSettingsResource
	ResourceId string
}

func diagnosticSettings() *schema.Table {
	return &schema.Table{
		Name:                 "azure_monitor_diagnostic_settings",
		Resolver:             fetchDiagnosticSettings,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/monitor/diagnostic-settings/list?tabs=HTTP#diagnosticsettingsresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_diagnostic_settings", client.Namespacemicrosoft_insights),
		Transform:            transformers.TransformWithStruct(&diagnosticSettingsWrapper{}, transformers.WithPrimaryKeys("ID"), transformers.WithUnwrapStructFields("DiagnosticSettingsResource")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func isResourceTypeNotSupported(err error) bool {
	var azureErr *azcore.ResponseError
	if errors.As(err, &azureErr) {
		return azureErr != nil && azureErr.ErrorCode == "ResourceTypeNotSupported"
	}
	return false
}

func fetchDiagnosticSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	r := parent.Item.(*armresources.GenericResourceExpanded)
	svc, err := armmonitor.NewDiagnosticSettingsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(*r.ID, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			if isResourceTypeNotSupported(err) {
				break
			}
			return err
		}
		for _, ds := range p.Value {
			res <- diagnosticSettingsWrapper{
				ds,
				*r.ID,
			}
		}
	}
	return nil
}
