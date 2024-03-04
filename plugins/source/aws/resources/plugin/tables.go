package plugin

import (
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/appstream"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudfront"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticsearch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elastictranscoder"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ram"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/rds"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ssm"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func getTables() schema.Tables {
	t := []*schema.Table{
		appstream.Images(),
		cloudfront.CachePolicies(),
		cloudfront.Distributions(),
		cloudfront.Functions(),
		cloudfront.OriginAccessIdentities(),
		cloudfront.OriginRequestPolicies(),
		cloudfront.ResponseHeaderPolicies(),
		ec2.AccountAttributes(),
		ec2.AvailabilityZones(),
		ec2.ByoipCidrs(),
		ec2.CapacityReservations(),
		ec2.CustomerGateways(),
		ec2.DHCPOptions(),
		ec2.EgressOnlyInternetGateways(),
		ec2.Eips(),
		ec2.FlowLogs(),
		ec2.Hosts(),
		ec2.Images(),
		ec2.InstanceConnectEndpoints(),
		ec2.Instances(),
		ec2.InstanceStatuses(),
		ec2.InstanceTypes(),
		ec2.InternetGateways(),
		ec2.KeyPairs(),
		ec2.LaunchTemplates(),
		ec2.ManagedPrefixLists(),
		ec2.NatGateways(),
		ec2.NetworkAcls(),
		ec2.NetworkInterfaces(),
		ec2.RegionalConfigs(),
		ec2.Regions(),
		ec2.ReservedInstances(),
		ec2.RouteTables(),
		ec2.SecurityGroups(),
		ec2.SpotFleetRequests(),
		ec2.SpotInstanceRequests(),
		ec2.Subnets(),
		ec2.TransitGateways(),
		ec2.VpcEndpointConnections(),
		ec2.VpcEndpoints(),
		ec2.VpcEndpointServiceConfigurations(),
		ec2.VpcEndpointServices(),
		ec2.VpcPeeringConnections(),
		ec2.Vpcs(),
		ec2.VpnConnections(),
		ec2.VpnGateways(),
		elasticache.ReservedCacheNodesOfferings(),
		elasticache.ServiceUpdates(),
		elasticsearch.Packages(),
		elasticsearch.Versions(),
		elastictranscoder.Presets(),
		emr.ReleaseLabels(),
		iam.Accounts(),
		iam.AccountAuthorizationDetails(),
		iam.CredentialReports(),
		iam.Groups(),
		iam.InstanceProfiles(),
		iam.OpenidConnectIdentityProviders(),
		iam.PasswordPolicies(),
		iam.Policies(),
		iam.Roles(),
		iam.SamlIdentityProviders(),
		iam.ServerCertificates(),
		iam.Users(),
		iam.VirtualMfaDevices(),
		ram.ResourceTypes(),
		rds.Certificates(),
		rds.ClusterParameterGroups(),
		rds.Clusters(),
		rds.ClusterSnapshots(),
		rds.DbParameterGroups(),
		rds.DbProxies(),
		rds.DbSecurityGroups(),
		rds.DbSnapshots(),
		rds.EngineVersions(),
		rds.Events(),
		rds.EventSubscriptions(),
		rds.Instances(),
		rds.OptionGroups(),
		rds.ReservedInstances(),
		rds.SubnetGroups(),
		s3.AccessPoints(),
		s3.Accounts(),
		s3.Buckets(),
		s3.MultiRegionAccessPoints(),
		ssm.PatchBaselines(),
	}
	if err := transformers.TransformTables(t); err != nil {
		panic(err)
	}
	transformTitles := titleTransformer()
	for _, table := range t {
		schema.AddCqIDs(table)
		transformTitles(table)
		if err := validateTagsIsJSON(table); err != nil {
			panic(err)
		}
	}
	return t
}

func validateTagsIsJSON(table *schema.Table) error {
	for _, col := range table.Columns {
		if col.Name == "tags" && col.Type != types.ExtensionTypes.JSON {
			return fmt.Errorf("column %s in table %s must be of type %s", col.Name, table.Name, types.ExtensionTypes.JSON)
		}
	}
	for _, rel := range table.Relations {
		if err := validateTagsIsJSON(rel); err != nil {
			return err
		}
	}

	return nil
}
