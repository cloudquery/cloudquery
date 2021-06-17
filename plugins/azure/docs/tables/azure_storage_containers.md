
# Table: azure_storage_containers
Azure storage container
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|account_id|uuid|Azure storage account id|
|version|text|The version of the deleted blob container|
|deleted|boolean|Indicates whether the blob container was deleted|
|deleted_time|timestamp without time zone|Blob container deletion time|
|remaining_retention_days|integer|Remaining retention days for soft deleted blob container|
|default_encryption_scope|text|Default the container to use specified encryption scope for all writes|
|deny_encryption_scope_override|boolean|Block override of encryption scope from the container default|
|public_access|text|Specifies whether data in the container may be accessed publicly and the level of access Possible values include: 'PublicAccessContainer', 'PublicAccessBlob', 'PublicAccessNone'|
|last_modified_time|timestamp without time zone|Returns the date and time the container was last modified|
|lease_status|text|The lease status of the container Possible values include: 'LeaseStatusLocked', 'LeaseStatusUnlocked'|
|lease_state|text|Lease state of the container Possible values include: 'LeaseStateAvailable', 'LeaseStateLeased', 'LeaseStateExpired', 'LeaseStateBreaking', 'LeaseStateBroken'|
|lease_duration|text|Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased Possible values include: 'Infinite', 'Fixed'|
|metadata|jsonb|A name-value pair to associate with the container as metadata|
|immutability_policy|jsonb|The ImmutabilityPolicy property of the container|
|legal_hold|jsonb|The LegalHold property of the container|
|has_legal_hold|boolean|The hasLegalHold public property is set to true by SRP if there are at least one existing tag The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account|
|has_immutability_policy|boolean|The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container|
|etag|text|Resource Etag|
|resource_id|text|Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}|
|name|text|The name of the resource|
|type|text|The type of the resource Eg "MicrosoftCompute/virtualMachines" or "MicrosoftStorage/storageAccounts"|
