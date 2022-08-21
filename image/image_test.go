package image

import (
	"testing"
)

func TestWidthAndHeight(t *testing.T) {
	tests := []struct {
		name   string
		size   string
		want_w string
		want_h string
	}{
		{
			name:   "Check if parsed appropriate size",
			size:   "200px,200px",
			want_w: "200px",
			want_h: "200px",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, h := widthAndHeight(tt.size)
			if w != tt.want_w || h != tt.want_h {
				t.Errorf("widthAndHeight() = %s %s, want %s %s", w, h, tt.want_w, tt.want_h)
			}
		})
	}
}
