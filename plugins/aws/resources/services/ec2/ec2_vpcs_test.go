// +build integration

package ec2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEc2Vpcs(t *testing.T) {
	client.AWSTestHelper(t, Ec2Vpcs(),
		"./snapshots")
}

func TestIntegrationEc2VpcPeeringConnections(t *testing.T) {
	client.AWSTestHelper(t, Ec2VpcPeeringConnections(),
		"./snapshots")
}

func TestIntegrationEc2VpcEndpoints(t *testing.T) {
	client.AWSTestHelper(t, Ec2VpcEndpoints(),
		"./snapshots")
}

func TestIntegrationEc2TransitGateways(t *testing.T) {
	client.AWSTestHelper(t, Ec2TransitGateways(),
		"./snapshots")
}

func TestIntegrationEc2Subnets(t *testing.T) {
	client.AWSTestHelper(t, Ec2Subnets(),
		"./snapshots")
}

func TestIntegrationEc2SecurityGroups(t *testing.T) {
	client.AWSTestHelper(t, Ec2SecurityGroups(),
		"./snapshots")
}

func TestIntegrationEc2RouteTables(t *testing.T) {
	client.AWSTestHelper(t, Ec2RouteTables(),
		"./snapshots")
}

func TestIntegrationEc2NetworkAcls(t *testing.T) {
	client.AWSTestHelper(t, Ec2NetworkAcls(),
		"./snapshots")
}

func TestIntegrationEc2NatGateways(t *testing.T) {
	client.AWSTestHelper(t, Ec2NatGateways(),
		"./snapshots")
}

func TestIntegrationEc2InternetGateways(t *testing.T) {
	client.AWSTestHelper(t, Ec2InternetGateways(),
		"./snapshots")
}
