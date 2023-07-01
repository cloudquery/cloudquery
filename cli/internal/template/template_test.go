package template

import "testing"


func TestTemplate(t *testing.T) {
	res, err := ReplaceVariables("test", map[string]any{"test": "test"})
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	if res != "test" {
		t.Fatalf("expected test, got %s", res)
	}

	res, err = ReplaceVariables("@@plugins", map[string]any{"plugins": "test"})
	if err != nil {
		t.Error(err)
	}
	if res != "test" {
		t.Fatalf("expected test, got %s", res)
	}

	res, err = ReplaceVariables("@@plugins.aws", map[string]any{"plugins": map[string]any{"aws": "test"}})
	if err != nil {
		t.Error(err)
	}
	if res != "test" {
		t.Fatalf("expected test, got %s", res)
	}

	_, err = ReplaceVariables("@@plugins.aw", map[string]any{"plugins": map[string]any{"aws": "test"}})
	if err == nil {
		t.Error("expected error")
	}
}