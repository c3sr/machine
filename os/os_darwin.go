package osinfo

import (
	"os/exec"
	"strings"
)

func localOSVersion() string {
	command := exec.Command("bash", "-c", `sw_vers | grep ProductVersion | cut -d$'\t' -f2`)
	output, err := command.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}
