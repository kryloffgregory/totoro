package main

import (
	"log"
	"net/rpc"
	"os"
	"os/user"

	"github.com/kryloffgregory/totoro/server/api"
)


func main() {
	client, err := rpc.Dial("unix", "/tmp/totoro")
	if err != nil {
		log.Fatal(err)
	}
	usr, err:=user.Current()
	if err!=nil {
		log.Fatal(err)
	}
	args:=os.Args[1:]
	switch args[0] {
	case "install":
		lib:=args[1]
		version:=""
		if len(args) >= 3{
			version = args[2]
		}
		req:=&api.InstallParams{
			Package: lib,
			Version: version,
			User:    usr.Uid,
		}
		repl:=&api.InstallReply{}
		err = client.Call("Listener.Install", req, repl)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v", repl)
	case "remove":
		lib:=args[1]
		req:=&api.RemoveParams{
			Package: lib,
			User:    usr.Uid,
		}
		repl:=&api.RemoveReply{}
		err = client.Call("Listener.Remove", req, repl)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v", repl)

	}

}