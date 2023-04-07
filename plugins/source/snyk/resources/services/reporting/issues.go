package reporting

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/internal/legacy"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"golang.org/x/sync/errgroup"
)

const issuesTableName = "snyk_reporting_issues"

func Issues() *schema.Table {
	return &schema.Table{
		Name:        issuesTableName,
		Description: `https://snyk.docs.apiary.io/#reference/reporting-api/get-list-of-latest-issues`,
		Resolver:    fetchIssues,
		Multiplex:   nil,
		Transform: transformers.TransformWithStruct(
			&legacy.ListReportingIssueResult{},
		),
		Columns: schema.ColumnList{
			client.OrganizationID,
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Issue.ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
					NotNull:    true,
				},
			},
		},
	}
}

func fetchIssues(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	req := legacy.ListReportingIssuesRequest{
		Page:    1,
		PerPage: 1000,
		SortBy:  "severity",
		Order:   "desc",
		GroupBy: "issue",
	}
	total := 0
	resp, err := c.LegacyClient.ListLatestReportingIssues(ctx, req)
	if err != nil {
		return err
	}
	if resp == nil || len(resp.Results) == 0 {
		return nil
	}
	res <- resp.Results
	total = resp.Total
	pages := total / req.PerPage
	if (total % req.PerPage) > 0 {
		pages++
	}
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(10) // limit the number of concurrent requests to 10

	for i := 2; i <= pages; i++ {
		r := req
		r.Page = i
		g.Go(func() error {
			return c.RetryOnError(gctx, issuesTableName, func() error {
				issues, err := c.LegacyClient.ListLatestReportingIssues(ctx, r)
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
