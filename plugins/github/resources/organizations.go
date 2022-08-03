package resources

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/go-github/v45/github"
)

//go:generate cq-gen --resource  --config organizations.hcl --output .
func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "github_organizations",
		Description: "Organization represents a GitHub organization account.",
		Resolver:    fetchOrganizations,
		Multiplex:   client.OrgMultiplex,
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
		Relations: []*schema.Table{
			{
				Name:        "github_organization_members",
				Description: "User represents a GitHub user.",
				Resolver:    fetchOrganizationMembers,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"org", "id"}},
				Columns: []schema.Column{
					{
						Name:        "organization_cq_id",
						Description: "Unique CloudQuery ID of github_organizations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "org",
						Description: "The Github Organization of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveOrg,
					},
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
						Name:     "gravatar_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GravatarID"),
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
						Name: "hireable",
						Type: schema.TypeBool,
					},
					{
						Name: "bio",
						Type: schema.TypeString,
					},
					{
						Name: "twitter_username",
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
						Name:     "suspended_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SuspendedAt.Time"),
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
					{
						Name: "site_admin",
						Type: schema.TypeBool,
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
						Name: "two_factor_authentication",
						Type: schema.TypeBool,
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
						Name: "ldap_dn",
						Type: schema.TypeString,
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
						Name:     "following_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("FollowingURL"),
					},
					{
						Name:     "followers_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("FollowersURL"),
					},
					{
						Name:     "gists_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GistsURL"),
					},
					{
						Name:     "organizations_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("OrganizationsURL"),
					},
					{
						Name:     "received_events_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ReceivedEventsURL"),
					},
					{
						Name:     "repos_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ReposURL"),
					},
					{
						Name:     "starred_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("StarredURL"),
					},
					{
						Name:     "subscriptions_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubscriptionsURL"),
					},
					{
						Name:        "text_matches",
						Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
						Type:        schema.TypeJSON,
						Resolver:    resolveOrganizationMembersTextMatches,
					},
					{
						Name:        "permissions",
						Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
						Type:        schema.TypeJSON,
					},
					{
						Name: "role_name",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "github_organization_member_membership",
						Description: "Membership represents the status of a user's membership in an organization or team.",
						Resolver:    fetchOrganizationMemberMemberships,
						Columns: []schema.Column{
							{
								Name:        "organization_member_cq_id",
								Description: "Unique CloudQuery ID of github_organization_members table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "url",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("URL"),
							},
							{
								Name:        "state",
								Description: "State is the user's status within the organization or team. Possible values are: \"active\", \"pending\"",
								Type:        schema.TypeString,
							},
							{
								Name:        "role",
								Description: "Role identifies the user's role within the organization or team. Possible values for organization membership:     member - non-owner organization member     admin - organization owner  Possible values for team membership are:     member - a normal member of the team     maintainer - a team maintainer",
								Type:        schema.TypeString,
							},
							{
								Name:        "organization_url",
								Description: "For organization membership, the API URL of the organization.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("OrganizationURL"),
							},
						},
					},
				},
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
		return diag.WrapError(err)
	}
	res <- org
	return nil
}
func fetchOrganizationMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListMembersOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	for {
		members, resp, err := c.Github.Organizations.ListMembers(ctx, c.Org, opts)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- members
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func resolveOrganizationMembersTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.User)
	j, err := json.Marshal(u.TextMatches)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func fetchOrganizationMemberMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	m := parent.Item.(*github.User)

	membership, _, err := c.Github.Organizations.GetOrgMembership(ctx, *m.Name, c.Org)
	if err != nil {
		return diag.WrapError(err)
	}
	res <- membership
	return nil
}
