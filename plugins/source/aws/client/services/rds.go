// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/rds.go -source=rds.go RdsClient
type RdsClient interface {
	DescribeAccountAttributes(context.Context, *rds.DescribeAccountAttributesInput, ...func(*rds.Options)) (*rds.DescribeAccountAttributesOutput, error)
	DescribeBlueGreenDeployments(context.Context, *rds.DescribeBlueGreenDeploymentsInput, ...func(*rds.Options)) (*rds.DescribeBlueGreenDeploymentsOutput, error)
	DescribeCertificates(context.Context, *rds.DescribeCertificatesInput, ...func(*rds.Options)) (*rds.DescribeCertificatesOutput, error)
	DescribeDBClusterBacktracks(context.Context, *rds.DescribeDBClusterBacktracksInput, ...func(*rds.Options)) (*rds.DescribeDBClusterBacktracksOutput, error)
	DescribeDBClusterEndpoints(context.Context, *rds.DescribeDBClusterEndpointsInput, ...func(*rds.Options)) (*rds.DescribeDBClusterEndpointsOutput, error)
	DescribeDBClusterParameterGroups(context.Context, *rds.DescribeDBClusterParameterGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBClusterParameterGroupsOutput, error)
	DescribeDBClusterParameters(context.Context, *rds.DescribeDBClusterParametersInput, ...func(*rds.Options)) (*rds.DescribeDBClusterParametersOutput, error)
	DescribeDBClusterSnapshotAttributes(context.Context, *rds.DescribeDBClusterSnapshotAttributesInput, ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterSnapshots(context.Context, *rds.DescribeDBClusterSnapshotsInput, ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBClusters(context.Context, *rds.DescribeDBClustersInput, ...func(*rds.Options)) (*rds.DescribeDBClustersOutput, error)
	DescribeDBEngineVersions(context.Context, *rds.DescribeDBEngineVersionsInput, ...func(*rds.Options)) (*rds.DescribeDBEngineVersionsOutput, error)
	DescribeDBInstanceAutomatedBackups(context.Context, *rds.DescribeDBInstanceAutomatedBackupsInput, ...func(*rds.Options)) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error)
	DescribeDBInstances(context.Context, *rds.DescribeDBInstancesInput, ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error)
	DescribeDBLogFiles(context.Context, *rds.DescribeDBLogFilesInput, ...func(*rds.Options)) (*rds.DescribeDBLogFilesOutput, error)
	DescribeDBParameterGroups(context.Context, *rds.DescribeDBParameterGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBParameterGroupsOutput, error)
	DescribeDBParameters(context.Context, *rds.DescribeDBParametersInput, ...func(*rds.Options)) (*rds.DescribeDBParametersOutput, error)
	DescribeDBProxies(context.Context, *rds.DescribeDBProxiesInput, ...func(*rds.Options)) (*rds.DescribeDBProxiesOutput, error)
	DescribeDBProxyEndpoints(context.Context, *rds.DescribeDBProxyEndpointsInput, ...func(*rds.Options)) (*rds.DescribeDBProxyEndpointsOutput, error)
	DescribeDBProxyTargetGroups(context.Context, *rds.DescribeDBProxyTargetGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBProxyTargetGroupsOutput, error)
	DescribeDBProxyTargets(context.Context, *rds.DescribeDBProxyTargetsInput, ...func(*rds.Options)) (*rds.DescribeDBProxyTargetsOutput, error)
	DescribeDBSecurityGroups(context.Context, *rds.DescribeDBSecurityGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBSecurityGroupsOutput, error)
	DescribeDBSnapshotAttributes(context.Context, *rds.DescribeDBSnapshotAttributesInput, ...func(*rds.Options)) (*rds.DescribeDBSnapshotAttributesOutput, error)
	DescribeDBSnapshots(context.Context, *rds.DescribeDBSnapshotsInput, ...func(*rds.Options)) (*rds.DescribeDBSnapshotsOutput, error)
	DescribeDBSubnetGroups(context.Context, *rds.DescribeDBSubnetGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBSubnetGroupsOutput, error)
	DescribeEngineDefaultClusterParameters(context.Context, *rds.DescribeEngineDefaultClusterParametersInput, ...func(*rds.Options)) (*rds.DescribeEngineDefaultClusterParametersOutput, error)
	DescribeEngineDefaultParameters(context.Context, *rds.DescribeEngineDefaultParametersInput, ...func(*rds.Options)) (*rds.DescribeEngineDefaultParametersOutput, error)
	DescribeEventCategories(context.Context, *rds.DescribeEventCategoriesInput, ...func(*rds.Options)) (*rds.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *rds.DescribeEventSubscriptionsInput, ...func(*rds.Options)) (*rds.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *rds.DescribeEventsInput, ...func(*rds.Options)) (*rds.DescribeEventsOutput, error)
	DescribeExportTasks(context.Context, *rds.DescribeExportTasksInput, ...func(*rds.Options)) (*rds.DescribeExportTasksOutput, error)
	DescribeGlobalClusters(context.Context, *rds.DescribeGlobalClustersInput, ...func(*rds.Options)) (*rds.DescribeGlobalClustersOutput, error)
	DescribeOptionGroupOptions(context.Context, *rds.DescribeOptionGroupOptionsInput, ...func(*rds.Options)) (*rds.DescribeOptionGroupOptionsOutput, error)
	DescribeOptionGroups(context.Context, *rds.DescribeOptionGroupsInput, ...func(*rds.Options)) (*rds.DescribeOptionGroupsOutput, error)
	DescribeOrderableDBInstanceOptions(context.Context, *rds.DescribeOrderableDBInstanceOptionsInput, ...func(*rds.Options)) (*rds.DescribeOrderableDBInstanceOptionsOutput, error)
	DescribePendingMaintenanceActions(context.Context, *rds.DescribePendingMaintenanceActionsInput, ...func(*rds.Options)) (*rds.DescribePendingMaintenanceActionsOutput, error)
	DescribeReservedDBInstances(context.Context, *rds.DescribeReservedDBInstancesInput, ...func(*rds.Options)) (*rds.DescribeReservedDBInstancesOutput, error)
	DescribeReservedDBInstancesOfferings(context.Context, *rds.DescribeReservedDBInstancesOfferingsInput, ...func(*rds.Options)) (*rds.DescribeReservedDBInstancesOfferingsOutput, error)
	DescribeSourceRegions(context.Context, *rds.DescribeSourceRegionsInput, ...func(*rds.Options)) (*rds.DescribeSourceRegionsOutput, error)
	DescribeValidDBInstanceModifications(context.Context, *rds.DescribeValidDBInstanceModificationsInput, ...func(*rds.Options)) (*rds.DescribeValidDBInstanceModificationsOutput, error)
	ListTagsForResource(context.Context, *rds.ListTagsForResourceInput, ...func(*rds.Options)) (*rds.ListTagsForResourceOutput, error)
}
