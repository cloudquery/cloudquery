module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.20

require (
	github.com/apache/arrow/go/v13 112f94971882
	github.com/cloudquery/plugin-pb-go v1.9.3
	github.com/cloudquery/plugin-sdk/v4 v4.5.5
	github.com/rs/zerolog v1.29.0
)

replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 10df4b9d1986

require (
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
