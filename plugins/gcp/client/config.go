package client

// Config defines Provider Configuration
type Config struct {
	ProjectFilter         string   `hcl:"project_filter,optional"`
	ProjectIDs            []string `hcl:"project_ids,optional"`
	ServiceAccountKeyJSON string   `hcl:"service_account_key_json,optional"`
}

func (c Config) Example() string {
	return `configuration {
				// Optional. Filter as described https://cloud.google.com/sdk/gcloud/reference/projects/list --filter
				// project_filter = ""
				// Optional. If not specified either using all projects accessible.
				// project_ids = [<CHANGE_THIS_TO_YOUR_PROJECT_ID>]
			}`
}
