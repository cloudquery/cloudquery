module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.22.4

require (
	github.com/apache/arrow-go/v18 v18.0.0
	github.com/cloudquery/plugin-pb-go v1.26.4
	github.com/cloudquery/plugin-sdk/v4 v4.72.6
	github.com/rs/zerolog v1.33.0
)

require (
	google.golang.org/genproto/googleapis/rpc 1f4bbc51befe // indirect
)
