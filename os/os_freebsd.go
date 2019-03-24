package osinfo

import (
	"os/exec"
	"strings"
)

func localOSVersion() string {
	command := exec.Command("uname", "-r")
	output, err := command.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}
