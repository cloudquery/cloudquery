
# Table: gcp_iam_service_accounts
An IAM service account A service account is an account for an application or a virtual machine (VM) instance, not a person You can use a service account to call Google APIs To learn more, read the overview of service accounts (https://cloudgooglecom/iam/help/service-accounts/overview) When you create a service account, you specify the project ID that owns the service account, as well as a name that must be unique within the project IAM uses these values to create an email address that identifies the service account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|description|text|A user-specified, human-readable description of the service account The maximum length is 256 UTF-8 bytes|
|disabled|boolean|Whether the service account is disabled|
|display_name|text|A user-specified, human-readable name for the service account The maximum length is 100 UTF-8 bytes|
|email|text|The email address of the service account|
|name|text|The resource name of the service account In one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` OR `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` OR `projects/-/serviceAccounts/{UNIQUE_ID}|
|oauth2_client_id|text|The OAuth 20 client ID for the service account|
|project_id|text|The ID of the project that owns the service account|
|unique_id|text|The unique, stable numeric ID for the service account Each service account retains its unique ID even if you delete the service account For example, if you delete a service account, then create a new service account with the same name, the new service account has a different unique ID than the deleted service account|
