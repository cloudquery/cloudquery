package plugin

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	Version = "Development"
)

var awsExceptions = map[string]string{
	"accessanalyzer":         "AWS Identity and Access Management (IAM) Access Analyzer",
	"acm":                    "Amazon Certificate Manager (ACM)",
	"acmpca":                 "AWS Certificate Manager Private Certificate Authority (ACM PCA)",
	"amp":                    "Amazon Managed Service for Prometheus (AMP)",
	"apigateway":             "Amazon API Gateway",
	"apigatewayv2":           "Amazon API Gateway v2",
	"applicationautoscaling": "Application Auto Scaling",
	"apprunner":              "AWS App Runner",
	"appstream":              "Amazon AppStream",
	"arn":                    "Amazon Resource Name (ARN)",
	"arns":                   "Amazon Resource Names (ARNs)",
	"autoscaling":            "Auto Scaling",
	"autoscalingplans":       "Auto Scaling Plans",
	"aws":                    "", // remove "AWS" from names, because in most cases it will be replaced with either Amazon or AWS
	"byoip":                  "Bring your own IP addresses (BYOIP)",
	"cloudformation":         "AWS CloudFormation",
	"cloudhsm":               "AWS CloudHSM",
	"cloudhsmv2":             "AWS CloudHSM v2",
	"cloudtrail":             "AWS CloudTrail",
	"codeartifact":           "AWS CodeArtifact",
	"codebuild":              "AWS CodeBuild",
	"codecommit":             "AWS CodeCommit",
	"computeoptimizer":       "Compute Optimizer",
	"costexplorer":           "AWS Cost Explorer",
	"detective":              "Amazon Detective",
	"directconnect":          "AWS Direct Connect",
	"docdb":                  "Amazon DocumentDB",
	"dynamodb":               "Amazon DynamoDB",
	"dynamodbstreams":        "Amazon DynamoDB",
	"ebs":                    "Amazon Elastic Block Store (EBS)",
	"ec2":                    "Amazon Elastic Compute Cloud (EC2)",
	"ecr":                    "Amazon Elastic Container Registry (ECR)",
	"ecs":                    "Amazon Elastic Container Service (ECS)",
	"efs":                    "Amazon Elastic File System (EFS)",
	"eks":                    "Amazon Elastic Kubernetes Service (EKS)",
	"elasticbeanstalk":       "AWS Elastic Beanstalk",
	"elastictranscoder":      "Amazon Elastic Transcoder",
	"elb":                    "Amazon Elastic Load Balancer (ELB)",
	"elbv1":                  "Amazon Elastic Load Balancer (ELB) v1",
	"elbv2":                  "Amazon Elastic Load Balancer (ELB) v2",
	"emr":                    "Amazon EMR",
	"eventbridge":            "Amazon EventBridge",
	"frauddetector":          "Amazon Fraud Detector",
	"fsx":                    "Amazon FSx",
	"guardduty":              "Amazon GuardDuty",
	"identitystore":          "Identity Store",
	"iot":                    "AWS IoT",
	"kms":                    "AWS Key Management Service (AWS KMS)",
	"lambda":                 "AWS Lambda",
	"mq":                     "Amazon MQ",
	"mwaa":                   "Amazon MWAA",
	"nat":                    "NAT",
	"qldb":                   "Quantum Ledger Database (QLDB)",
	"quicksight":             "QuickSight",
	"rds":                    "Amazon Relational Database Service (RDS)",
	"route53":                "Amazon Route 53",
	"route53resolver":        "Amazon Route 53 Resolver",
	"resiliencehub":          "AWS Resilience Hub",
	"sagemaker":              "Amazon SageMaker",
	"secretsmanager":         "AWS Secrets Manager",
	"securityhub":            "AWS Security Hub",
	"servicecatalog":         "AWS Service Catalog",
	"servicediscovery":       "AWS Cloud Map",
	"ses":                    "Amazon Simple Email Service (SES)",
	"signer":                 "AWS Signer",
	"ssm":                    "AWS Systems Manager (SSM)",
	"wellarchitected":        "AWS Well-Architected",
	"xray":                   "AWS X-Ray",
}

func titleTransformer(table *schema.Table) {
	if table.Title == "" {
		exceptions := make(map[string]string)
		for k, v := range source.DefaultTitleExceptions {
			exceptions[k] = v
		}
		for k, v := range awsExceptions {
			exceptions[k] = v
		}
		csr := caser.New(caser.WithCustomExceptions(exceptions))
		t := csr.ToTitle(table.Name)
		table.Title = strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
	}
	for _, rel := range table.Relations {
		titleTransformer(rel)
	}
}

func AWS() *plugin.Plugin {
	return plugin.NewPlugin(
		"aws",
		Version,
		New,
	)
}
