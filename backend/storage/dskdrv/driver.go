// Package dskdrv implements the storage driver using local disks as the storage medium
package dskdrv

import (
	"github.com/peterbourgon/diskv"
	"fmt"
	"strings"
	"github.com/andrebq/the-core/backend/storage"
)

type (
	// Driver implements the driver interface
	Driver struct {
		db *diskv.Diskv
	}
)

// New returns a new driver using the given folder as base
func New(base string) *Driver {
	drv := &Driver{}
	drv.db = diskv.New(diskv.Options{
		BasePath: base,
		Transform: transformFn,
		Compression: diskv.NewZlibCompressionLevel(9),
	})
	return drv
}

// Put files
func (d *Driver) Put(id storage.ContentID, content storage.Content) error {
	return d.db.Write(idToKey(id), content)
}

// Get files
func (d *Driver) Get(id storage.ContentID) (storage.Content, error) {
	value, err := d.db.Read(idToKey(id))
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, storage.ContentNotFoundErr
	}
	return storage.Content(value), nil
}

// Exists content
func (d *Driver) Exists(id storage.ContentID) (bool, error) {
	return d.db.Has(idToKey(id)), nil
}

func transformFn(s string) []string {
	return strings.Split(s, "_")[:2]
}

func idToKey(id storage.ContentID) string {
	return fmt.Sprintf("%s_%s_%s",
		strings.ToLower(id.Alg),
		strings.ToLower(id.ID[:2]),
		strings.ToLower(id.ID[2:]))
}
