package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID := range client.ServicesManager.services {
		l = append(l, client.withAccountID(accountID))
	}
	return l
}

func AccountRegionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID := range client.ServicesManager.services {
		for _, region := range client.regions {
			l = append(l, client.withAccountIDAndRegion(accountID, region))
		}
	}
	return l
}

var AllNamespaces = []string{ // this is only used in applicationautoscaling
	"comprehend", "rds", "sagemaker", "appstream", "elasticmapreduce", "dynamodb", "lambda", "ecs", "cassandra", "ec2", "neptune", "kafka", "custom-resource", "elasticache",
}

func AccountRegionNamespaceMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID := range client.ServicesManager.services {
		for _, region := range client.regions {
			for _, ns := range AllNamespaces {
				l = append(l, client.withAccountIDRegionAndNamespace(accountID, region, ns))
			}
		}
	}
	return l
}
