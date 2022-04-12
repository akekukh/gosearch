package geom

import "testing"

func TestGeom_Distance(t *testing.T) {
	tests := []struct {
		name         string
		lat1         float64
		long1        float64
		lat2         float64
		long2        float64
		wantDistance float64
	}{
		{
			name:         "#1",
			lat1:         1,
			long1:        1,
			lat2:         4,
			long2:        5,
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := Distance(tt.lat1, tt.long1, tt.lat2, tt.long2); gotDistance != tt.wantDistance {
				t.Errorf("Geom.Distance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
