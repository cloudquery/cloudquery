package client

import (
	"testing"
)

func Test_getPrecisionAndScale(t *testing.T) {
	tests := []struct {
		name          string
		dataType      string
		wantPrecision int32
		wantScale     int32
	}{
		{
			name:          "should return default precision and scale when not provided",
			dataType:      "number",
			wantPrecision: 38,
			wantScale:     0,
		},
		{
			name:          "should return default scale when only precision is provided",
			dataType:      "number(10)",
			wantPrecision: 10,
			wantScale:     0,
		},
		{
			name:          "should return precision and scale when precision and scale are provided",
			dataType:      "number(10,2)",
			wantPrecision: 10,
			wantScale:     2,
		},
		{
			name:          "should return default precision when * is provided",
			dataType:      "number(*,3)",
			wantPrecision: 38,
			wantScale:     3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrecision, gotScale := getPrecisionAndScale(tt.dataType)
			if gotPrecision != tt.wantPrecision {
				t.Errorf("getPrecisionAndScale() gotPrecision = %v, want %v", gotPrecision, tt.wantPrecision)
			}
			if gotScale != tt.wantScale {
				t.Errorf("getPrecisionAndScale() gotScale = %v, want %v", gotScale, tt.wantScale)
			}
		})
	}
}
