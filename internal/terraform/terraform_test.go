package terraform

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed testdata/tf_show_ec2.json
var terraformShowOutput []byte

var testTerraformResource = Resource{
	Address:       "module.demo.aws_kms_key.ec2_kms_key",
	Mode:          "managed",
	Type:          "aws_kms_key",
	Name:          "ec2_kms_key",
	ProviderName:  "registry.terraform.io/hashicorp/aws",
	SchemaVersion: 0,
	Values: map[string]interface{}{
		"arn":                                "arn:aws:kms:us-east-1:1234567:key/12341234-234-234-234-123412341234",
		"bypass_policy_lockout_safety_check": false,
		"customer_master_key_spec":           "SYMMETRIC_DEFAULT",
		"deletion_window_in_days":            nil,
		"another_attribute":                  []interface{}{},
		"another_attribute2":                 []int{123},
		"tags": map[string]interface{}{
			"Name": "first-resource",
		},
		"tags_all": map[string]interface{}{},
	},
}

var testTerraformShowOutput = ShowOutput{
	FormatVersion:    "1.0",
	TerraformVersion: "1.1.8",
	Values: ShowOutputValues{
		RootModule: RootModule{
			ChildModules: []ChildModule{
				{
					Address: "module.demo",
					Resources: []Resource{
						testTerraformResource,
						testTerraformResource,
					},
					ChildModules: []ChildModule{{
						Address: "module.demo.module.deny_all_sg",
						Resources: []Resource{
							testTerraformResource,
						}}},
				}}},
	},
}

func TestLoadTerraformShowOutput(t *testing.T) {
	o, err := LoadTerraformShowOutput(terraformShowOutput)
	if err != nil {
		t.Fatal(err)
	}
	content, err := json.MarshalIndent(testTerraformShowOutput, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(content))
	if diff := cmp.Diff(o, testTerraformShowOutput); diff != "" {
		t.Errorf("Config mismatch (-want +got):\n%s", diff)
	}
}

func TestGetAllResourcesByProvider(t *testing.T) {
	o, err := LoadTerraformShowOutput(terraformShowOutput)
	if err != nil {
		t.Fatal(err)
	}
	resources := o.GetAllResourcesByProvider()
	if len(resources) != 2 {
		t.Errorf("Expected 2 providers, got %d", len(resources))
	}
	if len(resources["registry.terraform.io/hashicorp/aws"]) != 2 {
		t.Errorf("Expected 2 aws resources, got %d", len(resources["registry.terraform.io/hashicorp/aws"]))
	}
}
