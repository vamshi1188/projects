package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"alaram-cli/bio"
)

const (
	appDirName      = ".sleepalarm"
	profileFileName = "profile.json"
	historyFileName = "history.json"
)

// SleepRecord represents a nightly sleep record for history tracking.
type SleepRecord struct {
	Date       string `json:"date"`        // ISO date of sleep start
	Bedtime    string `json:"bedtime"`     // HH:MM
	SleepStart string `json:"sleep_start"` // HH:MM
	WakeTime   string `json:"wake_time"`   // HH:MM
	TotalMin   int    `json:"total_min"`
	Notes      string `json:"notes,omitempty"`
	RecordedAt string `json:"recorded_at"`
}

func homeConfigDir() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(hd, appDirName)
	return dir, nil
}

func ensureConfigDir() (string, error) {
	dir, err := homeConfigDir()
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0700); err != nil {
			return "", err
		}
	}
	return dir, nil
}

func profilePath() (string, error) {
	dir, err := ensureConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, profileFileName), nil
}

func historyPath() (string, error) {
	dir, err := ensureConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, historyFileName), nil
}

// SaveProfile writes the profile to disk and stamps CreatedAt.
func SaveProfile(p bio.Profile) error {
	p.CreatedAt = time.Now()
	path, err := profilePath()
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(p)
}

// LoadProfile reads the profile from disk.
func LoadProfile() (bio.Profile, error) {
	var p bio.Profile
	path, err := profilePath()
	if err != nil {
		return p, err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return p, err
	}
	if err := json.Unmarshal(b, &p); err != nil {
		return p, err
	}
	return p, nil
}

// AppendHistory appends a record to history.json
func AppendHistory(rec SleepRecord) error {
	path, err := historyPath()
	if err != nil {
		return err
	}
	var arr []SleepRecord
	if b, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(b, &arr)
	}
	arr = append(arr, rec)
	buf, err := json.MarshalIndent(arr, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, buf, 0600)
}

// LoadHistory returns the history records.
func LoadHistory() ([]SleepRecord, error) {
	path, err := historyPath()
	if err != nil {
		return nil, err
	}
	var arr []SleepRecord
	if b, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(b, &arr)
	}
	return arr, nil
}

// ComputeHistoryAverages computes simple averages; currently returns avg total sleep minutes as second value.
func ComputeHistoryAverages() (avgLatency float64, avgTotalSleep float64, _ error) {
	recs, err := LoadHistory()
	if err != nil {
		return 0, 0, err
	}
	if len(recs) == 0 {
		return 0, 0, errors.New("no history")
	}
	total := 0
	for _, r := range recs {
		total += r.TotalMin
	}
	return 0, float64(total) / float64(len(recs)), nil
}
