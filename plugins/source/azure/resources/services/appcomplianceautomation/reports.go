package appcomplianceautomation

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Reports() *schema.Table {
	return &schema.Table{
		Name:                 "azure_appcomplianceautomation_reports",
		Resolver:             fetchReports,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/appcompliance/reports/list?tabs=HTTP#reportresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_appcomplianceautomation_reports", client.Namespacemicrosoft_appcomplianceautomation),
		Transform:            transformers.TransformWithStruct(&armappcomplianceautomation.ReportResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchReports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armappcomplianceautomation.NewReportsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
