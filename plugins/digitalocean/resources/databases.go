package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Databases() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_databases",
		Description:  "Database represents a DigitalOcean managed database product",
		Resolver:     fetchDatabases,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "A unique ID that can be used to identify and reference a database cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "A unique, human-readable name referring to a database cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine",
				Description: "A slug representing the database engine used for the cluster. The possible values are: \"pg\" for PostgreSQL, \"mysql\" for MySQL, \"redis\" for Redis, and \"mongodb\" for MongoDB.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EngineSlug"),
			},
			{
				Name:        "version",
				Description: "A string representing the version of the database engine in use for the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VersionSlug"),
			},
			{
				Name:        "connection_uri",
				Description: "A connection string in the format accepted by the `psql` command. This is provided as a convenience and should be able to be constructed by the other attributes.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Connection.URI"),
			},
			{
				Name:        "connection_database",
				Description: "The name of the default database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Connection.Database"),
			},
			{
				Name:        "connection_host",
				Description: "The FQDN pointing to the database cluster's current primary node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Connection.Host"),
			},
			{
				Name:        "connection_port",
				Description: "The port on which the database cluster is listening.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Connection.Port"),
			},
			{
				Name:        "connection_user",
				Description: "The default user for the database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Connection.User"),
			},
			{
				Name:        "connection_password",
				Description: "The randomly generated password for the default user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Connection.Password"),
			},
			{
				Name:        "connection_ssl",
				Description: "A boolean value indicating if the connection should be made over SSL.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Connection.SSL"),
			},
			{
				Name:        "private_connection_uri",
				Description: "A connection string in the format accepted by the `psql` command. This is provided as a convenience and should be able to be constructed by the other attributes.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateConnection.URI"),
			},
			{
				Name:        "private_connection_database",
				Description: "The name of the default database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateConnection.Database"),
			},
			{
				Name:        "private_connection_host",
				Description: "The FQDN pointing to the database cluster's current primary node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateConnection.Host"),
			},
			{
				Name:        "private_connection_port",
				Description: "The port on which the database cluster is listening.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("PrivateConnection.Port"),
			},
			{
				Name:        "private_connection_user",
				Description: "The default user for the database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateConnection.User"),
			},
			{
				Name:        "private_connection_password",
				Description: "The randomly generated password for the default user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateConnection.Password"),
			},
			{
				Name:        "private_connection_ssl",
				Description: "A boolean value indicating if the connection should be made over SSL.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PrivateConnection.SSL"),
			},
			{
				Name:        "num_nodes",
				Description: "The number of nodes in the database cluster.",
				Type:        schema.TypeBigInt,
			},
			{
				Name: "size_slug",
				Type: schema.TypeString,
			},
			{
				Name:        "db_names",
				Description: "An array of strings containing the names of databases created in the database cluster.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DBNames"),
			},
			{
				Name: "region_slug",
				Type: schema.TypeString,
			},
			{
				Name:        "status",
				Description: "A string representing the current status of the database cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "maintenance_window_day",
				Description: "The day of the week on which to apply maintenance updates.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MaintenanceWindow.Day"),
			},
			{
				Name:        "maintenance_window_hour",
				Description: "The hour in UTC at which maintenance updates will be applied in 24 hour format.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MaintenanceWindow.Hour"),
			},
			{
				Name:        "maintenance_window_pending",
				Description: "A boolean value indicating whether any maintenance is scheduled to be performed in the next window.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("MaintenanceWindow.Pending"),
			},
			{
				Name:          "maintenance_window_description",
				Description:   "A list of strings, each containing information about a pending maintenance update.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("MaintenanceWindow.Description"),
				IgnoreInTests: true,
			},
			{
				Name:        "created_at",
				Description: "A time value given in ISO8601 combined date and time format that represents when the database cluster was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "private_network_uuid",
				Description: "A string specifying the UUID of the VPC to which the database cluster will be assigned. If excluded, the cluster when creating a new database cluster, it will be assigned to your account's default VPC for the region.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PrivateNetworkUUID"),
			},
			{
				Name:          "tags",
				Description:   "An array of tags that have been applied to the database cluster.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_database_users",
				Description: "DatabaseUser represents a user in the database",
				Resolver:    fetchDatabaseUsers,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of a database user.",
						Type:        schema.TypeString,
					},
					{
						Name:        "role",
						Description: "A string representing the database user's role. The value will be either\n\"primary\" or \"normal\".\n",
						Type:        schema.TypeString,
					},
					{
						Name:        "my_sql_settings_auth_plugin",
						Description: "A string specifying the authentication method to be used for connections\nto the MySQL user account. The valid values are `mysql_native_password`\nor `caching_sha2_password`. If excluded when creating a new user, the\ndefault for the version of MySQL in use will be used. As of MySQL 8.0, the\ndefault is `caching_sha2_password`.\n",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("MySQLSettings.AuthPlugin"),
					},
				},
			},
			{
				Name:        "digitalocean_database_backups",
				Description: "DatabaseBackup represents a database backup.",
				Resolver:    fetchDatabaseBackups,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "created_at",
						Description: "A time value given in ISO8601 combined date and time format at which the backup was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "size_gigabytes",
						Description: "The size of the database backup in GBs.",
						Type:        schema.TypeFloat,
					},
				},
			},
			{
				Name:        "digitalocean_database_replicas",
				Description: "DatabaseReplica represents a read-only replica of a particular database",
				Resolver:    fetchDatabaseReplicas,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name to give the read-only replicating",
						Type:        schema.TypeString,
					},
					{
						Name:        "connection_uri",
						Description: "A connection string in the format accepted by the `psql` command. This is provided as a convenience and should be able to be constructed by the other attributes.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Connection.URI"),
					},
					{
						Name:        "connection_database",
						Description: "The name of the default database.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Connection.Database"),
					},
					{
						Name:        "connection_host",
						Description: "The FQDN pointing to the database cluster's current primary node.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Connection.Host"),
					},
					{
						Name:        "connection_port",
						Description: "The port on which the database cluster is listening.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Connection.Port"),
					},
					{
						Name:        "connection_user",
						Description: "The default user for the database.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Connection.User"),
					},
					{
						Name:        "connection_password",
						Description: "The randomly generated password for the default user.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Connection.Password"),
					},
					{
						Name:        "connection_ssl",
						Description: "A boolean value indicating if the connection should be made over SSL.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Connection.SSL"),
					},
					{
						Name:        "private_connection_uri",
						Description: "A connection string in the format accepted by the `psql` command. This is provided as a convenience and should be able to be constructed by the other attributes.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateConnection.URI"),
					},
					{
						Name:        "private_connection_database",
						Description: "The name of the default database.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateConnection.Database"),
					},
					{
						Name:        "private_connection_host",
						Description: "The FQDN pointing to the database cluster's current primary node.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateConnection.Host"),
					},
					{
						Name:        "private_connection_port",
						Description: "The port on which the database cluster is listening.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("PrivateConnection.Port"),
					},
					{
						Name:        "private_connection_user",
						Description: "The default user for the database.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateConnection.User"),
					},
					{
						Name:        "private_connection_password",
						Description: "The randomly generated password for the default user.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateConnection.Password"),
					},
					{
						Name:        "private_connection_ssl",
						Description: "A boolean value indicating if the connection should be made over SSL.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("PrivateConnection.SSL"),
					},
					{
						Name:        "region",
						Description: "A slug identifier for the region where the read-only replica will be located. If excluded, the replica will be placed in the same region as the cluster.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "A string representing the current status of the database cluster.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "A time value given in ISO8601 combined date and time format that represents when the database cluster was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "private_network_uuid",
						Description: "A string specifying the UUID of the VPC to which the read-only replica will be assigned. If excluded, the replica will be assigned to your account's default VPC for the region.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateNetworkUUID"),
					},
					{
						Name:          "tags",
						Description:   "A flat array of tag names as strings to apply to the read-only replica after it is created. Tag names can either be existing or new tags.",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:     "digitalocean_database_firewall_rules",
				Resolver: fetchDatabaseFirewallRules,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "uuid",
						Description: "A unique ID for the firewall rule itself.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UUID"),
					},
					{
						Name:        "cluster_uuid",
						Description: "A unique ID for the database cluster to which the rule is applied.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ClusterUUID"),
					},
					{
						Name:        "type",
						Description: "The type of resource that the firewall rule allows to access the database cluster.",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "A time value given in ISO8601 combined date and time format that represents when the firewall rule was created.",
						Type:        schema.TypeTimestamp,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		databases, resp, err := svc.DoClient.Databases.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- databases
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
func fetchDatabaseUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	db := parent.Item.(godo.Database)
	res <- db.Users
	return nil
}
func fetchDatabaseBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	db := parent.Item.(godo.Database)
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		backups, resp, err := svc.DoClient.Databases.ListBackups(ctx, db.ID, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- backups
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
func fetchDatabaseReplicas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	db := parent.Item.(godo.Database)
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		replicas, resp, err := svc.DoClient.Databases.ListReplicas(ctx, db.ID, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- replicas
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
func fetchDatabaseFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	db := parent.Item.(godo.Database)
	svc := meta.(*client.Client)
	rules, _, err := svc.DoClient.Databases.GetFirewallRules(ctx, db.ID)
	if err != nil {
		return err
	}
	res <- rules
	return nil
}
