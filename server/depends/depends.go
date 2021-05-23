package depends

import (
	"os/exec"
	"strings"
)

func GetRDepends(lib string) ([]string, error) {
	bytes, _ := exec.Command("apt-rdepends", "-rsv", lib).Output()
	lines := string(bytes)
	depStrings := strings.Split(lines, "\n")
	return depStrings, nil
}
