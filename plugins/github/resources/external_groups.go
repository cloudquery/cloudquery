package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/go-github/v45/github"
)

//go:generate cq-gen --resource  --config external_groups.hcl --output .
func ExternalGroups() *schema.Table {
	return &schema.Table{
		Name:        "github_external_groups",
		Description: "ExternalGroup represents an external group.",
		Resolver:    fetchExternalGroups,
		Multiplex:   client.OrgMultiplex,
		IgnoreError: client.IgnoreError,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"org", "group_id"}},
		Columns: []schema.Column{
			{
				Name:        "org",
				Description: "The Github Organization of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
			},
			{
				Name:     "group_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("GroupID"),
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
		},
		Relations: []*schema.Table{
			{
				Name:        "github_external_group_teams",
				Description: "ExternalGroupTeam represents a team connected to an external group.",
				Resolver:    fetchExternalGroupTeams,
				Columns: []schema.Column{
					{
						Name:        "external_group_cq_id",
						Description: "Unique CloudQuery ID of github_external_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "team_id",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("TeamID"),
					},
					{
						Name: "team_name",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:        "github_external_group_members",
				Description: "ExternalGroupMember represents a member of an external group.",
				Resolver:    fetchExternalGroupMembers,
				Columns: []schema.Column{
					{
						Name:        "external_group_cq_id",
						Description: "Unique CloudQuery ID of github_external_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "member_id",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("MemberID"),
					},
					{
						Name: "member_login",
						Type: schema.TypeString,
					},
					{
						Name: "member_name",
						Type: schema.TypeString,
					},
					{
						Name: "member_email",
						Type: schema.TypeString,
					},
				},
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
			return diag.WrapError(err)
		}
		res <- groups.Groups
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func fetchExternalGroupTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	g := parent.Item.(*github.ExternalGroup)
	res <- g.Teams
	return nil
}
func fetchExternalGroupMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	g := parent.Item.(*github.ExternalGroup)
	res <- g.Members
	return nil
}
