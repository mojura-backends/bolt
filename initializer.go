package bolt

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/mojura/backend"
)

var defaultOpts = &bolt.Options{Timeout: 1 * time.Second}

// New will create a new initializer with default options
func New() backend.Initializer {
	return NewWithOpts(defaultOpts)
}

// NewWithOpts will create a new initializer with the provided opts
func NewWithOpts(opts *bolt.Options) backend.Initializer {
	var i Initializer
	i.opts = opts
	return &i
}

// Initializer initializes a boltDB backend
type Initializer struct {
	opts *bolt.Options
}

// New will return a new instance of Backend
func (i *Initializer) New(filename string) (b backend.Backend, err error) {
	var be Backend
	if be.db, err = bolt.Open(filename, 0644, i.opts); err != nil {
		return
	}

	b = &be
	return
}
