//go:build integration
// +build integration

package ec2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationEc2Vpcs(t *testing.T) {
	client.AWSTestHelper(t, Ec2Vpcs())
}

func TestIntegrationEc2VpcPeeringConnections(t *testing.T) {
	client.AWSTestHelper(t, Ec2VpcPeeringConnections())
}

func TestIntegrationEc2VpcEndpoints(t *testing.T) {
	client.AWSTestHelper(t, Ec2VpcEndpoints())
}

func TestIntegrationEc2TransitGateways(t *testing.T) {
	client.AWSTestHelper(t, Ec2TransitGateways())
}

func TestIntegrationEc2Subnets(t *testing.T) {
	client.AWSTestHelper(t, Ec2Subnets())
}

func TestIntegrationEc2SecurityGroups(t *testing.T) {
	client.AWSTestHelper(t, Ec2SecurityGroups())
}

func TestIntegrationEc2RouteTables(t *testing.T) {
	client.AWSTestHelper(t, Ec2RouteTables())
}

func TestIntegrationEc2NetworkAcls(t *testing.T) {
	client.AWSTestHelper(t, Ec2NetworkAcls())
}

func TestIntegrationEc2NatGateways(t *testing.T) {
	client.AWSTestHelper(t, Ec2NatGateways())
}

func TestIntegrationEc2InternetGateways(t *testing.T) {
	client.AWSTestHelper(t, Ec2InternetGateways())
}
