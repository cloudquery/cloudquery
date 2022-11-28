# Table: github_repositories



The composite primary key for this table is (**org**, **id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|id (PK)|Int|
|created_at|Timestamp|
|pushed_at|Timestamp|
|updated_at|Timestamp|
|node_id|String|
|owner|JSON|
|name|String|
|full_name|String|
|description|String|
|homepage|String|
|code_of_conduct|JSON|
|default_branch|String|
|master_branch|String|
|html_url|String|
|clone_url|String|
|git_url|String|
|mirror_url|String|
|ssh_url|String|
|svn_url|String|
|language|String|
|fork|Bool|
|forks_count|Int|
|network_count|Int|
|open_issues_count|Int|
|open_issues|Int|
|stargazers_count|Int|
|subscribers_count|Int|
|watchers_count|Int|
|watchers|Int|
|size|Int|
|auto_init|Bool|
|parent|JSON|
|source|JSON|
|template_repository|JSON|
|organization|JSON|
|permissions|JSON|
|allow_rebase_merge|Bool|
|allow_update_branch|Bool|
|allow_squash_merge|Bool|
|allow_merge_commit|Bool|
|allow_auto_merge|Bool|
|allow_forking|Bool|
|delete_branch_on_merge|Bool|
|use_squash_pr_title_as_default|Bool|
|topics|StringArray|
|archived|Bool|
|disabled|Bool|
|license|JSON|
|private|Bool|
|has_issues|Bool|
|has_wiki|Bool|
|has_pages|Bool|
|has_projects|Bool|
|has_downloads|Bool|
|is_template|Bool|
|license_template|String|
|gitignore_template|String|
|security_and_analysis|JSON|
|team_id|Int|
|url|String|
|archive_url|String|
|assignees_url|String|
|blobs_url|String|
|branches_url|String|
|collaborators_url|String|
|comments_url|String|
|commits_url|String|
|compare_url|String|
|contents_url|String|
|contributors_url|String|
|deployments_url|String|
|downloads_url|String|
|events_url|String|
|forks_url|String|
|git_commits_url|String|
|git_refs_url|String|
|git_tags_url|String|
|hooks_url|String|
|issue_comment_url|String|
|issue_events_url|String|
|issues_url|String|
|keys_url|String|
|labels_url|String|
|languages_url|String|
|merges_url|String|
|milestones_url|String|
|notifications_url|String|
|pulls_url|String|
|releases_url|String|
|stargazers_url|String|
|statuses_url|String|
|subscribers_url|String|
|subscription_url|String|
|tags_url|String|
|trees_url|String|
|teams_url|String|
|text_matches|JSON|
|visibility|String|
|role_name|String|