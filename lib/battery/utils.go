package battery

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func matchName(reg string, str string) string {
	match := regexp.MustCompile(reg).FindStringSubmatch(str)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func matchPercent(reg string, str string) int {
	match := regexp.MustCompile(reg).FindStringSubmatch(str)
	if len(match) > 1 {
		p, err := strconv.Atoi(match[1])
		if err == nil {
			return p
		}
	}
	return -1
}

func getStdout(command string, arg ...string) string {
	out, err := exec.Command(command, arg...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func getLines(str string) []string {
	return strings.Split(str, "\n")
}
