package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RegulatoryComplianceStandards() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_regulatory_compliance_standards",
		Resolver:    fetchRegulatoryComplianceStandards,
		Description: "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/regulatory-compliance-standards/list?tabs=HTTP#regulatorycompliancestandard",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_regulatory_compliance_standards", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.RegulatoryComplianceStandard{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchRegulatoryComplianceStandards(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewRegulatoryComplianceStandardsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
