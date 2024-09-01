package store

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ebanfa/skeleton/pkg/types"
)

// StoreMetaData contains metadata for a store.
type StoreMetaData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// MultiStoreImpl is a concrete implementation of the MultiStore interface.
type MultiStoreImpl struct {
	types.Store                         // Embedding Store to satisfy the Store interface
	stores       map[string]types.Store // Map to store metadata of stores
	mutex        sync.RWMutex
	storeFactory StoreFactory
}

// NewMultiStore creates a new instance of MultiStoreImpl with the provided store options.
func NewMultiStore(store types.Store, storeFactory StoreFactory) (types.MultiStore, error) {
	// Return a new instance of MultiStoreImpl with the embedded Store instance,
	// along with other necessary fields initialized
	return &MultiStoreImpl{
		Store:        store,                        // Embed the Store instance to satisfy the Store interface
		stores:       make(map[string]types.Store), // Initialize the map to store metadata of stores
		storeFactory: storeFactory,
	}, nil
}

// GetStore returns the store with the given namespace.
// If the store doesn't exist, it returns an error.
func (ms *MultiStoreImpl) GetStore(namespace []byte) types.Store {
	// Lock the mutex to prevent concurrent access to the map
	ms.mutex.Lock()
	defer ms.mutex.Unlock() // Unlock the mutex when the function exits

	// Access the map using the namespace converted to a string as the key
	store, ok := ms.stores[string(namespace)]

	// Check if the store exists
	if !ok {
		// If the store does not exist, return an error
		return nil
	}

	// If the store exists, return it
	return store
}

// GetStoreCount returns the total number of stores in the multistore.
func (ms *MultiStoreImpl) GetStoreCount() int {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	return len(ms.stores)
}

// CreateStore creates and initializes a new store with the given namespace and options.
// If a store with the same namespace already exists, it returns an error.
func (ms *MultiStoreImpl) CreateStore(namespace string) (types.Store, bool, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	if !IsValidStoreName(namespace) {
		return nil, false, fmt.Errorf("invalid store name provided: %s", namespace)
	}

	ns := GenerateStoreId(namespace)

	store, exists := ms.stores[ns]
	if exists {
		return store, false, nil
	}

	// Create a new StoreImpl instance with the provided database
	store, err := ms.storeFactory.CreateStore(string(namespace))
	if err != nil {
		return nil, false, err
	}

	ms.stores[ns] = store

	return store, true, nil
}

// Load loads the latest versioned database from disk.
func (ms *MultiStoreImpl) Load() (int64, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	// Clear existing stores
	ms.stores = make(map[string]types.Store)

	// Load the database
	version, err := ms.Store.Load()
	if err != nil {
		return version, err
	}

	// Retrieve metadata for each store from the database
	err = ms.Store.Iterate(func(key, value []byte) bool {
		// Deserialize store metadata from value
		var meta StoreMetaData
		err := json.Unmarshal(value, &meta)
		if err != nil {
			// Handle error if metadata deserialization fails
			return false // Stop iteration
		}
		// Create and initialize store based on metadata
		store, err := ms.storeFactory.CreateStore(string(key))
		if err != nil {
			// Handle error if store creation fails
			return false // Stop iteration
		}
		// Add store to the stores map
		ms.stores[string(key)] = store
		return false // Continue iteration
	})

	if err != nil {
		return version, err
	}

	return version, nil
}

// SaveVersion saves a new version of the database to disk.
func (ms *MultiStoreImpl) SaveVersion() ([]byte, int64, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	// Serialize store metadata and store them in the database
	for id, store := range ms.stores {
		meta := StoreMetaData{
			Id:   id,
			Name: store.Name(), // Assuming String method returns the name of the store
			Path: store.Path(), // Assuming Path method returns the path of the store
		}
		// Serialize store metadata
		metaJSON, err := json.Marshal(meta)
		if err != nil {
			return nil, 0, err
		}
		// Store serialized metadata in the database
		err = ms.Store.Set([]byte(id), metaJSON)
		if err != nil {
			return nil, 0, err
		}
	}

	// Save the versioned database
	data, version, err := ms.Store.SaveVersion()
	if err != nil {
		return nil, version, err
	}

	return data, version, nil
}
