package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationRdsInstances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsInstances(), []string{"aws_rds_instances.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_rds_instances",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("rds-%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}

func TestIntegrationRdsSubnetGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsSubnetGroups(), []string{"aws_rds_db_subnet_groups.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_rds_subnet_groups",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("rds_db_subnet%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("rds_db_subnet%s%s", res.Prefix, res.Suffix),
					},
				},
			},
		}
	})
}

func TestIntegrationRdsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.RdsClusters(), []string{"aws_rds_clusters.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_rds_clusters",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"db_cluster_identifier": fmt.Sprintf("rdscluster%s", res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"db_cluster_identifier": fmt.Sprintf("rdscluster%s", res.Suffix),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_rds_cluster_db_cluster_members",
					ForeignKeyName: "cluster_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"db_instance_identifier": fmt.Sprintf("rdsclusterdb%s", res.Suffix),
							},
						},
					},
				},
			},
		}
	})
}
