package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/codegen/recipes"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/spaces"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/digitalocean/godo"
)

var Resources = []*recipes.Resource{
	{
		Service: "account",
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
		Service: "cdn",
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
		Service: "balance",
	},
	{
		Service: "certificates",
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
		SubService: "backups",
	},
	{
		Service:    "databases",
		SubService: "replicas",
	},
	{
		Service:    "databases",
		SubService: "firewall_rules",
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
				Type:     schema.TypeString,
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
		Service:    "droplets",
		SubService: "neighbors",
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
		Service:   "registry",
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
		Service:    "registry",
		SubService: "repositories",
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
	},
	{
		Service: "snapshots",
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
		Struct:    spaces.WrappedBucket{},
		Multiplex: "client.SpacesRegionMultiplex",
		Imports:   []string{"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"},
		Relations: []string{"Cors()"},
	},
	{
		Service:    "spaces",
		SubService: "cors",
		Struct:     types.CORSRule{},
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
		Service:    "storage",
		SubService: "volumes",
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
