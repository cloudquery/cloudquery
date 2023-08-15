package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/salesforce/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var notSupportedBulkObjects = map[string]bool{
	"WorkOrderLineItemStatus":          true,
	"WorkOrderStatus":                  true,
	"UndecidedEventRelation":           true,
	"TaskWhoRelation":                  true,
	"TaskStatus":                       true,
	"TaskPriority":                     true,
	"SolutionStatus":                   true,
	"RecentlyViewed":                   true,
	"PartnerRole":                      true,
	"OrderStatus":                      true,
	"In_App_Checklist_Settings__Share": true,
	"FieldSecurityClassification":      true,
	"EventWhoRelation":                 true,
	"DeclinedEventRelation":            true,
	"ContractStatus":                   true,
	"CaseStatus":                       true,
	"AcceptedEventRelation":            true,
}

type describeResponse struct {
	Fields []map[string]any `json:"fields"`
}

func fetchObjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	var describeRes describeResponse
	if err := c.Get(ctx, c.HTTPDataEndpoint+"/sobjects/"+c.Object+"/describe", &describeRes); err != nil {
		return err
	}

	compoundFields := make(map[string]bool)
	for _, field := range describeRes.Fields {
		// compound fields cannot be fetched in bulk but they are available
		// in the root fields so that's just a duplication
		if field["compoundFieldName"] != nil {
			compoundFields[field["compoundFieldName"].(string)] = true
		}
	}

	fieldsSupportedInBulk := true

	fields := make([]string, 0, len(describeRes.Fields))
	for _, field := range describeRes.Fields {
		if compoundFields[field["name"].(string)] {
			continue
		}
		fieldType := field["type"].(string)
		if fieldType == "base64" {
			fieldsSupportedInBulk = false
		}
		fields = append(fields, field["name"].(string))
	}

	// There are two APIs that can be used to fetch data from Salesforce:
	// - Bulk API: https://developer.salesforce.com/docs/atlas.en-us.api_bulk_v2.meta/api_bulk_v2/intro_bulk_api.htm
	// - Query API: https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_query.htm
	// Bulk API is faster but it doesn't support all the data types. If the object
	// is not supported by the Bulk API or if the object is supported but it has
	// fields that are not supported by the Bulk API then we fallback to the Query API.
	if notSupportedBulkObjects[c.Object] || !fieldsSupportedInBulk {
		return fetchQueryApi(ctx, c, fields, res)
	}

	jobId, err := createQueryJob(ctx, c, fields)
	if err != nil {
		return err
	}

	return fetchJobResults(ctx, c, jobId, fields, res)
}
