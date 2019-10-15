package pmjp

import (
	"testing"
	"time"
)

func TestToPmjp(t *testing.T) {
	testcases := []struct {
		inputYear  int
		inputMonth int
		inputDay   int
		expected   string
	}{
		{1, 1, 1, ""},
		{1885, 12, 23, "伊藤博文"},
		{1912, 12, 21, "桂太郎"},
		{1992, 11, 12, "宮澤喜一"},
		{2019, 5, 1, "安倍晋三"},
	}

	for _, tt := range testcases {
		_t := time.Date(tt.inputYear, time.Month(tt.inputMonth), tt.inputDay, 0, 0, 0, 0, time.UTC)
		actual := ToPmjp(_t)
		if actual != tt.expected {
			t.Fatalf("want %v, but got %v", tt.expected, actual)
		}
	}
}

func TestToPmjpNow(t *testing.T) {
	expected := "安倍晋三"
	actual := ToPmjp(time.Now())

	if expected != actual {
		t.Fatalf("want %v, but got %v", expected, actual)
	}
}
