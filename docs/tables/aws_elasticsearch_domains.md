
# Table: aws_elasticsearch_domains
The current status of an Elasticsearch domain.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|arn|text|The Amazon resource name (ARN) of an Elasticsearch domain|
|id|text|The unique identifier for the specified Elasticsearch domain.  This member is required.|
|name|text|The name of an Elasticsearch domain|
|cluster_cold_storage_options_enabled|boolean|True to enable cold storage for an Elasticsearch domain.|
|cluster_dedicated_master_count|integer|Total number of dedicated master nodes, active and on standby, for the cluster.|
|cluster_dedicated_master_enabled|boolean|A boolean value to indicate whether a dedicated master node is enabled|
|cluster_dedicated_master_type|text|The instance type for a dedicated master node.|
|cluster_instance_count|integer|The number of instances in the specified domain cluster.|
|cluster_instance_type|text|The instance type for an Elasticsearch cluster|
|cluster_warm_count|integer|The number of warm nodes in the cluster.|
|cluster_warm_enabled|boolean|True to enable warm storage.|
|cluster_warm_type|text|The instance type for the Elasticsearch cluster's warm nodes.|
|cluster_zone_awareness_config_availability_zone_count|integer|An integer value to indicate the number of availability zones for a domain when zone awareness is enabled|
|cluster_zone_awareness_enabled|boolean|A boolean value to indicate whether zone awareness is enabled|
|access_policies|text|IAM access policy as a JSON-formatted string.|
|advanced_options|jsonb|Specifies the status of the AdvancedOptions|
|advanced_security_enabled|boolean|True if advanced security is enabled.|
|advanced_security_internal_user_database_enabled|boolean|True if the internal user database is enabled.|
|advanced_security_saml_enabled|boolean|True if SAML is enabled.|
|advanced_security_saml_idp_entity_id|text|The unique Entity ID of the application in SAML Identity Provider.|
|advanced_security_saml_roles_key|text|The Metadata of the SAML application in XML format.|
|advanced_security_options_saml_options_roles_key|text|The key used for matching the SAML Roles attribute.|
|advanced_security_saml_session_timeout_minutes|integer|The duration, in minutes, after which a user session becomes inactive.|
|advanced_security_saml_subject_key|text|The key used for matching the SAML Subject attribute.|
|auto_tune_error_message|text|Specifies the error message while enabling or disabling the Auto-Tune.|
|auto_tune_options_state|text|Specifies the AutoTuneState for the Elasticsearch domain.|
|cognito_enabled|boolean|Specifies the option to enable Cognito for Kibana authentication.|
|cognito_identity_pool_id|text|Specifies the Cognito identity pool ID for Kibana authentication.|
|cognito_role_arn|text|Specifies the role ARN that provides Elasticsearch permissions for accessing Cognito resources.|
|cognito_user_pool_id|text|Specifies the Cognito user pool ID for Kibana authentication.|
|created|boolean|The domain creation status|
|deleted|boolean|The domain deletion status|
|domain_endpoint_custom|text|Specify the fully qualified domain for your custom endpoint.|
|domain_endpoint_custom_certificate_arn|text|Specify ACM certificate ARN for your custom endpoint.|
|domain_endpoint_custom_enabled|boolean|Specify if custom endpoint should be enabled for the Elasticsearch domain.|
|domain_endpoint_enforce_https|boolean|Specify if only HTTPS endpoint should be enabled for the Elasticsearch domain.|
|domain_endpoint_tls_security_policy|text|Specify the TLS security policy that needs to be applied to the HTTPS endpoint of Elasticsearch domain.|
|ebs_enabled|boolean|Specifies whether EBS-based storage is enabled.|
|ebs_iops|integer|Specifies the IOPD for a Provisioned IOPS EBS volume (SSD).|
|ebs_volume_size|integer|Integer to specify the size of an EBS volume.|
|ebs_volume_type|text|Specifies the volume type for EBS-based storage.|
|elasticsearch_version|text||
|encryption_at_rest_enabled|boolean|Specifies the option to enable Encryption At Rest.|
|encryption_at_rest_kms_key_id|text|Specifies the KMS Key ID for Encryption At Rest options.|
|endpoint|text|The Elasticsearch domain endpoint that you use to submit index and search requests.|
|endpoints|jsonb|Map containing the Elasticsearch domain endpoints used to submit index and search requests|
|log_publishing_options|jsonb|Log publishing options for the given domain.|
|node_to_node_encryption_enabled|boolean|Specify true to enable node-to-node encryption.|
|processing|boolean|The status of the Elasticsearch domain configuration|
|service_software_automated_update_date|timestamp without time zone|Timestamp, in Epoch time, until which you can manually request a service software update|
|service_software_cancellable|boolean|True if you are able to cancel your service software version update|
|service_software_current_version|text|The current service software version that is present on the domain.|
|service_software_description|text|The description of the UpdateStatus.|
|service_software_new_version|text|The new service software version if one is available.|
|service_software_optional_deployment|boolean|True if a service software is never automatically updated|
|service_software_update_available|boolean|True if you are able to update you service software version|
|service_software_update_status|text|The status of your service software update|
|snapshot_options_automated_snapshot_start_hour|integer|Specifies the time, in UTC format, when the service takes a daily automated snapshot of the specified Elasticsearch domain|
|upgrade_processing|boolean|The status of an Elasticsearch domain version upgrade|
|vpc_availability_zones|text[]|The availability zones for the Elasticsearch domain|
|vpc_security_group_ids|text[]|Specifies the security groups for VPC endpoint.|
|vpc_subnet_ids|text[]|Specifies the subnets for VPC endpoint.|
|vpc_vpc_id|text|The VPC Id for the Elasticsearch domain|
