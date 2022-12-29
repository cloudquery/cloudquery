// Code generated by codegen; DO NOT EDIT.

package appengine

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Apps() *schema.Table {
	return &schema.Table{
		Name:        "gcp_appengine_apps",
		Description: `https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#Application`,
		Resolver:    fetchApps,
		Multiplex:   client.ProjectMultiplexEnabledServices("appengine.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "dispatch_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DispatchRules"),
			},
			{
				Name:     "auth_domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthDomain"),
			},
			{
				Name:     "location_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LocationId"),
			},
			{
				Name:     "code_bucket",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CodeBucket"),
			},
			{
				Name:     "default_cookie_expiration",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProtoDuration("DefaultCookieExpiration"),
			},
			{
				Name:     "serving_status",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("ServingStatus"),
			},
			{
				Name:     "default_hostname",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultHostname"),
			},
			{
				Name:     "default_bucket",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultBucket"),
			},
			{
				Name:     "service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceAccount"),
			},
			{
				Name:     "iap",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Iap"),
			},
			{
				Name:     "gcr_domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GcrDomain"),
			},
			{
				Name:     "database_type",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("DatabaseType"),
			},
			{
				Name:     "feature_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FeatureSettings"),
			},
		},
	}
}
