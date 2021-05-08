package convert

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestLabelsWithNestedBlock(t *testing.T) {
	input := `
block "label_one" "label_two" {
	nested_block { }
}`

	expected := `{
	"block": {
		"label_one": {
			"label_two": [
				{
					"nested_block": [
						{}
					]
				}
			]
		}
	}
}`

	convertedBytes, err := Bytes([]byte(input), "", Options{})
	if err != nil {
		t.Fatal("parse bytes:", err)
	}

	compareTest(t, convertedBytes, expected)
}

func TestSingleBlock(t *testing.T) {
	input := `
block "label_one" {
	attribute = "value"
}
`

	expected := `{
	"block": {
		"label_one": [
			{
				"attribute": "value"
			}
		]
	}
}`

	convertedBytes, err := Bytes([]byte(input), "", Options{})
	if err != nil {
		t.Fatal("parse bytes:", err)
	}

	compareTest(t, convertedBytes, expected)
}

func TestMultipleBlocks(t *testing.T) {
	input := `
block "label_one" {
	attribute = "value"
}
block "label_one" {
	attribute = "value_two"
}
`

	expected := `{
	"block": {
		"label_one": [
			{
				"attribute": "value"
			},
			{
				"attribute": "value_two"
			}
		]
	}
}`

	convertedBytes, err := Bytes([]byte(input), "", Options{})
	if err != nil {
		t.Fatal("parse bytes:", err)
	}

	compareTest(t, convertedBytes, expected)
}

func TestConversion(t *testing.T) {
	const input = `
locals {
	test3 = 1 + 2
	test1 = "hello"
	test2 = 5
	arr = [1, 2, 3, 4]
	hyphen-test = 3
	temp = "${1 + 2} %{if local.test2 < 3}\"4\n\"%{endif}"
	temp2 = "${"hi"} there"
		quoted = "\"quoted\""
		squoted = "'quoted'"
	x = -10
	y = -x
	z = -(1 + 4)
}
locals {
	other = {
		num = local.test2 + 5
		thing = [for x in local.arr: x * 2]
		"${local.test3}" = 4
		3 = 1
		"local.test1" = 89
		"a.b.c[\"hi\"][3].*" = 3
		loop = "This has a for loop: %{for x in local.arr}x,%{endfor}"
		a.b.c = "True"
	}
}
locals {
	heredoc = <<-EOF
		This is a heredoc template.
		It references ${local.other.3}
	EOF
	simple = "${4 - 2}"
	cond = test3 > 2 ? 1: 0
	heredoc2 = <<EOF
		Another heredoc, that
		doesn't remove indentation
		${local.other.3}
		%{if true ? false : true}"gotcha"\n%{else}4%{endif}
	EOF
}
data "terraform_remote_state" "remote" {
	backend = "s3"
	config = {
		profile = var.profile
		region  = var.region
		bucket  = "mybucket"
		key     = "mykey"
	}
}
variable "profile" {}
variable "region" {
	default = "us-east-1"
}
`

	const expected = `{
	"data": {
		"terraform_remote_state": {
			"remote": [
				{
					"backend": "s3",
					"config": {
						"bucket": "mybucket",
						"key": "mykey",
						"profile": "${var.profile}",
						"region": "${var.region}"
					}
				}
			]
		}
	},
	"locals": [
		{
			"arr": [
				1,
				2,
				3,
				4
			],
			"hyphen-test": 3,
			"quoted": "\"quoted\"",
			"squoted": "'quoted'",
			"temp": "${1 + 2} %{if local.test2 \u003c 3}\"4\n\"%{endif}",
			"temp2": "hi there",
			"test1": "hello",
			"test2": 5,
			"test3": "${1 + 2}",
			"x": -10,
			"y": "${-x}",
			"z": "${-(1 + 4)}"
		},
		{
			"other": {
				"${local.test3}": 4,
				"3": 1,
				"a.b.c": "True",
				"a.b.c[\"hi\"][3].*": 3,
				"local.test1": 89,
				"loop": "This has a for loop: %{for x in local.arr}x,%{endfor}",
				"num": "${local.test2 + 5}",
				"thing": "${[for x in local.arr: x * 2]}"
			}
		},
		{
			"cond": "${test3 \u003e 2 ? 1: 0}",
			"heredoc": "This is a heredoc template.\nIt references ${local.other.3}\n",
			"heredoc2": "\t\tAnother heredoc, that\n\t\tdoesn't remove indentation\n\t\t${local.other.3}\n\t\t%{if true ? false : true}\"gotcha\"\\n%{else}4%{endif}\n",
			"simple": "${4 - 2}"
		}
	],
	"variable": {
		"profile": [
			{}
		],
		"region": [
			{
				"default": "us-east-1"
			}
		]
	}
}`

	convertedBytes, err := Bytes([]byte(input), "", Options{})
	if err != nil {
		t.Fatal("parse bytes:", err)
	}

	compareTest(t, convertedBytes, expected)
}

func TestSimplify(t *testing.T) {
	input := `locals {
		a = split("-", "xyx-abc-def")
		x = 1 + 2
		y = pow(2,3)
		t = "x=${4+abs(2-3)*parseint("02",16)}"
		j = jsonencode({
			a = "a"
			b = 5
		})
		with_vars = x + 1
	}`

	expected := `{
	"locals": [
		{
			"a": [
				"xyx",
				"abc",
				"def"
			],
			"j": "{\"a\":\"a\",\"b\":5}",
			"t": "x=6",
			"with_vars": "${x + 1}",
			"x": 3,
			"y": 8
		}
	]
}`

	convertedBytes, err := Bytes([]byte(input), "", Options{Simplify: true})
	if err != nil {
		t.Fatal("parse bytes:", err)
	}

	compareTest(t, convertedBytes, expected)
}

func TestEndOfFileExpr(t *testing.T) {
	input := `inputs = merge(
		{},
		foo().inputs
	)`
	expected := `{
	"inputs": "${merge(\n\t\t{},\n\t\tfoo().inputs\n\t)}"
}`

	convertedBytes, err := Bytes([]byte(input), "", Options{})
	if err != nil {
		t.Fatal("parse bytes:", err)
	}

	compareTest(t, convertedBytes, expected)
}

func compareTest(t *testing.T, input []byte, expected string) {
	var indented bytes.Buffer
	if err := json.Indent(&indented, input, "", "\t"); err != nil {
		t.Fatal("indent:", err)
	}

	actual := indented.String()
	if actual != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, actual)
	}
}