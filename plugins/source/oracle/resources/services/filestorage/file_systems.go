package filestorage

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/filestorage"
)

func FileSystems() *schema.Table {
	return &schema.Table{
		Name:      "oracle_filestorage_file_systems",
		Resolver:  fetchFileSystems,
		Multiplex: client.AvailibilityDomainCompartmentMultiplex,
		Transform: transformers.TransformWithStruct(&filestorage.FileSystemSummary{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn,
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}
