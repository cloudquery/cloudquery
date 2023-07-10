package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var customExceptions = map[string]string{
	"pagerduty": "PagerDuty",
}

func titleTransformer(table *schema.Table) error {
	if table.Title != "" {
		// respect hard-coded titles
		return nil
	}
	exceptions := make(map[string]string)
	for k, v := range docs.DefaultTitleExceptions {
		exceptions[k] = v
	}
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	table.Title = csr.ToTitle(table.Name)
	return nil
}
