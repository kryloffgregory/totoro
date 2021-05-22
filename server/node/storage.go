package node

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var bucket = "nodes"

func init() {
	var err error
	db, err = bolt.Open("./db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err!=nil {
		log.Fatal(err)
	}

	if err:=db.Update(func(tx *bolt.Tx) error {
		_, err:=tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	}); err!=nil {
		log.Fatal(err)
	}
}

func UpsertNode(node *Node) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		jsonNode,err:=json.Marshal(node)
		if err!=nil {
			return err
		}
		return b.Put([]byte(node.LibName), jsonNode)
	})
}

func GetNode(pack string) (*Node, error) {
	node:= &Node{}
	if err:=db.View(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte(bucket))
		v := b.Get([]byte(pack))
		return json.Unmarshal(v, node)
	}); err!=nil {
		return nil, err
	}

	return node, nil
}

func DeleteNode(pack string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte(bucket))
		return b.Delete([]byte(pack))
	})
}