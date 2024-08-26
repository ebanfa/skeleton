package system

import (
	"fmt"
	"sync"

	"github.com/ebanfa/skeleton/pkg/common"
)

// ComponentCreationInfo encapsulates the necessary information to create a component.
type ComponentCreationInfo struct {
	Config  *ComponentConfig
	Factory ComponentFactoryInterface
}

// ComponentRegistrarInterface defines the registry functionality for components and factories.
type ComponentRegistrarInterface interface {
	// GetComponentsByType retrieves components of the specified type.
	// It returns a list of components and an error if the type is not found or other error.
	GetComponentsByType(componentType ComponentType) []ComponentInterface

	// GetComponent retrieves the component with the specified ID.
	// It returns the component and an error if the component ID is not found or other error.
	GetComponent(id string) (ComponentInterface, error)

	// GetAllComponents returns a list of all registered components.
	GetAllComponents() []ComponentInterface

	// GetFactory retrieves the factory with the specified ID.
	// It returns the factory and an error if the factory ID is not found or other error.
	GetFactory(id string) (ComponentFactoryInterface, error)

	// RegisterFactory registers a factory with the given ID.
	// It returns an error if the registration fails.
	RegisterFactory(ctx *common.Context, id string, factory ComponentFactoryInterface) error

	// UnregisterFactory unregisters a factory with the specified ID.
	// It returns an error if the ID is not found or other error.
	UnregisterFactory(ctx *common.Context, id string) error

	// CreateComponent creates a component with the given ID and factory ID.
	// It returns the created component and an error if the creation fails.
	CreateComponent(ctx *common.Context, config *ComponentConfig) (ComponentInterface, error)

	// RemoveComponent removes a component with the specified ID from the registry.
	// It returns an error if the ID is not found or other error.
	RemoveComponent(ctx *common.Context, id string) error
}

// ComponentRegistrar defines the registry functionality for components and factories.
type ComponentRegistrar struct {
	ComponentRegistrarInterface
	factoriesMutex  sync.RWMutex
	componentsMutex sync.RWMutex
	factories       map[string]ComponentFactoryInterface
	components      map[string]ComponentInterface
}

// NewComponentRegistrar creates a new instance of ComponentRegistrar.
func NewComponentRegistrar() *ComponentRegistrar {
	return &ComponentRegistrar{
		factories:  make(map[string]ComponentFactoryInterface),
		components: make(map[string]ComponentInterface),
	}
}

// GetComponent retrieves the component with the specified ID.
func (cr *ComponentRegistrar) GetComponent(id string) (ComponentInterface, error) {
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
func (cr *ComponentRegistrar) GetComponentsByType(componentType ComponentType) []ComponentInterface {
	// Lock the components mutex for reading to prevent concurrent access while reading
	cr.componentsMutex.RLock()
	defer cr.componentsMutex.RUnlock()

	// Initialize an empty slice to store components of the specified type
	components := []ComponentInterface{}

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
func (cr *ComponentRegistrar) GetAllComponents() []ComponentInterface {
	// Lock the components mutex for reading to prevent concurrent access while reading
	cr.componentsMutex.RLock()
	defer cr.componentsMutex.RUnlock()

	// Initialize an empty slice to store all registered components
	allComponents := make([]ComponentInterface, 0, len(cr.components))

	// Iterate over all registered components
	for _, component := range cr.components {
		// Add each component to the slice
		allComponents = append(allComponents, component)
	}

	// Return the slice of all components
	return allComponents
}

// GetFactory retrieves the factory with the specified ID.
func (cr *ComponentRegistrar) GetFactory(id string) (ComponentFactoryInterface, error) {
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
func (cr *ComponentRegistrar) GetAllFactories() []ComponentFactoryInterface {
	// Lock the factories mutex for reading to prevent concurrent access while reading
	cr.factoriesMutex.RLock()
	defer cr.factoriesMutex.RUnlock()

	// Initialize an empty slice to store all registered factories
	allFactories := make([]ComponentFactoryInterface, 0, len(cr.factories))

	// Iterate over all registered factories
	for _, factory := range cr.factories {
		// Add each factory to the slice
		allFactories = append(allFactories, factory)
	}

	// Return the slice of all factories
	return allFactories
}

// CreateComponent creates and registers a new instance of the component.
func (cr *ComponentRegistrar) CreateComponent(ctx *common.Context, config *ComponentConfig) (ComponentInterface, error) {
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
func (cr *ComponentRegistrar) RegisterFactory(ctx *common.Context, id string, factory ComponentFactoryInterface) error {
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
