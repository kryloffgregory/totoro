package execute

import (
	"fmt"
	"os/exec"
	"strings"

	"accessModel/server/issue"
)

func Execute(command issue.Command) ([]byte, error) {
	switch command.Type {
	case issue.CommandTypeInstall:
		var packName string
		if command.InstallPayload.Version == "" {
			packName = command.InstallPayload.PackageName
		} else {
			packName = command.InstallPayload.PackageName+"="+command.InstallPayload.Version
		}
		fmt.Println(packName)
		//return exec.Command("ls", "-l").Output()
		cmd:=exec.Command("brew", "info", packName)
		fmt.Println(cmd.String())
		return cmd.CombinedOutput()
		//return exec.Command("apt", "-y","install", packName, "--no-upgrade").Output()
	case issue.CommandTypeRemove:
		return exec.Command("apt", "-y", "remove", "--purge", command.RemovePayload.PackageName).Output()
	default:
		panic("unknown command type")
	}
}

func ExecuteString(command string) (string, error) {
	tokens :=strings.Split(command, " ")
	bytes, _:= exec.Command(tokens[0], tokens[1:]...).CombinedOutput()
	return string(bytes), nil
}
