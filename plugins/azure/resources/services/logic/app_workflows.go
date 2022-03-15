package logic

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	"github.com/cloudquery/cq-provider-azure/client"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func LogicAppWorkflows() *schema.Table {
	return &schema.Table{
		Name:         "azure_logic_app_workflows",
		Description:  "Azure Logic App Workflow",
		Resolver:     fetchLogicAppWorkflows,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "provisioning_state",
				Description: "Gets the provisioning state. Possible values include: 'WorkflowProvisioningStateNotSpecified', 'WorkflowProvisioningStateAccepted', 'WorkflowProvisioningStateRunning', 'WorkflowProvisioningStateReady', 'WorkflowProvisioningStateCreating', 'WorkflowProvisioningStateCreated', 'WorkflowProvisioningStateDeleting', 'WorkflowProvisioningStateDeleted', 'WorkflowProvisioningStateCanceled', 'WorkflowProvisioningStateFailed', 'WorkflowProvisioningStateSucceeded', 'WorkflowProvisioningStateMoving', 'WorkflowProvisioningStateUpdating', 'WorkflowProvisioningStateRegistering', 'WorkflowProvisioningStateRegistered', 'WorkflowProvisioningStateUnregistering', 'WorkflowProvisioningStateUnregistered', 'WorkflowProvisioningStateCompleted', 'WorkflowProvisioningStateRenewing', 'WorkflowProvisioningStatePending', 'WorkflowProvisioningStateWaiting', 'WorkflowProvisioningStateInProgress'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProvisioningState"),
			},
			{
				Name:        "created_time",
				Description: "Gets the created time.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CreatedTime.Time"),
			},
			{
				Name:        "changed_time",
				Description: "Gets the changed time.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ChangedTime.Time"),
			},
			{
				Name:        "state",
				Description: "The state. Possible values include: 'WorkflowStateNotSpecified', 'WorkflowStateCompleted', 'WorkflowStateEnabled', 'WorkflowStateDisabled', 'WorkflowStateDeleted', 'WorkflowStateSuspended'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State"),
			},
			{
				Name:        "version",
				Description: "Gets the version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Version"),
			},
			{
				Name:        "access_endpoint",
				Description: "Gets the access endpoint.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccessEndpoint"),
			},
			{
				Name:        "endpoints_configuration",
				Description: "The endpoints configuration.",
				Type:        schema.TypeJSON,
				Resolver:    endpointsConfigurationResolver,
			},
			{
				Name:          "access_control",
				Description:   "The access control configuration.",
				Type:          schema.TypeJSON,
				Resolver:      accessControlResolver,
				IgnoreInTests: true,
			},
			{
				Name:        "sku_name",
				Description: "The sku name. Possible values include: 'SkuNameNotSpecified', 'SkuNameFree', 'SkuNameShared', 'SkuNameBasic', 'SkuNameStandard', 'SkuNamePremium'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:          "sku_plan_id",
				Description:   "The reference to plan resource id.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Sku.Plan.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "sku_plan_name",
				Description:   "The reference to plan resource name.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Sku.Plan.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "sku_plan_type",
				Description:   "The reference to plan resource type.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Sku.Plan.Type"),
				IgnoreInTests: true,
			},
			{
				Name:          "integration_account_id",
				Description:   "The integration account id.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("IntegrationAccount.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "integration_account_name",
				Description:   "The integration account name.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("IntegrationAccount.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "integration_account_type",
				Description:   "The integration account type.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("IntegrationAccount.Type"),
				IgnoreInTests: true,
			},
			{
				Name:          "integration_service_environment_id",
				Description:   "The integration service environment id.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("IntegrationServiceEnvironment.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "integration_service_environment_name",
				Description:   "The integration service environment name.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("IntegrationServiceEnvironment.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "integration_service_environment_type",
				Description:   "The integration service environment type.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("IntegrationServiceEnvironment.Type"),
				IgnoreInTests: true,
			},
			{
				Name:        "definition",
				Description: "The definition.",
				Type:        schema.TypeJSON,
				Resolver:    definitionResolver,
			},
			{
				Name:        "parameters",
				Description: "The parameters.",
				Type:        schema.TypeJSON,
				Resolver:    parametersResolver,
			},
			{
				Name:        "identity_type",
				Description: "Type of managed service identity. The type 'SystemAssigned' includes an implicitly created identity. The type 'None' will remove any identities from the resource. Possible values include: 'ManagedServiceIdentityTypeSystemAssigned', 'ManagedServiceIdentityTypeUserAssigned', 'ManagedServiceIdentityTypeNone'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:          "identity_tenant_id",
				Description:   "Tenant of managed service identity.",
				Type:          schema.TypeUUID,
				Resolver:      schema.PathResolver("Identity.TenantID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_principal_id",
				Description:   "Principal Id of managed service identity.",
				Type:          schema.TypeUUID,
				Resolver:      schema.PathResolver("Identity.PrincipalID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_user_assigned_identities",
				Description:   "The list of user assigned identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}",
				Type:          schema.TypeJSON,
				Resolver:      identityUserAssignedIdentitiesResolver,
				IgnoreInTests: true,
			},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Type"),
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location"),
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Tags"),
			},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the workflow.",
				Type:        schema.TypeJSON,
				Resolver:    fetchDiagnosticSettingsResolver,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLogicAppWorkflows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Logic.Workflows
	var top int32 = 100
	response, err := svc.ListBySubscription(ctx, &top, "")
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err = response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}

func endpointsConfigurationResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	workflow := resource.Item.(logic.Workflow)
	if workflow.EndpointsConfiguration == nil {
		return nil
	}
	b, err := json.Marshal(*workflow.EndpointsConfiguration)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func accessControlResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	workflow := resource.Item.(logic.Workflow)
	if workflow.AccessControl == nil {
		return nil
	}
	b, err := json.Marshal(*workflow.AccessControl)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func definitionResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	definition := resource.Item.(logic.Workflow).Definition
	if definition == nil {
		definition = make(map[string]interface{})
	}
	b, err := json.Marshal(definition)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func parametersResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	workflow := resource.Item.(logic.Workflow)
	if workflow.Parameters == nil {
		return nil
	}
	b, err := json.Marshal(workflow.Parameters)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func identityUserAssignedIdentitiesResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	workflow := resource.Item.(logic.Workflow)
	if workflow.Identity == nil || workflow.Identity.UserAssignedIdentities == nil {
		return nil
	}
	b, err := json.Marshal(workflow.Identity.UserAssignedIdentities)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func fetchDiagnosticSettingsResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	id := *resource.Item.(logic.Workflow).ID
	svc := meta.(*client.Client).Services().Logic.DiagnosticSettings
	response, err := svc.List(ctx, id)
	if err != nil {
		return err
	}
	diagnosticSettings := []map[string]interface{}{}
	for _, v := range *response.Value {
		diagnosticSetting := make(map[string]interface{})
		if v.ID != nil {
			diagnosticSetting["id"] = v.ID
		}
		if v.Name != nil {
			diagnosticSetting["name"] = v.Name
		}
		if v.Type != nil {
			diagnosticSetting["type"] = v.Type
		}
		if v.DiagnosticSettings != nil {
			diagnosticSetting["properties"] = v.DiagnosticSettings
		}
		diagnosticSettings = append(diagnosticSettings, diagnosticSetting)
	}
	b, err := json.Marshal(diagnosticSettings)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
