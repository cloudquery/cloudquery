package issues

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

func Issues() *schema.Table {
	return &schema.Table{
		Name:        "github_issues",
		Description: "Issue represents a GitHub issue on a repository.  Note: As far as the GitHub API is concerned, every pull request is an issue, but not every issue is a pull request",
		Resolver:    fetchIssues,
		Multiplex:   client.OrgMultiplex,
		IgnoreError: client.IgnoreError,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"org", "id"}},
		Columns: []schema.Column{
			{
				Name:        "org",
				Description: "The Github Organization of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name: "number",
				Type: schema.TypeInt,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "locked",
				Type: schema.TypeBool,
			},
			{
				Name: "title",
				Type: schema.TypeString,
			},
			{
				Name: "body",
				Type: schema.TypeString,
			},
			{
				Name: "author_association",
				Type: schema.TypeString,
			},
			{
				Name: "user",
				Type: schema.TypeJSON,
			},
			{
				Name: "assignee",
				Type: schema.TypeJSON,
			},
			{
				Name: "comments",
				Type: schema.TypeInt,
			},
			{
				Name: "closed_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "created_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "updated_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "closed_by",
				Type: schema.TypeJSON,
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "comments_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CommentsURL"),
			},
			{
				Name:     "events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventsURL"),
			},
			{
				Name:     "labels_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelsURL"),
			},
			{
				Name:     "repository_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryURL"),
			},
			{
				Name: "milestone",
				Type: schema.TypeJSON,
			},
			{
				Name: "pull_request_links",
				Type: schema.TypeJSON,
			},
			{
				Name: "repository",
				Type: schema.TypeJSON,
			},
			{
				Name: "reactions",
				Type: schema.TypeJSON,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "assignees",
				Type: schema.TypeJSON,
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeID"),
			},
			{
				Name:        "text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "active_lock_reason",
				Description: "ActiveLockReason is populated only when LockReason is provided while locking the issue. Possible values are: \"off-topic\", \"too heated\", \"resolved\", and \"spam\".",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIssues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.IssueListOptions{
		Filter: "all",
		State:  "all",
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	for {
		issues, resp, err := c.Github.Issues.ListByOrg(ctx, c.Org, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- issues
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
