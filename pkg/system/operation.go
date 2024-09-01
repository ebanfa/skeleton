package system

import (
	"errors"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/component"
	"github.com/ebanfa/skeleton/pkg/types"
)

// BaseSystemOperation.
type BaseSystemOperation struct {
	BaseSystemComponent // Embedding BaseComponent
}

// Type returns the type of the component.
func (bo *BaseSystemOperation) Type() types.ComponentType {
	return types.BasicComponentType
}

func NewBaseSystemOperation(id, name, description string) *BaseSystemOperation {
	return &BaseSystemOperation{
		BaseSystemComponent: BaseSystemComponent{
			BaseComponent: component.BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Execute performs the operation with the given context and input parameters,
// and returns any output or error encountered.
func (bo *BaseSystemOperation) Execute(ctx *common.Context, input *types.SystemOperationInput) (*types.SystemOperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	return nil, errors.New("operation not implemented")
}
