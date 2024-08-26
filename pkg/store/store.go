package store

import (
	"errors"
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/asaskevich/govalidator"
	"github.com/ebanfa/skeleton/pkg/common"
)

// StoreOptions contains options for configuring a store.
type StoreOptions struct {
	DatabaseFactory DatabaseFactory `valid:"-"`
	Name            string          `valid:"required"`
	Path            string          `valid:"required"`
}

// NewStoreOptions creates a new instance of StoreOptions with the provided parameters.
func NewStoreOptions(databaseFactory DatabaseFactory, name, path string) *StoreOptions {
	return &StoreOptions{
		DatabaseFactory: databaseFactory,
		Name:            name,
		Path:            path,
	}
}

// Validate checks the validity of the StoreOptions struct.
func (so *StoreOptions) Validate() bool {
	// Using govalidator.ValidateStruct to perform validation
	result, _ := govalidator.ValidateStruct(so)
	return result
}

// Store represents a database store.
type Store interface {
	Database

	// Name returns the name of the store.
	Name() string

	// Path returns the path of the store.
	Path() string
}

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

func GenererateStorageInfo(name string, databasesDir string) (string, string) {
	databaseID := GenerateStoreId(name)

	databasePath := filepath.Join(databasesDir, databaseID+".db")
	return databaseID, databasePath
}

func GenerateStoreId(name string) string {
	return common.HashSHA256(name)
}

func CreateMultiStore(name, databasesDir string, storeFactory StoreFactory) (MultiStore, error) {

	internalStore, err := storeFactory.CreateStore(name)

	if err != nil {
		return nil, err
	}

	multiStore, err := NewMultiStore(internalStore, storeFactory)
	if err != nil {
		return nil, err
	}

	return multiStore, nil
}

func IsValidStoreName(s string) bool {
	// Check if the string is empty
	if len(s) == 0 {
		return false
	}

	// Check if the string consists only of whitespace or punctuation characters
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			// If the character is not a letter or number, return false
			return false
		}
	}

	// If the string passes all checks, return true
	return true
}
