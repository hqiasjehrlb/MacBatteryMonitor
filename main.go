package main

import (
	"MacBatteryMonitor/lib/battery"
	"fmt"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	done := make(chan bool)
	go setInterval(checkDevicesBattery, 5*60*1000)
	go setInterval(checkInternalBattery, 2*60*1000)
	notify("Battery monitor running")
	<-done
}

func setInterval(callback func(), interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			callback()
		}
	}
}

func notify(msg string) {
	fmt.Println(dateStr(), "Notify:", msg)
	beeep.Notify("Battery level info", msg, "")
}

func checkInternalBattery() {
	info := battery.GetInternalBatteryInfo()
	fmt.Println(dateStr(), "CheckInternalBattery:", info.Power, "-", info.Percentage)
	if info.Percentage >= 0 {
		if info.Percentage < 40 && info.Power != "AC" {
			notify("Internal battery - " + strconv.Itoa(info.Percentage) + "% battery low")
		}
	}
}

func checkDevicesBattery() {
	infos := battery.GetDeviceBatteryInfo()
	for _, info := range infos {
		fmt.Println(dateStr(), "CheckDeviceBattery:", info.Device, "-", info.Percentage)
		if info.Percentage >= 0 {
			if info.Percentage < 40 {
				notify(info.Device + " - " + strconv.Itoa(info.Percentage) + "% battery low")
			}
			if info.Percentage > 80 {
				notify(info.Device + " - " + strconv.Itoa(info.Percentage) + "%")
			}
		}
	}
}

func dateStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}
