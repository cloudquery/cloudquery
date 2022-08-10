
# Table: azure_logic_app_workflows
Azure Logic App Workflow
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|provisioning_state|text|Gets the provisioning state. Possible values include: 'WorkflowProvisioningStateNotSpecified', 'WorkflowProvisioningStateAccepted', 'WorkflowProvisioningStateRunning', 'WorkflowProvisioningStateReady', 'WorkflowProvisioningStateCreating', 'WorkflowProvisioningStateCreated', 'WorkflowProvisioningStateDeleting', 'WorkflowProvisioningStateDeleted', 'WorkflowProvisioningStateCanceled', 'WorkflowProvisioningStateFailed', 'WorkflowProvisioningStateSucceeded', 'WorkflowProvisioningStateMoving', 'WorkflowProvisioningStateUpdating', 'WorkflowProvisioningStateRegistering', 'WorkflowProvisioningStateRegistered', 'WorkflowProvisioningStateUnregistering', 'WorkflowProvisioningStateUnregistered', 'WorkflowProvisioningStateCompleted', 'WorkflowProvisioningStateRenewing', 'WorkflowProvisioningStatePending', 'WorkflowProvisioningStateWaiting', 'WorkflowProvisioningStateInProgress'|
|created_time|timestamp without time zone|Gets the created time.|
|changed_time|timestamp without time zone|Gets the changed time.|
|state|text|The state. Possible values include: 'WorkflowStateNotSpecified', 'WorkflowStateCompleted', 'WorkflowStateEnabled', 'WorkflowStateDisabled', 'WorkflowStateDeleted', 'WorkflowStateSuspended'|
|version|text|Gets the version.|
|access_endpoint|text|Gets the access endpoint.|
|endpoints_configuration|jsonb|The endpoints configuration.|
|access_control|jsonb|The access control configuration.|
|sku_name|text|The sku name. Possible values include: 'SkuNameNotSpecified', 'SkuNameFree', 'SkuNameShared', 'SkuNameBasic', 'SkuNameStandard', 'SkuNamePremium'|
|sku_plan_id|text|The reference to plan resource id.|
|sku_plan_name|text|The reference to plan resource name.|
|sku_plan_type|text|The reference to plan resource type.|
|integration_account_id|text|The integration account id.|
|integration_account_name|text|The integration account name.|
|integration_account_type|text|The integration account type.|
|integration_service_environment_id|text|The integration service environment id.|
|integration_service_environment_name|text|The integration service environment name.|
|integration_service_environment_type|text|The integration service environment type.|
|definition|jsonb|The definition.|
|parameters|jsonb|The parameters.|
|identity_type|text|Type of managed service identity. The type 'SystemAssigned' includes an implicitly created identity. The type 'None' will remove any identities from the resource. Possible values include: 'ManagedServiceIdentityTypeSystemAssigned', 'ManagedServiceIdentityTypeUserAssigned', 'ManagedServiceIdentityTypeNone'|
|identity_tenant_id|uuid|Tenant of managed service identity.|
|identity_principal_id|uuid|Principal Id of managed service identity.|
|identity_user_assigned_identities|jsonb|The list of user assigned identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}|
|id|text|Fully qualified resource ID for the resource|
|name|text|The name of the resource.|
|type|text|The type of the resource.|
|location|text|The geo-location where the resource lives|
|tags|jsonb|Resource tags.|
|diagnostic_settings|jsonb|A list of active diagnostic settings for the workflow.|
