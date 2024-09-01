package types

import (
	"github.com/ebanfa/skeleton/pkg/common"
)

// SystemComponentInterface represents a component in the system.
type SystemComponentInterface interface {
	ComponentInterface

	// Initialize initializes the module.
	// Returns an error if the initialization fails.
	Initialize(ctx *common.Context, system SystemInterface) error
}

// SystemServiceInterface represents a service within the system.
type SystemServiceInterface interface {
	StartableInterface
	SystemComponentInterface
}

// SystemOperationInput represents the input data for an operation.
type SystemOperationInput struct {
	// Data is the input data for the operation.
	Data interface{}
}

// SystemOperationOutput represents the response data from an operation.
type SystemOperationOutput struct {
	// Data is the response data from the operation.
	Data interface{}
}

// SystemOperation represents a unit of work that can be executed.
type SystemOperationInterface interface {
	SystemComponentInterface

	// Execute performs the operation with the given context and input parameters,
	// and returns any output or error encountered.
	Execute(ctx *common.Context, input *SystemOperationInput) (*SystemOperationOutput, error)
}

// SystemInterface represents the core system in the application.
type SystemInterface interface {
	BootableInterface
	StartableInterface

	// Logger returns the system logger.
	Logger() common.LoggerInterface

	// EventBus returns the system event bus.
	EventBus() common.EventBusInterface

	// Configuration returns the system configuration.
	Configuration() *Configuration

	// ComponentRegistry returns the component registry
	ComponentRegistry() ComponentRegistrarInterface

	// MultiStore returns the multi-level store used by the system, which is a store of stores.
	MultiStore() MultiStore

	// PluginManager returns the plugin manager
	PluginManager() PluginManagerInterface

	// ExecuteOperation executes the operation with the given ID and input data.
	// Returns the output of the operation and an error if the operation is not found or if execution fails.
	ExecuteOperation(ctx *common.Context, operationID string, data *SystemOperationInput) (*SystemOperationOutput, error)

	// StartService starts the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	StartService(ctx *common.Context, serviceID string) error

	// StopService stops the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	StopService(ctx *common.Context, serviceID string) error

	// RestartService restarts the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	RestartService(ctx *common.Context, serviceID string) error
}

// System status.
type SystemStatusType int

const (
	SystemInitializedType SystemStatusType = iota
	SystemStartedType
	SystemStoppedType
)
