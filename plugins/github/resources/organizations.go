package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource  --config organizations.hcl --output .
func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "github_organizations",
		Description: "Organization represents a GitHub organization account.",
		Multiplex:   client.OrgMultiplex,
		IgnoreError: client.IgnoreError,
		Resolver:    fetchOrganizations,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name: "login",
				Type: schema.TypeString,
			},
			{
				Name:     "id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ID"),
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
				Type: schema.TypeBigInt,
			},
			{
				Name: "public_gists",
				Type: schema.TypeBigInt,
			},
			{
				Name: "followers",
				Type: schema.TypeBigInt,
			},
			{
				Name: "following",
				Type: schema.TypeBigInt,
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
				Type: schema.TypeBigInt,
			},
			{
				Name: "owned_private_repos",
				Type: schema.TypeBigInt,
			},
			{
				Name: "private_gists",
				Type: schema.TypeBigInt,
			},
			{
				Name: "disk_usage",
				Type: schema.TypeBigInt,
			},
			{
				Name: "collaborators",
				Type: schema.TypeBigInt,
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
				Name:     "plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Plan.Name"),
			},
			{
				Name:     "plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Plan.Space"),
			},
			{
				Name:     "plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Plan.Collaborators"),
			},
			{
				Name:     "plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Plan.PrivateRepos"),
			},
			{
				Name:     "plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Plan.FilledSeats"),
			},
			{
				Name:     "plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Plan.Seats"),
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
		return err
	}
	res <- org
	return nil
}
