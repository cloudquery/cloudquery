package tableoptions

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	accessanalyzertypes "github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	cloudtrailtypes "github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"

	inspector2types "github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/cloudquery/plugin-sdk/v3/caser"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTableOptionsValidate(t *testing.T) {
	tOpts := TableOptions{}
	err := tOpts.Validate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tOpts.CloudTrailEvents = &CloudtrailAPIs{
		LookupEventsOpts: []CustomLookupEventsOpts{
			{
				LookupEventsInput: cloudtrail.LookupEventsInput{
					EndTime:          nil,
					EventCategory:    "",
					LookupAttributes: nil,
					MaxResults:       nil,
					NextToken:        aws.String("123"),
					StartTime:        nil,
				},
			},
		},
	}
	err = tOpts.Validate()
	if err == nil {
		t.Fatal("expected error validating cloud_trail_events, got nil")
	}
}

// TestTableOptionsUnmarshal tests that the TableOptions struct can be unmarshaled from JSON using
// snake_case keys.
func TestTableOptionsUnmarshal(t *testing.T) {
	tableOpts := TableOptions{}
	err := faker.FakeObject(&tableOpts)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(tableOpts)
	if err != nil {
		t.Fatal(err)
	}
	m := map[string]any{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		t.Fatal(err)
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToSnake)
	nb, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var got TableOptions
	err = json.Unmarshal(nb, &got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(tableOpts, got, cmpopts.IgnoreUnexported(
		accessanalyzer.ListFindingsInput{},
		accessanalyzertypes.SortCriteria{},
		accessanalyzertypes.Criterion{},
		cloudtrail.LookupEventsInput{},
		cloudtrailtypes.LookupAttribute{},
		inspector2.ListFindingsInput{},
		inspector2types.StringFilter{},
		inspector2types.DateFilter{},
		inspector2types.NumberFilter{},
		inspector2types.PortRangeFilter{},
		inspector2types.MapFilter{},
		inspector2types.PackageFilter{},
		inspector2types.FilterCriteria{},
		inspector2types.SortCriteria{},
	)); diff != "" {
		t.Fatalf("mismatch between objects after loading from snake case json: %v", diff)
	}
}
