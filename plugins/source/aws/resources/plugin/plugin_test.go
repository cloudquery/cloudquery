package plugin

import (
	"testing"
)

func TestAWS(t *testing.T) {
	// Note: this test is simple, but serves as a smoke test.
	// The AWS() call below also catches duplicate columns and other issues
	// that may have been missed if mock tests are incomplete.
	p := AWS()
	name := p.Name()
	if name != "aws" {
		t.Errorf("Name() = %q, want %q", name, "aws")
	}
}

// This test ensures that all tables have a unique description.
func TestAWSTableDescriptions(t *testing.T) {
	descriptions := make(map[string]bool)
	p := AWS()
	tables := p.Tables()
	for _, table := range tables {
		if ignoreTable(table.Name) {
			continue
		}
		if _, ok := descriptions[table.Description]; ok || table.Description == "" {
			t.Errorf("description for table exists %s", table.Name)
		} else {
			descriptions[table.Description] = true
		}
	}
}

func ignoreTable(tableName string) bool {
	tablesToIgnore := map[string]bool{
		"aws_resiliencehub_suggested_resiliency_policies": true,
		// TODO: Remove this once we breakup S3 Bucket into multiple tables rather than a single composite table
		"aws_s3_buckets": true,
	}
	_, ok := tablesToIgnore[tableName]
	return ok

}
