
# Table: gcp_cloudbilling_accounts

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|display_name|text|The display name given to the billing account, such as `My Billing Account`|
|master_billing_account|text|If this account is a subaccount (https://cloudgooglecom/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through|
|name|text|The resource name of the billing account.|
|open|boolean|True if the billing account is open|
|project_billing_enabled|boolean|True if the project is associated with an open billing account, to which usage on the project is charged|
|project_name|text|The resource name for the `ProjectBillingInfo`; has the form `projects/{project_id}/billingInfo`|
|project_id|text|The ID of the project that this `ProjectBillingInfo` represents, such as `tokyo-rain-123`|
