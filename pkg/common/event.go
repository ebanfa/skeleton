package common

import (
	"fmt"

	"github.com/asaskevich/EventBus"
)

const (
	// EventTypeDataExtracted represents an event emitted when data is extracted from a source blockchain.
	EventTypeDataExtracted string = "data_extracted"

	// EventTypeDataTransformed represents an event emitted when data is transformed by a pipeline stage.
	EventTypeDataTransformed string = "data_transformed"

	// EventTypeDataLoaded represents an event emitted when data is loaded into a target blockchain.
	EventTypeDataLoaded string = "data_loaded"
)

// Event represents an event within the system.
type Event struct {
	// Type is the type or identifier of the event.
	Type string

	// Data is the payload or data associated with the event.
	Data interface{}
}

// EventHandler defines the signature for an event handler function.
type EventHandler func(event Event)

// BusSubscriptionParams represents the parameters for subscribing to an event topic.
type BusSubscriptionParams struct {
	// Topic is the event topic to subscribe to.
	Topic string

	// EventHandler is the function that will handle the events for the subscribed topic.
	EventHandler EventHandler
}

// BusSubscriber defines subscription-related bus behavior.
type BusSubscriber interface {
	// Subscribe subscribes to an event topic with the given parameters.
	Subscribe(params BusSubscriptionParams) error

	// SubscribeAsync subscribes to an event topic asynchronously with the given parameters.
	SubscribeAsync(params BusSubscriptionParams, transactional bool) error

	// SubscribeOnce subscribes to an event topic for a single event occurrence with the given parameters.
	SubscribeOnce(params BusSubscriptionParams) error

	// SubscribeOnceAsync subscribes to an event topic asynchronously for a single event occurrence with the given parameters.
	SubscribeOnceAsync(params BusSubscriptionParams) error

	// Unsubscribe unsubscribes from an event topic with the given parameters.
	Unsubscribe(params BusSubscriptionParams) error
}

// BusPublisher defines publishing-related bus behavior.
type BusPublisher interface {
	// Publish publishes an event to the event bus.
	Publish(event Event)
}

// BusController defines bus control behavior (checking handler's presence, synchronization).
type BusController interface {
	// HasCallback checks if a handler is registered for the given topic.
	HasCallback(topic string) bool

	// WaitAsync blocks until all asynchronous operations are completed.
	WaitAsync()
}

// EventBusInterface englobes global (subscribe, publish, control) bus behavior.
type EventBusInterface interface {
	BusController
	BusSubscriber
	BusPublisher
}

// SystemEventBus is a concrete implementation of the EventBusInterface.
type SystemEventBus struct {
	bus      EventBus.Bus           // Underlying third-party EventBus instance
	handlers map[string]interface{} // Map to store function references used for subscription
}

// NewSystemEventBus creates a new instance of the SystemEventBus.
func NewSystemEventBus() EventBusInterface {
	return &SystemEventBus{
		bus:      EventBus.New(),
		handlers: make(map[string]interface{}),
	}
}

// Subscribe subscribes to an event topic with the given parameters.
func (eb *SystemEventBus) Subscribe(params BusSubscriptionParams) error {
	// Create a wrapper function that constructs the Event and calls the provided EventHandler
	handler := func(args ...interface{}) {
		event := Event{
			Type: params.Topic,
			Data: args[0],
		}
		params.EventHandler(event)
	}

	// Store the function reference for later use in Unsubscribe
	eb.handlers[params.Topic] = handler

	// Subscribe to the underlying EventBus using the wrapper function
	return eb.bus.Subscribe(params.Topic, handler)
}

// SubscribeAsync subscribes to an event topic asynchronously with the given parameters.
func (eb *SystemEventBus) SubscribeAsync(params BusSubscriptionParams, transactional bool) error {
	// Create a wrapper function that constructs the Event and calls the provided EventHandler
	handler := func(args ...interface{}) {
		event := Event{
			Type: params.Topic,
			Data: args[0],
		}
		params.EventHandler(event)
	}

	// Store the function reference for later use in Unsubscribe
	eb.handlers[params.Topic] = handler

	// Subscribe asynchronously to the underlying EventBus using the wrapper function
	return eb.bus.SubscribeAsync(params.Topic, handler, transactional)
}

// SubscribeOnce subscribes to an event topic for a single event occurrence with the given parameters.
func (eb *SystemEventBus) SubscribeOnce(params BusSubscriptionParams) error {
	// Create a wrapper function that constructs the Event and calls the provided EventHandler
	handler := func(args ...interface{}) {
		event := Event{
			Type: params.Topic,
			Data: args[0],
		}
		params.EventHandler(event)
	}

	// Store the function reference for later use in Unsubscribe
	eb.handlers[params.Topic] = handler

	// Subscribe once to the underlying EventBus using the wrapper function
	return eb.bus.SubscribeOnce(params.Topic, handler)
}

// SubscribeOnceAsync subscribes to an event topic asynchronously for a single event occurrence with the given parameters.
func (eb *SystemEventBus) SubscribeOnceAsync(params BusSubscriptionParams) error {
	// Create a wrapper function that constructs the Event and calls the provided EventHandler
	handler := func(args ...interface{}) {
		event := Event{
			Type: params.Topic,
			Data: args[0],
		}
		params.EventHandler(event)
	}

	// Store the function reference for later use in Unsubscribe
	eb.handlers[params.Topic] = handler

	// Subscribe once asynchronously to the underlying EventBus using the wrapper function
	return eb.bus.SubscribeOnceAsync(params.Topic, handler)
}

// Unsubscribe unsubscribes from an event topic with the given parameters.
func (eb *SystemEventBus) Unsubscribe(params BusSubscriptionParams) error {
	// Retrieve the function reference used for subscription
	handler, ok := eb.handlers[params.Topic]
	if !ok {
		return fmt.Errorf("no handler found for topic %s", params.Topic)
	}

	// Unsubscribe from the underlying EventBus using the stored function reference
	err := eb.bus.Unsubscribe(params.Topic, handler)
	if err == nil {
		// Remove the function reference from the handlers map if unsubscription succeeded
		delete(eb.handlers, params.Topic)
	}
	return err
}

// Publish publishes an event to the event bus.
func (eb *SystemEventBus) Publish(event Event) {
	eb.bus.Publish(event.Type, event.Data)
}

// HasCallback checks if a handler is registered for the given topic.
func (eb *SystemEventBus) HasCallback(topic string) bool {
	return eb.bus.HasCallback(topic)
}

// WaitAsync blocks until all asynchronous operations are completed.
func (eb *SystemEventBus) WaitAsync() {
	eb.bus.WaitAsync()
}
