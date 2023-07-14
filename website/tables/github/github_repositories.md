# Table: github_repositories

This table shows data for Github Repositories.

The composite primary key for this table is (**org**, **id**).

## Relations

The following tables depend on github_repositories:
  - [github_releases](github_releases)
  - [github_repository_branches](github_repository_branches)
  - [github_repository_dependabot_alerts](github_repository_dependabot_alerts)
  - [github_repository_dependabot_secrets](github_repository_dependabot_secrets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|id (PK)|`int64`|
|node_id|`utf8`|
|owner|`json`|
|name|`utf8`|
|full_name|`utf8`|
|description|`utf8`|
|homepage|`utf8`|
|code_of_conduct|`json`|
|default_branch|`utf8`|
|master_branch|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|pushed_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|html_url|`utf8`|
|clone_url|`utf8`|
|git_url|`utf8`|
|mirror_url|`utf8`|
|ssh_url|`utf8`|
|svn_url|`utf8`|
|language|`utf8`|
|fork|`bool`|
|forks_count|`int64`|
|network_count|`int64`|
|open_issues_count|`int64`|
|open_issues|`int64`|
|stargazers_count|`int64`|
|subscribers_count|`int64`|
|watchers_count|`int64`|
|watchers|`int64`|
|size|`int64`|
|auto_init|`bool`|
|parent|`json`|
|source|`json`|
|template_repository|`json`|
|organization|`json`|
|permissions|`json`|
|allow_rebase_merge|`bool`|
|allow_update_branch|`bool`|
|allow_squash_merge|`bool`|
|allow_merge_commit|`bool`|
|allow_auto_merge|`bool`|
|allow_forking|`bool`|
|delete_branch_on_merge|`bool`|
|use_squash_pr_title_as_default|`bool`|
|squash_merge_commit_title|`utf8`|
|squash_merge_commit_message|`utf8`|
|merge_commit_title|`utf8`|
|merge_commit_message|`utf8`|
|topics|`list<item: utf8, nullable>`|
|archived|`bool`|
|disabled|`bool`|
|license|`json`|
|private|`bool`|
|has_issues|`bool`|
|has_wiki|`bool`|
|has_pages|`bool`|
|has_projects|`bool`|
|has_downloads|`bool`|
|has_discussions|`bool`|
|is_template|`bool`|
|license_template|`utf8`|
|gitignore_template|`utf8`|
|security_and_analysis|`json`|
|team_id|`int64`|
|url|`utf8`|
|archive_url|`utf8`|
|assignees_url|`utf8`|
|blobs_url|`utf8`|
|branches_url|`utf8`|
|collaborators_url|`utf8`|
|comments_url|`utf8`|
|commits_url|`utf8`|
|compare_url|`utf8`|
|contents_url|`utf8`|
|contributors_url|`utf8`|
|deployments_url|`utf8`|
|downloads_url|`utf8`|
|events_url|`utf8`|
|forks_url|`utf8`|
|git_commits_url|`utf8`|
|git_refs_url|`utf8`|
|git_tags_url|`utf8`|
|hooks_url|`utf8`|
|issue_comment_url|`utf8`|
|issue_events_url|`utf8`|
|issues_url|`utf8`|
|keys_url|`utf8`|
|labels_url|`utf8`|
|languages_url|`utf8`|
|merges_url|`utf8`|
|milestones_url|`utf8`|
|notifications_url|`utf8`|
|pulls_url|`utf8`|
|releases_url|`utf8`|
|stargazers_url|`utf8`|
|statuses_url|`utf8`|
|subscribers_url|`utf8`|
|subscription_url|`utf8`|
|tags_url|`utf8`|
|trees_url|`utf8`|
|teams_url|`utf8`|
|text_matches|`json`|
|visibility|`utf8`|
|role_name|`utf8`|