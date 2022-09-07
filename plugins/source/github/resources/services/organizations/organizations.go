package organizations

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "github_organizations",
		Description: "Organization represents a GitHub organization account.",
		Resolver:    fetchOrganizations,
		Multiplex:   client.OrgMultiplex,
		Columns: []schema.Column{
			{
				Name:            "id",
				Type:            schema.TypeInt,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "login",
				Type: schema.TypeString,
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeID"),
			},
			{
				Name:     "avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvatarURL"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "company",
				Type: schema.TypeString,
			},
			{
				Name: "blog",
				Type: schema.TypeString,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name: "email",
				Type: schema.TypeString,
			},
			{
				Name: "twitter_username",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "public_repos",
				Type: schema.TypeInt,
			},
			{
				Name: "public_gists",
				Type: schema.TypeInt,
			},
			{
				Name: "followers",
				Type: schema.TypeInt,
			},
			{
				Name: "following",
				Type: schema.TypeInt,
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
				Name: "total_private_repos",
				Type: schema.TypeInt,
			},
			{
				Name: "owned_private_repos",
				Type: schema.TypeInt,
			},
			{
				Name: "private_gists",
				Type: schema.TypeInt,
			},
			{
				Name: "disk_usage",
				Type: schema.TypeInt,
			},
			{
				Name: "collaborators",
				Type: schema.TypeInt,
			},
			{
				Name: "billing_email",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "plan",
				Type: schema.TypeJSON,
			},
			{
				Name: "two_factor_requirement_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "is_verified",
				Type: schema.TypeBool,
			},
			{
				Name: "has_organization_projects",
				Type: schema.TypeBool,
			},
			{
				Name: "has_repository_projects",
				Type: schema.TypeBool,
			},
			{
				Name:        "default_repo_permission",
				Description: "DefaultRepoPermission can be one of: \"read\", \"write\", \"admin\", or \"none\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "default_repo_settings",
				Description: "DefaultRepoSettings can be one of: \"read\", \"write\", \"admin\", or \"none\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "members_can_create_repos",
				Description: "MembersCanCreateRepos default value is true and is only used in Organizations.Edit.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "members_can_create_public_repos",
				Description: "https://developer.github.com/changes/2019-12-03-internal-visibility-changes/#rest-v3-api",
				Type:        schema.TypeBool,
			},
			{
				Name: "members_can_create_private_repos",
				Type: schema.TypeBool,
			},
			{
				Name: "members_can_create_internal_repos",
				Type: schema.TypeBool,
			},
			{
				Name:        "members_can_fork_private_repos",
				Description: "MembersCanForkPrivateRepos toggles whether organization members can fork private organization repositories.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "members_allowed_repository_creation_type",
				Description: "MembersAllowedRepositoryCreationType denotes if organization members can create repositories and the type of repositories they can create",
				Type:        schema.TypeString,
			},
			{
				Name:        "members_can_create_pages",
				Description: "MembersCanCreatePages toggles whether organization members can create GitHub Pages sites.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "members_can_create_public_pages",
				Description: "MembersCanCreatePublicPages toggles whether organization members can create public GitHub Pages sites.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "members_can_create_private_pages",
				Description: "MembersCanCreatePrivatePages toggles whether organization members can create private GitHub Pages sites.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("URL"),
			},
			{
				Name:     "events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventsURL"),
			},
			{
				Name:     "hooks_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HooksURL"),
			},
			{
				Name:     "issues_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IssuesURL"),
			},
			{
				Name:     "members_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MembersURL"),
			},
			{
				Name:     "members",
				Type:     schema.TypeJSON,
				Resolver: resolveOrganizationMembers,
			},
			{
				Name:     "public_members_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicMembersURL"),
			},
			{
				Name:     "repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReposURL"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	org, _, err := c.Github.Organizations.Get(ctx, c.Org)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- org
	return nil
}
func resolveOrganizationMembers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	c := meta.(*client.Client)
	opts := &github.ListMembersOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	var orgMembers []*Member
	for {
		members, resp, err := c.Github.Organizations.ListMembers(ctx, c.Org, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		for _, member := range members {
			membership, _, err := c.Github.Organizations.GetOrgMembership(ctx, *member.Login, c.Org)
			if err != nil {
				return errors.WithStack(err)
			}
			orgMembers = append(orgMembers, &Member{
				User:       member,
				Membership: membership,
			})
		}
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	data, err := json.Marshal(orgMembers)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(column.Name, data))
}
