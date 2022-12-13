package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/resources/services/core"
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() []*schema.Table {
	return []*schema.Table{core.Projects()}
}
