package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Assessments() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_assessments",
		Resolver:    fetchAssessments,
		Description: "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/assessments/list?tabs=HTTP#securityassessment",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_assessments", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.AssessmentResponse{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			subAssessments(),
		},
	}
}

func fetchAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewAssessmentsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager("/subscriptions/"+cl.SubscriptionId, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
