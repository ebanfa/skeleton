package store

import "fmt"

type StoreFactory interface {
	// CreateStoreInternal creates a new store.
	CreateStore(name string) (Store, error)
}

type StoreFactoryImpl struct {
	databasesDir string
	dbFactory    DatabaseFactory
}

func NewStoreFactory(databasesDir string, dbFactory DatabaseFactory) StoreFactory {
	return &StoreFactoryImpl{
		dbFactory:    dbFactory,
		databasesDir: databasesDir,
	}
}

func (f StoreFactoryImpl) CreateStore(name string) (Store, error) {
	fmt.Printf("Creating store name:%s databasesDir:%s\n", name, f.databasesDir)
	// Generate storage path and Id
	// Define the database path within the .nova directory
	databaseID, databasePath := GenererateStorageInfo(name, f.databasesDir)
	fmt.Printf("GenererateStorageInfo: databasePath:%s databaseID:%s databasesDir:%s\n", databasePath, databaseID, f.databasesDir)

	// Create the store using the internal function
	return f.createStoreInternal(databaseID, databasePath)
}

// CreateStoreInternal creates a new store with the given database ID and path using the provided database factory.
// It creates the database at the specified path and returns a store initialized with the database.
func (f StoreFactoryImpl) createStoreInternal(name, databasePath string) (Store, error) {
	// Create the database using the factory
	database, err := f.dbFactory.CreateDatabase(name, databasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	return NewStoreImpl(name, databasePath, database)
}
