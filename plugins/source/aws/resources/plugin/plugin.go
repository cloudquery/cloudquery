package plugin

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "Development"
)

var awsExceptions = map[string]string{
	"accessanalyzer":    "AWS Identity and Access Management (IAM) Access Analyzer",
	"acm":               "Amazon Certificate Manager (ACM)",
	"amp":               "Amazon Managed Service for Prometheus (AMP)",
	"apigateway":        "Amazon API Gateway",
	"apigatewayv2":      "Amazon API Gateway v2",
	"apprunner":         "AWS App Runner",
	"appstream":         "Amazon AppStream",
	"arn":               "Amazon Resource Name (ARN)",
	"arns":              "Amazon Resource Names (ARNs)",
	"aws":               "", // remove "AWS" from names, because in most cases it will be replaced with either Amazon or AWS
	"autoscaling":       "Auto Scaling",
	"autoscalingplans":  "Auto Scaling Plans",
	"byoip":             "Bring your own IP addresses (BYOIP)",
	"cloudhsm":          "AWS CloudHSM",
	"cloudhsmv2":        "AWS CloudHSM v2",
	"cloudformation":    "AWS CloudFormation",
	"cloudtrail":        "AWS CloudTrail",
	"computeoptimizer":  "Compute Optimizer",
	"directconnect":     "AWS Direct Connect",
	"docdb":             "Amazon DocumentDB",
	"dynamodb":          "Amazon DynamoDB",
	"ebs":               "Amazon Elastic Block Store (EBS)",
	"ec2":               "Amazon Elastic Compute Cloud (EC2)",
	"ecr":               "Amazon Elastic Container Registry (ECR)",
	"ecs":               "Amazon Elastic Container Service (ECS)",
	"efs":               "Amazon Elastic File System (EFS)",
	"eks":               "Amazon Elastic Kubernetes Service (EKS)",
	"elasticbeanstalk":  "AWS Elastic Beanstalk",
	"elastictranscoder": "Amazon Elastic Transcoder",
	"elb":               "Amazon Elastic Load Balancer (ELB)",
	"elbv1":             "Amazon Elastic Load Balancer (ELB) v1",
	"elbv2":             "Amazon Elastic Load Balancer (ELB) v2",
	"emr":               "Amazon EMR",
	"eventbridge":       "Amazon EventBridge",
	"frauddetector":     "Amazon Fraud Detector",
	"fsx":               "Amazon FSx",
	"guardduty":         "Amazon GuardDuty",
	"identitystore":     "Identity Store",
	"iot":               "AWS IoT",
	"kms":               "AWS Key Management Service (AWS KMS)",
	"lambda":            "AWS Lambda",
	"mq":                "Amazon MQ",
	"mwaa":              "Amazon MWAA",
	"nat":               "NAT",
	"qldb":              "Quantum Ledger Database (QLDB)",
	"quicksight":        "QuickSight",
	"rds":               "Amazon Relational Database Service (RDS)",
	"resiliencehub":     "AWS Resilience Hub",
	"sagemaker":         "Amazon SageMaker",
	"secretsmanager":    "AWS Secrets Manager",
	"securityhub":       "AWS Security Hub",
	"servicecatalog":    "AWS Service Catalog",
	"ses":               "Amazon Simple Email Service (SES)",
	"xray":              "AWS X-Ray",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := make(map[string]string)
	for k, v := range source.DefaultTitleExceptions {
		exceptions[k] = v
	}
	for k, v := range awsExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}

func AWS() *source.Plugin {
	return source.NewPlugin(
		"aws",
		Version,
		tables(),
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
}
