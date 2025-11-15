package core
package core

import (
	"fmt"
	"time"

	"alaram-cli/bio"
)

// SimulateSmartWake samples each minute in [targetWake-windowMin .. targetWake] and picks
// the minute with highest LightnessProbability. Returns best wake minute and a reason string.
func SimulateSmartWake(sleepStart time.Duration, targetWake time.Duration, windowMin int, cycleLen int) (time.Duration, string) {
	windowStart := targetWake - time.Duration(windowMin)*time.Minute
	if windowStart < 0 {
		windowStart += 24 * time.Hour
	}
	bestProb := -1.0
	var bestMinute time.Duration
	for i := 0; i <= windowMin; i++ {
		cand := targetWake - time.Duration(windowMin-i)*time.Minute
		minsSinceSleep := int((cand - sleepStart).Minutes())
		if minsSinceSleep < 0 {
			minsSinceSleep += 24 * 60
		}
		cycleIndex := minsSinceSleep / cycleLen
		offset := minsSinceSleep % cycleLen
		prob := bio.LightnessProbability(cycleIndex, offset, cycleLen)
		prob = prob*0.9 + 0.1*float64(i)/float64(windowMin)
		if prob > bestProb {
			bestProb = prob
			bestMinute = cand
		}
	}
	reason := fmt.Sprintf("selected minute has estimated lightness %.2f (higher is lighter sleep). window %d min", bestProb, windowMin)
	return bestMinute, reason
}
