package system

import (
	"errors"

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

// ComponentFactoryInterface is responsible for creating
type ComponentFactoryInterface interface {
	// CreateComponent creates a new instance of the component.
	// Returns the created component and an error if the creation fails.
	CreateComponent(config *ComponentConfig) (ComponentInterface, error)
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

// BaseComponent represents a concrete implementation of the ComponentInterface.
type BaseComponent struct {
	ComponentInterface
	Id   string
	Nm   string
	Desc string
}

func NewComponentImpl(Id, Nm, Desc string) *BaseComponent {
	return &BaseComponent{Id: Id, Nm: Nm, Desc: Desc}
}

// ID returns the unique identifier of the component.
func (bc *BaseComponent) ID() string {
	return bc.Id
}

// Name returns the Nm of the component.
func (bc *BaseComponent) Name() string {
	return bc.Nm
}

// Type returns the type of the component.
func (bc *BaseComponent) Type() ComponentType {
	return BasicComponentType
}

// Description returns the Desc of the component.
func (bc *BaseComponent) Description() string {
	return bc.Desc
}

// BaseSystemComponent.
type BaseSystemComponent struct {
	BaseComponent // Embedding BaseComponent
	System        SystemInterface
}

// Type returns the type of the component.
func (bo *BaseSystemComponent) Type() ComponentType {
	return SystemComponentType
}

func NewBaseSystemComponent(id, name, description string) *BaseSystemComponent {
	return &BaseSystemComponent{
		BaseComponent: BaseComponent{
			Id:   id,
			Nm:   name,
			Desc: description,
		},
	}
}

// Initialize initializes the module.
// Returns an error if the initialization fails.
func (bo *BaseSystemComponent) Initialize(ctx *common.Context, system SystemInterface) error {
	bo.System = system
	return nil
}

// BaseSystemOperation.
type BaseSystemOperation struct {
	BaseSystemComponent // Embedding BaseComponent
}

// Type returns the type of the component.
func (bo *BaseSystemOperation) Type() ComponentType {
	return BasicComponentType
}

func NewBaseSystemOperation(id, name, description string) *BaseSystemOperation {
	return &BaseSystemOperation{
		BaseSystemComponent: BaseSystemComponent{
			BaseComponent: BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Execute performs the operation with the given context and input parameters,
// and returns any output or error encountered.
func (bo *BaseSystemOperation) Execute(ctx *common.Context, input *SystemOperationInput) (*SystemOperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	return nil, errors.New("operation not implemented")
}

// BaseSystemService.
type BaseSystemService struct {
	BaseSystemComponent
}

// Type returns the type of the component.
func (bo *BaseSystemService) Type() ComponentType {
	return ServiceType
}

// NewBaseSystemService creates a new instance of BaseSystemService.
func NewBaseSystemService(id, name, description string) *BaseSystemService {
	return &BaseSystemService{
		BaseSystemComponent: BaseSystemComponent{
			BaseComponent: BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Start starts the component.
// Returns an error if the start operation fails.
func (bo *BaseSystemService) Start(ctx *common.Context) error {
	// Start the service component
	return errors.New("service not implemented")
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (bo *BaseSystemService) Stop(ctx *common.Context) error {
	// Stop the service component
	return errors.New("service not implemented")
}
