package homebrew

import (
	"context"
	"testing"
)

func TestClient_GetInstalls(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	installs, err := c.GetInstalls(ctx, Days30)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	checkInstalls(t, installs)
}

func TestClient_GetInstallOnRequestEvents(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	installs, err := c.GetInstallOnRequestEvents(ctx, Days30)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	checkInstalls(t, installs)
}

func checkInstalls(t *testing.T, installs Installs) {
	if installs.TotalCount == 0 {
		t.Errorf("expected installs to be greater than 0, got %d", installs.TotalCount)
	}
	if len(installs.Items) == 0 {
		t.Errorf("expected installs to be greater than 0, got %d", len(installs.Items))
	}
	it := installs.Items[0]
	if it.Count <= 0 {
		t.Errorf("expected installs[0].Count to be greater than 0, got %d", it.Count)
	}
	if it.Percent <= 0 {
		t.Errorf("expected installs[0].Percent to be greater than 0, got %.2f", it.Percent)
	}
	if it.Formula == "" {
		t.Errorf("expected installs[0].Formula to be not empty, got %s", it.Formula)
	}
	if it.Number <= 0 {
		t.Errorf("expected installs[0].Number to be greater than 0, got %d", it.Number)
	}
}

func TestClient_GetCaskInstalls(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	installs, err := c.GetCaskInstalls(ctx, Days30)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	checkCaskInstalls(t, installs)
}

func checkCaskInstalls(t *testing.T, installs CaskInstalls) {
	if installs.TotalCount == 0 {
		t.Errorf("expected installs to be greater than 0, got %d", installs.TotalCount)
	}
	if len(installs.Items) == 0 {
		t.Errorf("expected installs to be greater than 0, got %d", len(installs.Items))
	}
	it := installs.Items[0]
	if it.Count <= 0 {
		t.Errorf("expected installs[0].Count to be greater than 0, got %d", it.Count)
	}
	if it.Percent <= 0 {
		t.Errorf("expected installs[0].Percent to be greater than 0, got %.2f", it.Percent)
	}
	if it.Cask == "" {
		t.Errorf("expected installs[0].Cask to be not empty, got %s", it.Cask)
	}
	if it.Number <= 0 {
		t.Errorf("expected installs[0].Number to be greater than 0, got %d", it.Number)
	}
}

func TestClient_GetBuildErrors(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	errors, err := c.GetBuildErrors(ctx, Days30)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	checkBuildErrors(t, errors)
}

func checkBuildErrors(t *testing.T, errors BuildErrors) {
	if errors.TotalCount == 0 {
		t.Errorf("expected installs to be greater than 0, got %d", errors.TotalCount)
	}
	if len(errors.Items) == 0 {
		t.Errorf("expected installs to be greater than 0, got %d", len(errors.Items))
	}
	it := errors.Items[0]
	if it.Count <= 0 {
		t.Errorf("expected installs[0].Count to be greater than 0, got %d", it.Count)
	}
	if it.Percent <= 0 {
		t.Errorf("expected installs[0].Percent to be greater than 0, got %.2f", it.Percent)
	}
	if it.Formula == "" {
		t.Errorf("expected installs[0].Formula to be not empty, got %s", it.Formula)
	}
	if it.Number <= 0 {
		t.Errorf("expected installs[0].Number to be greater than 0, got %d", it.Number)
	}
}
