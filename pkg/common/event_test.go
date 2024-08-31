package common_test

import (
	"sync"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/common"
)

var _ = Describe("SystemEventBus", func() {
	var (
		eventBus common.EventBusInterface
	)

	BeforeEach(func() {
		eventBus = common.NewSystemEventBus()
	})

	Describe("Subscribe and Publish", func() {
		It("should successfully subscribe to and receive published events", func() {
			receivedEvent := make(chan common.Event, 1)
			handler := func(event common.Event) {
				receivedEvent <- event
			}

			err := eventBus.Subscribe(common.BusSubscriptionParams{
				Topic:        "test_topic",
				EventHandler: handler,
			})
			Expect(err).NotTo(HaveOccurred())

			publishedEvent := common.Event{
				Type: "test_topic",
				Data: "test_data",
			}
			eventBus.Publish(publishedEvent)

			var received common.Event
			Eventually(receivedEvent).Should(Receive(&received))
			Expect(received).To(Equal(publishedEvent))
		})
	})

	Describe("SubscribeAsync", func() {
		It("should successfully subscribe asynchronously and receive published events", func() {
			var wg sync.WaitGroup
			wg.Add(1)

			receivedEvent := make(chan common.Event, 1)
			handler := func(event common.Event) {
				defer wg.Done()
				receivedEvent <- event
			}

			err := eventBus.SubscribeAsync(common.BusSubscriptionParams{
				Topic:        "async_topic",
				EventHandler: handler,
			}, false)
			Expect(err).NotTo(HaveOccurred())

			publishedEvent := common.Event{
				Type: "async_topic",
				Data: "async_data",
			}
			eventBus.Publish(publishedEvent)

			wg.Wait()
			var received common.Event
			Eventually(receivedEvent).Should(Receive(&received))
			Expect(received).To(Equal(publishedEvent))
		})
	})

	Describe("SubscribeOnce", func() {
		It("should receive only one event when subscribed once", func() {
			receivedEvents := make(chan common.Event, 2)
			handler := func(event common.Event) {
				receivedEvents <- event
			}

			err := eventBus.SubscribeOnce(common.BusSubscriptionParams{
				Topic:        "once_topic",
				EventHandler: handler,
			})
			Expect(err).NotTo(HaveOccurred())

			event1 := common.Event{Type: "once_topic", Data: "data1"}
			event2 := common.Event{Type: "once_topic", Data: "data2"}

			eventBus.Publish(event1)
			eventBus.Publish(event2)

			Eventually(receivedEvents).Should(Receive(Equal(event1)))
			Consistently(receivedEvents).ShouldNot(Receive())
		})
	})

	Describe("SubscribeOnceAsync", func() {
		It("should receive only one event asynchronously when subscribed once", func() {
			var wg sync.WaitGroup
			wg.Add(1)

			receivedEvents := make(chan common.Event, 2)
			handler := func(event common.Event) {
				defer wg.Done()
				receivedEvents <- event
			}

			err := eventBus.SubscribeOnceAsync(common.BusSubscriptionParams{
				Topic:        "once_async_topic",
				EventHandler: handler,
			})
			Expect(err).NotTo(HaveOccurred())

			event1 := common.Event{Type: "once_async_topic", Data: "async_data1"}
			event2 := common.Event{Type: "once_async_topic", Data: "async_data2"}

			eventBus.Publish(event1)
			eventBus.Publish(event2)

			wg.Wait()
			Eventually(receivedEvents).Should(Receive(Equal(event1)))
			Consistently(receivedEvents).ShouldNot(Receive())
		})
	})

	Describe("Unsubscribe", func() {
		It("should successfully unsubscribe from a topic", func() {
			receivedEvents := make(chan common.Event, 1)
			handler := func(event common.Event) {
				receivedEvents <- event
			}

			params := common.BusSubscriptionParams{
				Topic:        "unsub_topic",
				EventHandler: handler,
			}

			err := eventBus.Subscribe(params)
			Expect(err).NotTo(HaveOccurred())

			event := common.Event{Type: "unsub_topic", Data: "unsub_data"}
			eventBus.Publish(event)

			Eventually(receivedEvents).Should(Receive(Equal(event)))

			err = eventBus.Unsubscribe(params)
			Expect(err).NotTo(HaveOccurred())

			eventBus.Publish(event)
			Consistently(receivedEvents).ShouldNot(Receive())
		})

		It("should return an error when unsubscribing from a non-existent topic", func() {
			params := common.BusSubscriptionParams{
				Topic:        "non_existent_topic",
				EventHandler: func(event common.Event) {},
			}

			err := eventBus.Unsubscribe(params)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("no handler found for topic"))
		})
	})

	Describe("HasCallback", func() {
		It("should return true for a subscribed topic", func() {
			err := eventBus.Subscribe(common.BusSubscriptionParams{
				Topic:        "callback_topic",
				EventHandler: func(event common.Event) {},
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(eventBus.HasCallback("callback_topic")).To(BeTrue())
		})

		It("should return false for an unsubscribed topic", func() {
			Expect(eventBus.HasCallback("non_existent_topic")).To(BeFalse())
		})
	})

	Describe("WaitAsync", func() {
		It("should wait for all asynchronous operations to complete", func() {
			var wg sync.WaitGroup
			wg.Add(1)

			err := eventBus.SubscribeAsync(common.BusSubscriptionParams{
				Topic: "wait_async_topic",
				EventHandler: func(event common.Event) {
					time.Sleep(100 * time.Millisecond)
					wg.Done()
				},
			}, false)
			Expect(err).NotTo(HaveOccurred())

			event := common.Event{Type: "wait_async_topic", Data: "async_wait_data"}
			eventBus.Publish(event)

			eventBus.WaitAsync()
			wg.Wait() // This should not block as WaitAsync should have ensured completion
		})
	})
})
