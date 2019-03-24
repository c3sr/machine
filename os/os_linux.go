package osinfo

import (
	"os/exec"
	"strings"
)

func localOSVersion() string {
	command := exec.Command("bash", "-c", `cat /etc/os-release | grep 'VERSION=' | cut -d'=' -f2`)
	output, err := command.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}
