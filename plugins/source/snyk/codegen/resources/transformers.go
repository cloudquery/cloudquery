package resources

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/caser"
)

var csr = caser.New(
	caser.WithCustomInitialisms(
		map[string]bool{
			"ACL":  true,
			"DERP": true,
			"SSH":  true,
		},
	),
)

func nameTransformer(f reflect.StructField) (string, error) {
	return csr.ToSnake(f.Name), nil
}
