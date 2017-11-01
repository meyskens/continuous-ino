package storage

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func openBucket(tx *bolt.Tx, name string) *bolt.Bucket {
	root, err := tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		log.Fatal("could not create root bucket:", err)
	}

	return root
}

func (s *Storage) getNextIDIn(bucket string) (id uint64) {
	s.db.Update(func(tx *bolt.Tx) error {
		bucket := openBucket(tx, bucket)
		id, _ = bucket.NextSequence()
		return nil
	})
	return
}

func (s *Storage) saveData(bucket string, id uint64, data []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := openBucket(tx, bucket)

		// Persist bytes to users bucket.
		return bucket.Put([]byte(fmt.Sprintf("%v", id)), data)
	})
}

func (s *Storage) getData(bucket string, id uint64) (data []byte) {
	s.db.Update(func(tx *bolt.Tx) error { // we do create if not exists on a bucket here!
		bucket := openBucket(tx, bucket)
		// Persist bytes to users bucket.
		data = bucket.Get([]byte(fmt.Sprintf("%v", id)))
		return nil
	})
	return
}
func (s *Storage) getAll(bucket string) (data [][]byte) {
	data = [][]byte{}

	s.db.Update(func(tx *bolt.Tx) error { // we do create if not exists on a bucket here!
		bucket := openBucket(tx, bucket)
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			data = append(data, v)
		}
		return nil
	})
	return
}
