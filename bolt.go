package bolt

import (
	"github.com/gdbu/bolt"
	"github.com/mojura/backend"
)

var (
	_ backend.Initializer = (*Initializer)(nil)
	_ backend.Backend     = (*Backend)(nil)
	_ backend.Transaction = (*Transaction)(nil)
	_ backend.Bucket      = (*Bucket)(nil)
	_ backend.Cursor      = (*bolt.Cursor)(nil)
)
