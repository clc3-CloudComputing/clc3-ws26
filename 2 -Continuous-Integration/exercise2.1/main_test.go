package main

import (
	"testing"
	"time"
)

func TestRoundToNearestMinute(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
		want time.Time
	}{
		{
			name: "round down",
			in:   time.Date(2026, time.January, 1, 1, 1, 15, 0, time.UTC),
			want: time.Date(2026, time.January, 1, 1, 1, 0, 0, time.UTC),
		},
		{
			name: "round up",
			in:   time.Date(2026, time.January, 1, 1, 1, 30, 0, time.UTC),
			want: time.Date(2026, time.January, 1, 1, 2, 0, 0, time.UTC),
		},
		{
			name: "roll over hour",
			in:   time.Date(2026, time.January, 1, 1, 59, 30, 0, time.UTC),
			want: time.Date(2026, time.January, 1, 2, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundToNearestMinute(tt.in); !got.Equal(tt.want) {
				t.Errorf("roundToNearestMinute(%v) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}
}
