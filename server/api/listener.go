package api

import (
	"fmt"
	"log"

	"github.com/kryloffgregory/totoro/server/config"
	"github.com/kryloffgregory/totoro/server/git"
	"github.com/kryloffgregory/totoro/server/node"
)

type Listener struct {
	nodeManager *node.Manager
	userMapping map[string]string
}

func NewListener() *Listener {
	mapping, err := config.GetUserMapping()
	if err != nil {
		panic(err)
	}
	return &Listener{
		nodeManager: node.NewManager(),
		userMapping: mapping.Mapping,
	}
}

type InstallReply struct {
	State string
}

type InstallParams struct {
	Package string
	Version string
	User    string
}

type RemoveReply struct {
	State string
}

type RemoveParams struct {
	Package string
	User    string
}

type InterestParams struct {
	Package string
	User    string
}

type InterestReply struct {
	State string
}

func (l *Listener) Install(params *InstallParams, reply *InstallReply) error {
	log.Printf("Install request: %v", params)

	affected, err := l.nodeManager.GetAffectedForNodeUpdate(params.Package)
	if err != nil {
		return err
	}


	url := git.CreatePR("vasya2048", "apt -y install --no-upgrade "+params.Package, affected)
	reply.State = url
	return nil
}

func (l *Listener) Remove(params *RemoveParams, reply *RemoveReply) error {
	log.Printf("Remove request: %v", params)

	affected, err := l.nodeManager.GetAffectedForNodeDelete(params.Package)
	if err != nil {
		return err
	}

	url := git.CreatePR("vasya2048", "apt -y remove --purge"+params.Package, affected)
	reply.State = url
	return nil
}

func (l *Listener) Interest(params *InterestParams, reply *InterestReply) error {
	fmt.Printf("Interest request: %v", params)

	return l.nodeManager.AddAffected(params.Package, params.User)
}
