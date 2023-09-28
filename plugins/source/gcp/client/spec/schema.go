package spec

import (
	"log"

	"github.com/cloudquery/codegen/jsonschema"
)

var jsonSchema string

func init() {
	data, err := jsonschema.Generate(new(Spec))
	if err != nil {
		log.Fatal(err)
	}
	jsonSchema = string(data)
}

func JSONSchema() string {
	return jsonSchema
}
