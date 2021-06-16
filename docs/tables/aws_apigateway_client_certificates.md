
# Table: aws_apigateway_client_certificates
Represents a client certificate used to configure client-side SSL authentication while sending requests to the integration endpoint.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|client_certificate_id|text|The identifier of the client certificate.|
|created_date|timestamp without time zone|The timestamp when the client certificate was created.|
|description|text|The description of the client certificate.|
|expiration_date|timestamp without time zone|The timestamp when the client certificate will expire.|
|pem_encoded_certificate|text|The PEM-encoded public key of the client certificate, which can be used to configure certificate authentication in the integration endpoint .|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
