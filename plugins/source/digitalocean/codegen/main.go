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
		Service: "accounts",
		Struct:  godo.Account{},
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
		Service: "cdns",
		Struct:  godo.CDN{},
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
		Service: "billing_history",
		Struct:  godo.BillingHistoryEntry{},
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
		Service:    "monitoring",
		SubService: "alert_policies",
		Struct:     godo.AlertPolicy{},
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
		Service: "balances",
		Struct:  godo.Balance{},
	},
	{
		Service: "certificates",
		Struct:  godo.Certificate{},
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
		Service:   "databases",
		Struct:    godo.Database{},
		Relations: []string{"FirewallRules()", "Replicas()", "Backups()"},
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
		Service:    "databases",
		Struct:     godo.DatabaseBackup{},
		SubService: "backups",
	},
	{
		Service:    "databases",
		SubService: "replicas",
		Struct:     godo.DatabaseReplica{},
	},
	{
		Service:    "databases",
		SubService: "firewall_rules",
		Struct:     godo.DatabaseFirewallRule{},
	},
	{
		Service:   "domains",
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
		Service:    "domains",
		SubService: "records",
		Struct:     godo.DomainRecord{},
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
		Service:   "droplets",
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
				Type:     schema.TypeStringArray,
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
		Service:    "droplets",
		SubService: "neighbors",
		Struct:     &droplets.NeighborWrapper{},
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
		Service: "firewalls",
		Struct:  godo.Firewall{},
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
		Service: "floating_ips",
		Struct:  godo.FloatingIP{},
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
		Service: "images",
		Struct:  godo.Image{},
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
		Service: "keys",
		Struct:  godo.Key{},
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
		Service:    "projects",
		SubService: "resources",
		Struct:     godo.ProjectResource{},
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
		Service:   "registries",
		Struct:    &godo.Registry{},
		Relations: []string{"Repositories()"},
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
		Service:    "registries",
		SubService: "repositories",
		Struct:     &godo.Repository{},
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
		Service: "sizes",
		Struct:  godo.Size{},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "slug",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Slug")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		SkipFields: []string{"Slug"},
	},
	{
		Service: "snapshots",
		Struct:  godo.Snapshot{},
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
		Service:      "spaces",
		Struct:       spaces.WrappedBucket{},
		PostResolver: "resolveSpaceAttributes",
		Multiplex:    "client.SpacesRegionMultiplex",
		Imports:      []string{"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"},
		Relations:    []string{"Cors()"},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "acls",
				Type:     schema.TypeJSON,
				Resolver: `schema.PathResolver("ACLs")`,
			},
		},
		SkipFields: []string{"ACLs"},
	},
	{
		Service:    "spaces",
		SubService: "cors",
		Struct:     types.CORSRule{},
		// 'id' is not a PK - it can be `null` when arriving from the upstream API.
	},
	{
		Service:    "storage",
		SubService: "volumes",
		Struct:     godo.Volume{},
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
		Service:   "vpcs",
		Struct:    &godo.VPC{},
		Relations: []string{"Members()"},
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
		Service:    "vpcs",
		SubService: "members",
		Struct:     &godo.VPCMember{},
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
