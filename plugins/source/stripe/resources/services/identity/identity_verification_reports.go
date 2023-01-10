package identity

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func IdentityVerificationReports() *schema.Table {
	return &schema.Table{
		Name:        "stripe_identity_verification_reports",
		Description: `https://stripe.com/docs/api/identity_verification_reports`,
		Transform:   transformers.TransformWithStruct(&stripe.IdentityVerificationReport{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchIdentityVerificationReports,

		Columns: []schema.Column{
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

func fetchIdentityVerificationReports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.IdentityVerificationReports.List(&stripe.IdentityVerificationReportListParams{})
	for it.Next() {
		res <- it.IdentityVerificationReport()
	}
	return it.Err()
}
