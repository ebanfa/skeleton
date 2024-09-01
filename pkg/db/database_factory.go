package db

import (
	"cosmossdk.io/log"
	"github.com/cosmos/iavl"
	"github.com/cosmos/iavl/db"
	"github.com/ebanfa/skeleton/pkg/types"
)

// DatabaseFactory is an interface for creating databases.
type DatabaseFactory interface {
	// CreateDatabase creates and initializes a database instance with the given name and path.
	CreateDatabase(name, path string) (types.Database, error)
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
func (f *IAVLDatabaseFactory) CreateDatabase(name, path string) (types.Database, error) {
	// Initialize the LevelDB instance
	ldb, err := db.NewGoLevelDB(name, path)
	if err != nil {
		return nil, err
	}

	// Initialize the IAVLDB instance
	iavlTree := iavl.NewMutableTree(db.NewPrefixDB(ldb, []byte("s/k:main/")), 100, false, log.NewNopLogger())
	iavlDB := NewIAVLDatabase(iavlTree)

	return iavlDB, nil
}
