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
		actual := ToPmjp(tt.inputYear, tt.inputMonth, tt.inputDay)
		if actual != tt.expected {
			t.Fatalf("want %v, but got %v", tt.expected, actual)
		}
	}
}

func TestToPmjpFromTime(t *testing.T) {
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
		actual := ToPmjpFromTime(_t)
		if actual != tt.expected {
			t.Fatalf("want %v, but got %v", tt.expected, actual)
		}
	}
}

func TestToPmjpNow(t *testing.T) {
	expected := "安倍晋三"
	actual := ToPmjpFromTime(time.Now())

	if expected != actual {
		t.Fatalf("want %v, but got %v", expected, actual)
	}
}

func TestCalcTenure(t *testing.T) {
	testcases := []struct {
		inputName string
		expected  int
	}{
		{"", 0},
		{"東久邇宮稔彦王", 54},
		{"田中角栄", 886},
		{"吉田茂", 2616},
	}

	for _, tt := range testcases {
		actual := CalcTenure(tt.inputName)
		if actual != tt.expected {
			t.Fatalf("%s's tenure: want %v, but got %v", tt.inputName, tt.expected, actual)
		}
	}
}
