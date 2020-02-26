package battery

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// DeviceBatteryInfo - bluetooth device battery info object
type DeviceBatteryInfo struct {
	Device     string
	Percentage int
}

// GetDeviceBatteryInfo - return bluetooth devices battery infos
func GetDeviceBatteryInfo() []DeviceBatteryInfo {
	var infos []DeviceBatteryInfo
	out, _ := exec.Command("ioreg", "-k", "BatteryPercent", "-r").Output()
	ar := strings.Split(string(out), "    {\n")
	for _, val := range ar {
		var info DeviceBatteryInfo
		info.Percentage = -1

		lines := strings.Split(val, "\n")
		for _, line := range lines {
			match := regexp.MustCompile(`"Product" = "(.*)"`).FindStringSubmatch(line)
			if len(match) > 1 {
				info.Device = match[1]
			}
			match = regexp.MustCompile(`"BatteryPercent" = ([0-9]+)`).FindStringSubmatch(line)
			if len(match) > 1 {
				p, err := strconv.Atoi(match[1])
				if err == nil {
					info.Percentage = p
				}
			}
		}
		if info.Percentage >= 0 {
			infos = append(infos, info)
		}
	}
	return infos
}
