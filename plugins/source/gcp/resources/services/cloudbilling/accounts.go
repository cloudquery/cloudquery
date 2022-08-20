package cloudbilling

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/cloudbilling/v1"
)

type BillingAccountWrapper struct {
	*cloudbilling.BillingAccount
	*cloudbilling.ProjectBillingInfo
}

const maxGoroutines = 10

//go:generate cq-gen --resource accounts --config gen.hcl --output .
func Accounts() *schema.Table {
	return &schema.Table{
		Name:          "gcp_cloudbilling_accounts",
		Resolver:      fetchBillingAccounts,
		Multiplex:     client.ProjectMultiplex,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "name"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "display_name",
				Description: "The display name given to the billing account, such as `My Billing Account`",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingAccount.DisplayName"),
			},
			{
				Name:        "master_billing_account",
				Description: "If this account is a subaccount (https://cloudgooglecom/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingAccount.MasterBillingAccount"),
			},
			{
				Name:        "name",
				Description: "The resource name of the billing account.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BillingAccount.Name"),
			},
			{
				Name:        "open",
				Description: "True if the billing account is open",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BillingAccount.Open"),
			},
			{
				Name:        "project_billing_enabled",
				Description: "True if the project is associated with an open billing account, to which usage on the project is charged",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ProjectBillingInfo.BillingEnabled"),
			},
			{
				Name:        "project_name",
				Description: "The resource name for the `ProjectBillingInfo`; has the form `projects/{project_id}/billingInfo`",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProjectBillingInfo.Name"),
			},
			{
				Name:        "project_id",
				Description: "The ID of the project that this `ProjectBillingInfo` represents, such as `tokyo-rain-123`",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProjectBillingInfo.ProjectId"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBillingAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.CloudBilling.BillingAccounts.List().PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		errs, ctx := errgroup.WithContext(ctx)
		errs.SetLimit(maxGoroutines)
		for _, b := range output.BillingAccounts {
			func(account cloudbilling.BillingAccount) {
				errs.Go(func() error {
					return fetchProjectBillingInfo(ctx, res, c, account)
				})
			}(*b)
		}
		err = errs.Wait()
		if err != nil {
			return errors.WithStack(err)
		}
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func fetchProjectBillingInfo(ctx context.Context, res chan<- interface{}, c *client.Client, b cloudbilling.BillingAccount) error {
	nextPageToken := ""
	for {
		output, err := c.Services.CloudBilling.BillingAccounts.Projects.List(b.Name).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		for _, p := range output.ProjectBillingInfo {
			if p.ProjectId == c.ProjectId {
				res <- BillingAccountWrapper{
					&b,
					p,
				}
				return nil
			}
		}
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
