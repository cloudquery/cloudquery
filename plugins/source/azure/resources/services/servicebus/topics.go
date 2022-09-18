// Auto generated code - DO NOT EDIT.

package servicebus

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
)

func topics() *schema.Table {
	return &schema.Table{
		Name:     "azure_servicebus_topics",
		Resolver: fetchServicebusTopics,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "size_in_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeInBytes"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "accessed_at",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessedAt"),
			},
			{
				Name:     "subscription_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SubscriptionCount"),
			},
			{
				Name:     "count_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CountDetails"),
			},
			{
				Name:     "default_message_time_to_live",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultMessageTimeToLive"),
			},
			{
				Name:     "max_size_in_megabytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxSizeInMegabytes"),
			},
			{
				Name:     "max_message_size_in_kilobytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxMessageSizeInKilobytes"),
			},
			{
				Name:     "requires_duplicate_detection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RequiresDuplicateDetection"),
			},
			{
				Name:     "duplicate_detection_history_time_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DuplicateDetectionHistoryTimeWindow"),
			},
			{
				Name:     "enable_batched_operations",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableBatchedOperations"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "support_ordering",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SupportOrdering"),
			},
			{
				Name:     "auto_delete_on_idle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoDeleteOnIdle"),
			},
			{
				Name:     "enable_partitioning",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnablePartitioning"),
			},
			{
				Name:     "enable_express",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableExpress"),
			},
			{
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
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
			authorizationRules(),
		},
	}
}

func fetchServicebusTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Servicebus.Topics

	namespace := parent.Item.(servicebus.SBNamespace)
	resourceDetails, err := client.ParseResourceID(*namespace.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.ListByNamespace(ctx, resourceDetails.ResourceGroup, *namespace.Name, nil, nil)

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
