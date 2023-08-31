package aiplatform

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/genproto/googleapis/cloud/location"
)

func filterLocation(parent *schema.Resource) bool {
	locationId := parent.Item.(*location.Location).LocationId
	// ListLocations return these, but they don't map to a valid endpoint from https://cloud.google.com/vertex-ai/docs/reference/rest
	return locationId == "us" || locationId == "eu"
}

func filterStudiesLocation(parent *schema.Resource) bool {
	if filterLocation(parent) {
		return true
	}
	locationId := parent.Item.(*location.Location).LocationId
	// The endpoints for these locations return 503 status code
	toFilter := []string{"asia-southeast2", "europe-central2", "me-west1", "us-south1", "us-west3"}
	for _, f := range toFilter {
		if locationId == f {
			return true
		}
	}
	return false
}

func filterIndexesLocations(parent *schema.Resource) bool {
	if filterLocation(parent) {
		return true
	}
	locationId := parent.Item.(*location.Location).LocationId
	// This locations returns `FailedPrecondition desc = Matching Engine is not supported in region europe-west4`
	return locationId == "europe-west4"
}
