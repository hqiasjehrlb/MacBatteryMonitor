package battery

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// InternalBatteryInfo - internal battery info object
type InternalBatteryInfo struct {
	Power      string
	Percentage int
}

// GetInternalBatteryInfo - return the MacBood internal battery info
func GetInternalBatteryInfo() InternalBatteryInfo {
	var info InternalBatteryInfo
	info.Percentage = -1
	out, _ := exec.Command("pmset", "-g", "batt").Output()
	lines := strings.Split(string(out), "\n")
	for idx, line := range lines {
		if idx == 0 {
			match := regexp.MustCompile("'(.*) Power'").FindStringSubmatch(line)
			if len(match) > 1 {
				info.Power = match[1]
			}
		} else if idx == 1 {
			match := regexp.MustCompile("([0-9]+)%").FindStringSubmatch(line)
			if len(match) > 1 {
				p, err := strconv.Atoi(match[1])
				if err == nil {
					info.Percentage = p
				}
			}
		}
	}

	return info
}
