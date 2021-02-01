package bolt

import (
	"github.com/gdbu/bolt"
	"github.com/mojura/backend"
)

// Bucket represents a boltDB bucket
type Bucket struct {
	b *bolt.Bucket
}

// Get will get a value
func (b *Bucket) Get(key []byte) (value []byte) {
	return b.b.Get(key)
}

// Put will put a value
func (b *Bucket) Put(key, value []byte) error {
	return b.b.Put(key, value)
}

// Delete will delete a value
func (b *Bucket) Delete(key []byte) error {
	return b.b.Delete(key)
}

// Cursor will return a new cursor
func (b *Bucket) Cursor() backend.Cursor {
	return b.b.Cursor()
}

// GetBucket will get a bucket
func (b *Bucket) GetBucket(key []byte) (bkt backend.Bucket) {
	bucket := b.b.Bucket(key)
	if bucket == nil {
		return
	}

	return &Bucket{b: bucket}
}

// GetOrCreateBucket will get or create a bucket
func (b *Bucket) GetOrCreateBucket(key []byte) (bkt backend.Bucket, err error) {
	var bucket *bolt.Bucket
	if bucket, err = b.b.CreateBucketIfNotExists(key); err != nil {
		return
	}

	bkt = &Bucket{b: bucket}
	return
}

// DeleteBucket will get a bucket
func (b *Bucket) DeleteBucket(key []byte) (err error) {
	return b.b.DeleteBucket(key)
}

// ForEach will iterate through all the entries within a bucket
func (b *Bucket) ForEach(fn func(key, value []byte) error) (err error) {
	return b.b.ForEach(fn)
}
