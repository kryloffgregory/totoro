package api

import (
	"fmt"

	"github.com/kryloffgregory/totoro/server/git"
	"github.com/kryloffgregory/totoro/server/node"
)

type Listener struct {
	nodeManager *node.Manager
}

func NewListener() *Listener {
	return &Listener{nodeManager:node.NewManager()}
}

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

type InterestParams struct {
	Package string
	User string
}

type InterestReply struct {
	State string

}

func(l *Listener) Install(params *InstallParams, reply *InstallReply) error {
	fmt.Printf("Install request: %v", params)

	url:=git.CreatePR("vasya2048", "apt -y install --no-upgrade "+params.Package, []string{"kryloffgregory"})
	reply.State = url
	return nil
}

func (l *Listener) Remove(params *RemoveParams, reply *RemoveReply) error {
	fmt.Printf("Remove request: %v", params)

	url:=git.CreatePR("vasya2048", "apt -y remove --purge" + params.Package, []string{"kryloffgregory"})
	reply.State = url
	return nil
}

func (l *Listener) Interest(params *InterestParams, reply *InterestReply) error {
	fmt.Printf("Interest request: %v", params)

	l.nodeManager.AddAffected(params.Package, params.User)
	return nil
}

