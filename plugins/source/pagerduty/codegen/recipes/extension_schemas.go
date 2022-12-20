package recipes

import "github.com/PagerDuty/go-pagerduty"

func ExtensionSchemaResources() []*Resource {
	return []*Resource{
		{
			SubService:  "extension_schemas",
			Description: "https://developer.pagerduty.com/api-reference/6eef27c5b452f-list-extension-schemas",
			Struct:      pagerduty.ExtensionSchema{},
			PKColumns:   []string{"id"},

			ListOptionsStructNameOverride: "ListExtensionSchemaOptions",
			ResponseFieldOverride:         "ExtensionSchemas",
			ResponseStructOverride:        "ListExtensionSchemaResponse",
		},
	}
}
