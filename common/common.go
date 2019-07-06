package common

import (
	"os"
	"path/filepath"
)

// ShardCount is amount of shards a file should be split into.
const ShardCount int = 10

// MaxFileSize is the maximum size a file on the network can be (in bytes).
const MaxFileSize int = 1000

// CreateDirIfDoesNotExist creates a directory if it does not already exist.
func CreateDirIfDoesNotExist(dir string) error {
	dir = filepath.FromSlash(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
