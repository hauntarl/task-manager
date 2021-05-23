package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var (
	db     *bolt.DB
	bucket = []byte("task")
)

type Task struct {
	key int
	val string
}

func Init(path string) error {
	var err error
	if db, err = bolt.Open(path, 0600,
		&bolt.Options{Timeout: 1 * time.Second}); err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	})
}
