package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func subAssessments() *schema.Table {
	return &schema.Table{
		Name:                 "azure_security_sub_assessments",
		Resolver:             fetchSubAssessments,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/sub-assessments/list?tabs=HTTP#securitysubassessment",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_security_sub_assessments", client.Namespacemicrosoft_security),
		Transform:            transformers.TransformWithStruct(&armsecurity.SubAssessment{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSubAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	item := parent.Item.(*armsecurity.AssessmentResponse)
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewSubAssessmentsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager("/subscriptions/"+cl.SubscriptionId, *item.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
