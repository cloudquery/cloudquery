package cloudbilling

import (
	"context"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/cloudbilling/v1"
)

const MAX_GOROUTINES = 10

//go:generate cq-gen --resource accounts --config gen.hcl --output .
func Accounts() *schema.Table {
	return &schema.Table{
		Name:          "gcp_cloudbilling_accounts",
		Resolver:      fetchBillingAccounts,
		Multiplex:     client.ProjectMultiplex,
		IgnoreError:   client.IgnoreErrorHandler,
		DeleteFilter:  client.DeleteProjectFilter,
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
	var sem = semaphore.NewWeighted(int64(MAX_GOROUTINES))
	for {
		call := c.Services.CloudBilling.BillingAccounts.List().PageToken(nextPageToken)
		list, err := c.RetryingDo(ctx, call)
		if err != nil {
			return err
		}
		output := list.(*cloudbilling.ListBillingAccountsResponse)
		errs, ctx := errgroup.WithContext(ctx)
		for _, b := range output.BillingAccounts {
			if err := sem.Acquire(ctx, 1); err != nil {
				return diag.WrapError(err)
			}
			func(account cloudbilling.BillingAccount) {
				errs.Go(func() error {
					defer sem.Release(1)
					return fetchProjectBillingInfo(ctx, res, c, account)
				})
			}(*b)
		}
		err = errs.Wait()
		if err != nil {
			return diag.WrapError(err)
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

type BillingAccountWrapper struct {
	*cloudbilling.BillingAccount
	*cloudbilling.ProjectBillingInfo
}

func fetchProjectBillingInfo(ctx context.Context, res chan<- interface{}, c *client.Client, b cloudbilling.BillingAccount) error {
	nextPageToken := ""
	for {
		projectsCall := c.Services.CloudBilling.BillingAccounts.Projects.List(b.Name).PageToken(nextPageToken)
		projectsList, err := c.RetryingDo(ctx, projectsCall)
		if err != nil {
			return err
		}
		output := projectsList.(*cloudbilling.ListProjectBillingInfoResponse)
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
