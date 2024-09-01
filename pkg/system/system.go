package system

import (
	"fmt"
	"sync"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/types"
)

// SystemImpl represents the core system in the application.
type SystemImpl struct {
	types.SystemInterface
	mutex         sync.RWMutex
	configuration *types.Configuration
	componentReg  types.ComponentRegistrarInterface
	logger        common.LoggerInterface
	eventBus      common.EventBusInterface
	pluginManager types.PluginManagerInterface
	status        types.SystemStatusType
	store         types.MultiStore
}

// NewSystem creates a new instance of the SystemImpl.
func NewSystem(
	logger common.LoggerInterface,
	eventBus common.EventBusInterface,
	configuration *types.Configuration,
	pluginManager types.PluginManagerInterface,
	componentReg types.ComponentRegistrarInterface,
	store types.MultiStore) *SystemImpl {
	return &SystemImpl{
		logger:        logger,
		eventBus:      eventBus,
		componentReg:  componentReg,
		configuration: configuration,
		pluginManager: pluginManager,
		status:        types.SystemStoppedType,
		store:         store,
	}
}

// Logger returns the system logger.
func (s *SystemImpl) Logger() common.LoggerInterface {
	return s.logger
}

// EventBus returns the system event bus.
func (s *SystemImpl) EventBus() common.EventBusInterface {
	return s.eventBus
}

// Configuration returns the system configuration.
func (s *SystemImpl) Configuration() *types.Configuration {
	return s.configuration
}

// ComponentRegistry returns the component registry.
func (s *SystemImpl) ComponentRegistry() types.ComponentRegistrarInterface {
	return s.componentReg
}

// MultiStore returns the multistore
func (s *SystemImpl) MultiStore() types.MultiStore {
	return s.store
}

// ComponentRegistry returns the component registry.
func (s *SystemImpl) PluginManager() types.PluginManagerInterface {
	return s.pluginManager
}

// Initialize initializes the system component by executing the initialize operation.
func (s *SystemImpl) Initialize(ctx *common.Context) error {
	// Override this function to customize system initialization

	s.status = types.SystemInitializedType

	return s.pluginManager.Initialize(ctx, s)
}

// Start starts the system component along with all registered services.
func (s *SystemImpl) Start(ctx *common.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.status != types.SystemInitializedType {
		return types.ErrSystemNotInitialized
	}

	if err := s.pluginManager.StartPlugins(ctx); err != nil {
		// Log the error, but continue stopping other services
		s.logger.Log(common.LevelError, "Error starting plugin:", err)
		return err
	}
	s.status = types.SystemStartedType
	return nil
}

// Stop stops the system component along with all registered services.
func (s *SystemImpl) Stop(ctx *common.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.status != types.SystemStartedType {
		return types.ErrSystemNotStarted
	}
	// Retrieve all components of type ServiceType
	components := s.ComponentRegistry().GetComponentsByType(types.ServiceType)

	// Iterate over each service component and start it
	for _, service := range components {
		// Check if the component implements SystemServiceInterface
		systemService, ok := service.(types.SystemServiceInterface)
		if !ok {
			return fmt.Errorf("failed to start service: component %v is not a service", service)
		}

		// Stop the service
		if err := systemService.Stop(ctx); err != nil {
			// Log the error, but continue stopping other services
			s.logger.Log(common.LevelError, "Error stopping service:", err)
		}
	}

	s.status = types.SystemStoppedType
	return nil
}

// ExecuteOperation executes the operation with the given ID and input data.
// Returns the output of the operation and an error if the operation is not found or if execution fails.
func (s *SystemImpl) ExecuteOperation(ctx *common.Context, operationID string, data *types.SystemOperationInput) (*types.SystemOperationOutput, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Retrieve the operation by its ID
	component, err := s.ComponentRegistry().GetComponent(operationID)
	if err != nil {
		return nil, err
	}

	// Check if the component implements Operation interface
	operation, ok := component.(types.SystemOperationInterface)
	if !ok {
		return nil, fmt.Errorf("failed to execute operation: component %v is not an operation", operation)
	}
	// Execute the operation
	return operation.Execute(ctx, data)
}

// StartService starts the service with the given ID.
// Returns an error if the service ID is not found or other error
func (s *SystemImpl) StartService(ctx *common.Context, serviceID string) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Retrieve the service by its ID
	component, err := s.ComponentRegistry().GetComponent(serviceID)
	if err != nil {
		return err
	}
	// Check if the component implements SystemServiceInterface interface
	service, ok := component.(types.SystemServiceInterface)
	if !ok {
		return fmt.Errorf("failed to start service: component %v is not a service", service)
	}

	// Start the service
	return service.Start(ctx)
}

// StopService stops the service with the given ID.
// Returns an error if the service ID is not found or other error.
func (s *SystemImpl) StopService(ctx *common.Context, serviceID string) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Retrieve the service by its ID
	component, err := s.ComponentRegistry().GetComponent(serviceID)
	if err != nil {
		return err
	}
	// Check if the component implements SystemServiceInterface interface
	service, ok := component.(types.SystemServiceInterface)
	if !ok {
		return fmt.Errorf("failed to start service: component %v is not a service", service)
	}

	// Start the service
	return service.Stop(ctx)
}

// RestartService restarts the service with the given ID.
// Returns an error if the service ID is not found or other error.
func (s *SystemImpl) RestartService(ctx *common.Context, serviceID string) error {
	// Stop the service first
	if err := s.StopService(ctx, serviceID); err != nil {
		return err
	}

	// Start the service
	return s.StartService(ctx, serviceID)
}
