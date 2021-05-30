package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db    *bolt.DB         // database pointer?
	Btodo = []byte("todo") // where all the new tasks go
	Bdone = []byte("done") // where all completed tasks go
)

// Task is a Go representation of key, value data for given todo
type Task struct {
	Key uint64
	Val string
}

// Init initializes the database and creates a bucket to store tasks
func Init(path string) error {
	var err error
	if db, err = bolt.Open(
		path, 0600,
		// with boltdb comes a restriction where simultaneous access to
		// database is not allowed, which puts the later instance into a
		// deadlock until the database is freed up, this might take an
		// indefinite amount of time and instead of making the user wait the
		// whole way through, we provide an option stating that, wait 1 whole
		// second before giving up.
		&bolt.Options{Timeout: 1 * time.Second},
	); err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) (err error) {
		_, err = tx.CreateBucketIfNotExists(Btodo)
		if err != nil {
			return
		}
		_, err = tx.CreateBucketIfNotExists(Bdone)
		return
	})
}

// AddTask creates a new unique id and inserts the task into database
func AddTask(name string) (key uint64, err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Btodo)
		key, _ = b.NextSequence()
		// another restriction with boltdb is that, it will only accepts values
		// in the form of slice of bytes, hence we need to encode/decode key.
		return b.Put(itob(key), []byte(name))
	})
	return
}

// ReadTasks reads all the tasks from database and generated Go representation
// for the same
func ReadTasks(buc []byte) (tasks []Task, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(buc)
		c := b.Cursor() // helps you iterate through each element in database
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{Key: btoi(k), Val: string(v)})
		}
		return nil
	})
	return
}

// DeleteTask removes a task from database using given key
func DeleteTask(key uint64) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Btodo)
		return b.Delete(itob(key))
	})
	return
}

func CompleteTask(t Task) (err error) {
	err = DeleteTask(t.Key)
	if err != nil {
		return
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bdone)
		return b.Put(itob(t.Key), []byte(t.Val))
	})
	return
}

// converts int to byte stream
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// decodes int value from given byte stream
func btoi(b []byte) uint64 { return binary.BigEndian.Uint64(b) }
