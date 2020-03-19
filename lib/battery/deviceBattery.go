package battery

import (
	"strings"
)

// DeviceBatteryInfo - bluetooth device battery info object
type DeviceBatteryInfo struct {
	Device     string
	Percentage int
	Charing    bool
}

// GetDeviceBatteryInfo - return bluetooth devices battery infos
func GetDeviceBatteryInfo() []DeviceBatteryInfo {
	var infos []DeviceBatteryInfo
	for _, val := range strings.Split(getStdout("ioreg", "-k", "BatteryPercent", "-r"), "    {\n") {
		info := DeviceBatteryInfo{"", -1, false}
		for _, line := range getLines(val) {
			if info.Device == "" {
				info.Device = matchName(`"Product" = "(.*)"`, line)
			}
			if info.Percentage < 0 {
				info.Percentage = matchPercent(`"BatteryPercent" = ([0-9]+)`, line)
			}
			if info.Charing == false {
				charing := matchPercent(`"BatteryStatusFlags" = ([0-9]+)`, line)
				info.Charing = charing == 3
			}
			if info.Device != "" && info.Percentage >= 0 {
				infos = append(infos, info)
				break
			}
		}
	}
	return infos
}
