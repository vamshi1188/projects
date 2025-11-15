package bio

import (
	"strings"
	"time"
)

// Chronotype represents a user's daily preference.
type Chronotype string

const (
	ChronoEarly   Chronotype = "early"
	ChronoNeutral Chronotype = "neutral"
	ChronoLate    Chronotype = "late"
)

// Profile captures user bio preferences and baselines.
type Profile struct {
	Age           int       `json:"age"`
	Chronotype    string    `json:"chronotype"`    // early|neutral|late
	AvgCycleMin   int       `json:"avg_cycle_min"` // user baseline
	AvgLatencyMin int       `json:"avg_latency_min"`
	CreatedAt     time.Time `json:"created_at"`
}

// AdjustedCycle returns adjusted cycle length (minutes) based on profile and a baseline.
func AdjustedCycle(profile *Profile, baselineCycle int) int {
	adj := baselineCycle
	if profile != nil {
		switch strings.ToLower(profile.Chronotype) {
		case string(ChronoEarly):
			adj -= 3
		case string(ChronoLate):
			adj += 5
		}
		if profile.Age > 60 {
			adj += 8
		} else if profile.Age >= 40 {
			adj += 4
		}
		if profile.AvgCycleMin > 0 {
			// blend user baseline with system baseline
			adj = (adj + profile.AvgCycleMin) / 2
		}
	}
	if adj < 70 {
		adj = 70
	}
	if adj > 120 {
		adj = 120
	}
	return adj
}

// AdjustedLatency returns adjusted sleep latency (minutes) based on profile and a baseline.
func AdjustedLatency(profile *Profile, baselineLatency int) int {
	lat := baselineLatency
	if profile != nil {
		switch strings.ToLower(profile.Chronotype) {
		case string(ChronoLate):
			lat += 5
		case string(ChronoEarly):
			lat -= 2
		}
		// older age tends to shorter deep sleep but not necessarily latency; small increase.
		if profile.Age >= 60 {
			lat += 3
		}
		if profile.AvgLatencyMin > 0 {
			lat = (lat + profile.AvgLatencyMin) / 2
		}
	}
	if lat < 1 {
		lat = 1
	}
	if lat > 120 {
		lat = 120
	}
	return lat
}

// LightnessProbability estimates probability [0..1] of being in light sleep for a given cycle index and offset.
func LightnessProbability(cycleIndex int, offsetMin int, cycleLen int) float64 {
	// base by cycle index
	base := 0.4 + 0.12*float64(cycleIndex) // cycleIndex 0 -> 0.4, 3 -> 0.76
	if base > 0.95 {
		base = 0.95
	}
	// within cycle: offsetMin ranges 0..cycleLen
	frac := float64(offsetMin) / float64(cycleLen)
	within := 0.1
	if frac > 0.66 {
		within = 0.6 + 0.4*(frac-0.66)/(0.34) // ramp near end into REM
	} else if frac > 0.33 {
		within = 0.3 + 0.3*(frac-0.33)/(0.33)
	} else {
		within = 0.05 + 0.25*(frac/0.33)
	}
	prob := base*0.6 + within*0.4
	if prob < 0 {
		prob = 0
	}
	if prob > 1 {
		prob = 1
	}
	return prob
}
