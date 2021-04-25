package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func DeleteProjectFilter(meta schema.ClientMeta) []interface{} {
	client := meta.(*Client)
	return []interface{}{"project_id", client.ProjectId}
}
