package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:        "github_repositories",
		Description: "Repository represents a GitHub repository.",
		Resolver:    fetchRepositories,
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
				Name: "owner",
				Type: schema.TypeJSON,
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
				Name: "code_of_conduct",
				Type: schema.TypeJSON,
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
				Type: schema.TypeInt,
			},
			{
				Name: "network_count",
				Type: schema.TypeInt,
			},
			{
				Name: "open_issues_count",
				Type: schema.TypeInt,
			},
			{
				Name:        "open_issues",
				Description: "Deprecated: Replaced by OpenIssuesCount",
				Type:        schema.TypeInt,
			},
			{
				Name: "stargazers_count",
				Type: schema.TypeInt,
			},
			{
				Name: "subscribers_count",
				Type: schema.TypeInt,
			},
			{
				Name:        "watchers_count",
				Description: "Deprecated: Replaced by StargazersCount",
				Type:        schema.TypeInt,
			},
			{
				Name:        "watchers",
				Description: "Deprecated: Replaced by StargazersCount",
				Type:        schema.TypeInt,
			},
			{
				Name: "size",
				Type: schema.TypeInt,
			},
			{
				Name: "auto_init",
				Type: schema.TypeBool,
			},
			{
				Name: "parent",
				Type: schema.TypeJSON,
			},
			{
				Name: "source",
				Type: schema.TypeJSON,
			},
			{
				Name: "template_repository",
				Type: schema.TypeJSON,
			},
			{
				Name: "organization",
				Type: schema.TypeJSON,
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
				Name: "license",
				Type: schema.TypeJSON,
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
				Name: "security_and_analysis",
				Type: schema.TypeJSON,
			},
			{
				Name:        "team_id",
				Description: "Creating an organization repository",
				Type:        schema.TypeInt,
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
