package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db     *bolt.DB
	bucket = []byte("task")
)

type Task struct {
	key uint64
	val string
}

func Init(path string) error {
	var err error
	if db, err = bolt.Open(
		path, 0600,
		&bolt.Options{Timeout: 1 * time.Second},
	); err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	})
}

func AddTask(name string) (id uint64, err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		buc := tx.Bucket(bucket)
		id, _ = buc.NextSequence()
		return buc.Put(itob(id), []byte(name))
	})
	return
}

func ReadTasks() (tasks []Task, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		buc := tx.Bucket(bucket)
		cur := buc.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			tasks = append(tasks, Task{key: btoi(k), val: string(v)})
		}
		return nil
	})
	return
}

func DeleteTask(key uint64) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		buc := tx.Bucket(bucket)
		return buc.Delete(itob(key))
	})
	return
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func btoi(b []byte) uint64 { return binary.BigEndian.Uint64(b) }
