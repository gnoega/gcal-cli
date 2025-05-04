package timeutils

import (
	"testing"
	"time"
)

func TestConverToGoLayout(t *testing.T) {

	testCase := []struct {
		title, value, expected string
	}{
		{"%d-%m-%Y to 01-02-2006", "%d-%m-%Y", "02-01-2006"},
		{"%d-%m-%Y %H:%M to 01-02-2006 15:04", "%d-%m-%Y %H:%M", "02-01-2006 15:04"},
		{"%A %d %B %Y to Monday 02 01 2006", "%A %d %B %Y", "Monday 02 January 2006"},
	}

	for _, tc := range testCase {
		t.Run(tc.title, func(t *testing.T) {
			result := ConvertToGoLayout(tc.value)
			if result != tc.expected {
				t.Errorf("expected %s, but got %s", tc.expected, result)
			}
		})
	}

}

func TestDaysIn(t *testing.T) {
	testCase := []struct {
		month    time.Month
		year     int
		expected uint32
		label    string
	}{
		{time.January, 2025, 31, "January has 31 days"},
		{time.February, 2025, 28, "February non-leap year"},
		{time.February, 2024, 29, "February leap year"},
		{time.March, 2025, 31, "March has 31 days"},
		{time.April, 2025, 30, "April has 30 days"},
		{time.May, 2025, 31, "May has 31 days"},
		{time.June, 2025, 30, "June has 30 days"},
		{time.July, 2025, 31, "July has 31 days"},
		{time.August, 2025, 31, "August has 31 days"},
		{time.September, 2025, 30, "September has 30 days"},
		{time.October, 2025, 31, "October has 31 days"},
		{time.November, 2025, 30, "November has 30 days"},
		{time.December, 2025, 31, "December has 31 days"},
	}

	for _, tc := range testCase {
		got := DaysIn(tc.month, tc.year)

		if got != int(tc.expected) {
			t.Errorf("%s: got %d days, expected to have %d", tc.label, got, tc.expected)
		}
	}

}
