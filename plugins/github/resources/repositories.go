package resources

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/go-github/v45/github"
)

//go:generate cq-gen --resource  --config respositories.hcl --output .
func Repositories() *schema.Table {
	return &schema.Table{
		Name:        "github_repositories",
		Description: "Repository represents a GitHub repository.",
		Resolver:    fetchRepositories,
		Multiplex:   client.OrgMultiplex,
		IgnoreError: client.IgnoreError,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
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
				Name:     "owner_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Login"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.ID"),
			},
			{
				Name:     "owner_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.NodeID"),
			},
			{
				Name:     "owner_avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.AvatarURL"),
			},
			{
				Name:     "owner_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.HTMLURL"),
			},
			{
				Name:     "owner_gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.GravatarID"),
			},
			{
				Name:     "owner_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Name"),
			},
			{
				Name:     "owner_company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Company"),
			},
			{
				Name:     "owner_blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Blog"),
			},
			{
				Name:     "owner_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Location"),
			},
			{
				Name:     "owner_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Email"),
			},
			{
				Name:     "owner_hireable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Owner.Hireable"),
			},
			{
				Name:     "owner_bio",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Bio"),
			},
			{
				Name:     "owner_twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.TwitterUsername"),
			},
			{
				Name:     "owner_public_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.PublicRepos"),
			},
			{
				Name:     "owner_public_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.PublicGists"),
			},
			{
				Name:     "owner_followers",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.Followers"),
			},
			{
				Name:     "owner_following",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.Following"),
			},
			{
				Name:     "owner_created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Owner.CreatedAt.Time"),
			},
			{
				Name:     "owner_updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Owner.UpdatedAt.Time"),
			},
			{
				Name:     "owner_suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Owner.SuspendedAt.Time"),
			},
			{
				Name:     "owner_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Type"),
			},
			{
				Name:     "owner_site_admin",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Owner.SiteAdmin"),
			},
			{
				Name:     "owner_total_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.TotalPrivateRepos"),
			},
			{
				Name:     "owner_owned_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.OwnedPrivateRepos"),
			},
			{
				Name:     "owner_private_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.PrivateGists"),
			},
			{
				Name:     "owner_disk_usage",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.DiskUsage"),
			},
			{
				Name:     "owner_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.Collaborators"),
			},
			{
				Name:     "owner_two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Owner.TwoFactorAuthentication"),
			},
			{
				Name:     "owner_plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Plan.Name"),
			},
			{
				Name:     "owner_plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.Plan.Space"),
			},
			{
				Name:     "owner_plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.Plan.Collaborators"),
			},
			{
				Name:     "owner_plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.Plan.PrivateRepos"),
			},
			{
				Name:     "owner_plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.Plan.FilledSeats"),
			},
			{
				Name:     "owner_plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Owner.Plan.Seats"),
			},
			{
				Name:     "owner_ldap_dn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.LdapDn"),
			},
			{
				Name:        "owner_url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Owner.URL"),
			},
			{
				Name:     "owner_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.EventsURL"),
			},
			{
				Name:     "owner_following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.FollowingURL"),
			},
			{
				Name:     "owner_followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.FollowersURL"),
			},
			{
				Name:     "owner_gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.GistsURL"),
			},
			{
				Name:     "owner_organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.OrganizationsURL"),
			},
			{
				Name:     "owner_received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.ReceivedEventsURL"),
			},
			{
				Name:     "owner_repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.ReposURL"),
			},
			{
				Name:     "owner_starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.StarredURL"),
			},
			{
				Name:     "owner_subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.SubscriptionsURL"),
			},
			{
				Name:        "owner_text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveRepositoriesOwnerTextMatches,
			},
			{
				Name:        "owner_permissions",
				Description: "Permissions and RoleName identify the permissions and role that a user has on a given repository",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Owner.Permissions"),
			},
			{
				Name:     "owner_role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.RoleName"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "full_name",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "homepage",
				Type: schema.TypeString,
			},
			{
				Name:     "code_of_conduct_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CodeOfConduct.Name"),
			},
			{
				Name:     "code_of_conduct_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CodeOfConduct.Key"),
			},
			{
				Name:     "code_of_conduct_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CodeOfConduct.URL"),
			},
			{
				Name:     "code_of_conduct_body",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CodeOfConduct.Body"),
			},
			{
				Name: "default_branch",
				Type: schema.TypeString,
			},
			{
				Name: "master_branch",
				Type: schema.TypeString,
			},
			{
				Name:     "created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt.Time"),
			},
			{
				Name:     "pushed_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("PushedAt.Time"),
			},
			{
				Name:     "updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt.Time"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "clone_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloneURL"),
			},
			{
				Name:     "git_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GitURL"),
			},
			{
				Name:     "mirror_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MirrorURL"),
			},
			{
				Name:     "ssh_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SSHURL"),
			},
			{
				Name:     "svn_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SVNURL"),
			},
			{
				Name: "language",
				Type: schema.TypeString,
			},
			{
				Name: "fork",
				Type: schema.TypeBool,
			},
			{
				Name: "forks_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "network_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "open_issues_count",
				Type: schema.TypeBigInt,
			},
			{
				Name:        "open_issues",
				Description: "Deprecated: Replaced by OpenIssuesCount",
				Type:        schema.TypeBigInt,
			},
			{
				Name: "stargazers_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "subscribers_count",
				Type: schema.TypeBigInt,
			},
			{
				Name:        "watchers_count",
				Description: "Deprecated: Replaced by StargazersCount",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "watchers",
				Description: "Deprecated: Replaced by StargazersCount",
				Type:        schema.TypeBigInt,
			},
			{
				Name: "size",
				Type: schema.TypeBigInt,
			},
			{
				Name: "auto_init",
				Type: schema.TypeBool,
			},
			{
				Name:     "parent",
				Type:     schema.TypeBigInt,
				Resolver: resolveRepositoriesParent,
			},
			{
				Name:     "source",
				Type:     schema.TypeBigInt,
				Resolver: resolveRepositoriesSource,
			},
			{
				Name:     "template_repository",
				Type:     schema.TypeBigInt,
				Resolver: resolveRepositoriesTemplateRepository,
			},
			{
				Name:     "organization_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Login"),
			},
			{
				Name:     "organization_id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.ID"),
			},
			{
				Name:     "organization_node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.NodeID"),
			},
			{
				Name:     "organization_avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.AvatarURL"),
			},
			{
				Name:     "organization_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.HTMLURL"),
			},
			{
				Name:     "organization_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Name"),
			},
			{
				Name:     "organization_company",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Company"),
			},
			{
				Name:     "organization_blog",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Blog"),
			},
			{
				Name:     "organization_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Location"),
			},
			{
				Name:     "organization_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Email"),
			},
			{
				Name:     "organization_twitter_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.TwitterUsername"),
			},
			{
				Name:     "organization_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Description"),
			},
			{
				Name:     "organization_public_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.PublicRepos"),
			},
			{
				Name:     "organization_public_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.PublicGists"),
			},
			{
				Name:     "organization_followers",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.Followers"),
			},
			{
				Name:     "organization_following",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.Following"),
			},
			{
				Name:     "organization_created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Organization.CreatedAt"),
			},
			{
				Name:     "organization_updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Organization.UpdatedAt"),
			},
			{
				Name:     "organization_total_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.TotalPrivateRepos"),
			},
			{
				Name:     "organization_owned_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.OwnedPrivateRepos"),
			},
			{
				Name:     "organization_private_gists",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.PrivateGists"),
			},
			{
				Name:     "organization_disk_usage",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.DiskUsage"),
			},
			{
				Name:     "organization_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.Collaborators"),
			},
			{
				Name:     "organization_billing_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.BillingEmail"),
			},
			{
				Name:     "organization_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Type"),
			},
			{
				Name:     "organization_plan_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.Plan.Name"),
			},
			{
				Name:     "organization_plan_space",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.Plan.Space"),
			},
			{
				Name:     "organization_plan_collaborators",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.Plan.Collaborators"),
			},
			{
				Name:     "organization_plan_private_repos",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.Plan.PrivateRepos"),
			},
			{
				Name:     "organization_plan_filled_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.Plan.FilledSeats"),
			},
			{
				Name:     "organization_plan_seats",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Organization.Plan.Seats"),
			},
			{
				Name:     "organization_two_factor_requirement_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Organization.TwoFactorRequirementEnabled"),
			},
			{
				Name:     "organization_is_verified",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Organization.IsVerified"),
			},
			{
				Name:     "organization_has_organization_projects",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Organization.HasOrganizationProjects"),
			},
			{
				Name:     "organization_has_repository_projects",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Organization.HasRepositoryProjects"),
			},
			{
				Name:        "organization_default_repo_permission",
				Description: "DefaultRepoPermission can be one of: \"read\", \"write\", \"admin\", or \"none\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Organization.DefaultRepoPermission"),
			},
			{
				Name:        "organization_default_repo_settings",
				Description: "DefaultRepoSettings can be one of: \"read\", \"write\", \"admin\", or \"none\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Organization.DefaultRepoSettings"),
			},
			{
				Name:        "organization_members_can_create_repos",
				Description: "MembersCanCreateRepos default value is true and is only used in Organizations.Edit.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Organization.MembersCanCreateRepos"),
			},
			{
				Name:        "organization_members_can_create_public_repos",
				Description: "https://developer.github.com/changes/2019-12-03-internal-visibility-changes/#rest-v3-api",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Organization.MembersCanCreatePublicRepos"),
			},
			{
				Name:     "organization_members_can_create_private_repos",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Organization.MembersCanCreatePrivateRepos"),
			},
			{
				Name:     "organization_members_can_create_internal_repos",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Organization.MembersCanCreateInternalRepos"),
			},
			{
				Name:        "organization_members_can_fork_private_repos",
				Description: "MembersCanForkPrivateRepos toggles whether organization members can fork private organization repositories.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Organization.MembersCanForkPrivateRepos"),
			},
			{
				Name:        "organization_members_allowed_repository_creation_type",
				Description: "MembersAllowedRepositoryCreationType denotes if organization members can create repositories and the type of repositories they can create",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Organization.MembersAllowedRepositoryCreationType"),
			},
			{
				Name:        "organization_members_can_create_pages",
				Description: "MembersCanCreatePages toggles whether organization members can create GitHub Pages sites.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Organization.MembersCanCreatePages"),
			},
			{
				Name:        "organization_members_can_create_public_pages",
				Description: "MembersCanCreatePublicPages toggles whether organization members can create public GitHub Pages sites.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Organization.MembersCanCreatePublicPages"),
			},
			{
				Name:        "organization_members_can_create_private_pages",
				Description: "MembersCanCreatePrivatePages toggles whether organization members can create private GitHub Pages sites.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Organization.MembersCanCreatePrivatePages"),
			},
			{
				Name:        "organization_url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Organization.URL"),
			},
			{
				Name:     "organization_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.EventsURL"),
			},
			{
				Name:     "organization_hooks_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.HooksURL"),
			},
			{
				Name:     "organization_issues_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.IssuesURL"),
			},
			{
				Name:     "organization_members_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.MembersURL"),
			},
			{
				Name:     "organization_public_members_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.PublicMembersURL"),
			},
			{
				Name:     "organization_repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Organization.ReposURL"),
			},
			{
				Name: "permissions",
				Type: schema.TypeJSON,
			},
			{
				Name: "allow_rebase_merge",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_update_branch",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_squash_merge",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_merge_commit",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_auto_merge",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_forking",
				Type: schema.TypeBool,
			},
			{
				Name: "delete_branch_on_merge",
				Type: schema.TypeBool,
			},
			{
				Name:     "use_squash_pr_title_as_default",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("UseSquashPRTitleAsDefault"),
			},
			{
				Name: "topics",
				Type: schema.TypeStringArray,
			},
			{
				Name: "archived",
				Type: schema.TypeBool,
			},
			{
				Name: "disabled",
				Type: schema.TypeBool,
			},
			{
				Name:     "license_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("License.Key"),
			},
			{
				Name:     "license_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("License.Name"),
			},
			{
				Name:     "license_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("License.URL"),
			},
			{
				Name:     "license_spdx_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("License.SPDXID"),
			},
			{
				Name:     "license_html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("License.HTMLURL"),
			},
			{
				Name:     "license_featured",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("License.Featured"),
			},
			{
				Name:     "license_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("License.Description"),
			},
			{
				Name:     "license_implementation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("License.Implementation"),
			},
			{
				Name:     "license_permissions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("License.Permissions"),
			},
			{
				Name:     "license_conditions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("License.Conditions"),
			},
			{
				Name:     "license_limitations",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("License.Limitations"),
			},
			{
				Name:     "license_body",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("License.Body"),
			},
			{
				Name:        "private",
				Description: "Additional mutable fields when creating and editing a repository",
				Type:        schema.TypeBool,
			},
			{
				Name: "has_issues",
				Type: schema.TypeBool,
			},
			{
				Name: "has_wiki",
				Type: schema.TypeBool,
			},
			{
				Name: "has_pages",
				Type: schema.TypeBool,
			},
			{
				Name: "has_projects",
				Type: schema.TypeBool,
			},
			{
				Name: "has_downloads",
				Type: schema.TypeBool,
			},
			{
				Name: "is_template",
				Type: schema.TypeBool,
			},
			{
				Name: "license_template",
				Type: schema.TypeString,
			},
			{
				Name: "gitignore_template",
				Type: schema.TypeString,
			},
			{
				Name:     "security_and_analysis_advanced_security_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityAndAnalysis.AdvancedSecurity.Status"),
			},
			{
				Name:     "security_and_analysis_secret_scanning_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityAndAnalysis.SecretScanning.Status"),
			},
			{
				Name:        "team_id",
				Description: "Creating an organization repository",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("TeamID"),
			},
			{
				Name:        "url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("URL"),
			},
			{
				Name:     "archive_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ArchiveURL"),
			},
			{
				Name:     "assignees_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssigneesURL"),
			},
			{
				Name:     "blobs_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BlobsURL"),
			},
			{
				Name:     "branches_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BranchesURL"),
			},
			{
				Name:     "collaborators_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CollaboratorsURL"),
			},
			{
				Name:     "comments_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CommentsURL"),
			},
			{
				Name:     "commits_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CommitsURL"),
			},
			{
				Name:     "compare_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CompareURL"),
			},
			{
				Name:     "contents_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContentsURL"),
			},
			{
				Name:     "contributors_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContributorsURL"),
			},
			{
				Name:     "deployments_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentsURL"),
			},
			{
				Name:     "downloads_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DownloadsURL"),
			},
			{
				Name:     "events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventsURL"),
			},
			{
				Name:     "forks_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ForksURL"),
			},
			{
				Name:     "git_commits_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GitCommitsURL"),
			},
			{
				Name:     "git_refs_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GitRefsURL"),
			},
			{
				Name:     "git_tags_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GitTagsURL"),
			},
			{
				Name:     "hooks_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HooksURL"),
			},
			{
				Name:     "issue_comment_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IssueCommentURL"),
			},
			{
				Name:     "issue_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IssueEventsURL"),
			},
			{
				Name:     "issues_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IssuesURL"),
			},
			{
				Name:     "keys_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeysURL"),
			},
			{
				Name:     "labels_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelsURL"),
			},
			{
				Name:     "languages_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LanguagesURL"),
			},
			{
				Name:     "merges_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MergesURL"),
			},
			{
				Name:     "milestones_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MilestonesURL"),
			},
			{
				Name:     "notifications_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NotificationsURL"),
			},
			{
				Name:     "pulls_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PullsURL"),
			},
			{
				Name:     "releases_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReleasesURL"),
			},
			{
				Name:     "stargazers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StargazersURL"),
			},
			{
				Name:     "statuses_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusesURL"),
			},
			{
				Name:     "subscribers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscribersURL"),
			},
			{
				Name:     "subscription_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionURL"),
			},
			{
				Name:     "tags_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TagsURL"),
			},
			{
				Name:     "trees_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TreesURL"),
			},
			{
				Name:     "teams_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TeamsURL"),
			},
			{
				Name:        "text_matches",
				Description: "TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata",
				Type:        schema.TypeJSON,
				Resolver:    resolveRepositoriesTextMatches,
			},
			{
				Name:        "visibility",
				Description: "Visibility is only used for Create and Edit endpoints",
				Type:        schema.TypeString,
			},
			{
				Name:        "role_name",
				Description: "RoleName is only returned by the API 'check team permissions for a repository'. See: teams.go (IsTeamRepoByID) https://docs.github.com/en/rest/teams/teams#check-team-permissions-for-a-repository",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	}
	for {
		repos, resp, err := c.Github.Repositories.ListByOrg(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- repos
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func resolveRepositoriesOwnerTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.Repository)
	j, err := json.Marshal(u.Owner.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveRepositoriesParent(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.Repository)
	if u.Parent == nil {
		return nil
	}
	return resource.Set(c.Name, u.Parent.ID)
}
func resolveRepositoriesSource(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.Repository)
	if u.Source == nil {
		return nil
	}
	return resource.Set(c.Name, u.Source.ID)
}
func resolveRepositoriesTemplateRepository(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.Repository)
	if u.TemplateRepository == nil {
		return nil
	}
	return resource.Set(c.Name, u.TemplateRepository.ID)
}
func resolveRepositoriesTextMatches(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	u := resource.Item.(*github.Repository)
	j, err := json.Marshal(u.TextMatches)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
