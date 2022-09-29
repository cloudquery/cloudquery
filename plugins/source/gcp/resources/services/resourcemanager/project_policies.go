// Code generated by codegen; DO NOT EDIT.

package resourcemanager

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ProjectPolicies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_resourcemanager_project_policies",
		Resolver:  fetchProjectPolicies,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "audit_configs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AuditConfigs"),
			},
			{
				Name:     "bindings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Bindings"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
		},
	}
}
