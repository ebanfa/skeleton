package plugin

import (
	"fmt"
	"sync"

	"errors"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/types"
)

// PluginManager represents functionality for managing plugins.
type PluginManager struct {
	types.PluginManagerInterface
	mu      sync.RWMutex                     // Mutex for synchronizing access to plugins map
	plugins map[string]types.PluginInterface // Map to store plugins by ID
	started bool                             // Flag to track whether the plugins have been started
	System  types.SystemInterface
}

// NewPluginManager creates a new instance of PluginManager.
func NewPluginManager() types.PluginManagerInterface {
	return &PluginManager{
		plugins: make(map[string]types.PluginInterface),
	}
}

// Initialize initializes the manager.
// Returns an error if the initialization fails.
func (m *PluginManager) Initialize(ctx *common.Context, system types.SystemInterface) error {
	m.System = system
	return nil
}

// AddPlugin adds a plugin to the plugin manager, initializes it, and registers its resources.
func (m *PluginManager) AddPlugin(ctx *common.Context, plugin types.PluginInterface) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if the plugin with the same ID already exists
	if _, exists := m.plugins[plugin.ID()]; exists {
		return fmt.Errorf("plugin with ID %s already exists", plugin.ID())
	}

	// Initialize the plugin
	if err := plugin.Initialize(ctx, m.System); err != nil {
		return fmt.Errorf("failed to initialize plugin %s: %w", plugin.ID(), err)
	}

	// Register resources for the plugin
	if err := plugin.RegisterResources(ctx); err != nil {
		return fmt.Errorf("failed to register resources for plugin %s: %w", plugin.ID(), err)
	}

	// Add the plugin to the plugins map
	m.plugins[plugin.ID()] = plugin
	return nil
}

// RemovePlugin removes a plugin from the plugin manager.
func (m *PluginManager) RemovePlugin(plugin types.PluginInterface) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Get the ID of the plugin to remove
	id := plugin.ID()

	// Check if the plugin exists
	if _, exists := m.plugins[id]; !exists {
		return fmt.Errorf("plugin with ID %s not found", id)
	}

	// Remove the plugin from the plugins map
	delete(m.plugins, id)
	return nil
}

// GetPlugin returns the plugin with the given ID.
func (m *PluginManager) GetPlugin(id string) (types.PluginInterface, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Retrieve the plugin from the plugins map
	plugin, exists := m.plugins[id]
	if !exists {
		return nil, fmt.Errorf("plugin with ID %s not found", id)
	}

	return plugin, nil
}

// StartPlugins starts all plugins managed by the plugin manager.
func (m *PluginManager) StartPlugins(ctx *common.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if the plugins have already been started
	if m.started {
		return nil
	}

	// Iterate through all plugins and start each one
	var errs []error
	for _, plugin := range m.plugins {
		if err := plugin.Start(ctx); err != nil {
			errs = append(errs, fmt.Errorf("error starting plugin %s: %w", plugin.ID(), err))
		}
	}

	// Update the started flag
	m.started = true

	// Check if there were any errors starting plugins
	if len(errs) > 0 {
		return fmt.Errorf("errors starting plugins: %v", errs)
	}

	return nil
}

// StopPlugins stops all plugins managed by the plugin manager.
func (m *PluginManager) StopPlugins(ctx *common.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if the plugins have not been started
	if !m.started {
		return nil
	}

	// Iterate through all plugins and stop each one
	var errs []error
	for _, plugin := range m.plugins {
		if err := plugin.Stop(ctx); err != nil {
			errs = append(errs, fmt.Errorf("error stopping plugin %s: %w", plugin.ID(), err))
		}
	}

	// Update the started flag
	m.started = false

	// Check if there were any errors stopping plugins
	if len(errs) > 0 {
		return fmt.Errorf("errors stopping plugins: %v", errs)
	}

	return nil
}

// DiscoverPlugins discovers available plugins within the system.
func (m *PluginManager) DiscoverPlugins(ctx *common.Context) ([]types.PluginInterface, error) {
	// Implement logic to discover available plugins
	return nil, errors.New("not implemented")
}

// LoadRemotePlugin loads a plugin from a remote source.
func (m *PluginManager) LoadRemotePlugin(ctx *common.Context, pluginURL string) (types.PluginInterface, error) {
	// Implement logic to load a plugin from a remote source
	return nil, errors.New("not implemented")
}
