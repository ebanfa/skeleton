package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/types"
)

// LoadConfigurationFromFile loads the configuration from a file at the given path.
func LoadConfigurationFromFile(filePath string, target interface{}) error {
	// Read the configuration file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %v", err)
	}

	// Unmarshal the JSON data into the target struct
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal configuration data: %v", err)
	}

	return nil
}

func StartService(
	ctx *common.Context,
	system types.SystemInterface,
	config *types.ComponentConfig) error {

	registrar := system.ComponentRegistry()

	component, err := registrar.CreateComponent(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to start service. Could not create component %s", config.ID)
	}

	service, ok := component.(types.SystemServiceInterface)
	if !ok {
		return fmt.Errorf("failed to start service. Component %s is not a system service", component.ID())
	}

	// Initialize the service
	if err := service.Initialize(ctx, system); err != nil {
		return fmt.Errorf("failed to initialize service: %s %v", component.ID(), err)
	}

	return service.Start(ctx)
}

func StopService(ctx *common.Context, system types.SystemInterface, id string) error {
	// Retrieve the BuildService component from the ComponentRegistry
	component, err := system.ComponentRegistry().GetComponent(id)
	if err != nil {
		return fmt.Errorf("failed to stop build service. Service not found: %v", err)
	}

	// Check if the retrieved component implements the SystemServiceInterface
	service, ok := component.(types.SystemServiceInterface)
	if !ok {
		return errors.New("failed to stop service. Service component is not a system service")
	}

	return service.Stop(ctx)
}

func RegisterComponent(ctx *common.Context, system types.SystemInterface, config *types.ComponentConfig, factory types.ComponentFactoryInterface) error {
	registrar := system.ComponentRegistry()
	system.Logger().Logf(common.LevelDebug, "Registering component %s with factory ID %s", config.ID, config.FactoryID)
	// Register the factory
	err := registrar.RegisterFactory(ctx, config.FactoryID, factory)
	if err != nil {
		return fmt.Errorf("failed to register component factory: %w", err)
	}

	// Create the component
	component, err := registrar.CreateComponent(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to create component: %s", config.ID)
	}

	// Initial system services and operations
	switch v := component.(type) {
	case types.SystemOperationInterface:
		v.Initialize(ctx, system)
	case types.SystemServiceInterface:
		v.Initialize(ctx, system)
	default:

	}

	return nil
}
