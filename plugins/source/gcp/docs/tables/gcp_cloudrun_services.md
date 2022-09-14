
# Table: gcp_cloudrun_services
Service acts as a top-level container that manages a set of Routes and Configurations which implement a network service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text||
|api_version|text|The API version for this call such as "servingknativedev/v1"|
|kind|text|The kind of resource, in this case "Service"|
|metadata_annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata|
|metadata_creation_timestamp|text|CreationTimestamp is a timestamp representing the server time when this object was created|
|metadata_generation|bigint|A sequence number representing a specific generation of the desired state|
|metadata_labels|jsonb|Map of string keys and values that can be used to organize and categorize (scope and select) objects|
|metadata_name|text|Name must be unique within a namespace, within a Cloud Run region|
|metadata_namespace|text|Namespace defines the space within each name must be unique, within a Cloud Run region|
|metadata_resource_version|text|Optional|
|metadata_self_link|text|SelfLink is a URL representing this object Populated by the system|
|metadata_uid|text|UID is the unique in time and space value for this object|
|spec_template_metadata_annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata|
|spec_template_metadata_creation_timestamp|text|CreationTimestamp is a timestamp representing the server time when this object was created|
|spec_template_metadata_generation|bigint|A sequence number representing a specific generation of the desired state|
|spec_template_metadata_labels|jsonb|Map of string keys and values that can be used to organize and categorize (scope and select) objects|
|spec_template_metadata_name|text|Name must be unique within a namespace, within a Cloud Run region|
|spec_template_metadata_namespace|text|Namespace defines the space within each name must be unique, within a Cloud Run region|
|spec_template_metadata_resource_version|text|Optional|
|spec_template_metadata_self_link|text|SelfLink is a URL representing this object Populated by the system|
|spec_template_metadata_uid|text|UID is the unique in time and space value for this object|
|spec_template_spec_container_concurrency|bigint|Optional|
|spec_template_spec_service_account_name|text|Email address of the IAM service account associated with the revision of the service|
|spec_template_spec_timeout_seconds|bigint|TimeoutSeconds holds the max duration the instance is allowed for responding to a request|
|status_address_url|text||
|status_latest_created_revision_name|text|From ConfigurationStatus LatestCreatedRevisionName is the last revision that was created from this Service's Configuration|
|status_latest_ready_revision_name|text|From ConfigurationStatus LatestReadyRevisionName holds the name of the latest Revision stamped out from this Service's Configuration that has had its "Ready" condition become "True"|
|status_observed_generation|bigint|ObservedGeneration is the 'Generation' of the Route that was last processed by the controller|
|status_url|text|From RouteStatus|
