package recipes

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func OSS() []*Resource {
	bucketName := codegen.ColumnDefinition{
		Name:     "name",
		Type:     schema.TypeString,
		Resolver: `client.ResolveParentColumn("Name")`,
		Options: schema.ColumnCreationOptions{
			PrimaryKey: true,
		},
	}
	updateDate := codegen.ColumnDefinition{
		Name:     "update_date",
		Type:     schema.TypeString,
		Resolver: `client.ResolveUpdateDate()`,
		Options: schema.ColumnCreationOptions{
			PrimaryKey: true,
		},
	}
	return []*Resource{
		{
			Service:    "oss",
			SubService: "buckets",
			Struct:     new(oss.BucketProperties),
			TableName:  "oss_buckets",
			Multiplex:  "",
			SkipFields: []string{"XMLName"},
			Relations:  []string{"BucketStats()"},
			PKColumns:  []string{"name"},
		},
		{
			Service:      "oss",
			SubService:   "bucket_stats",
			Multiplex:    "", // we skip multiplexing here as it's a relation
			Struct:       new(oss.BucketStat),
			TableName:    "oss_bucket_stats",
			SkipFields:   []string{"XMLName"},
			ExtraColumns: []codegen.ColumnDefinition{bucketName, updateDate},
		},
	}
}
