package component

import "github.com/ebanfa/skeleton/pkg/types"

// BaseComponent represents a concrete implementation of the ComponentInterface.
type BaseComponent struct {
	types.ComponentInterface
	Id   string
	Nm   string
	Desc string
}

func NewComponentImpl(Id, Nm, Desc string) *BaseComponent {
	return &BaseComponent{Id: Id, Nm: Nm, Desc: Desc}
}

// ID returns the unique identifier of the component.
func (bc *BaseComponent) ID() string {
	return bc.Id
}

// Name returns the Nm of the component.
func (bc *BaseComponent) Name() string {
	return bc.Nm
}

// Type returns the type of the component.
func (bc *BaseComponent) Type() types.ComponentType {
	return types.BasicComponentType
}

// Description returns the Desc of the component.
func (bc *BaseComponent) Description() string {
	return bc.Desc
}
