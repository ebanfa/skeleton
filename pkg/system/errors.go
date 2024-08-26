package system

import "errors"

// Custom errors
var (
	ErrComponentNil                  = errors.New("component is nil")
	ErrComponentAlreadyExist         = errors.New("component already exists")
	ErrFactoryNotFound               = errors.New("factory not found")
	ErrComponentFactoryNil           = errors.New("component factory is nil")
	ErrComponentFactoryAlreadyExists = errors.New("component factory already exists")
	ErrComponentNotFound             = errors.New("component not found")
	ErrServiceAlreadyExists          = errors.New("service already exists")
	ErrServiceNotRegistered          = errors.New("service not registered")
	ErrOperationNotRegistered        = errors.New("operation not registered")
	ErrOperationAlreadyExists        = errors.New("operation already exists")
	ErrSystemNotInitialized          = errors.New("system not initialized")
	ErrSystemNotStarted              = errors.New("system not started")
	ErrSystemNotStopped              = errors.New("system not stopped")
	ErrComponentTypeNotFound         = errors.New("component type not found")
)
