package component

import (
	"fmt"
	"sync"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/types"
)

// ComponentRegistrar defines the registry functionality for components and factories.
type ComponentRegistrar struct {
	types.ComponentRegistrarInterface
	factoriesMutex  sync.RWMutex
	componentsMutex sync.RWMutex
	factories       map[string]types.ComponentFactoryInterface
	components      map[string]types.ComponentInterface
}

// NewComponentRegistrar creates a new instance of ComponentRegistrar.
func NewComponentRegistrar() *ComponentRegistrar {
	return &ComponentRegistrar{
		factories:  make(map[string]types.ComponentFactoryInterface),
		components: make(map[string]types.ComponentInterface),
	}
}

// GetComponent retrieves the component with the specified ID.
func (cr *ComponentRegistrar) GetComponent(id string) (types.ComponentInterface, error) {
	cr.componentsMutex.RLock()
	defer cr.componentsMutex.RUnlock()

	// Check if the component exists
	component, exists := cr.components[id]
	if !exists {
		return nil, fmt.Errorf("component with ID %s not found", id)
	}
	return component, nil
}

// GetComponentsByType retrieves components of the specified type.
func (cr *ComponentRegistrar) GetComponentsByType(componentType types.ComponentType) []types.ComponentInterface {
	// Lock the components mutex for reading to prevent concurrent access while reading
	cr.componentsMutex.RLock()
	defer cr.componentsMutex.RUnlock()

	// Initialize an empty slice to store components of the specified type
	components := []types.ComponentInterface{}

	// Iterate over all registered components
	for _, component := range cr.components {
		// Check if the type of the component matches the specified type
		if component.Type() == componentType {
			// If the type matches, add the component to the slice
			components = append(components, component)
		}
	}

	// Return the slice of components
	return components
}

// GetAllComponents returns a list of all registered components.
func (cr *ComponentRegistrar) GetAllComponents() []types.ComponentInterface {
	// Lock the components mutex for reading to prevent concurrent access while reading
	cr.componentsMutex.RLock()
	defer cr.componentsMutex.RUnlock()

	// Initialize an empty slice to store all registered components
	allComponents := make([]types.ComponentInterface, 0, len(cr.components))

	// Iterate over all registered components
	for _, component := range cr.components {
		// Add each component to the slice
		allComponents = append(allComponents, component)
	}

	// Return the slice of all components
	return allComponents
}

// GetFactory retrieves the factory with the specified ID.
func (cr *ComponentRegistrar) GetFactory(id string) (types.ComponentFactoryInterface, error) {
	// Lock the factories mutex for reading to prevent concurrent access while reading
	cr.factoriesMutex.RLock()
	defer cr.factoriesMutex.RUnlock()

	// Check if the factory exists
	factory, exists := cr.factories[id]
	if !exists {
		return nil, fmt.Errorf("factory with ID %s not found", id)
	}
	return factory, nil
}

// GetAllFactories returns a list of all registered component factories.
func (cr *ComponentRegistrar) GetAllFactories() []types.ComponentFactoryInterface {
	// Lock the factories mutex for reading to prevent concurrent access while reading
	cr.factoriesMutex.RLock()
	defer cr.factoriesMutex.RUnlock()

	// Initialize an empty slice to store all registered factories
	allFactories := make([]types.ComponentFactoryInterface, 0, len(cr.factories))

	// Iterate over all registered factories
	for _, factory := range cr.factories {
		// Add each factory to the slice
		allFactories = append(allFactories, factory)
	}

	// Return the slice of all factories
	return allFactories
}

// CreateComponent creates and registers a new instance of the component.
func (cr *ComponentRegistrar) CreateComponent(ctx *common.Context, config *types.ComponentConfig) (types.ComponentInterface, error) {
	cr.factoriesMutex.RLock()
	defer cr.factoriesMutex.RUnlock()

	// Check if the factory exists
	factory, exists := cr.factories[config.FactoryID]
	if !exists {
		return nil, fmt.Errorf("factory with ID %s not found", config.FactoryID)
	}

	// Use the factory to create the component
	component, err := factory.CreateComponent(config)
	if err != nil {
		return nil, err
	}

	// Register the component
	cr.componentsMutex.Lock()
	defer cr.componentsMutex.Unlock()
	cr.components[component.ID()] = component

	return component, nil
}

// RegisterFactory registers a factory with the given ID.
func (cr *ComponentRegistrar) RegisterFactory(ctx *common.Context, id string, factory types.ComponentFactoryInterface) error {
	cr.factoriesMutex.Lock()
	defer cr.factoriesMutex.Unlock()

	// Check if the factory already exists
	if _, exists := cr.factories[id]; exists {
		return fmt.Errorf("factory with ID %s already exists", id)
	}

	// Register the factory
	cr.factories[id] = factory
	return nil
}

// UnregisterComponent unregisters the component with the specified ID.
func (cr *ComponentRegistrar) UnregisterComponent(ctx *common.Context, id string) error {
	cr.componentsMutex.Lock()
	defer cr.componentsMutex.Unlock()

	// Check if the component exists
	if _, exists := cr.components[id]; !exists {
		return fmt.Errorf("component with ID %s not found", id)
	}

	// Unregister the component
	delete(cr.components, id)
	return nil
}

// UnregisterFactory unregisters the factory with the specified ID.
func (cr *ComponentRegistrar) UnregisterFactory(ctx *common.Context, id string) error {
	cr.factoriesMutex.Lock()
	defer cr.factoriesMutex.Unlock()

	// Check if the factory exists
	if _, exists := cr.factories[id]; !exists {
		return fmt.Errorf("factory with ID %s not found", id)
	}

	// Unregister the factory
	delete(cr.factories, id)

	// Remove components created from this factory
	cr.componentsMutex.Lock()
	defer cr.componentsMutex.Unlock()
	for key := range cr.components {
		if key == id {
			delete(cr.components, key)
		}
	}
	return nil
}
