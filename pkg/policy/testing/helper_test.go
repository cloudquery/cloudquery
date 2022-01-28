package testing

import (
	"testing"
)

func TestManager_Execute(t *testing.T) {
	// Skip test for now since github is annoying
	TestPolicy(t, "../../../.cq/policies/github.com/cloudquery-policies/aws", "foundational_security/ec2/EC2.18", "../../../database-data")

}
