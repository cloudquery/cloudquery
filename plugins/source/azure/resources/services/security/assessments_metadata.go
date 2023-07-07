package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AssessmentsMetadata() *schema.Table {
	return &schema.Table{
		Name:                 "azure_security_assessments_metadata",
		Resolver:             fetchAssessmentsMetadata,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/assessments-metadata/list?tabs=HTTP#securityassessmentmetadata",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_security_assessments_metadata", client.Namespacemicrosoft_security),
		Transform:            transformers.TransformWithStruct(&armsecurity.AssessmentMetadataResponse{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionIDPK},
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
