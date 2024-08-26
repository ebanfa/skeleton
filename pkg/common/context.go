package common

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Context is a custom wrapper around the standard context.Context. It provides
// additional functionality such as logging, tracing, and custom cancelation logic.
type Context struct {
	// Embed the standard context.Context to inherit its methods and behavior.
	context.Context

	// mu is a mutex to protect concurrent access to the context values.
	mu sync.RWMutex

	// values is a map to store custom key-value pairs associated with this context.
	values map[interface{}]interface{}

	// PluginPaths is a slice of paths to search for plugins.
	PluginPaths []string

	// RemotePluginLocations is a slice of remote locations (e.g., GitHub repositories) to download plugins from.
	RemotePluginLocations []string
}

// WithValue returns a new Context with the given key-value pair associated with it.
// It overrides the standard context.WithValue method to provide thread-safe access
// to the context values.
func (c *Context) WithValue(key, value interface{}) *Context {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.values == nil {
		c.values = make(map[interface{}]interface{})
	}

	newCtx := &Context{
		Context: c.Context,
		values:  make(map[interface{}]interface{}, len(c.values)+1),
	}

	for k, v := range c.values {
		newCtx.values[k] = v
	}
	newCtx.values[key] = value

	return newCtx
}

// Value returns the value associated with the given key in the context.
// It overrides the standard context.Value method to provide thread-safe access
// to the context values.
func (c *Context) Value(key interface{}) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.values == nil {
		return nil
	}

	return c.values[key]
}

// WithPluginPaths returns a new Context with the given plugin paths.
func (c *Context) WithPluginPaths(paths ...string) *Context {
	newCtx := &Context{
		Context:     c.Context,
		values:      c.values,
		PluginPaths: make([]string, len(paths)),
	}

	copy(newCtx.PluginPaths, paths)

	return newCtx
}

// WithTraceID returns a new Context with the given traceID associated with it.
func (c *Context) WithTraceID(traceID string) *Context {
	newCtx := &Context{
		Context: c.Context,
		values:  c.values,
	}

	newCtx.values["traceID"] = traceID
	return newCtx
}

// WithRemotePluginLocations returns a new Context with the given remote plugin locations.
func (c *Context) WithRemotePluginLocations(locations ...string) *Context {
	newCtx := &Context{
		Context:               c.Context,
		values:                c.values,
		RemotePluginLocations: append([]string(nil), locations...),
	}

	return newCtx
}

// Background returns a non-nil, empty Context. It is similar to the standard
// context.Background() function but returns a custom Context type.
func Background() *Context {
	return &Context{
		Context: context.Background(),
		values:  make(map[interface{}]interface{}),
	}
}

// Background returns a non-nil, empty Context. It is similar to the standard
// context.Background() function but returns a custom Context type.
func WithContext(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
		values:  make(map[interface{}]interface{}),
	}
}

// WithTimeout returns a new Context with the given timeout duration.
// It is similar to the standard context.WithTimeout() function but returns
// a custom Context type and logs the timeout event.
func WithTimeout(parent *Context, timeout time.Duration) (*Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent.Context, timeout)
	newCtx := &Context{
		Context: ctx,
		values:  parent.values,
	}

	go func() {
		<-ctx.Done()
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Print(fmt.Sprintf("Request timed out after %s", timeout))
		}
	}()

	return newCtx, cancel
}
