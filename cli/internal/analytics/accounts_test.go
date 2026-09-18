package analytics

import (
	"fmt"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

func recordWithColumn(t *testing.T, column string, values []*string) arrow.RecordBatch {
	t.Helper()

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "_cq_id", Type: arrow.BinaryTypes.String},
		{Name: column, Type: arrow.BinaryTypes.String, Nullable: true},
	}, nil)
	builder := array.NewRecordBuilder(memory.DefaultAllocator, schema)
	defer builder.Release()

	for i, value := range values {
		builder.Field(0).(*array.StringBuilder).Append(fmt.Sprintf("id-%d", i))
		if value == nil {
			builder.Field(1).(*array.StringBuilder).AppendNull()
			continue
		}
		builder.Field(1).(*array.StringBuilder).Append(*value)
	}
	return builder.NewRecordBatch()
}

func strPtr(s string) *string { return &s }

func TestNewIDCollectorSourceSupport(t *testing.T) {
	for _, sourcePath := range []string{"cloudquery/aws", "cloudquery/gcp", "cloudquery/azure", "cloudquery/k8s", "cloudquery/github", "cloudquery/cloudflare"} {
		if NewIDCollector(sourcePath) == nil {
			t.Errorf("expected a collector for %q", sourcePath)
		}
	}
	for _, sourcePath := range []string{"cloudquery/test", "cloudquery/postgresql", "acme/aws", "acme/github", "aws", ""} {
		if NewIDCollector(sourcePath) != nil {
			t.Errorf("expected no collector for %q", sourcePath)
		}
	}
}

func TestIDCollectorColumnPerSource(t *testing.T) {
	for _, tc := range []struct {
		sourcePath string
		table      string
		column     string
	}{
		{"cloudquery/aws", "aws_lambda_functions", "account_id"},
		{"cloudquery/gcp", "gcp_compute_instances", "project_id"},
		{"cloudquery/azure", "azure_compute_virtual_machines", "subscription_id"},
		{"cloudquery/k8s", "k8s_core_pods", "cloud_cluster_id"},
		{"cloudquery/cloudflare", "cloudflare_zones", "account_id"},
		{"cloudquery/github", "github_issues", "org"},
	} {
		t.Run(tc.sourcePath, func(t *testing.T) {
			collector := NewIDCollector(tc.sourcePath)
			record := recordWithColumn(t, tc.column, []*string{strPtr("one"), strPtr("two")})
			defer record.Release()
			collector.Observe(tc.table, record)

			summaries := collector.Summaries()
			if got := summaries[AccountDimension].Count; got != 2 {
				t.Fatalf("got count %d, want 2", got)
			}
			if got := len(summaries[AccountDimension].Hashes); got != 2 {
				t.Fatalf("got %d hashes, want 2", got)
			}
		})
	}
}

func TestIDCollectorIgnoresUnrelatedColumn(t *testing.T) {
	collector := NewIDCollector("cloudquery/aws")
	record := recordWithColumn(t, "project_id", []*string{strPtr("one")})
	defer record.Release()
	collector.Observe("aws_lambda_functions", record)

	if summaries := collector.Summaries(); summaries != nil {
		t.Fatalf("got %v, want no summaries", summaries)
	}
}

func TestIDCollectorGitHubRepositories(t *testing.T) {
	collector := NewIDCollector("cloudquery/github")

	orgs := recordWithColumn(t, "org", []*string{strPtr("cloudquery"), strPtr("cloudquery")})
	defer orgs.Release()
	collector.Observe("github_issues", orgs)

	repos := recordWithColumn(t, "id", []*string{strPtr("1"), strPtr("2"), strPtr("3")})
	defer repos.Release()
	collector.Observe("github_repositories", repos)

	other := recordWithColumn(t, "id", []*string{strPtr("99")})
	defer other.Release()
	collector.Observe("github_issues", other)

	summaries := collector.Summaries()
	if got := summaries[AccountDimension].Count; got != 1 {
		t.Errorf("got %d accounts, want 1", got)
	}
	if got := summaries[RepositoryDimension].Count; got != 3 {
		t.Errorf("got %d repositories, want 3", got)
	}
}

