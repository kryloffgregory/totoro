package execute

import (
	"os/exec"
	"strings"
)

func ExecuteString(command string) (string, error) {
	tokens := strings.Split(command, " ")
	bytes, _ := exec.Command(tokens[0], tokens[1:]...).CombinedOutput()
	return string(bytes), nil
}
