
# Table: aws_ec2_customer_gateways
Describes a customer gateway.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|bgp_asn|text|The customer gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).|
|certificate_arn|text|The Amazon Resource Name (ARN) for the customer gateway certificate.|
|customer_gateway_id|text|The ID of the customer gateway.|
|device_name|text|The name of customer gateway device.|
|ip_address|text|The Internet-routable IP address of the customer gateway's outside interface.|
|state|text|The current state of the customer gateway (pending | available | deleting | deleted).|
|tags|jsonb|Any tags assigned to the customer gateway.|
|type|text|The type of VPN connection the customer gateway supports (ipsec.|