func TestIDCollectorDeduplicatesAcrossRecords(t *testing.T) {
	collector := NewIDCollector("cloudquery/aws")
	for range 3 {
		record := recordWithColumn(t, "account_id", []*string{strPtr("111111111111"), strPtr("222222222222")})
		collector.Observe("aws_lambda_functions", record)
		record.Release()
	}

	summary := collector.Summaries()[AccountDimension]
	if summary.Count != 2 {
		t.Fatalf("got count %d, want 2", summary.Count)
	}
	if summary.Truncated || summary.CountIsFloor {
		t.Error("got a clipped summary, want an exact one")
	}
}

func TestIDCollectorSkipsNullAndEmpty(t *testing.T) {
	collector := NewIDCollector("cloudquery/aws")
	record := recordWithColumn(t, "account_id", []*string{nil, strPtr(""), strPtr("111111111111")})
	defer record.Release()
	collector.Observe("aws_lambda_functions", record)

	summary := collector.Summaries()[AccountDimension]
	if summary.Count != 1 {
		t.Fatalf("got count %d, want 1", summary.Count)
	}
}

func TestIDCollectorCapsHashesAndReportsCount(t *testing.T) {
	collector := NewIDCollector("cloudquery/aws")
	values := make([]*string, 0, maxHashesPerDimension+10)
	for i := range maxHashesPerDimension + 10 {
		values = append(values, strPtr(fmt.Sprintf("account-%d", i)))
	}
	record := recordWithColumn(t, "account_id", values)
	defer record.Release()
	collector.Observe("aws_lambda_functions", record)

	summary := collector.Summaries()[AccountDimension]
	if got := len(summary.Hashes); got != maxHashesPerDimension {
		t.Errorf("got %d hashes, want %d", got, maxHashesPerDimension)
	}
	if summary.Count != maxHashesPerDimension+10 {
		t.Errorf("got count %d, want %d", summary.Count, maxHashesPerDimension+10)
	}
	if !summary.Truncated {
		t.Error("got truncated false, want true")
	}
	if summary.CountIsFloor {
		t.Error("got count_is_floor true, want false: the count is exact below the tracking cap")
	}
}

func TestIDCollectorStopsTrackingAtCap(t *testing.T) {
	collector := NewIDCollector("cloudquery/aws")
	values := make([]*string, 0, maxTrackedPerDimension+5)
	for i := range maxTrackedPerDimension + 5 {
		values = append(values, strPtr(fmt.Sprintf("account-%d", i)))
	}
	record := recordWithColumn(t, "account_id", values)
	defer record.Release()
	collector.Observe("aws_lambda_functions", record)

	summary := collector.Summaries()[AccountDimension]
	if summary.Count != maxTrackedPerDimension {
		t.Errorf("got count %d, want %d", summary.Count, maxTrackedPerDimension)
	}
	if !summary.CountIsFloor {
		t.Error("got count_is_floor false, want true once tracking stops")
	}
}

func TestNilIDCollectorIsUsable(t *testing.T) {
	var collector *IDCollector
	record := recordWithColumn(t, "account_id", []*string{strPtr("one")})
	defer record.Release()

	collector.Observe("aws_lambda_functions", record)
	if summaries := collector.Summaries(); summaries != nil {
		t.Fatalf("got %v, want nil", summaries)
	}
}

func TestHashIDIsStableAndOpaque(t *testing.T) {
	const accountID = "111111111111"
	hash := hashID(accountID)

	if hash != hashID(accountID) {
		t.Error("hash is not stable across calls")
	}
	if len(hash) != hashLength {
		t.Errorf("got length %d, want %d", len(hash), hashLength)
	}
	if hash == accountID {
		t.Error("hash returned the raw value")
	}
	if hash == hashID("222222222222") {
		t.Error("distinct values produced the same hash")
	}
}
