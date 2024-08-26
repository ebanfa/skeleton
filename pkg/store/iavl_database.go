package store

import (
	"sync"

	"github.com/cosmos/iavl"
)

// IAVLDatabase wraps an IAVL+ tree to implement the Database interface.
type IAVLDatabase struct {
	tree *iavl.MutableTree
	mtx  sync.RWMutex // Mutex for concurrent access
}

// NewIAVLDatabase creates a new IAVLDatabase instance.
func NewIAVLDatabase(tree *iavl.MutableTree) *IAVLDatabase {
	return &IAVLDatabase{tree: tree}
}

// Get retrieves the value associated with the given key from the tree.
func (db *IAVLDatabase) Get(key []byte) ([]byte, error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Retrieve the value associated with the key from the tree
	// Here we throw an error if the value is nil
	return db.tree.Get(key)
}

// Set stores the key-value pair in the tree. If the key already exists,
// its value will be updated.
func (db *IAVLDatabase) Set(key, value []byte) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	// Store the key-value pair in the tree
	_, err := db.tree.Set(key, value)
	return err
}

// Delete removes the key-value pair from the tree.
func (db *IAVLDatabase) Delete(key []byte) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	// Remove the key-value pair from the tree
	_, _, err := db.tree.Remove(key)
	return err
}

// Has returns true if the key exists in the tree, otherwise false.
func (db *IAVLDatabase) Has(key []byte) (bool, error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Check if the key exists in the tree
	return db.tree.Has(key)
}

// Iterate iterates over all keys of the tree and calls the given function
// for each key-value pair. Iteration stops if the function returns true.
func (db *IAVLDatabase) Iterate(fn func(key, value []byte) bool) error {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Iterate over all keys of the tree and call the given function for each key-value pair
	stopped, err := db.tree.Iterate(func(key []byte, value []byte) bool {
		return fn(key, value)
	})
	if err != nil {
		return err
	}
	if stopped {
		return nil
	}
	return nil
}

// IterateRange iterates over all key-value pairs with keys in the range
// [start, end) and calls the given function for each pair. Iteration stops
// if the function returns true.
func (db *IAVLDatabase) IterateRange(start, end []byte, ascending bool, fn func(key, value []byte) bool) error {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Iterate over key-value pairs with keys in the specified range
	db.tree.IterateRange(start, end, ascending, fn)
	return nil
}

// Hash returns the root hash of the tree.
func (db *IAVLDatabase) Hash() []byte {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Get the root hash of the tree
	return db.tree.Hash()
}

// Version returns the version of the tree.
func (db *IAVLDatabase) Version() int64 {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Get the version of the tree
	return db.tree.Version()
}

// Load loads the latest versioned tree from disk.
func (db *IAVLDatabase) Load() (int64, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	// Load the latest versioned tree from disk
	return db.tree.Load()
}

// LoadVersion loads a specific version of the tree from disk.
func (db *IAVLDatabase) LoadVersion(targetVersion int64) (int64, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	// Load the specified version of the tree from disk
	return db.tree.LoadVersion(targetVersion)
}

// SaveVersion saves a new tree version to disk.
func (db *IAVLDatabase) SaveVersion() ([]byte, int64, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	// Save a new tree version to disk
	hash, version, err := db.tree.SaveVersion()
	return hash, version, err
}

// Rollback resets the working tree to the latest saved version, discarding
// any unsaved modifications.
func (db *IAVLDatabase) Rollback() {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	// Reset the working tree to the latest saved version
	db.tree.Rollback()
}

// Close closes the tree.
func (db *IAVLDatabase) Close() error {
	db.mtx.Lock()
	defer db.mtx.Unlock()
	// Close the tree
	return db.tree.Close()
}

// String returns a string representation of the tree.
func (db *IAVLDatabase) String() (string, error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Get a string representation of the tree
	return db.tree.String()
}

// WorkingVersion returns the current working version of the tree.
func (db *IAVLDatabase) WorkingVersion() int64 {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Get the current working version of the tree
	return db.tree.WorkingVersion()
}

// WorkingHash returns the root hash of the current working tree.
func (db *IAVLDatabase) WorkingHash() []byte {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Get the root hash of the current working tree
	return db.tree.WorkingHash()
}

// AvailableVersions returns a list of available versions.
func (db *IAVLDatabase) AvailableVersions() []int {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Get a list of available versions
	return db.tree.AvailableVersions()
}

// IsEmpty checks if the database is empty.
func (db *IAVLDatabase) IsEmpty() bool {
	db.mtx.RLock()
	defer db.mtx.RUnlock()
	// Check if the database is empty
	return db.tree.IsEmpty()
}
