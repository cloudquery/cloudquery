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
	latestIssuesTableName = "snyk_reporting_latest_issues"
	maxRequests           = 10 // limit the number of concurrent requests to 10
)

func LatestIssues() *schema.Table {
	return &schema.Table{
		Name:        latestIssuesTableName,
		Description: `https://snyk.docs.apiary.io/#reference/reporting-api/get-list-of-latest-issues`,
		Resolver:    fetchLatestIssues,
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

func fetchLatestIssues(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
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
	err = c.RetryOnError(ctx, latestIssuesTableName, func() error {
		resp, _, err = c.Client.Reporting.ListLatestIssues(ctx, c.OrganizationID, req)
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
			return c.RetryOnError(gctx, latestIssuesTableName, func() error {
				issues, _, err := c.Client.Reporting.ListLatestIssues(ctx, c.OrganizationID, r)
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
