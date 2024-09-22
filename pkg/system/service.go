package system

import (
	"errors"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/component"
	"github.com/ebanfa/skeleton/pkg/types"
)

// BaseSystemService.
type BaseSystemService struct {
	types.SystemServiceInterface
	BaseSystemComponent
}

// Type returns the type of the component.
func (bo *BaseSystemService) Type() types.ComponentType {
	return types.ServiceType
}

// NewBaseSystemService creates a new instance of BaseSystemService.
func NewBaseSystemService(id, name, description string) *BaseSystemService {
	return &BaseSystemService{
		BaseSystemComponent: BaseSystemComponent{
			BaseComponent: component.BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Start starts the component.
// Returns an error if the start operation fails.
func (bo *BaseSystemService) Start(ctx *common.Context) error {
	// Start the service component
	return errors.New("service not implemented")
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (bo *BaseSystemService) Stop(ctx *common.Context) error {
	// Stop the service component
	return errors.New("service not implemented")
}
