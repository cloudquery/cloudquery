package resources

import (
	"context"
	"encoding/json"

	"github.com/google/go-github/v45/github"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource  --config issues.hcl --output .
func Issues() *schema.Table {
	return &schema.Table{
		Name:        "github_issues",
		Description: "Issue represents a GitHub issue on a repository.  Note: As far as the GitHub API is concerned, every pull request is an issue, but not every issue is a pull request",
		Resolver:    fetchIssues,
		Multiplex:   client.OrgMultiplex,
		IgnoreError: client.IgnoreError,
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
				Name: "number",
				Type: schema.TypeBigInt,
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
				Name:     "user_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Login"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.ID"),
			},
			{
				Name:     "user_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.NodeID"),
			},
			{
				Name:     "user_avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.AvatarURL"),
			},
			{
				Name:     "user_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.HTMLURL"),
			},
			{
				Name:     "user_gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.GravatarID"),
			},
			{
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Name"),
			},
			{
				Name:     "user_company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Company"),
			},
			{
				Name:     "user_blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Blog"),
			},
			{
				Name:     "user_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Location"),
			},
			{
				Name:     "user_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Email"),
			},
			{
				Name:     "user_hireable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("User.Hireable"),
			},
			{
				Name:     "user_bio",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Bio"),
			},
			{
				Name:     "user_twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.TwitterUsername"),
			},
			{
				Name:     "user_public_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.PublicRepos"),
			},
			{
				Name:     "user_public_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.PublicGists"),
			},
			{
				Name:     "user_followers",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.Followers"),
			},
			{
				Name:     "user_following",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.Following"),
			},
			{
				Name:     "user_created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("User.CreatedAt.Time"),
			},
			{
				Name:     "user_updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("User.UpdatedAt.Time"),
			},
			{
				Name:     "user_suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("User.SuspendedAt.Time"),
			},
			{
				Name:     "user_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Type"),
			},
			{
				Name:     "user_site_admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("User.SiteAdmin"),
			},
			{
				Name:     "user_total_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.TotalPrivateRepos"),
			},
			{
				Name:     "user_owned_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.OwnedPrivateRepos"),
			},
			{
				Name:     "user_private_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.PrivateGists"),
			},
			{
				Name:     "user_disk_usage",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.DiskUsage"),
			},
			{
				Name:     "user_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.Collaborators"),
			},
			{
				Name:     "user_two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("User.TwoFactorAuthentication"),
			},
			{
				Name:     "user_plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.Plan.Name"),
			},
			{
				Name:     "user_plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.Plan.Space"),
			},
			{
				Name:     "user_plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.Plan.Collaborators"),
			},
			{
				Name:     "user_plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.Plan.PrivateRepos"),
			},
			{
				Name:     "user_plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.Plan.FilledSeats"),
			},
			{
				Name:     "user_plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("User.Plan.Seats"),
			},
			{
				Name:     "user_ldap_dn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.LdapDn"),
			},
			{
				Name:        "user_url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("User.URL"),
			},
			{
				Name:     "user_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.EventsURL"),
			},
			{
				Name:     "user_following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.FollowingURL"),
			},
			{
				Name:     "user_followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.FollowersURL"),
			},
			{
				Name:     "user_gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.GistsURL"),
			},
			{
				Name:     "user_organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.OrganizationsURL"),
			},
			{
				Name:     "user_received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.ReceivedEventsURL"),
			},
			{
				Name:     "user_repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.ReposURL"),
			},
			{
				Name:     "user_starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.StarredURL"),
			},
			{
				Name:     "user_subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.SubscriptionsURL"),
			},
			{
				Name:        "user_text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveIssuesUserTextMatches,
			},
			{
				Name:        "user_permissions",
				Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("User.Permissions"),
			},
			{
				Name:     "user_role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User.RoleName"),
			},
			{
				Name:     "assignee_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Login"),
			},
			{
				Name:     "assignee_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.ID"),
			},
			{
				Name:     "assignee_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.NodeID"),
			},
			{
				Name:     "assignee_avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.AvatarURL"),
			},
			{
				Name:     "assignee_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.HTMLURL"),
			},
			{
				Name:     "assignee_gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.GravatarID"),
			},
			{
				Name:     "assignee_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Name"),
			},
			{
				Name:     "assignee_company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Company"),
			},
			{
				Name:     "assignee_blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Blog"),
			},
			{
				Name:     "assignee_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Location"),
			},
			{
				Name:     "assignee_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Email"),
			},
			{
				Name:     "assignee_hireable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Assignee.Hireable"),
			},
			{
				Name:     "assignee_bio",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Bio"),
			},
			{
				Name:     "assignee_twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.TwitterUsername"),
			},
			{
				Name:     "assignee_public_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.PublicRepos"),
			},
			{
				Name:     "assignee_public_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.PublicGists"),
			},
			{
				Name:     "assignee_followers",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.Followers"),
			},
			{
				Name:     "assignee_following",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.Following"),
			},
			{
				Name:     "assignee_created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Assignee.CreatedAt.Time"),
			},
			{
				Name:     "assignee_updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Assignee.UpdatedAt.Time"),
			},
			{
				Name:     "assignee_suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Assignee.SuspendedAt.Time"),
			},
			{
				Name:     "assignee_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Type"),
			},
			{
				Name:     "assignee_site_admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Assignee.SiteAdmin"),
			},
			{
				Name:     "assignee_total_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.TotalPrivateRepos"),
			},
			{
				Name:     "assignee_owned_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.OwnedPrivateRepos"),
			},
			{
				Name:     "assignee_private_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.PrivateGists"),
			},
			{
				Name:     "assignee_disk_usage",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.DiskUsage"),
			},
			{
				Name:     "assignee_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.Collaborators"),
			},
			{
				Name:     "assignee_two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Assignee.TwoFactorAuthentication"),
			},
			{
				Name:     "assignee_plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.Plan.Name"),
			},
			{
				Name:     "assignee_plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.Plan.Space"),
			},
			{
				Name:     "assignee_plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.Plan.Collaborators"),
			},
			{
				Name:     "assignee_plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.Plan.PrivateRepos"),
			},
			{
				Name:     "assignee_plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.Plan.FilledSeats"),
			},
			{
				Name:     "assignee_plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Assignee.Plan.Seats"),
			},
			{
				Name:     "assignee_ldap_dn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.LdapDn"),
			},
			{
				Name:        "assignee_url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Assignee.URL"),
			},
			{
				Name:     "assignee_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.EventsURL"),
			},
			{
				Name:     "assignee_following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.FollowingURL"),
			},
			{
				Name:     "assignee_followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.FollowersURL"),
			},
			{
				Name:     "assignee_gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.GistsURL"),
			},
			{
				Name:     "assignee_organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.OrganizationsURL"),
			},
			{
				Name:     "assignee_received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.ReceivedEventsURL"),
			},
			{
				Name:     "assignee_repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.ReposURL"),
			},
			{
				Name:     "assignee_starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.StarredURL"),
			},
			{
				Name:     "assignee_subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.SubscriptionsURL"),
			},
			{
				Name:        "assignee_text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveIssuesAssigneeTextMatches,
			},
			{
				Name:        "assignee_permissions",
				Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Assignee.Permissions"),
			},
			{
				Name:     "assignee_role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Assignee.RoleName"),
			},
			{
				Name: "comments",
				Type: schema.TypeBigInt,
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
				Name:     "closed_by_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Login"),
			},
			{
				Name:     "closed_by_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.ID"),
			},
			{
				Name:     "closed_by_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.NodeID"),
			},
			{
				Name:     "closed_by_avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.AvatarURL"),
			},
			{
				Name:     "closed_by_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.HTMLURL"),
			},
			{
				Name:     "closed_by_gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.GravatarID"),
			},
			{
				Name:     "closed_by_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Name"),
			},
			{
				Name:     "closed_by_company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Company"),
			},
			{
				Name:     "closed_by_blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Blog"),
			},
			{
				Name:     "closed_by_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Location"),
			},
			{
				Name:     "closed_by_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Email"),
			},
			{
				Name:     "closed_by_hireable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ClosedBy.Hireable"),
			},
			{
				Name:     "closed_by_bio",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Bio"),
			},
			{
				Name:     "closed_by_twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.TwitterUsername"),
			},
			{
				Name:     "closed_by_public_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.PublicRepos"),
			},
			{
				Name:     "closed_by_public_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.PublicGists"),
			},
			{
				Name:     "closed_by_followers",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.Followers"),
			},
			{
				Name:     "closed_by_following",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.Following"),
			},
			{
				Name:     "closed_by_created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ClosedBy.CreatedAt.Time"),
			},
			{
				Name:     "closed_by_updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ClosedBy.UpdatedAt.Time"),
			},
			{
				Name:     "closed_by_suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ClosedBy.SuspendedAt.Time"),
			},
			{
				Name:     "closed_by_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Type"),
			},
			{
				Name:     "closed_by_site_admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ClosedBy.SiteAdmin"),
			},
			{
				Name:     "closed_by_total_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.TotalPrivateRepos"),
			},
			{
				Name:     "closed_by_owned_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.OwnedPrivateRepos"),
			},
			{
				Name:     "closed_by_private_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.PrivateGists"),
			},
			{
				Name:     "closed_by_disk_usage",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.DiskUsage"),
			},
			{
				Name:     "closed_by_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.Collaborators"),
			},
			{
				Name:     "closed_by_two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ClosedBy.TwoFactorAuthentication"),
			},
			{
				Name:     "closed_by_plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.Plan.Name"),
			},
			{
				Name:     "closed_by_plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.Plan.Space"),
			},
			{
				Name:     "closed_by_plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.Plan.Collaborators"),
			},
			{
				Name:     "closed_by_plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.Plan.PrivateRepos"),
			},
			{
				Name:     "closed_by_plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.Plan.FilledSeats"),
			},
			{
				Name:     "closed_by_plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ClosedBy.Plan.Seats"),
			},
			{
				Name:     "closed_by_ldap_dn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.LdapDn"),
			},
			{
				Name:        "closed_by_url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClosedBy.URL"),
			},
			{
				Name:     "closed_by_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.EventsURL"),
			},
			{
				Name:     "closed_by_following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.FollowingURL"),
			},
			{
				Name:     "closed_by_followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.FollowersURL"),
			},
			{
				Name:     "closed_by_gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.GistsURL"),
			},
			{
				Name:     "closed_by_organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.OrganizationsURL"),
			},
			{
				Name:     "closed_by_received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.ReceivedEventsURL"),
			},
			{
				Name:     "closed_by_repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.ReposURL"),
			},
			{
				Name:     "closed_by_starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.StarredURL"),
			},
			{
				Name:     "closed_by_subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.SubscriptionsURL"),
			},
			{
				Name:        "closed_by_text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveIssuesClosedByTextMatches,
			},
			{
				Name:        "closed_by_permissions",
				Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ClosedBy.Permissions"),
			},
			{
				Name:     "closed_by_role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClosedBy.RoleName"),
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
				Name:     "milestone_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.URL"),
			},
			{
				Name:     "milestone_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.HTMLURL"),
			},
			{
				Name:     "milestone_labels_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.LabelsURL"),
			},
			{
				Name:     "milestone_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.ID"),
			},
			{
				Name:     "milestone_number",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Number"),
			},
			{
				Name:     "milestone_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.State"),
			},
			{
				Name:     "milestone_title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Title"),
			},
			{
				Name:     "milestone_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Description"),
			},
			{
				Name:     "milestone_creator_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Login"),
			},
			{
				Name:     "milestone_creator_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.ID"),
			},
			{
				Name:     "milestone_creator_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.NodeID"),
			},
			{
				Name:     "milestone_creator_avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.AvatarURL"),
			},
			{
				Name:     "milestone_creator_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.HTMLURL"),
			},
			{
				Name:     "milestone_creator_gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.GravatarID"),
			},
			{
				Name:     "milestone_creator_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Name"),
			},
			{
				Name:     "milestone_creator_company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Company"),
			},
			{
				Name:     "milestone_creator_blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Blog"),
			},
			{
				Name:     "milestone_creator_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Location"),
			},
			{
				Name:     "milestone_creator_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Email"),
			},
			{
				Name:     "milestone_creator_hireable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Milestone.Creator.Hireable"),
			},
			{
				Name:     "milestone_creator_bio",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Bio"),
			},
			{
				Name:     "milestone_creator_twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.TwitterUsername"),
			},
			{
				Name:     "milestone_creator_public_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.PublicRepos"),
			},
			{
				Name:     "milestone_creator_public_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.PublicGists"),
			},
			{
				Name:     "milestone_creator_followers",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.Followers"),
			},
			{
				Name:     "milestone_creator_following",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.Following"),
			},
			{
				Name:     "milestone_creator_created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Milestone.Creator.CreatedAt.Time"),
			},
			{
				Name:     "milestone_creator_updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Milestone.Creator.UpdatedAt.Time"),
			},
			{
				Name:     "milestone_creator_suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Milestone.Creator.SuspendedAt.Time"),
			},
			{
				Name:     "milestone_creator_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Type"),
			},
			{
				Name:     "milestone_creator_site_admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Milestone.Creator.SiteAdmin"),
			},
			{
				Name:     "milestone_creator_total_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.TotalPrivateRepos"),
			},
			{
				Name:     "milestone_creator_owned_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.OwnedPrivateRepos"),
			},
			{
				Name:     "milestone_creator_private_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.PrivateGists"),
			},
			{
				Name:     "milestone_creator_disk_usage",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.DiskUsage"),
			},
			{
				Name:     "milestone_creator_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.Collaborators"),
			},
			{
				Name:     "milestone_creator_two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Milestone.Creator.TwoFactorAuthentication"),
			},
			{
				Name:     "milestone_creator_plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.Plan.Name"),
			},
			{
				Name:     "milestone_creator_plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.Plan.Space"),
			},
			{
				Name:     "milestone_creator_plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.Plan.Collaborators"),
			},
			{
				Name:     "milestone_creator_plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.Plan.PrivateRepos"),
			},
			{
				Name:     "milestone_creator_plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.Plan.FilledSeats"),
			},
			{
				Name:     "milestone_creator_plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.Creator.Plan.Seats"),
			},
			{
				Name:     "milestone_creator_ldap_dn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.LdapDn"),
			},
			{
				Name:        "milestone_creator_url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Milestone.Creator.URL"),
			},
			{
				Name:     "milestone_creator_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.EventsURL"),
			},
			{
				Name:     "milestone_creator_following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.FollowingURL"),
			},
			{
				Name:     "milestone_creator_followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.FollowersURL"),
			},
			{
				Name:     "milestone_creator_gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.GistsURL"),
			},
			{
				Name:     "milestone_creator_organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.OrganizationsURL"),
			},
			{
				Name:     "milestone_creator_received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.ReceivedEventsURL"),
			},
			{
				Name:     "milestone_creator_repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.ReposURL"),
			},
			{
				Name:     "milestone_creator_starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.StarredURL"),
			},
			{
				Name:     "milestone_creator_subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.SubscriptionsURL"),
			},
			{
				Name:        "milestone_creator_text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveIssuesMilestoneCreatorTextMatches,
			},
			{
				Name:        "milestone_creator_permissions",
				Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Milestone.Creator.Permissions"),
			},
			{
				Name:     "milestone_creator_role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.Creator.RoleName"),
			},
			{
				Name:     "milestone_open_issues",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.OpenIssues"),
			},
			{
				Name:     "milestone_closed_issues",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Milestone.ClosedIssues"),
			},
			{
				Name:     "milestone_created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Milestone.CreatedAt"),
			},
			{
				Name:     "milestone_updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Milestone.UpdatedAt"),
			},
			{
				Name:     "milestone_closed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Milestone.ClosedAt"),
			},
			{
				Name:     "milestone_due_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Milestone.DueOn"),
			},
			{
				Name:     "milestone_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Milestone.NodeID"),
			},
			{
				Name:     "pull_request_links_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PullRequestLinks.URL"),
			},
			{
				Name:     "pull_request_links_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PullRequestLinks.HTMLURL"),
			},
			{
				Name:     "pull_request_links_diff_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PullRequestLinks.DiffURL"),
			},
			{
				Name:     "pull_request_links_patch_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PullRequestLinks.PatchURL"),
			},
			{
				Name:     "repository_id",
				Type:     schema.TypeBigInt,
				Resolver: resolveIssuesRepositoryId,
			},
			{
				Name:     "reactions_total_count",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Reactions.TotalCount"),
			},
			{
				Name:     "reactions_plus_one",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Reactions.PlusOne"),
			},
			{
				Name:     "reactions_laugh",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Reactions.Laugh"),
			},
			{
				Name:     "reactions_confused",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Reactions.Confused"),
			},
			{
				Name:     "reactions_heart",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Reactions.Heart"),
			},
			{
				Name:     "reactions_hooray",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Reactions.Hooray"),
			},
			{
				Name:     "reactions_rocket",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Reactions.Rocket"),
			},
			{
				Name:     "reactions_eyes",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Reactions.Eyes"),
			},
			{
				Name:     "reactions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Reactions.URL"),
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
				Resolver:    resolveIssuesTextMatches,
			},
			{
				Name:        "active_lock_reason",
				Description: "ActiveLockReason is populated only when LockReason is provided while locking the issue. Possible values are: \"off-topic\", \"too heated\", \"resolved\", and \"spam\".",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "github_issue_labels",
				Description: "Label represents a GitHub label on an Issue",
				Resolver:    fetchIssueLabels,
				Columns: []schema.Column{
					{
						Name:        "issue_cq_id",
						Description: "Unique CloudQuery ID of github_issues table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name:     "url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("URL"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "color",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "default",
						Type: schema.TypeBool,
					},
					{
						Name:     "node_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("NodeID"),
					},
				},
			},
			{
				Name:        "github_issue_assignees",
				Description: "User represents a GitHub user.",
				Resolver:    fetchIssueAssignees,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"issue_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "issue_cq_id",
						Description: "Unique CloudQuery ID of github_issues table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "issue_id",
						Description: "The id of the issue",
						Type:        schema.TypeBigInt,
						Resolver:    schema.ParentResourceFieldResolver("id"),
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
						Resolver:    resolveIssueAssigneesTextMatches,
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
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	for {
		issues, resp, err := c.Github.Issues.ListByOrg(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- issues
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func resolveIssuesUserTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*github.Issue)
	if i.User == nil {
		return nil
	}
	j, err := json.Marshal(i.User.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveIssuesAssigneeTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*github.Issue)
	if i.Assignee == nil {
		return nil
	}
	j, err := json.Marshal(i.Assignee.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveIssuesClosedByTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*github.Issue)
	if i.ClosedBy == nil {
		return nil
	}
	j, err := json.Marshal(i.ClosedBy.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveIssuesMilestoneCreatorTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*github.Issue)
	if i.Milestone == nil {
		return nil
	}
	if i.Milestone.Creator == nil {
		return nil
	}
	j, err := json.Marshal(i.Milestone.Creator.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}

func resolveIssuesRepositoryId(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*github.Issue)
	// Shouldn't occur, but for type safety
	if i.Repository == nil {
		return nil
	}
	return resource.Set(c.Name, i.Repository.ID)
}
func resolveIssuesTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*github.Issue)
	j, err := json.Marshal(i.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func fetchIssueLabels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	i := parent.Item.(*github.Issue)
	res <- i.Labels
	return nil
}
func fetchIssueAssignees(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	i := parent.Item.(*github.Issue)
	res <- i.Assignees
	return nil
}
func resolveIssueAssigneesTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*github.User)
	j, err := json.Marshal(i.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
