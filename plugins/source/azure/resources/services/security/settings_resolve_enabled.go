package security

import (
	"context"

	security "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/plugin-sdk/schema"
)

func resolveEnabled(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	enabled := true
	item := resource.Item
	switch t := item.(type) {
	case *security.AlertSyncSettings:
		enabled = *t.Properties.Enabled
	case *security.DataExportSettings:
		enabled = *t.Properties.Enabled
	}
	return resource.Set(c.Name, enabled)
}
