package component

import (
	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/types"
)

func RegisterComponentFactory(system types.SystemInterface, factoryConfig types.FactoryConfig) error {
	ctx := &common.Context{}

	// Register the factory
	err := system.ComponentRegistry().RegisterFactory(ctx, factoryConfig.FactoryId, factoryConfig.Factory)
	if err != nil {
		return err
	}

	// Create and register each component
	for _, id := range factoryConfig.ComponentIDs {
		config := &types.ComponentConfig{
			ID:          id,
			Name:        id,
			Description: id,
			FactoryID:   factoryConfig.FactoryId,
		}

		_, err := system.ComponentRegistry().CreateComponent(ctx, config)
		if err != nil {
			return err
		}
	}

	return nil

}
