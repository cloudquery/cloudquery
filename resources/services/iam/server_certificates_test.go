// +build integration

package iam

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationIamServerCertificates(t *testing.T) {
	client.AWSTestHelper(t, IamServerCertificates())
}
