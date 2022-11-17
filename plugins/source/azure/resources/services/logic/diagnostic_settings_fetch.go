// Code generated by codegen; DO NOT EDIT.

package logic

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	logic "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDiagnosticSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Monitor

	workflow := parent.Item.(*logic.Workflow)
	id, err := arm.ParseResourceID(*workflow.ID)
	if err != nil {
		return err
	}

	pager := svc.DiagnosticSettingsClient.NewListPager(id.String(), nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
