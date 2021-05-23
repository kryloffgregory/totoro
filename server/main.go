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

func main() {
	l, err := net.Listen("unix", "/tmp/totoro")
	if err != nil {
		log.Fatal(err)
	}

	err=os.Chmod("/tmp/totoro", 0777)
	if err!=nil{
		log.Fatal(err)
	}

	err=rpc.Register(new(api.Listener))
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
