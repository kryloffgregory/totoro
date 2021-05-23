package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"

	"github.com/kryloffgregory/totoro/server/api"
	"github.com/kryloffgregory/totoro/server/git"
)

const socketAddr = "/tmp/totoro/sock"

func main() {
	if err:=os.Remove(socketAddr); err!=nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", socketAddr)
	if err != nil {
		log.Fatal(err)
	}

	err=os.Chmod(socketAddr, 0777)
	if err!=nil{
		log.Fatal(err)
	}

	err=rpc.Register(api.NewListener())
	if err!=nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err:=git.ProcessPRs(); err!=nil {
				log.Println(fmt.Sprintf("Error occured while processing prs: %v", err))
			}
			time.Sleep(time.Second*10)
		}
	}()

	rpc.Accept(l)

}
