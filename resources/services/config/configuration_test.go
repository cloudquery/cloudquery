//go:build integration
// +build integration

package config

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationConfigConfigurationRecorders(t *testing.T) {
	client.AWSTestHelper(t, ConfigConfigurationRecorders())
}

func TestIntegrationConfigConformancePack(t *testing.T) {
	client.AWSTestHelper(t, ConfigConformancePack())
}
