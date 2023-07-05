module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.19

require (
	github.com/apache/arrow/go/v13 5a06b2ec2a8e
	github.com/cloudquery/plugin-pb-go v1.6.0
	github.com/cloudquery/plugin-sdk/v3 v3.10.6
	github.com/rs/zerolog v1.29.0
)

replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 88d5dc2ed455

require (
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
