
# Table: azure_iothub_hub_authorization_policies
SharedAccessSignatureAuthorizationRule the properties of an IoT hub shared access policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|key_name|text|The name of the shared access policy.|
|primary_key|text|The primary key.|
|secondary_key|text|The secondary key.|
|rights|text|The permissions assigned to the shared access policy|
