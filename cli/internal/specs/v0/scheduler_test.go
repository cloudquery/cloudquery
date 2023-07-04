package specs

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestSchedulerJsonMarshalUnmarshal(t *testing.T) {
	b, err := json.Marshal(SchedulerDFS)
	if err != nil {
		t.Fatal("failed to marshal scheduler:", err)
	}
	var scheduler Scheduler
	if err := json.Unmarshal(b, &scheduler); err != nil {
		t.Fatal("failed to unmarshal scheduler:", err)
	}
	if scheduler != SchedulerDFS {
		t.Fatal("expected scheduler to be dfs, but got:", scheduler)
	}
}

func TestSchedulerYamlMarshalUnmarshal(t *testing.T) {
	b, err := yaml.Marshal(SchedulerDFS)
	if err != nil {
		t.Fatal("failed to marshal scheduler:", err)
	}
	var scheduler Scheduler
	if err := yaml.Unmarshal(b, &scheduler); err != nil {
		t.Fatal("failed to unmarshal scheduler:", err)
	}
	if scheduler != SchedulerDFS {
		t.Fatal("expected scheduler to be dfs, but got:", scheduler)
	}
}
