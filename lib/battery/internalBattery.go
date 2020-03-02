package battery

// InternalBatteryInfo - internal battery info object
type InternalBatteryInfo struct {
	Power      string // AC or DC
	Percentage int
}

// GetInternalBatteryInfo - return the MacBook internal battery info
func GetInternalBatteryInfo() InternalBatteryInfo {
	info := InternalBatteryInfo{"", -1}
	for idx, line := range getLines(getStdout("pmset", "-g", "batt")) {
		if idx == 0 {
			info.Power = matchName("'(.*) Power'", line)
		} else if idx == 1 {
			info.Percentage = matchPercent("([0-9]+)%", line)
		}
	}
	return info
}
