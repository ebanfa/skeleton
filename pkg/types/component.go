package types

import (
	"github.com/ebanfa/skeleton/pkg/common"
)

// ComponentType represents the type of a component.
type ComponentType int

const (
	// BasicComponentType represents the type of a basic component.
	BasicComponentType ComponentType = iota

	// SystemComponentType represents the type of a system component.
	SystemComponentType

	// OperationType represents the type of an operation component.
	OperationType

	// ServiceType represents the type of a service component.
	ServiceType

	// ModuleType represents the type of a module component.
	ApplicationComponentType
)

// ComponentInterface represents a generic component in the system.
type ComponentInterface interface {
	// ID returns the unique identifier of the component.
	ID() string

	// Name returns the name of the component.
	Name() string

	// Type returns the type of the component.
	Type() ComponentType

	// Description returns the description of the component.
	Description() string
}

// BootableComponentInterface represents a component that can be initialized and started.
type BootableInterface interface {
	// Initialize initializes the component.
	// Returns an error if the initialization fails.
	Initialize(ctx *common.Context) error
}

// Startable defines the interface for instances that can be started and stopped.
type StartableInterface interface {
	// Start starts the component.
	// Returns an error if the start operation fails.
	Start(ctx *common.Context) error

	// Stop stops the component.
	// Returns an error if the stop operation fails.
	Stop(ctx *common.Context) error
}

// ComponentFactoryInterface is responsible for creating
type ComponentFactoryInterface interface {
	// CreateComponent creates a new instance of the component.
	// Returns the created component and an error if the creation fails.
	CreateComponent(config *ComponentConfig) (ComponentInterface, error)
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

// ComponentCreationInfo encapsulates the necessary information to create a component.
type ComponentCreationInfo struct {
	Config  *ComponentConfig
	Factory ComponentFactoryInterface
}
