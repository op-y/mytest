package date

import (
	"testing"
)

// go test -run=TestGetWeekDayName -coverprofile=coverage.out date/data_test.go
// go tool cover -html=coverage.out
func TestGetWeekDayName(t *testing.T) {
	var tests = []struct {
		day      int
		expected string
	}{
		{0, "星期天"},
		{3, "星期三"},
		{8, "序号不在合理范围"},
	}
	for _, tt := range tests {
		result := GetWeekDayName(tt.day)
		if result != tt.expected {
			t.Errorf("GetWeekDayName(%d): got %s, want %s", tt.day, result, tt.expected)
		}
	}
}
