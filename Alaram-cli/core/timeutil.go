package core

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseFlexibleTime accepts formats like "23:00", "11:00PM", "11PM", "11:00 PM", "15:04:05".
func ParseFlexibleTime(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	layouts := []string{"15:04", "3:04PM", "3:04 PM", "3PM", "3 PM", "15:04:05"}
	var t time.Time
	var err error
	for _, l := range layouts {
		t, err = time.Parse(l, s)
		if err == nil {
			return t, nil
		}
	}
	// numeric HHMM
	if len(s) == 4 {
		if h, err1 := strconv.Atoi(s[:2]); err1 == nil {
			if m, err2 := strconv.Atoi(s[2:]); err2 == nil {
				return time.Date(1, 1, 1, h, m, 0, 0, time.Local), nil
			}
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse time: %s", s)
}

// ToDurationSinceMidnight converts time to duration since midnight.
func ToDurationSinceMidnight(t time.Time) time.Duration {
	return time.Duration(t.Hour())*time.Hour + time.Duration(t.Minute())*time.Minute
}

// FormatClockFromDuration returns HH:MM from duration since midnight (wraps 24h).
func FormatClockFromDuration(d time.Duration) string {
	d = d % (24 * time.Hour)
	if d < 0 {
		d += 24 * time.Hour
	}
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

func AbsFloat(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
