module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.19

require (
	github.com/apache/arrow/go/v13 v13.0.0-20230531201200-cbc17a98dfd9
	github.com/cloudquery/plugin-pb-go v1.0.9
	github.com/cloudquery/plugin-sdk/v3 v3.10.4
	github.com/rs/zerolog v1.29.0
)

replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 88d5dc2ed455

require (
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
