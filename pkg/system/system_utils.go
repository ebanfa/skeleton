package system

import (
	"errors"
	"fmt"

	"github.com/ebanfa/skeleton/pkg/common"
)

func StartService(
	ctx *common.Context,
	system SystemInterface,
	config *ComponentConfig) error {

	registrar := system.ComponentRegistry()

	component, err := registrar.CreateComponent(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to start service. Could not create component %s", config.ID)
	}

	service, ok := component.(SystemServiceInterface)
	if !ok {
		return fmt.Errorf("failed to start service. Component %s is not a system service", component.ID())
	}

	// Initialize the service
	if err := service.Initialize(ctx, system); err != nil {
		return fmt.Errorf("failed to initialize service: %s %v", component.ID(), err)
	}

	return service.Start(ctx)
}

func StopService(ctx *common.Context, system SystemInterface, id string) error {
	// Retrieve the BuildService component from the ComponentRegistry
	component, err := system.ComponentRegistry().GetComponent(id)
	if err != nil {
		return fmt.Errorf("failed to stop build service. Service not found: %v", err)
	}

	// Check if the retrieved component implements the SystemServiceInterface
	service, ok := component.(SystemServiceInterface)
	if !ok {
		return errors.New("failed to stop service. Service component is not a system service")
	}

	return service.Stop(ctx)
}

func RegisterComponent(ctx *common.Context, system SystemInterface, config *ComponentConfig, factory ComponentFactoryInterface) error {
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
	case SystemOperationInterface:
		v.Initialize(ctx, system)
	case SystemServiceInterface:
		v.Initialize(ctx, system)
	default:

	}

	return nil
}
