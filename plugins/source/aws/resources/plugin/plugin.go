package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "Development"
)

var awsExceptions = map[string]string{
	"accessanalyzer":   "Access Analyzer",
	"acm":              "ACM",
	"amp":              "AMP",
	"apigateway":       "API Gateway",
	"apigatewayv2":     "API Gateway v2",
	"arn":              "ARN",
	"arns":             "ARNs",
	"aws":              "AWS",
	"byoip":            "BYOIP",
	"directconnect":    "Direct Connect",
	"docdb":            "DocDB",
	"dynamodb":         "DynamoDB",
	"ebs":              "EBS",
	"ec2":              "EC2",
	"ecr":              "ECR",
	"ecs":              "ECS",
	"efs":              "EFS",
	"eks":              "EKS",
	"elasticbeanstalk": "Elastic Beanstalk",
	"elb":              "ELB",
	"elbv2":            "ELB v2",
	"emr":              "EMR",
	"frauddetector":    "Fraud Detector",
	"fsx":              "FSX",
	"identitystore":    "Identity Store",
	"kms":              "KMS",
	"mq":               "MQ",
	"mwaa":             "MWAA",
	"nat":              "NAT",
	"qldb":             "QLDB",
	"quicksight":       "QuickSight",
	"rds":              "RDS",
	"secretsmanager":   "Secrets Manager",
	"securityhub":      "Security Hub",
	"servicecatalog":   "Service Catalog",
	"xray":             "X-Ray",
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
	return csr.ToTitle(table.Name)
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
