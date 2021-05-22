package api

import (
	"fmt"

	"accessModel/server/execute"
	"accessModel/server/issue"
)

type Listener int
type InstallReply struct {
	State string
}


type InstallParams struct {
	Package string
	Version string
	User string
}

type RemoveReply struct {
	State string
}

type RemoveParams struct {
	Package string
	User string
}


func(l *Listener) Install(params *InstallParams, reply *InstallReply) error {
	fmt.Printf("Install request: %v", params)
	output, _ :=execute.Execute(issue.Command{
		Type:           issue.CommandTypeInstall,
		InstallPayload: &issue.InstallPayload{
			PackageName: params.Package,
			Version:     params.Version,
		},
	})
	reply.State = string(output)
	return nil
}

func (l *Listener) Remove(params *RemoveParams, reply *RemoveReply) error {
	fmt.Printf("Remove request: %v", params)
	execute.Execute(issue.Command{
		Type:           issue.CommandTypeRemove,
		RemovePayload:&issue.RemovePayload{
			PackageName:params.Package,
		},
	})
	reply = &RemoveReply{State:"34"}
	return nil
}

