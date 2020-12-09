package bolt

import (
	"github.com/gdbu/bolt"
	"github.com/mojura/backend"
)

// Backend is the interface wrapper for bolt DB to be used as a Backend
type Backend struct {
	db *bolt.DB
}

// Transaction will initialize a new read-write transaction
func (b *Backend) Transaction(fn func(backend.Transaction) error) (err error) {
	return b.db.Update(func(txn *bolt.Tx) (err error) {
		t := Transaction{txn: txn}
		return fn(&t)
	})
}

// ReadTransaction will initialize a new read-only transaction
func (b *Backend) ReadTransaction(fn func(backend.Transaction) error) (err error) {
	return b.db.View(func(txn *bolt.Tx) (err error) {
		t := Transaction{txn: txn}
		return fn(&t)
	})
}

// Close will close the underlying DB
func (b *Backend) Close() error {
	return b.db.Close()
}
