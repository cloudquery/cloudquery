package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/codegen/recipes"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/droplets"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/spaces"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/digitalocean/godo"
)

var Resources = []*recipes.Resource{
	{
		Service:      "account",
		Template:     "resource_get",
		MockTemplate: "resource_get_mock",
		Struct:       godo.Account{},
		ChildTable:   false,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("UUID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"UUID"},
	},
	{
		Service:      "cdn",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		Struct:       godo.CDN{},
		ChildTable:   false,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:      "billing_history",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		ResponsePath: ".BillingHistory",
		Struct:       godo.BillingHistoryEntry{},
		MockStruct:   godo.BillingHistory{},
		ChildTable:   false,
		MockWrapper:  true,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "invoice_id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("InvoiceID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"InvoiceID"},
	},
	{
		Service:      "monitoring",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		SubService:   "alert_policies",
		Struct:       godo.AlertPolicy{},
		ChildTable:   false,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("UUID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"UUID"},
	},

	{
		Service:      "balance",
		Template:     "resource_get",
		MockTemplate: "resource_get_mock",
		Struct:       godo.Balance{},
		ChildTable:   false,
	},
	{
		Service:      "certificates",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		Struct:       godo.Certificate{},
		ChildTable:   false,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:      "databases",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		Struct:       godo.Database{},
		Relations:    []string{"FirewallRules()", "Replicas()", "Backups()"},
		ChildTable:   false,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:      "databases",
		SubService:   "backups",
		Template:     "resource_list",
		MockTemplate: "resource_list_child_mock",
		Struct:       godo.DatabaseBackup{},
		ParentStruct: godo.Database{},
		ChildTable:   true,
	},
	{
		Service:      "databases",
		SubService:   "replicas",
		Template:     "resource_list",
		MockTemplate: "resource_list_child_mock",
		Struct:       godo.DatabaseReplica{},
		ParentStruct: godo.Database{},
		ChildTable:   true,
	},
	{
		Service:      "databases",
		SubService:   "firewall_rules",
		Template:     "resource_get",
		MockTemplate: "resource_get_child_mock",
		Struct:       godo.DatabaseFirewallRule{},
		ParentStruct: godo.Database{},
		ChildTable:   true,
	},
	{
		Service:   "domains",
		Template:  "resource_list",
		Struct:    godo.Domain{},
		Relations: []string{"Records()"},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Name")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"Name"},
	},
	{
		Service:      "domains",
		SubService:   "records",
		Template:     "resource_list",
		MockTemplate: "resource_list_child_mock",
		Args:         ", p.Name",
		Struct:       godo.DomainRecord{},
		ParentStruct: godo.Domain{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:   "droplets",
		Template:  "resource_list",
		Struct:    godo.Droplet{},
		Relations: []string{"Neighbors()"},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "backup_ids",
				Type:     schema.TypeIntArray,
				Resolver: `schema.PathResolver("BackupIDs")`,
			},
			{
				Name:     "snapshot_ids",
				Type:     schema.TypeIntArray,
				Resolver: `schema.PathResolver("SnapshotIDs")`,
			},
			{
				Name:     "volume_ids",
				Type:     schema.TypeIntArray,
				Resolver: `schema.PathResolver("VolumeIDs")`,
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"BackupIDs", "SnapshotIDs", "VolumeIDs", "ID"},
	},
	{
		Service:      "droplets",
		SubService:   "neighbors",
		Template:     "resource_list",
		MockTemplate: "resource_list_child_mock",
		Struct:       &droplets.NeighborWrapper{},
		ParentStruct: &godo.Droplet{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "neighbor_id",
				Type:     schema.TypeInt,
				Resolver: `schema.PathResolver("NeighborId")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"NeighborId"},
	},
	{
		Service:  "firewalls",
		Template: "resource_list",
		Struct:   godo.Firewall{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "droplet_ids",
				Type:     schema.TypeIntArray,
				Resolver: `schema.PathResolver("DropletIDs")`,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"DropletIDs", "ID"},
	},

	{
		Service:  "floating_ips",
		Template: "resource_list",
		Struct:   godo.FloatingIP{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "ip",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("IP")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"IP"},
	},
	{
		Service:  "images",
		Template: "resource_list",
		Struct:   godo.Image{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:  "keys",
		Template: "resource_list",
		Struct:   godo.Key{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:   "projects",
		Template:  "resource_list",
		Struct:    godo.Project{},
		Relations: []string{"Resources()"},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:      "projects",
		SubService:   "resources",
		Template:     "resource_list",
		MockTemplate: "resource_list_child_mock",
		Struct:       godo.ProjectResource{},
		ParentStruct: godo.Project{},
		ChildTable:   true,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "urn",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("URN")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"URN"},
	},
	{
		Service:      "registry",
		Template:     "resource_get",
		MockTemplate: "resource_get_mock",
		Struct:       &godo.Registry{},
		ChildTable:   false,
		Relations:    []string{"Repositories()"},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Name")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"Name"},
	},
	{
		Service:      "registry",
		SubService:   "repositories",
		Args:         ", p.Name",
		Template:     "resource_list",
		MockTemplate: "resource_list_child_mock",
		Struct:       &godo.Repository{},
		ParentStruct: &godo.Registry{},
		ChildTable:   true,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Name")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"Name"},
	},
	{
		Service:      "sizes",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		Struct:       godo.Size{},
	},
	{
		Service:      "snapshots",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		Struct:       godo.Snapshot{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:   "spaces",
		Template:  "resource_list",
		Struct:    spaces.WrappedBucket{},
		Multiplex: "client.SpacesRegionMultiplex",
		Imports:   []string{"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"},
		Relations: []string{"Cors()"},
	},
	{
		Service:      "spaces",
		SubService:   "cors",
		Template:     "resource_list",
		Struct:       types.CORSRule{},
		ParentStruct: spaces.WrappedBucket{},
		ChildTable:   true,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:      "storage",
		SubService:   "volumes",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		Struct:       godo.Volume{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "droplet_ids",
				Type:     schema.TypeIntArray,
				Resolver: `schema.PathResolver("DropletIDs")`,
			},
		},
		SkipFields: []string{"DropletIDs", "ID"},
	},
	{
		Service:      "vpcs",
		Template:     "resource_list",
		MockTemplate: "resource_list_mock",
		Struct:       &godo.VPC{},
		Relations:    []string{"Members()"},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("ID")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"ID"},
	},
	{
		Service:      "vpcs",
		SubService:   "members",
		Args:         ", p.ID, nil",
		Template:     "resource_list",
		Struct:       &godo.VPCMember{},
		ParentStruct: &godo.VPC{},
		ChildTable:   true,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "urn",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("URN")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"URN"},
	},
}

func main() {
	for _, r := range Resources {
		r.Generate()
	}
}
