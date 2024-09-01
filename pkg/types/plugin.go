package types

import "github.com/ebanfa/skeleton/pkg/common"

// PluginInterface represents a plugin in the system.
type PluginInterface interface {
	SystemServiceInterface

	// RegisterResources registers resources into the system.
	// Returns an error if resource registration fails.
	RegisterResources(ctx *common.Context) error
}

// PluginManagerInterface represents functionality for managing plugins.
type PluginManagerInterface interface {

	// Initialize initializes the manager.
	// Returns an error if the initialization fails.
	Initialize(ctx *common.Context, system SystemInterface) error

	// AddPlugin adds a plugin to the plugin manager.
	AddPlugin(ctx *common.Context, plugin PluginInterface) error

	// RemovePlugin removes a plugin from the plugin manager.
	RemovePlugin(plugin PluginInterface) error

	// GetPlugin returns the plugin with the given name.
	GetPlugin(name string) (PluginInterface, error)

	// StartPlugins starts all plugins managed by the plugin manager.
	StartPlugins(ctx *common.Context) error

	// StopPlugins stops all plugins managed by the plugin manager.
	StopPlugins(ctx *common.Context) error

	// DiscoverPlugins discovers available plugins within the system.
	DiscoverPlugins(ctx *common.Context) ([]PluginInterface, error)

	// LoadRemotePlugin loads a plugin from a remote source.
	LoadRemotePlugin(ctx *common.Context, pluginURL string) (PluginInterface, error)
}
