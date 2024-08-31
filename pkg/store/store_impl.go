package store

import (
	"errors"
)

// StoreImpl is a concrete implementation of the Store interface.
type StoreImpl struct {
	Database        // Embedding Database to satisfy the Database interface
	name     string // Name of the store
	path     string // Path of the store
}

// NewStoreImpl creates a new instance of StoreImpl with the provided StoreOptions object.
func NewStoreImpl(name, path string, database Database) (*StoreImpl, error) {
	if database == nil {
		return nil, errors.New("cannot create Store from nil database")
	}
	return &StoreImpl{
		Database: database,
		name:     name,
		path:     path,
	}, nil
}

// Name returns the name of the store.
func (s *StoreImpl) Name() string {
	return s.name
}

// Path returns the path of the store.
func (s *StoreImpl) Path() string {
	return s.path
}
