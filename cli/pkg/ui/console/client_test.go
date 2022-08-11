package console

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cloudquery/cloudquery/cli/internal/analytics"
	"github.com/google/uuid"
)

func TestCreateClient(t *testing.T) {
	_ = analytics.Init()

	_, filename, _, _ := runtime.Caller(0)
	fixtures := filepath.Join(filepath.Dir(filename), "fixtures")

	tests := []struct {
		name       string
		configPath string
		wantErr    bool
	}{
		{
			name:       "valid",
			configPath: filepath.Join(fixtures, "config.yml"),
			wantErr:    false,
		},
		{
			name:       "invalid config",
			configPath: filepath.Join(fixtures, "boom.yml"),
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := CreateClient(context.Background(), tt.configPath, false, nil, uuid.Nil); (err != nil) != tt.wantErr {
				t.Errorf("CreateClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
