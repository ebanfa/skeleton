package types

import "time"

// ComponentConfig represents the configuration for a component.
type ComponentConfig struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	FactoryID    string      `json:"factoryId"`
	CustomConfig interface{} // Custom configuration
}

type ServiceConfiguration struct {
	ComponentConfig
	RetryInterval time.Duration // Interval between retries
	// Other service-specific configuration optionsetl.ErrScheduledProcessNotFound
	CustomConfig interface{} // Custom configuration
}

// OperationConfiguration represents the configuration for an operation.
type OperationConfiguration struct {
	ComponentConfig
}

// Configuration represents the system configuration.
type Configuration struct {
	Debug        bool
	Verbose      bool
	Services     []*ServiceConfiguration   // Service configurations
	Operations   []*OperationConfiguration // Operation configurations
	CustomConfig interface{}
}
