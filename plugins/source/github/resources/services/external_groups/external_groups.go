package external_groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

func ExternalGroups() *schema.Table {
	return &schema.Table{
		Name:        "github_external_groups",
		Description: "ExternalGroup represents an external group.",
		Resolver:    fetchExternalGroups,
		Multiplex:   client.OrgMultiplex,
		Columns: []schema.Column{
			{
				Name:            "org",
				Description:     "The Github Organization of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveOrg,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "group_id",
				Type:            schema.TypeInt,
				Resolver:        schema.PathResolver("GroupID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "group_name",
				Type: schema.TypeString,
			},
			{
				Name:     "updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt.Time"),
			},
			{
				Name: "teams",
				Type: schema.TypeJSON,
			},
			{
				Name: "members",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchExternalGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListExternalGroupsOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	for {
		groups, resp, err := c.Github.Teams.ListExternalGroups(ctx, c.Org, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- groups.Groups
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
