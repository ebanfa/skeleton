package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

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
