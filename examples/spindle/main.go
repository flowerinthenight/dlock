package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/flowerinthenight/dlock"
)

func main() {
	// To run, update the database name, table name, and, optionally, the lock name.
	// It is assumed that your environment is able to authenticate to Spanner via
	// GOOGLE_APPLICATION_CREDENTIALS environment variable.
	db, err := spanner.NewClient(
		context.Background(),
		"projects/mobingi-main/instances/alphaus-prod/databases/main",
	)

	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	l := dlock.NewSpindleLock(&dlock.SpindleLockOptions{
		Client:   db,
		Table:    "testlease",
		Name:     "dlock",
		Duration: 1000,
		Logger:   log.New(ioutil.Discard, "", 0), // don't set if you want to see spindle logs
	})

	start := time.Now()
	l.Lock(context.Background())
	log.Printf("lock acquired after %v, do protected work...", time.Since(start))
	time.Sleep(time.Second * 5)
	l.Unlock()
}
