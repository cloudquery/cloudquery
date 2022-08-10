
# Table: azure_security_assessments
Assessment security assessment on a resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|display_name|text|User friendly display name of the assessment.|
|code|text|Programmatic code for the status of the assessment.|
|cause|text|Programmatic code for the cause of the assessment status.|
|description|text|Human readable description of the assessment status.|
|additional_data|jsonb|Additional data regarding the assessment.|
|azure_portal_uri|text|Link to assessment in Azure Portal.|
|metadata_display_name|text|User friendly display name of the assessment.|
|metadata_policy_definition_id|text|Azure resource ID of the policy definition that turns this assessment calculation on.|
|metadata_description|text|Human readable description of the assessment.|
|metadata_remediation_description|text|Human readable description of what you should do to mitigate this security issue.|
|metadata_categories|text[]||
|metadata_severity|text|The severity level of the assessment.|
|metadata_user_impact|text|The user impact of the assessment.|
|metadata_implementation_effort|text|The implementation effort required to remediate this assessment.|
|metadata_threats|text[]||
|metadata_preview|boolean|True if this assessment is in preview release status.|
|metadata_assessment_type|text|BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition|
|metadata_partner_data_partner_name|text|Name of the company of the partner.|
|metadata_partner_data_product_name|text|Name of the product of the partner that created the assessment.|
|partner_name|text|Name of the company of the partner|
|id|text|Resource Id.|
|name|text|Resource name.|
|type|text|Resource type.|
|resource_details|jsonb|Assessed resource details.|
