package network

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkSecurityGroups() *schema.Table {
	return &schema.Table{
		Name:          "azure_network_security_groups",
		Description:   "Azure network security group",
		Resolver:      fetchNetworkSecurityGroups,
		Multiplex:     client.SubscriptionMultiplex,
		DeleteFilter:  client.DeleteSubscriptionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the network security group resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SecurityGroupPropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the network security group resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SecurityGroupPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_security_group_security_rules",
				Description: "SecurityRule network security rule",
				Resolver:    fetchNetworkSecurityGroupSecurityRules,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "security_group_cq_id",
						Description: "Unique CloudQuery ID of azure_network_security_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "description",
						Description:   "A description for this rule Restricted to 140 chars",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SecurityRulePropertiesFormat.Description"),
						IgnoreInTests: true,
					},
					{
						Name:        "protocol",
						Description: "Network protocol this rule applies to Possible values include: 'SecurityRuleProtocolTCP', 'SecurityRuleProtocolUDP', 'SecurityRuleProtocolIcmp', 'SecurityRuleProtocolEsp', 'SecurityRuleProtocolAsterisk', 'SecurityRuleProtocolAh'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Protocol"),
					},
					{
						Name:        "source_port_range",
						Description: "The source port or range Integer or range between 0 and 65535 Asterisk '*' can also be used to match all ports",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.SourcePortRange"),
					},
					{
						Name:        "destination_port_range",
						Description: "The destination port or range Integer or range between 0 and 65535 Asterisk '*' can also be used to match all ports",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.DestinationPortRange"),
					},
					{
						Name:        "source_address_prefix",
						Description: "The CIDR or source IP range Asterisk '*' can also be used to match all source IPs Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used If this is an ingress rule, specifies where network traffic originates from",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.SourceAddressPrefix"),
					},
					{
						Name:        "source_address_prefixes",
						Description: "The CIDR or source IP ranges",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.SourceAddressPrefixes"),
					},
					{
						Name:        "destination_address_prefix",
						Description: "The destination address prefix CIDR or destination IP range Asterisk '*' can also be used to match all source IPs Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.DestinationAddressPrefix"),
					},
					{
						Name:        "destination_address_prefixes",
						Description: "The destination address prefixes CIDR or destination IP ranges",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.DestinationAddressPrefixes"),
					},
					{
						Name:        "source_port_ranges",
						Description: "The source port ranges",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.SourcePortRanges"),
					},
					{
						Name:        "destination_port_ranges",
						Description: "The destination port ranges",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.DestinationPortRanges"),
					},
					{
						Name:        "access",
						Description: "The network traffic is allowed or denied Possible values include: 'SecurityRuleAccessAllow', 'SecurityRuleAccessDeny'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Access"),
					},
					{
						Name:        "priority",
						Description: "The priority of the rule The value can be between 100 and 4096 The priority number must be unique for each rule in the collection The lower the priority number, the higher the priority of the rule",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Priority"),
					},
					{
						Name:        "direction",
						Description: "The direction of the rule The direction specifies if rule will be evaluated on incoming or outgoing traffic Possible values include: 'SecurityRuleDirectionInbound', 'SecurityRuleDirectionOutbound'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Direction"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the security rule resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "name",
						Description: "The name of the resource that is unique within a resource group This name can be used to access the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_network_security_group_flow_logs",
				Description: "FlowLog a flow log resource",
				Resolver:    fetchNetworkSecurityGroupFlowLogs,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "security_group_cq_id",
						Description: "Unique CloudQuery ID of azure_network_security_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "target_resource_id",
						Description: "ID of network security group to which flow log will be applied",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.TargetResourceID"),
					},
					{
						Name:          "target_resource_guid",
						Description:   "Guid of network security group to which flow log will be applied",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("FlowLogPropertiesFormat.TargetResourceGUID"),
						IgnoreInTests: true,
					},
					{
						Name:        "storage_id",
						Description: "ID of the storage account which is used to store the flow log",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.StorageID"),
					},
					{
						Name:        "enabled",
						Description: "Flag to enable/disable flow logging",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.Enabled"),
					},
					{
						Name:        "retention_policy_days",
						Description: "Number of days to retain flow log records",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.RetentionPolicy.Days"),
					},
					{
						Name:        "retention_policy_enabled",
						Description: "Flag to enable/disable retention",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.RetentionPolicy.Enabled"),
					},
					{
						Name:        "format_type",
						Description: "The file type of flow log Possible values include: 'JSON'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.Format.Type"),
					},
					{
						Name:        "format_version",
						Description: "The version (revision) of the flow log",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.Format.Version"),
					},
					{
						Name:        "flow_analytics_configuration_enabled",
						Description: "Flag to enable/disable traffic analytics for network watcher",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.FlowAnalyticsConfiguration.NetworkWatcherFlowAnalyticsConfiguration.Enabled"),
					},
					{
						Name:        "flow_analytics_configuration_workspace_id",
						Description: "The resource guid of the attached workspace for network watcher",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.FlowAnalyticsConfiguration.NetworkWatcherFlowAnalyticsConfiguration.WorkspaceID"),
					},
					{
						Name:        "flow_analytics_configuration_workspace_region",
						Description: "The location of the attached workspace for network watcher",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.FlowAnalyticsConfiguration.NetworkWatcherFlowAnalyticsConfiguration.WorkspaceRegion"),
					},
					{
						Name:        "flow_analytics_configuration_workspace_resource_id",
						Description: "Resource Id of the attached workspace for network watcher",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.FlowAnalyticsConfiguration.NetworkWatcherFlowAnalyticsConfiguration.WorkspaceResourceID"),
					},
					{
						Name:        "flow_analytics_configuration_traffic_analytics_interval",
						Description: "The interval in minutes which would decide how frequently TA service should do flow analytics for network watcher",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.FlowAnalyticsConfiguration.NetworkWatcherFlowAnalyticsConfiguration.TrafficAnalyticsInterval"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the flow log Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FlowLogPropertiesFormat.ProvisioningState"),
					},
					{
						Name:          "etag",
						Description:   "A unique read-only string that changes whenever the resource is updated",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:          "type",
						Description:   "Resource type",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "location",
						Description: "Resource location",
						Type:        schema.TypeString,
					},
					{
						Name:          "tags",
						Description:   "Resource tags",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:        "azure_network_security_group_default_security_rules",
				Description: "SecurityRule network security rule",
				Resolver:    fetchNetworkSecurityGroupDefaultSecurityRules,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "security_group_cq_id",
						Description: "Unique CloudQuery ID of azure_network_security_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "description",
						Description: "A description for this rule Restricted to 140 chars",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Description"),
					},
					{
						Name:        "protocol",
						Description: "Network protocol this rule applies to Possible values include: 'SecurityRuleProtocolTCP', 'SecurityRuleProtocolUDP', 'SecurityRuleProtocolIcmp', 'SecurityRuleProtocolEsp', 'SecurityRuleProtocolAsterisk', 'SecurityRuleProtocolAh'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Protocol"),
					},
					{
						Name:        "source_port_range",
						Description: "The source port or range Integer or range between 0 and 65535 Asterisk '*' can also be used to match all ports",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.SourcePortRange"),
					},
					{
						Name:        "destination_port_range",
						Description: "The destination port or range Integer or range between 0 and 65535 Asterisk '*' can also be used to match all ports",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.DestinationPortRange"),
					},
					{
						Name:        "source_address_prefix",
						Description: "The CIDR or source IP range Asterisk '*' can also be used to match all source IPs Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used If this is an ingress rule, specifies where network traffic originates from",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.SourceAddressPrefix"),
					},
					{
						Name:        "source_address_prefixes",
						Description: "The CIDR or source IP ranges",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.SourceAddressPrefixes"),
					},
					{
						Name:        "destination_address_prefix",
						Description: "The destination address prefix CIDR or destination IP range Asterisk '*' can also be used to match all source IPs Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.DestinationAddressPrefix"),
					},
					{
						Name:        "destination_address_prefixes",
						Description: "The destination address prefixes CIDR or destination IP ranges",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.DestinationAddressPrefixes"),
					},
					{
						Name:        "source_port_ranges",
						Description: "The source port ranges",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.SourcePortRanges"),
					},
					{
						Name:        "destination_port_ranges",
						Description: "The destination port ranges",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.DestinationPortRanges"),
					},
					{
						Name:        "access",
						Description: "The network traffic is allowed or denied Possible values include: 'SecurityRuleAccessAllow', 'SecurityRuleAccessDeny'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Access"),
					},
					{
						Name:        "priority",
						Description: "The priority of the rule The value can be between 100 and 4096 The priority number must be unique for each rule in the collection The lower the priority number, the higher the priority of the rule",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Priority"),
					},
					{
						Name:        "direction",
						Description: "The direction of the rule The direction specifies if rule will be evaluated on incoming or outgoing traffic Possible values include: 'SecurityRuleDirectionInbound', 'SecurityRuleDirectionOutbound'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.Direction"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the security rule resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SecurityRulePropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "name",
						Description: "The name of the resource that is unique within a resource group This name can be used to access the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchNetworkSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.SecurityGroups
	response, err := svc.ListAll(ctx)
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
func fetchNetworkSecurityGroupSecurityRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(network.SecurityGroup)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", parent.Item)
	}
	res <- *p.SecurityRules
	return nil
}
func fetchNetworkSecurityGroupFlowLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(network.SecurityGroup)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", parent.Item)
	}
	///subscriptions/ce681e53-1bc4-4218-b777-ddbfeec1986f/resourceGroups/example-resources/providers/Microsoft.Network/networkWatchers/acctestnw/flowLogs/Microsoft.Networkexample-resourcesacceptanceTestSecurityGroup1test
	if p.FlowLogs == nil {
		return nil
	}

	svc := meta.(*client.Client).Services().Network.Watchers
	for _, fl := range *p.FlowLogs {
		//parse flow log id and get required fields from it
		v := strings.Split(*fl.ID, "/")
		if len(v) != 11 {
			return fmt.Errorf("wrong format of flow logs id")
		}
		networkWatcherName := v[8]
		resourceGroup := v[4]
		name := v[10]

		//there is no API to get network.FlowLog directly so we fetch network.FlowLogInformation and fill network.FlowLog structure
		result, err := svc.GetFlowLogStatus(ctx, resourceGroup, networkWatcherName, network.FlowLogStatusParameters{TargetResourceID: p.ID})
		if err != nil {
			return err
		}
		client, ok := svc.(network.WatchersClient)
		if !ok {
			client = network.WatchersClient{} //use a dummy network.WatchersClient with unit tests
		}
		properties, err := result.Result(client)
		if err != nil {
			return err
		}

		fl.Name = &name
		if fl.FlowLogPropertiesFormat == nil {
			fl.FlowLogPropertiesFormat = &network.FlowLogPropertiesFormat{}
		}

		fl.Enabled = properties.Enabled
		fl.TargetResourceID = properties.TargetResourceID
		fl.StorageID = properties.StorageID
		fl.RetentionPolicy = properties.RetentionPolicy
		fl.Format = properties.Format
		fl.FlowAnalyticsConfiguration = properties.FlowAnalyticsConfiguration

		if properties.FlowAnalyticsConfiguration != nil && properties.FlowAnalyticsConfiguration.NetworkWatcherFlowAnalyticsConfiguration != nil {
			fl.Location = properties.FlowAnalyticsConfiguration.NetworkWatcherFlowAnalyticsConfiguration.WorkspaceRegion
		}
		res <- fl
	}
	return nil
}
func fetchNetworkSecurityGroupDefaultSecurityRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(network.SecurityGroup)
	if !ok {
		return fmt.Errorf("expected to have network.SecurityGroup but got %T", parent.Item)
	}
	res <- *p.DefaultSecurityRules
	return nil
}
