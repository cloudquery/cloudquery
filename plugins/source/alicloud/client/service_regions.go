package client

const ServiceECS = "ecs"

// ServiceRegions maps service name to a list of regions where the service is available.
// Extracted from https://api.aliyun.com/#/region
var ServiceRegions = map[string][]string{
	ServiceECS: {
		"cn-hangzhou",
		"cn-qingdao",
		"cn-beijing",
		"cn-zhangjiakou",
		"cn-huhehaote",
		"cn-shanghai",
		"cn-shenzhen",
		"cn-hongkong",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-southeast-3",
		"ap-southeast-5",
		"ap-northeast-1",
		"eu-west-1",
		"us-west-1",
		"us-east-1",
		"eu-central-1",
		"me-east-1",
		"ap-south-1",
		"cn-shanghai-finance-1",
		"cn-shenzhen-finance-1",
		"cn-chengdu",
		"cn-heyuan",
		"cn-wulanchabu",
		"rus-west-1",
		"cn-guangzhou",
		"ap-southeast-6",
		"cn-nanjing",
	},
}
