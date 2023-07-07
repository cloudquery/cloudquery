package reporting

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
	"golang.org/x/sync/errgroup"
)

const (
	issuesTableName = "snyk_reporting_issues"
)

func Issues() *schema.Table {
	return &schema.Table{
		Name:        issuesTableName,
		Description: `https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues`,
		Resolver:    fetchIssues,
		Multiplex:   client.ByOrganization,
		Transform: transformers.TransformWithStruct(
			&snyk.ListReportingIssueResult{},
		),
		Columns: schema.ColumnList{
			client.OrganizationID,
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
				NotNull:    true,
				Resolver:   schema.PathResolver("Issue.ID"),
			},
		},
	}
}

func fetchIssues(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	req := snyk.ListReportingIssuesRequest{
		Page:    1,
		PerPage: 1000,
		GroupBy: "issue",
	}
	total := 0
	var (
		resp *snyk.ListReportingIssuesResponse
		err  error
	)
	from := c.Spec.TableOptions.SnykReportingIssues.FromTime()
	to := c.Spec.TableOptions.SnykReportingIssues.ToTime()
	err = c.RetryOnError(ctx, issuesTableName, func() error {
		resp, _, err = c.Client.Reporting.ListIssues(ctx, c.OrganizationID, from, to, req)
		return err
	})
	if err != nil {
		return err
	}
	res <- resp.Results
	total = resp.Total
	pages := total / req.PerPage
	if (total % req.PerPage) > 0 {
		pages++
	}
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(maxRequests)

	for i := 2; i <= pages; i++ {
		r := req
		r.Page = i
		g.Go(func() error {
			return c.RetryOnError(gctx, issuesTableName, func() error {
				issues, _, err := c.Client.Reporting.ListIssues(ctx, c.OrganizationID, from, to, r)
				if err != nil {
					return err
				}
				res <- issues.Results
				return nil
			})
		})
	}
	return g.Wait()
}
