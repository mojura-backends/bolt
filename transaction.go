package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/mojura/backend"
)

// Transaction represents a boltDB transaction
type Transaction struct {
	txn *bolt.Tx
}

// GetBucket will get a bucket
func (t *Transaction) GetBucket(key []byte) (bkt backend.Bucket) {
	bucket := t.txn.Bucket(key)
	if bucket == nil {
		return
	}

	return &Bucket{b: bucket}
}

// GetOrCreateBucket will get or create a bucket
func (t *Transaction) GetOrCreateBucket(key []byte) (bkt backend.Bucket, err error) {
	var bucket *bolt.Bucket
	if bucket, err = t.txn.CreateBucketIfNotExists(key); err != nil {
		return
	}

	bkt = &Bucket{b: bucket}
	return
}
