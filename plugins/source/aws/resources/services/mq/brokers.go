// Code generated by codegen; DO NOT EDIT.

package mq

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Brokers() *schema.Table {
	return &schema.Table{
		Name:                "aws_mq_brokers",
		Description:         "https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers.html",
		Resolver:            fetchMqBrokers,
		PreResourceResolver: getMqBroker,
		Multiplex:           client.ServiceAccountRegionMultiplexer("mq"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BrokerArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "actions_required",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActionsRequired"),
			},
			{
				Name:     "authentication_strategy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthenticationStrategy"),
			},
			{
				Name:     "auto_minor_version_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoMinorVersionUpgrade"),
			},
			{
				Name:     "broker_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BrokerId"),
			},
			{
				Name:     "broker_instances",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BrokerInstances"),
			},
			{
				Name:     "broker_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BrokerName"),
			},
			{
				Name:     "broker_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BrokerState"),
			},
			{
				Name:     "configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Configurations"),
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "deployment_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentMode"),
			},
			{
				Name:     "encryption_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionOptions"),
			},
			{
				Name:     "engine_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineType"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "host_instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostInstanceType"),
			},
			{
				Name:     "ldap_server_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LdapServerMetadata"),
			},
			{
				Name:     "logs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Logs"),
			},
			{
				Name:     "maintenance_window_start_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MaintenanceWindowStartTime"),
			},
			{
				Name:     "pending_authentication_strategy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingAuthenticationStrategy"),
			},
			{
				Name:     "pending_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingEngineVersion"),
			},
			{
				Name:     "pending_host_instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingHostInstanceType"),
			},
			{
				Name:     "pending_ldap_server_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingLdapServerMetadata"),
			},
			{
				Name:     "pending_security_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PendingSecurityGroups"),
			},
			{
				Name:     "publicly_accessible",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PubliclyAccessible"),
			},
			{
				Name:     "security_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SecurityGroups"),
			},
			{
				Name:     "storage_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageType"),
			},
			{
				Name:     "subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubnetIds"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "users",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Users"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},

		Relations: []*schema.Table{
			BrokerConfigurations(),
			BrokerUsers(),
		},
	}
}
