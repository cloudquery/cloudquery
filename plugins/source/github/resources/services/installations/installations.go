package installations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

func Installations() *schema.Table {
	return &schema.Table{
		Name:        "github_installations",
		Description: "Installation represents a GitHub Apps installation.",
		Resolver:    fetchInstallations,
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
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeID"),
			},
			{
				Name:     "app_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AppID"),
			},
			{
				Name: "app_slug",
				Type: schema.TypeString,
			},
			{
				Name:     "target_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TargetID"),
			},
			{
				Name: "account",
				Type: schema.TypeJSON,
			},
			{
				Name:     "access_tokens_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessTokensURL"),
			},
			{
				Name:     "repositories_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoriesURL"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name: "target_type",
				Type: schema.TypeString,
			},
			{
				Name: "single_file_name",
				Type: schema.TypeString,
			},
			{
				Name: "repository_selection",
				Type: schema.TypeString,
			},
			{
				Name: "events",
				Type: schema.TypeStringArray,
			},
			{
				Name: "single_file_paths",
				Type: schema.TypeStringArray,
			},
			{
				Name: "permissions",
				Type: schema.TypeJSON,
			},
			{
				Name:     "created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt.Time"),
			},
			{
				Name:     "updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt.Time"),
			},
			{
				Name: "has_multiple_single_files",
				Type: schema.TypeBool,
			},
			{
				Name: "suspended_by",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchInstallations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{
		Page:    0,
		PerPage: 100,
	}
	for {
		installations, resp, err := c.Github.Organizations.ListInstallations(ctx, c.Org, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- installations.Installations
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
