package helpers

import (
	"github.com/jedib0t/go-pretty/v6/text"
	"reflect"
	"testing"
)

func Test_GetLessonColor(t *testing.T) {
	tests := []struct {
		name       string
		want_color text.Color
		lessonType string
	}{
		{
			name:       "Check if function return appropriate color",
			want_color: text.FgRed,
			lessonType: "ЛР",
		},
		{
			name:       "Check if function return appropriate color",
			want_color: text.FgWhite,
			lessonType: "ПР",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color := GetLessonColor(tt.lessonType)
			if !reflect.DeepEqual(color, tt.want_color) {
				t.Errorf("GetLessonColor() = %v, want %v", color, tt.want_color)
			}
		})
	}
}