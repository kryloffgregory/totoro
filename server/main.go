package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kryloffgregory/totoro/server/api"
	"github.com/kryloffgregory/totoro/server/git"
)

const socketAddr = "/tmp/totoro/sock"

func main() {
	l, err := net.Listen("unix", socketAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	err = os.Chmod(socketAddr, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = rpc.Register(api.NewListener())
	if err != nil {
		log.Fatal(err)
	}

	shutdownStart := make(chan bool, 1)
	shutdownEnd := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-shutdownStart:
				shutdownEnd <- true
			case <-time.After(time.Second * 10):
				log.Println("Processing PRs")
				if err := git.ProcessPRs(shutdownStart, shutdownEnd); err != nil {
					log.Println(fmt.Sprintf("Error occured while processing prs: %v", err))
				}
			}
		}
	}()

	go rpc.Accept(l)

	sigs := make(chan os.Signal, 1)
	signalCaught := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println(fmt.Sprintf("Signal caught: %v", sig))
		signalCaught <- true
	}()

	log.Println("Server started")
	<-signalCaught
	log.Println("Performing graceful shutdown")

	shutdownStart <- true
	select {
	case <-shutdownEnd:
		break
	case <-time.After(time.Second * 5):
		log.Fatal("Graceful shutdown timeout")
	}
}
