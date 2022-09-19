// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ManagedInstances() *schema.Table {
	return &schema.Table{
		Name:      "azure_sql_managed_instances",
		Resolver:  fetchSQLManagedInstances,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "managed_instance_create_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ManagedInstanceCreateMode"),
			},
			{
				Name:     "fully_qualified_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FullyQualifiedDomainName"),
			},
			{
				Name:     "administrator_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdministratorLogin"),
			},
			{
				Name:     "administrator_login_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdministratorLoginPassword"),
			},
			{
				Name:     "subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetID"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "license_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LicenseType"),
			},
			{
				Name:     "v_cores",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VCores"),
			},
			{
				Name:     "storage_size_in_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StorageSizeInGB"),
			},
			{
				Name:     "collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Collation"),
			},
			{
				Name:     "dns_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DNSZone"),
			},
			{
				Name:     "dns_zone_partner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DNSZonePartner"),
			},
			{
				Name:     "public_data_endpoint_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PublicDataEndpointEnabled"),
			},
			{
				Name:     "source_managed_instance_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceManagedInstanceID"),
			},
			{
				Name:     "restore_point_in_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RestorePointInTime"),
			},
			{
				Name:     "proxy_override",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProxyOverride"),
			},
			{
				Name:     "timezone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TimezoneID"),
			},
			{
				Name:     "instance_pool_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstancePoolID"),
			},
			{
				Name:     "maintenance_configuration_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MaintenanceConfigurationID"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpointConnections"),
			},
			{
				Name:     "minimal_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinimalTLSVersion"),
			},
			{
				Name:     "storage_account_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageAccountType"),
			},
			{
				Name:     "zone_redundant",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ZoneRedundant"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},

		Relations: []*schema.Table{
			managedDatabases(),
			managedInstanceVulnerabilityAssessments(),
			managedInstanceEncryptionProtectors(),
		},
	}
}

func fetchSQLManagedInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedInstances

	response, err := svc.List(ctx)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
