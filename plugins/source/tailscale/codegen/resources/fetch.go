package resources

import (
	"github.com/iancoleman/strcase"
)

func (r *Resource) fetcherName() string {
	return "fetch" + strcase.ToCamel(r.SubService)
}
