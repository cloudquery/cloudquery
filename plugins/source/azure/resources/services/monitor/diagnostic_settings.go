package monitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

type diagnosticSettingsWrapper struct {
	*armmonitor.DiagnosticSettingsResource
	ResourceId string
}

func DiagnosticSettings() *schema.Table {
	return &schema.Table{
		Name:      "azure_monitor_diagnostic_settings",
		Resolver:  fetchDiagnosticSettings,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_diagnostic_settings", client.Namespacemicrosoft_insights),
		Transform: transformers.TransformWithStruct(&diagnosticSettingsWrapper{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
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
	svc, err := armresources.NewClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, r := range p.Value {
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
		}
	}
	return nil
}
