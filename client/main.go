package main

import (
	"log"
	"net/rpc"
	"os/user"

	"accessModel/server/api"

	"github.com/davecgh/go-spew/spew"
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
	spew.Dump(usr)

		req:=&api.InstallParams{
			Package: "cowsay",
			Version: "",
			User:    usr.Uid,
		}
		repl:=&api.InstallReply{}
		err = client.Call("Listener.Install", req, repl)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v", repl)
}