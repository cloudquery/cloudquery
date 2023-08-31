package client

import (
	"testing"
)

func Test_getPrecisionAndScale(t *testing.T) {
	tests := []struct {
		name          string
		dataTypes     []string
		wantPrecision int32
		wantScale     int32
	}{
		{
			name:          "should return default precision and scale when not provided",
			dataTypes:     []string{"decimal", "numeric"},
			wantPrecision: 10,
			wantScale:     0,
		},
		{
			name:          "should return default scale when only precision is provided",
			dataTypes:     []string{"decimal(15)", "numeric(15)"},
			wantPrecision: 15,
			wantScale:     0,
		},
		{
			name:          "should return precision and scale when precision and scale are provided",
			dataTypes:     []string{"decimal(15,2)", "numeric(15,2)"},
			wantPrecision: 15,
			wantScale:     2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, dataType := range tt.dataTypes {
				gotPrecision, gotScale := getPrecisionAndScale(dataType)
				if gotPrecision != tt.wantPrecision {
					t.Errorf("getPrecisionAndScale() gotPrecision = %v, want %v", gotPrecision, tt.wantPrecision)
				}
				if gotScale != tt.wantScale {
					t.Errorf("getPrecisionAndScale() gotScale = %v, want %v", gotScale, tt.wantScale)
				}
			}
		})
	}
}
