package core

import (
	"fmt"
	"time"

	"alaram-cli/bio"
)

// SuggestAlarms prints alarm suggestions given a bedtime string and parameters.
func SuggestAlarms(bedtimeStr string, baselineLatency, baselineCycle, minCycles, maxCycles int, profile *bio.Profile, smartWindow int) {
	pt, err := ParseFlexibleTime(bedtimeStr)
	if err != nil {
		fmt.Println("Error parsing bedtime:", err)
		return
	}
	bedtime := ToDurationSinceMidnight(pt)
	adjCycle := bio.AdjustedCycle(profile, baselineCycle)
	adjLatency := bio.AdjustedLatency(profile, baselineLatency)
	sleepStart := bedtime + time.Duration(adjLatency)*time.Minute

	fmt.Println("Planned lights-off:", FormatClockFromDuration(bedtime))
	fmt.Printf("Estimated sleep start (bedtime + latency %d min): %s\n", adjLatency, FormatClockFromDuration(sleepStart))
	fmt.Printf("Using cycle length ~%d min (adjusted by profile)\n", adjCycle)
	fmt.Println()

	fmt.Println("Suggested alarm times (cycle-aligned):")
	var bestN int
	var bestDiff float64 = 1e9
	for n := minCycles; n <= maxCycles; n++ {
		wake := sleepStart + time.Duration(n*adjCycle)*time.Minute
		totalHours := float64(n*adjCycle) / 60.0
		fmt.Printf("  %d cycles -> %s  (%.1f hours total sleep)\n", n, FormatClockFromDuration(wake), totalHours)
		if diff := AbsFloat(totalHours - 7.5); diff < bestDiff {
			bestDiff = diff
			bestN = n
		}
	}
	if bestN != 0 {
		bestWake := sleepStart + time.Duration(bestN*adjCycle)*time.Minute
		fmt.Printf("\nRecommendation: %d cycles -> %s (closest to 7.5 h goal)\n", bestN, FormatClockFromDuration(bestWake))
		if smartWindow > 0 {
			bestMinute, reason := SimulateSmartWake(sleepStart, bestWake, smartWindow, adjCycle)
			fmt.Printf("SmartWake enabled: best moment within %d min window -> %s\n", smartWindow, FormatClockFromDuration(bestMinute))
			fmt.Println("  reason:", reason)
		} else {
			fmt.Println("Tip: enable smart-wake (--smart-window) to try waking during lighter sleep.")
		}
	}
}

// SuggestBedtimes prints bedtimes suggestions given a target wake time.
func SuggestBedtimes(wakeStr string, baselineLatency, baselineCycle, minCycles, maxCycles int, profile *bio.Profile, smartWindow int) {
	pt, err := ParseFlexibleTime(wakeStr)
	if err != nil {
		fmt.Println("Error parsing wake time:", err)
		return
	}
	wake := ToDurationSinceMidnight(pt)
	adjCycle := bio.AdjustedCycle(profile, baselineCycle)
	adjLatency := bio.AdjustedLatency(profile, baselineLatency)

	fmt.Println("Desired wake time:", FormatClockFromDuration(wake))
	fmt.Printf("Assuming sleep latency %d min and cycle %d min (adjusted by profile)\n", adjLatency, adjCycle)
	fmt.Println()

	fmt.Println("Suggested bedtimes (cycle-aligned):")
	for n := maxCycles; n >= minCycles; n-- {
		bed := wake - time.Duration(n*adjCycle)*time.Minute - time.Duration(adjLatency)*time.Minute
		totalMin := n * adjCycle
		fmt.Printf("  %d cycles -> go to bed at %s (sleep start ~%s) => %.1f hours sleep\n", n, FormatClockFromDuration(bed), FormatClockFromDuration(bed+time.Duration(adjLatency)*time.Minute), float64(totalMin)/60.0)
	}
	if minCycles <= 5 && maxCycles >= 5 {
		bed := wake - time.Duration(5*adjCycle)*time.Minute - time.Duration(adjLatency)*time.Minute
		fmt.Printf("\nRecommendation: 5 cycles -> go to bed at %s for ~7.5 hours of sleep\n", FormatClockFromDuration(bed))
		if smartWindow > 0 {
			bestMinute, reason := SimulateSmartWake(bed+time.Duration(adjLatency)*time.Minute, wake, smartWindow, adjCycle)
			fmt.Printf("SmartWake: best wake within %d min -> %s\n", smartWindow, FormatClockFromDuration(bestMinute))
			fmt.Println("  reason:", reason)
		}
	}
}
