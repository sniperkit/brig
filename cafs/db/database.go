package db

import (
	"errors"
	"io"
)

// TODO: Implement an actual fast KV store based on moss, boltdb or badger
//       if there is any performance problem later on.
//       For now, the filesystem based kv should suffice fine though.

var (
	// ErrNoSuchKey is returned when Get() was passed a non-existant key
	ErrNoSuchKey = errors.New("This key does not exist")
)

type Batch interface {
	// Put sets `val` at `key`.
	Put(val []byte, key ...string)

	// Clear all contents below and including `key`.
	Clear(key ...string)

	// Erase a key from the database.
	Erase(key ...string)

	// Flush the batch to the database.
	// Only now, all changes will be written to disk.
	Flush() error

	// Rollback will forget all changes without executing them.
	Rollback()
}

// Database is a key/value store that offers different buckets
// for storage. Keys are strings, values are arbitary untyped data.
// TODO: Write down assumptions made (single user database e.g.)
type Database interface {
	// Get retrievies the key `key` out of bucket.
	// If no such key exists, it will return (nil, ErrNoSuchKey)
	// If a badge is currently open, Get() shall still return the
	// most current value currently set by the last Put() call
	// to `key`.
	Get(key ...string) ([]byte, error)

	// Keys returns all keys
	Keys(fn func(key []string) error, prefix ...string) error

	// Batch returns a new Batch object, that will allow modifications
	// of the state. Batch() can be called recursive: The changes will
	// only be flushed to disk if batch.Flush() was called equal times
	// to the number Batch() was called.
	Batch() Batch

	// Export backups all database content to `w` in
	// an implemenation specific format that can be read by Import.
	Export(w io.Writer) error

	// Import reads a previously exported db dump by Export from `r`.
	// Existing keys might be overwritten if the dump also contains them.
	Import(r io.Reader) error

	// Close closes the database. Since I/O may happen, an error is returned.
	Close() error
}

func CopyBucket(db Database, from, to []string) error {
	batch := db.Batch()

	walker := func(key []string) error {
		base := key[len(key)-1]

		data, err := db.Get(key...)
		if err != nil {
			return err
		}

		newPath := make([]string, len(to)+1)
		copy(newPath, to)
		newPath[len(newPath)-1] = base
		batch.Put(data, newPath...)
		return nil
	}

	if err := db.Keys(walker, from...); err != nil {
		batch.Rollback()
		return err
	}

	return batch.Flush()
}
