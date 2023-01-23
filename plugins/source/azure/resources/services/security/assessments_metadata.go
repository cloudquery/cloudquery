package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AssessmentsMetadata() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_assessments_metadata",
		Resolver:    fetchAssessmentsMetadata,
		Description: "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/assessments-metadata/list?tabs=HTTP#securityassessmentmetadata",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_assessments_metadata", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.AssessmentMetadataResponse{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
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

func fetchAssessmentsMetadata(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewAssessmentsMetadataClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
