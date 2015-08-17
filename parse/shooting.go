package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type Shooting struct {
	Number         int      `json:"number"`
	Date           string   `json:"date"`
	AllegedShooter string   `json:"alleged_shooter"`
	Killed         int      `json:"killed"`
	Wounded        int      `json:"wounded"`
	City           string   `json:"city"`
	State          string   `json:"state"`
	References     []string `json:"references"`
}

func ShootingsFromTextArea(body string) (shootings []Shooting, err error) {
	statsThenLink := strings.Split(body, "|-\n")
	for _, section := range statsThenLink[1:] {
		if !strings.Contains(section, "<ref>") {
			// no <ref> indicates stats
			eachStat := strings.Split(section, "\n")
			if len(eachStat) < 6 {
				continue
			}
			var strStats []string
			for i, stat := range eachStat {
				strStats = append(strStats, strings.Trim(stat, "|"))
				if i == 5 {
					break
				}
			}
			shootingStats, err := buildStats(strStats)
			if err != nil {
				return []Shooting{}, err
			}
			shootings = append(shootings, shootingStats)
		} else {
			// <ref> presence means link
			refLine := strings.Trim((strings.Split(section, "\n"))[0], "|")
			shootings[len(shootings)-1].References = buildRefLines(refLine)
		}
	}
	return
}

func buildStats(stats []string) (shooting Shooting, err error) {
	if len(stats) < 6 || 6 < len(stats) {
		return Shooting{}, fmt.Errorf("Cannot build stats without exactly 6 - %d found", len(stats))
	}
	number, err := strconv.Atoi(strings.TrimSpace(stats[0]))
	if err != nil {
		return Shooting{}, err
	}
	killed, err := strconv.Atoi(strings.TrimSpace(stats[3]))
	if err != nil {
		return Shooting{}, err
	}
	wounded, err := strconv.Atoi(strings.TrimSpace(stats[4]))
	if err != nil {
		return Shooting{}, err
	}
	shooting.Number = number
	shooting.Date = stats[1]
	shooting.AllegedShooter = stats[2]
	shooting.Killed = killed
	shooting.Wounded = wounded
	city, state := cityState(stats[5])
	shooting.City = city
	shooting.State = state
	return
}

func cityState(s string) (string, string) {
	parts := strings.Split(s, ", ")
	if len(parts) < 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

func buildRefLines(str string) (refs []string) {
	eachRef := strings.Split(str, "</ref>")
	for _, ref := range eachRef {
		justURL := strings.Trim(strings.TrimSpace(ref), "<ref>")
		if justURL != "" {
			refs = append(refs, strings.Trim(strings.TrimSpace(ref), "<ref>"))
		}
	}
	return
}
