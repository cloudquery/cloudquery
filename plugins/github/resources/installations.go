package resources

import (
	"context"
	"encoding/json"

	"github.com/google/go-github/v45/github"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource  --config installations.hcl --output .
func Installations() *schema.Table {
	return &schema.Table{
		Name:        "github_installations",
		Description: "Installation represents a GitHub Apps installation.",
		Resolver:    fetchInstallations,
		Multiplex:   client.OrgMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "org",
				Description: "The Github Organization of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
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
				Name:     "app_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("AppID"),
			},
			{
				Name: "app_slug",
				Type: schema.TypeString,
			},
			{
				Name:     "target_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("TargetID"),
			},
			{
				Name:     "account_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Login"),
			},
			{
				Name:     "account_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.ID"),
			},
			{
				Name:     "account_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.NodeID"),
			},
			{
				Name:     "account_avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.AvatarURL"),
			},
			{
				Name:     "account_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.HTMLURL"),
			},
			{
				Name:     "account_gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.GravatarID"),
			},
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Name"),
			},
			{
				Name:     "account_company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Company"),
			},
			{
				Name:     "account_blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Blog"),
			},
			{
				Name:     "account_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Location"),
			},
			{
				Name:     "account_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Email"),
			},
			{
				Name:     "account_hireable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Account.Hireable"),
			},
			{
				Name:     "account_bio",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Bio"),
			},
			{
				Name:     "account_twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.TwitterUsername"),
			},
			{
				Name:     "account_public_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.PublicRepos"),
			},
			{
				Name:     "account_public_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.PublicGists"),
			},
			{
				Name:     "account_followers",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.Followers"),
			},
			{
				Name:     "account_following",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.Following"),
			},
			{
				Name:     "account_created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Account.CreatedAt.Time"),
			},
			{
				Name:     "account_updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Account.UpdatedAt.Time"),
			},
			{
				Name:     "account_suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Account.SuspendedAt.Time"),
			},
			{
				Name:     "account_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Type"),
			},
			{
				Name:     "account_site_admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Account.SiteAdmin"),
			},
			{
				Name:     "account_total_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.TotalPrivateRepos"),
			},
			{
				Name:     "account_owned_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.OwnedPrivateRepos"),
			},
			{
				Name:     "account_private_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.PrivateGists"),
			},
			{
				Name:     "account_disk_usage",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.DiskUsage"),
			},
			{
				Name:     "account_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.Collaborators"),
			},
			{
				Name:     "account_two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Account.TwoFactorAuthentication"),
			},
			{
				Name:     "account_plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Plan.Name"),
			},
			{
				Name:     "account_plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.Plan.Space"),
			},
			{
				Name:     "account_plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.Plan.Collaborators"),
			},
			{
				Name:     "account_plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.Plan.PrivateRepos"),
			},
			{
				Name:     "account_plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.Plan.FilledSeats"),
			},
			{
				Name:     "account_plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Account.Plan.Seats"),
			},
			{
				Name:     "account_ldap_dn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.LdapDn"),
			},
			{
				Name:        "account_url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Account.URL"),
			},
			{
				Name:     "account_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.EventsURL"),
			},
			{
				Name:     "account_following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.FollowingURL"),
			},
			{
				Name:     "account_followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.FollowersURL"),
			},
			{
				Name:     "account_gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.GistsURL"),
			},
			{
				Name:     "account_organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.OrganizationsURL"),
			},
			{
				Name:     "account_received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.ReceivedEventsURL"),
			},
			{
				Name:     "account_repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.ReposURL"),
			},
			{
				Name:     "account_starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.StarredURL"),
			},
			{
				Name:     "account_subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.SubscriptionsURL"),
			},
			{
				Name:        "account_text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveInstallationsAccountTextMatches,
			},
			{
				Name:        "account_permissions",
				Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Account.Permissions"),
			},
			{
				Name:     "account_role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.RoleName"),
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
				Name:     "permissions_actions",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Actions"),
			},
			{
				Name:     "permissions_administration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Administration"),
			},
			{
				Name:     "permissions_blocking",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Blocking"),
			},
			{
				Name:     "permissions_checks",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Checks"),
			},
			{
				Name:     "permissions_contents",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Contents"),
			},
			{
				Name:     "permissions_content_references",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.ContentReferences"),
			},
			{
				Name:     "permissions_deployments",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Deployments"),
			},
			{
				Name:     "permissions_emails",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Emails"),
			},
			{
				Name:     "permissions_environments",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Environments"),
			},
			{
				Name:     "permissions_followers",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Followers"),
			},
			{
				Name:     "permissions_issues",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Issues"),
			},
			{
				Name:     "permissions_metadata",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Metadata"),
			},
			{
				Name:     "permissions_members",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Members"),
			},
			{
				Name:     "permissions_organization_administration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.OrganizationAdministration"),
			},
			{
				Name:     "permissions_organization_hooks",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.OrganizationHooks"),
			},
			{
				Name:     "permissions_organization_plan",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.OrganizationPlan"),
			},
			{
				Name:     "permissions_organization_pre_receive_hooks",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.OrganizationPreReceiveHooks"),
			},
			{
				Name:     "permissions_organization_projects",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.OrganizationProjects"),
			},
			{
				Name:     "permissions_organization_secrets",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.OrganizationSecrets"),
			},
			{
				Name:     "permissions_organization_self_hosted_runners",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.OrganizationSelfHostedRunners"),
			},
			{
				Name:     "permissions_organization_user_blocking",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.OrganizationUserBlocking"),
			},
			{
				Name:     "permissions_packages",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Packages"),
			},
			{
				Name:     "permissions_pages",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Pages"),
			},
			{
				Name:     "permissions_pull_requests",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.PullRequests"),
			},
			{
				Name:     "permissions_repository_hooks",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.RepositoryHooks"),
			},
			{
				Name:     "permissions_repository_projects",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.RepositoryProjects"),
			},
			{
				Name:     "permissions_repository_pre_receive_hooks",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.RepositoryPreReceiveHooks"),
			},
			{
				Name:     "permissions_secrets",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Secrets"),
			},
			{
				Name:     "permissions_secret_scanning_alerts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.SecretScanningAlerts"),
			},
			{
				Name:     "permissions_security_events",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.SecurityEvents"),
			},
			{
				Name:     "permissions_single_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.SingleFile"),
			},
			{
				Name:     "permissions_statuses",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Statuses"),
			},
			{
				Name:     "permissions_team_discussions",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.TeamDiscussions"),
			},
			{
				Name:     "permissions_vulnerability_alerts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.VulnerabilityAlerts"),
			},
			{
				Name:     "permissions_workflows",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permissions.Workflows"),
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
				Name:     "suspended_by_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Login"),
			},
			{
				Name:     "suspended_by_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.ID"),
			},
			{
				Name:     "suspended_by_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.NodeID"),
			},
			{
				Name:     "suspended_by_avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.AvatarURL"),
			},
			{
				Name:     "suspended_by_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.HTMLURL"),
			},
			{
				Name:     "suspended_by_gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.GravatarID"),
			},
			{
				Name:     "suspended_by_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Name"),
			},
			{
				Name:     "suspended_by_company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Company"),
			},
			{
				Name:     "suspended_by_blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Blog"),
			},
			{
				Name:     "suspended_by_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Location"),
			},
			{
				Name:     "suspended_by_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Email"),
			},
			{
				Name:     "suspended_by_hireable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SuspendedBy.Hireable"),
			},
			{
				Name:     "suspended_by_bio",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Bio"),
			},
			{
				Name:     "suspended_by_twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.TwitterUsername"),
			},
			{
				Name:     "suspended_by_public_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.PublicRepos"),
			},
			{
				Name:     "suspended_by_public_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.PublicGists"),
			},
			{
				Name:     "suspended_by_followers",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.Followers"),
			},
			{
				Name:     "suspended_by_following",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.Following"),
			},
			{
				Name:     "suspended_by_created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SuspendedBy.CreatedAt.Time"),
			},
			{
				Name:     "suspended_by_updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SuspendedBy.UpdatedAt.Time"),
			},
			{
				Name:     "suspended_by_suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SuspendedBy.SuspendedAt.Time"),
			},
			{
				Name:     "suspended_by_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Type"),
			},
			{
				Name:     "suspended_by_site_admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SuspendedBy.SiteAdmin"),
			},
			{
				Name:     "suspended_by_total_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.TotalPrivateRepos"),
			},
			{
				Name:     "suspended_by_owned_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.OwnedPrivateRepos"),
			},
			{
				Name:     "suspended_by_private_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.PrivateGists"),
			},
			{
				Name:     "suspended_by_disk_usage",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.DiskUsage"),
			},
			{
				Name:     "suspended_by_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.Collaborators"),
			},
			{
				Name:     "suspended_by_two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SuspendedBy.TwoFactorAuthentication"),
			},
			{
				Name:     "suspended_by_plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.Plan.Name"),
			},
			{
				Name:     "suspended_by_plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.Plan.Space"),
			},
			{
				Name:     "suspended_by_plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.Plan.Collaborators"),
			},
			{
				Name:     "suspended_by_plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.Plan.PrivateRepos"),
			},
			{
				Name:     "suspended_by_plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.Plan.FilledSeats"),
			},
			{
				Name:     "suspended_by_plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("SuspendedBy.Plan.Seats"),
			},
			{
				Name:     "suspended_by_ldap_dn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.LdapDn"),
			},
			{
				Name:        "suspended_by_url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SuspendedBy.URL"),
			},
			{
				Name:     "suspended_by_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.EventsURL"),
			},
			{
				Name:     "suspended_by_following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.FollowingURL"),
			},
			{
				Name:     "suspended_by_followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.FollowersURL"),
			},
			{
				Name:     "suspended_by_gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.GistsURL"),
			},
			{
				Name:     "suspended_by_organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.OrganizationsURL"),
			},
			{
				Name:     "suspended_by_received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.ReceivedEventsURL"),
			},
			{
				Name:     "suspended_by_repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.ReposURL"),
			},
			{
				Name:     "suspended_by_starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.StarredURL"),
			},
			{
				Name:     "suspended_by_subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.SubscriptionsURL"),
			},
			{
				Name:        "suspended_by_text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveInstallationsSuspendedByTextMatches,
			},
			{
				Name:        "suspended_by_permissions",
				Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SuspendedBy.Permissions"),
			},
			{
				Name:     "suspended_by_role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SuspendedBy.RoleName"),
			},
			{
				Name:     "suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SuspendedAt.Time"),
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
			return err
		}
		res <- installations.Installations
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func resolveInstallationsAccountTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.Installation)
	if u.Account == nil {
		return nil
	}
	j, err := json.Marshal(u.Account.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveInstallationsSuspendedByTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.Installation)
	if u.SuspendedBy == nil {
		return nil
	}
	j, err := json.Marshal(u.SuspendedBy.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
