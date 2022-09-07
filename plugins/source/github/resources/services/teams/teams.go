package teams

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:        "github_teams",
		Description: "Team represents a team within a GitHub organization",
		Resolver:    fetchTeams,
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
				Name:            "id",
				Type:            schema.TypeInt,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name: "slug",
				Type: schema.TypeString,
			},
			{
				Name:        "permission",
				Description: "Permission specifies the default permission for repositories owned by the team.",
				Type:        schema.TypeString,
			},
			{
				Name:        "permissions",
				Description: "Permissions identifies the permissions that a team has on a given repository",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "privacy",
				Description: "Privacy identifies the level of privacy this team should have. Possible values are:     secret - only visible to organization owners and members of this team     closed - visible to all members of this organization Default is \"secret\".",
				Type:        schema.TypeString,
			},
			{
				Name:     "members",
				Type:     schema.TypeJSON,
				Resolver: resolveTeamMembers,
			},
			{
				Name: "members_count",
				Type: schema.TypeInt,
			},
			{
				Name: "repos_count",
				Type: schema.TypeInt,
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "members_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MembersURL"),
			},
			{
				Name:     "repositories",
				Type:     schema.TypeJSON,
				Resolver: resolveTeamRepositories,
			},
			{
				Name:     "repositories_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoriesURL"),
			},
			{
				Name: "parent",
				Type: schema.TypeJSON,
			},
			{
				Name:        "ldapdn",
				Description: "LDAPDN is only available in GitHub Enterprise and when the team membership is synchronized with LDAP.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LDAPDN"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{
		Page:    0,
		PerPage: 100,
	}
	for {
		repos, resp, err := c.Github.Teams.ListTeams(ctx, c.Org, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- repos
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}

func resolveTeamMembers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	t := resource.Item.(*github.Team)
	c := meta.(*client.Client)
	opts := &github.TeamListTeamMembersOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	orgId, err := strconv.Atoi(strings.Split(*t.MembersURL, "/")[4])
	if err != nil {
		return errors.WithStack(err)
	}
	var teamMembers []*github.User
	for {
		members, resp, err := c.Github.Teams.ListTeamMembersByID(ctx, int64(orgId), *t.ID, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		teamMembers = append(teamMembers, members...)
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	data, err := json.Marshal(teamMembers)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(column.Name, data))
}

func resolveTeamRepositories(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	t := resource.Item.(*github.Team)
	c := meta.(*client.Client)
	opts := &github.ListOptions{
		Page:    0,
		PerPage: 100,
	}
	orgId, err := strconv.Atoi(strings.Split(*t.MembersURL, "/")[4])
	if err != nil {
		return errors.WithStack(err)
	}
	var repositories []*github.Repository
	for {
		repos, resp, err := c.Github.Teams.ListTeamReposByID(ctx, int64(orgId), *t.ID, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		repositories = append(repositories, repos...)
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	data, err := json.Marshal(repositories)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(column.Name, data))
}
