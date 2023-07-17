module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.20

require (
	github.com/apache/arrow/go/v13 v13.0.0-20230630125530-5a06b2ec2a8e
	github.com/cloudquery/plugin-pb-go v1.8.0
	github.com/cloudquery/plugin-sdk/v4 v4.2.0
	github.com/rs/zerolog v1.29.0
)

replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230717001540-8e2219bec8ee

require (
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
