
# Table: aws_eks_clusters
An object representing an Amazon EKS cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the cluster.|
|certificate_authority_data|text|The Base64-encoded certificate data required to communicate with your cluster.|
|client_request_token|text|Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.|
|created_at|timestamp without time zone|The Unix epoch timestamp in seconds for when the cluster was created.|
|endpoint|text|The endpoint for your Kubernetes API server.|
|identity_oidc_issuer|text|The issuer URL for the OIDC identity provider.|
|kubernetes_network_config_service_ipv4_cidr|text|The CIDR block that Kubernetes service IP addresses are assigned from.|
|name|text|The name of the cluster.|
|platform_version|text|The platform version of your Amazon EKS cluster.|
|resources_vpc_config_cluster_security_group_id|text|The cluster security group that was created by Amazon EKS for the cluster.|
|resources_vpc_config_endpoint_private_access|boolean|This parameter indicates whether the Amazon EKS private API server endpoint is enabled.|
|resources_vpc_config_endpoint_public_access|boolean|This parameter indicates whether the Amazon EKS public API server endpoint is enabled.|
|resources_vpc_config_public_access_cidrs|text[]|The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint.|
|resources_vpc_config_security_group_ids|text[]|The security groups associated with the cross-account elastic network interfaces that are used to allow communication between your nodes and the Kubernetes control plane.|
|resources_vpc_config_subnet_ids|text[]|The subnets associated with your cluster.|
|resources_vpc_config_vpc_id|text|The VPC associated with your cluster.|
|role_arn|text|The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.|
|status|text|The current status of the cluster.|
|tags|jsonb|The metadata that you apply to the cluster to assist with categorization and organization.|
|version|text|The Kubernetes server version for the cluster.|
