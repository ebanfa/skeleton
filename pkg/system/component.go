package system

import (
	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/component"
	"github.com/ebanfa/skeleton/pkg/types"
)

// BaseSystemComponent.
type BaseSystemComponent struct {
	component.BaseComponent // Embedding BaseComponent
	System                  types.SystemInterface
}

// Type returns the type of the component.
func (bo *BaseSystemComponent) Type() types.ComponentType {
	return types.SystemComponentType
}

func NewBaseSystemComponent(id, name, description string) *BaseSystemComponent {
	return &BaseSystemComponent{
		BaseComponent: component.BaseComponent{
			Id:   id,
			Nm:   name,
			Desc: description,
		},
	}
}

// Initialize initializes the module.
// Returns an error if the initialization fails.
func (bo *BaseSystemComponent) Initialize(ctx *common.Context, system types.SystemInterface) error {
	bo.System = system
	return nil
}
