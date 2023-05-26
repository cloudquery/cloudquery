module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.19

require (
	github.com/cloudquery/plugin-sdk/v3 v3.6.4
	github.com/rs/zerolog v1.28.0
	github.com/apache/arrow/go/v13 v13.0.0-20230509040948-de6c3cd2b604
)

replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230525142029-2d32efeedad8
