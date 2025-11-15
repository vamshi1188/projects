package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"time"

	"alaram-cli/bio"
	"alaram-cli/core"
	"alaram-cli/storage"
)

// Execute is the entrypoint for the CLI, given process args.
func Execute(args []string) {
	if len(args) < 2 {
		printUsage()
		return
	}

	switch args[1] {
	case "suggest-alarms":
		fs := flag.NewFlagSet("suggest-alarms", flag.ExitOnError)
		bed := fs.String("bedtime", "", "Planned bedtime (e.g., 23:00 or 11:00PM). REQUIRED")
		cycle := fs.Int("cycle", core.DefaultCycleMinutes, "Baseline cycle length (min)")
		lat := fs.Int("latency", core.DefaultLatencyMinutes, "Baseline sleep latency (min)")
		minc := fs.Int("mincycles", core.DefaultMinCycles, "Minimum cycles to propose")
		maxc := fs.Int("maxcycles", core.DefaultMaxCycles, "Maximum cycles to propose")
		smart := fs.Int("smart-window", core.DefaultSmartWindowMin, "Smart-wake window in minutes (0 to disable)")
		profileUse := fs.Bool("use-profile", true, "Use saved profile adjustments if profile exists")
		_ = fs.Parse(args[2:])

		if *bed == "" {
			fmt.Println("Error: --bedtime is required")
			fs.Usage()
			return
		}
		var prof *bio.Profile
		if *profileUse {
			if p, err := storage.LoadProfile(); err == nil {
				prof = &p
			}
		}
		core.SuggestAlarms(*bed, *lat, *cycle, *minc, *maxc, prof, *smart)

	case "suggest-bedtimes":
		fs := flag.NewFlagSet("suggest-bedtimes", flag.ExitOnError)
		wake := fs.String("wake", "", "Desired wake time (e.g., 07:00 or 7:00AM). REQUIRED")
		cycle := fs.Int("cycle", core.DefaultCycleMinutes, "Baseline cycle length (min)")
		lat := fs.Int("latency", core.DefaultLatencyMinutes, "Baseline sleep latency (min)")
		minc := fs.Int("mincycles", core.DefaultMinCycles, "Minimum cycles to propose")
		maxc := fs.Int("maxcycles", core.DefaultMaxCycles, "Maximum cycles to propose")
		smart := fs.Int("smart-window", core.DefaultSmartWindowMin, "Smart-wake window in minutes (0 to disable)")
		profileUse := fs.Bool("use-profile", true, "Use saved profile adjustments if profile exists")
		_ = fs.Parse(args[2:])

		if *wake == "" {
			fmt.Println("Error: --wake is required")
			fs.Usage()
			return
		}
		var prof *bio.Profile
		if *profileUse {
			if p, err := storage.LoadProfile(); err == nil {
				prof = &p
			}
		}
		core.SuggestBedtimes(*wake, *lat, *cycle, *minc, *maxc, prof, *smart)

	case "save-profile":
		fs := flag.NewFlagSet("save-profile", flag.ExitOnError)
		age := fs.Int("age", 30, "User age (years)")
		chrono := fs.String("chronotype", string(bio.ChronoNeutral), "Chronotype: early|neutral|late")
		avgCycle := fs.Int("avg-cycle", 0, "Your measured average cycle in minutes (optional)")
		avgLatency := fs.Int("avg-latency", 0, "Your measured average latency in minutes (optional)")
		_ = fs.Parse(args[2:])
		chronoLower := strings.ToLower(*chrono)
		if chronoLower != string(bio.ChronoEarly) && chronoLower != string(bio.ChronoNeutral) && chronoLower != string(bio.ChronoLate) {
			fmt.Println("Invalid chronotype. Use early, neutral, or late.")
			return
		}
		p := bio.Profile{
			Age:           *age,
			Chronotype:    chronoLower,
			AvgCycleMin:   *avgCycle,
			AvgLatencyMin: *avgLatency,
		}
		if err := storage.SaveProfile(p); err != nil {
			fmt.Println("Failed to save profile:", err)
			return
		}
		fmt.Println("Profile saved:", p)

	case "show-profile":
		p, err := storage.LoadProfile()
		if err != nil {
			fmt.Println("No saved profile (use save-profile to create one). Error:", err)
			return
		}
		buf, _ := json.MarshalIndent(p, "", "  ")
		fmt.Println(string(buf))

	case "record-sleep":
		fs := flag.NewFlagSet("record-sleep", flag.ExitOnError)
		bed := fs.String("bedtime", "", "Bedtime (HH:MM)")
		sleepStart := fs.String("sleepstart", "", "Sleep start (HH:MM)")
		wake := fs.String("wake", "", "Wake time (HH:MM)")
		mins := fs.Int("minutes", 0, "Total minutes slept")
		notes := fs.String("notes", "", "Optional notes")
		_ = fs.Parse(args[2:])
		if *bed == "" || *sleepStart == "" || *wake == "" || *mins == 0 {
			fmt.Println("Error: --bedtime --sleepstart --wake --minutes are required")
			return
		}
		rec := storage.SleepRecord{
			Date:       time.Now().Format("2006-01-02"),
			Bedtime:    *bed,
			SleepStart: *sleepStart,
			WakeTime:   *wake,
			TotalMin:   *mins,
			Notes:      *notes,
			RecordedAt: time.Now().Format(time.RFC3339),
		}
		if err := storage.AppendHistory(rec); err != nil {
			fmt.Println("Failed to record sleep:", err)
			return
		}
		fmt.Println("Recorded sleep:", rec.Date, rec.TotalMin, "min")

	case "simulate-smartwake":
		fs := flag.NewFlagSet("simulate-smartwake", flag.ExitOnError)
		bed := fs.String("bedtime", "", "Planned bedtime (HH:MM)")
		lat := fs.Int("latency", core.DefaultLatencyMinutes, "Latency (min)")
		cycle := fs.Int("cycle", core.DefaultCycleMinutes, "Cycle length (min)")
		target := fs.String("target-wake", "", "Target wake time (HH:MM) to analyze inside window")
		window := fs.Int("window", core.DefaultSmartWindowMin, "Window length (min)")
		_ = fs.Parse(args[2:])
		if *bed == "" || *target == "" {
			fmt.Println("Error: --bedtime and --target-wake required")
			return
		}
		ptBed, err := core.ParseFlexibleTime(*bed)
		if err != nil {
			fmt.Println("bad bedtime:", err)
			return
		}
		ptTarget, err := core.ParseFlexibleTime(*target)
		if err != nil {
			fmt.Println("bad target:", err)
			return
		}
		sleepStart := core.ToDurationSinceMidnight(ptBed) + time.Duration(*lat)*time.Minute
		best, reason := core.SimulateSmartWake(sleepStart, core.ToDurationSinceMidnight(ptTarget), *window, *cycle)
		fmt.Printf("Best wake in window: %s\nReason: %s\n", core.FormatClockFromDuration(best), reason)

	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Println("Unknown command:", args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Println("SleepAlarm CLI - Team B (Bio-aware recommendations)")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  suggest-alarms   --bedtime=23:00 [--cycle=90 --latency=15 --mincycles=3 --maxcycles=6 --smart-window=20]")
	fmt.Println("  suggest-bedtimes --wake=07:00  [--cycle=90 --latency=15 --mincycles=3 --maxcycles=6 --smart-window=20]")
	fmt.Println("  save-profile     --age=30 --chronotype=neutral --avg-cycle=0 --avg-latency=0")
	fmt.Println("  show-profile")
	fmt.Println("  record-sleep     --bedtime=23:00 --sleepstart=23:15 --wake=07:00 --minutes=465")
	fmt.Println("  simulate-smartwake --bedtime=23:00 --latency=15 --cycle=90 --target-wake=06:45 --window=20")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ./sleepalarm suggest-alarms --bedtime=23:00 --smart-window=20")
	fmt.Println("  ./sleepalarm suggest-bedtimes --wake=07:00 --smart-window=20")
	fmt.Println("  ./sleepalarm save-profile --age=28 --chronotype=late --avg-cycle=92 --avg-latency=20")
	fmt.Println()
	fmt.Println("Data stored in ~/.sleepalarm (profile.json and history.json).")
}
