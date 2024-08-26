package store

import (
	cosLogApi "cosmossdk.io/log"
	"github.com/cosmos/iavl"
	dbm "github.com/cosmos/iavl/db"
)

// ReadOnlyDatabase provides methods for reading data from the database.
type ReadOnlyDatabase interface {
	// Get retrieves the value associated with the given key from the database.
	Get(key []byte) ([]byte, error)

	// Has checks if a key exists in the database.
	Has(key []byte) (bool, error)

	// Iterate iterates over all key-value pairs in the database and calls the given function for each pair.
	// Iteration stops if the function returns true.
	Iterate(fn func(key, value []byte) bool) error

	// IterateRange iterates over key-value pairs with keys in the specified range
	// and calls the given function for each pair. Iteration stops if the function returns true.
	IterateRange(start, end []byte, ascending bool, fn func(key, value []byte) bool) error

	// Hash returns the hash of the database.
	Hash() []byte

	// Version returns the version of the database.
	Version() int64

	// String returns a string representation of the database.
	String() (string, error)

	// WorkingVersion returns the current working version of the database.
	WorkingVersion() int64

	// WorkingHash returns the hash of the current working version of the database.
	WorkingHash() []byte

	// AvailableVersions returns a list of available versions.
	AvailableVersions() []int

	// IsEmpty checks if the database is empty.
	IsEmpty() bool
}

// MutableDatabase provides methods for modifying the database.
type MutableDatabase interface {
	// Set stores the key-value pair in the database. If the key already exists, its value will be updated.
	Set(key, value []byte) error

	// Delete removes the key-value pair from the database.
	Delete(key []byte) error
}

// VersionedDatabase provides methods for managing versions of the database.
type VersionedDatabase interface {
	// Load loads the latest versioned database from disk.
	Load() (int64, error)

	// LoadVersion loads a specific version of the database from disk.
	LoadVersion(targetVersion int64) (int64, error)

	// SaveVersion saves a new version of the database to disk.
	SaveVersion() ([]byte, int64, error)

	// Rollback resets the working database to the latest saved version, discarding any unsaved modifications.
	Rollback()
}

// Database combines all the interfaces for a complete database interface.
type Database interface {
	ReadOnlyDatabase
	MutableDatabase
	VersionedDatabase

	// Close closes the database.
	Close() error
}

// DatabaseFactory is an interface for creating databases.
type DatabaseFactory interface {
	// CreateDatabase creates and initializes a database instance with the given name and path.
	CreateDatabase(name, path string) (Database, error)
}

// IAVLDatabaseFactory is a concrete implementation of the DatabaseFactory interface
// that creates IAVL database instances.
type IAVLDatabaseFactory struct {
	DatabaseFactory
}

// NewIAVLDatabaseFactory creates a new instance of IAVLDatabaseFactory with the given DB creator function.
func NewIAVLDatabaseFactory() *IAVLDatabaseFactory {
	return &IAVLDatabaseFactory{}
}

// CreateDatabase creates and initializes an IAVL database instance with the given name and path.
func (f *IAVLDatabaseFactory) CreateDatabase(name, path string) (Database, error) {
	// Initialize the LevelDB instance
	ldb, err := dbm.NewGoLevelDB(name, path)
	if err != nil {
		return nil, err
	}

	// Initialize the IAVLDB instance
	iavlTree := iavl.NewMutableTree(dbm.NewPrefixDB(ldb, []byte("s/k:main/")), 100, false, cosLogApi.NewNopLogger())
	iavlDB := NewIAVLDatabase(iavlTree)

	return iavlDB, nil
}
