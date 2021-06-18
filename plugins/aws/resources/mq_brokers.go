package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MqBrokers() *schema.Table {
	return &schema.Table{
		Name:         "aws_mq_brokers",
		Resolver:     fetchMqBrokers,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "authentication_strategy",
				Description: "The authentication strategy used to secure the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "broker_arn",
				Description: "The Amazon Resource Name (ARN) of the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "broker_id",
				Description: "The unique ID that Amazon MQ generates for the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "broker_instances",
				Description: "A list of information about allocated brokers.",
				Type:        schema.TypeJSON,
				Resolver:    resolveMqBrokerBrokerInstances,
			},
			{
				Name:        "broker_name",
				Description: "The name of the broker",
				Type:        schema.TypeString,
			},
			{
				Name:        "broker_state",
				Description: "The status of the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created",
				Description: "The time when the broker was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "deployment_mode",
				Description: "The deployment mode of the broker.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encryption_options_use_aws_owned_key",
				Description: "Enables the use of an AWS owned CMK using AWS Key Management Service (KMS).",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EncryptionOptions.UseAwsOwnedKey"),
			},
			{
				Name:        "encryption_options_kms_key_id",
				Description: "The symmetric customer master key (CMK) to use for the AWS Key Management Service (KMS).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionOptions.KmsKeyId"),
			},
			{
				Name:        "engine_type",
				Description: "The type of broker engine.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The version of the broker engine",
				Type:        schema.TypeString,
			},
			{
				Name:        "host_instance_type",
				Description: "The broker's instance type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ldap_server_metadata",
				Description: "The metadata of the LDAP server used to authenticate and authorize connections to the broker.",
				Type:        schema.TypeJSON,
				Resolver:    resolveMqBrokerLdapServerMetadata,
			},
			{
				Name:        "logs",
				Description: "The list of information about logs currently enabled and pending to be deployed for the specified broker.",
				Type:        schema.TypeJSON,
				Resolver:    resolveMqBrokerLogs,
			},
			{
				Name:        "maintenance_window_start_time",
				Description: "The parameters that determine the WeeklyStartTime.",
				Type:        schema.TypeJSON,
				Resolver:    resolveMqBrokerMaintenanceWindowStartTime,
			},
			{
				Name:        "pending_authentication_strategy",
				Description: "The authentication strategy that will be applied when the broker is rebooted.",
				Type:        schema.TypeString,
			},
			{
				Name:        "pending_engine_version",
				Description: "The version of the broker engine to upgrade to",
				Type:        schema.TypeString,
			},
			{
				Name:        "pending_host_instance_type",
				Description: "The host instance type of the broker to upgrade to",
				Type:        schema.TypeString,
			},
			{
				Name:        "pending_ldap_server_metadata",
				Description: "The metadata of the LDAP server that will be used to authenticate and authorize connections to the broker once it is rebooted.",
				Type:        schema.TypeJSON,
				Resolver:    resolveMqBrokerPendingLdapServerMetadata,
			},
			{
				Name:        "pending_security_groups",
				Description: "The list of pending security groups to authorize connections to brokers.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "publicly_accessible",
				Description: "Enables connections from applications outside of the VPC that hosts the broker's subnets.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "security_groups",
				Description: "The list of security groups (1 minimum, 5 maximum) that authorizes connections to brokers.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "storage_type",
				Description: "The broker's storage type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_ids",
				Description: "The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "The list of all tags associated with this broker.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:         "aws_mq_broker_configurations",
				Resolver:     fetchMqBrokerConfigurations,
				Multiplex:    client.AccountRegionMultiplex,
				IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
				DeleteFilter: client.DeleteAccountRegionFilter,
				Columns: []schema.Column{
					{
						Name:        "broker_id",
						Description: "Unique ID of aws_mq_brokers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The AWS Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAWSAccount,
					},
					{
						Name:        "region",
						Description: "The AWS Region of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAWSRegion,
					},
					{
						Name:        "arn",
						Description: "The ARN of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "authentication_strategy",
						Description: "The authentication strategy associated with the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created",
						Description: "The date and time of the configuration revision.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "description",
						Description: "The description of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "engine_type",
						Description: "The type of broker engine.",
						Type:        schema.TypeString,
					},
					{
						Name:        "engine_version",
						Description: "The version of the broker engine.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_id",
						Description: "The unique ID that Amazon MQ generates for the configuration.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "latest_revision_created",
						Description: "The date and time of the configuration revision.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("LatestRevision.Created"),
					},
					{
						Name:        "latest_revision_description",
						Description: "The description of the configuration revision.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LatestRevision.Description"),
					},
					{
						Name:        "latest_revision",
						Description: "The revision number of the configuration.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LatestRevision.Revision"),
					},
					{
						Name:        "name",
						Description: "The name of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The list of all tags associated with this configuration.",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:         "aws_mq_broker_users",
				Resolver:     fetchMqBrokerUsers,
				Multiplex:    client.AccountRegionMultiplex,
				IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
				DeleteFilter: client.DeleteAccountRegionFilter,
				Columns: []schema.Column{
					{
						Name:        "broker_id",
						Description: "Unique ID of aws_mq_brokers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The AWS Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAWSAccount,
					},
					{
						Name:        "region",
						Description: "The AWS Region of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAWSRegion,
					},
					{
						Name:        "console_access",
						Description: "Enables access to the the ActiveMQ Web Console for the ActiveMQ user.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "groups",
						Description: "The list of groups (20 maximum) to which the ActiveMQ user belongs",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "pending",
						Description: "The status of the changes pending for the ActiveMQ user.",
						Type:        schema.TypeJSON,
						Resolver:    resolveMqBrokerUserPending,
					},
					{
						Name:        "username",
						Description: "The username of the ActiveMQ user.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMqBrokers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config mq.ListBrokersInput
	c := meta.(*client.Client)
	svc := c.Services().MQ
	for {
		response, err := svc.ListBrokers(ctx, &config, func(options *mq.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, bs := range response.BrokerSummaries {
			output, err := svc.DescribeBroker(ctx, &mq.DescribeBrokerInput{BrokerId: bs.BrokerId}, func(options *mq.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- output
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveMqBrokerBrokerInstances(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	broker, ok := resource.Item.(*mq.DescribeBrokerOutput)
	if !ok {
		return fmt.Errorf("not a DescribeBrokerOutput instance: %#v", resource.Item)
	}
	data, err := json.Marshal(broker.BrokerInstances)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveMqBrokerLdapServerMetadata(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	broker, ok := resource.Item.(*mq.DescribeBrokerOutput)
	if !ok {
		return fmt.Errorf("not a DescribeBrokerOutput instance: %#v", resource.Item)
	}
	data, err := json.Marshal(broker.LdapServerMetadata)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveMqBrokerLogs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	broker, ok := resource.Item.(*mq.DescribeBrokerOutput)
	if !ok {
		return fmt.Errorf("not a DescribeBrokerOutput instance: %#v", resource.Item)
	}
	data, err := json.Marshal(broker.Logs)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveMqBrokerMaintenanceWindowStartTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	broker, ok := resource.Item.(*mq.DescribeBrokerOutput)
	if !ok {
		return fmt.Errorf("not a DescribeBrokerOutput instance: %#v", resource.Item)
	}
	data, err := json.Marshal(broker.MaintenanceWindowStartTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveMqBrokerPendingLdapServerMetadata(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	broker, ok := resource.Item.(*mq.DescribeBrokerOutput)
	if !ok {
		return fmt.Errorf("not a DescribeBrokerOutput instance: %#v", resource.Item)
	}
	data, err := json.Marshal(broker.PendingLdapServerMetadata)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func fetchMqBrokerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().MQ
	for _, cfg := range broker.Configurations.History {
		input := mq.DescribeConfigurationInput{ConfigurationId: cfg.Id}
		output, err := svc.DescribeConfiguration(ctx, &input, func(options *mq.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output
	}
	return nil
}

func fetchMqBrokerUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().MQ
	for _, us := range broker.Users {
		input := mq.DescribeUserInput{
			BrokerId: broker.BrokerId,
			Username: us.Username,
		}
		output, err := svc.DescribeUser(ctx, &input, func(options *mq.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output
	}
	return nil
}

func resolveMqBrokerUserPending(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	user, ok := resource.Item.(*mq.DescribeUserOutput)
	if !ok {
		return fmt.Errorf("not a DescribeUserOutput instance: %#v", resource.Item)
	}
	data, err := json.Marshal(user.Pending)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
