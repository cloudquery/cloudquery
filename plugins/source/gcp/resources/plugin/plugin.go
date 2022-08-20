package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugins/source/gcp/resources/servicesv2/compute"
)

var (
	Version = "development"
)

const exampleConfig = `
kind: source
spec:
  tables: ["*"]
  spec:
    # Optional. List of folders to get projects from. Required permission: resourcemanager.projects.list
    # folder_ids:
    # 	- "organizations/<ORG_ID>"
    # 	- "folders/<FOLDER_ID>"
    # Optional. Maximum level of folders to recurse into
    # folders_max_depth: 5
    # Optional. If not specified either using all projects accessible.
    # project_ids:
    # 	- "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"
    # Optional. ServiceAccountKeyJSON passed as value instead of a file path, can be passed also via env: CQ_SERVICE_ACCOUNT_KEY_JSON
    # service_account_key_json: <YOUR_JSON_SERVICE_ACCOUNT_KEY_DATA>
    # Optional. GRPC Retry/backoff configuration, time units in seconds. Documented in https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md
    # backoff_base_delay: 1
    # backoff_multiplier: 1.6
    # backoff_max_delay: 120
    # backoff_jitter: 0.2
    # backoff_min_connect_timeout = 0
    # Optional. Max amount of retries for retrier, defaults to max 3 retries.
    # max_retries: 3
`

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"gcp",
		Version,
		[]*schema.Table{
			compute.ComputeAddresses(),
			compute.ComputeAutoscalers(),
			compute.ComputeBackendServices(),
			compute.ComputeDiskTypes(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
		plugins.WithClassifyError(client.ClassifyError),
	)
}
