package issue

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

var db *bolt.DB
var issuesBucket = "issues"

func init() {
	var err error
	db, err = bolt.Open("./db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err!=nil {
		log.Fatal(err)
	}

	if err:=db.Update(func(tx *bolt.Tx) error {
		_, err:=tx.CreateBucketIfNotExists([]byte(issuesBucket))
		return err
	}); err!=nil {
		log.Fatal(err)
	}
}

func UpsertIssue(issue *Issue) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(issuesBucket))
		jsonIssue,err:=json.Marshal(issue)
		if err!=nil {
			return err
		}
		return b.Put([]byte(issue.ID.String()), jsonIssue)
	})
}

func GetIssue(issueID uuid.UUID) (*Issue, error) {
	issue:= &Issue{}
	if err:=db.View(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte(issuesBucket))
		v := b.Get([]byte(issueID.String()))
		return json.Unmarshal(v, issue)
	}); err!=nil {
		return nil, err
	}

	return issue, nil
}

func DeleteIssue(issueID string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte(issuesBucket))
		return b.Delete([]byte(issueID))
	})
}