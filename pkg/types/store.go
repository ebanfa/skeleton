package types

// MultiStore is a multi-store interface that manages multiple key-value stores.
type MultiStore interface {
	Store

	// GetStoreCount returns the total number of stores in the multistore.
	GetStoreCount() int

	// GetStore returns the store with the given namespace.
	GetStore(namespace []byte) Store

	// Creates and adds a new store with the given namespace.
	// If a store with the same namespace already exists, it returns an error.
	CreateStore(namespace string) (Store, bool, error)
}

// Store represents a database store.
type Store interface {
	Database

	// Name returns the name of the store.
	Name() string

	// Path returns the path of the store.
	Path() string
}
