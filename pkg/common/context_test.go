package common_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/common"
)

// Replace with the actual import path

var _ = Describe("Context", func() {
	var (
		ctx *common.Context
	)

	BeforeEach(func() {
		ctx = common.Background()
	})

	Describe("WithValue", func() {
		It("should store and retrieve a value", func() {
			newCtx := ctx.WithValue("key", "value")
			Expect(newCtx.Value("key")).To(Equal("value"))
		})

		It("should not affect the parent context", func() {
			newCtx := ctx.WithValue("key", "value")
			Expect(ctx.Value("key")).To(BeNil())
			Expect(newCtx.Value("key")).To(Equal("value"))
		})

		It("should handle multiple key-value pairs", func() {
			newCtx := ctx.WithValue("key1", "value1").WithValue("key2", "value2")
			Expect(newCtx.Value("key1")).To(Equal("value1"))
			Expect(newCtx.Value("key2")).To(Equal("value2"))
		})
	})

	Describe("WithPluginPaths", func() {
		It("should store plugin paths", func() {
			paths := []string{"/path/1", "/path/2"}
			newCtx := ctx.WithPluginPaths(paths...)
			Expect(newCtx.PluginPaths).To(Equal(paths))
		})

		It("should not affect the parent context", func() {
			paths := []string{"/path/1", "/path/2"}
			newCtx := ctx.WithPluginPaths(paths...)
			Expect(ctx.PluginPaths).To(BeEmpty())
			Expect(newCtx.PluginPaths).To(Equal(paths))
		})
	})

	Describe("WithTraceID", func() {
		It("should store a trace ID", func() {
			traceID := "trace-123"
			newCtx := ctx.WithTraceID(traceID)
			Expect(newCtx.Value("traceID")).To(Equal(traceID))
		})

		It("should not affect the parent context", func() {
			traceID := "trace-123"
			newCtx := ctx.WithTraceID(traceID)
			Expect(ctx.Value("traceID")).To(BeNil())
			Expect(newCtx.Value("traceID")).To(Equal(traceID))
		})
	})

	Describe("WithRemotePluginLocations", func() {
		It("should store remote plugin locations", func() {
			locations := []string{"https://github.com/repo1", "https://github.com/repo2"}
			newCtx := ctx.WithRemotePluginLocations(locations...)
			Expect(newCtx.RemotePluginLocations).To(Equal(locations))
		})

		It("should not affect the parent context", func() {
			locations := []string{"https://github.com/repo1", "https://github.com/repo2"}
			newCtx := ctx.WithRemotePluginLocations(locations...)
			Expect(ctx.RemotePluginLocations).To(BeEmpty())
			Expect(newCtx.RemotePluginLocations).To(Equal(locations))
		})
	})

	Describe("Background", func() {
		It("should return a non-nil context", func() {
			bgCtx := common.Background()
			Expect(bgCtx).NotTo(BeNil())
		})

		It("should have an empty value map", func() {
			bgCtx := common.Background()
			Expect(bgCtx.Value("any-key")).To(BeNil())
		})
	})

	Describe("WithContext", func() {
		It("should create a new Context from a standard context.Context", func() {
			stdCtx := context.Background()
			newCtx := common.WithContext(stdCtx)
			Expect(newCtx).NotTo(BeNil())
			Expect(newCtx.Value("any-key")).To(BeNil())
		})
	})

	Describe("WithTimeout", func() {
		It("should create a context with a timeout", func() {
			timeout := 100 * time.Millisecond
			newCtx, cancel := common.WithTimeout(ctx, timeout)
			defer cancel()

			Expect(newCtx).NotTo(BeNil())

			//Expect(newCtx.Deadline()).To(BeTemporally("~", time.Now().Add(timeout), 10*time.Millisecond))
		})

		It("should cancel the context after the timeout", func() {
			timeout := 50 * time.Millisecond
			newCtx, cancel := common.WithTimeout(ctx, timeout)
			defer cancel()

			time.Sleep(2 * timeout)
			Expect(newCtx.Err()).To(Equal(context.DeadlineExceeded))
		})
	})

	Describe("Concurrent access", func() {
		It("should handle concurrent read/write operations safely", func() {
			const goroutines = 100
			done := make(chan bool)

			for i := 0; i < goroutines; i++ {
				go func(id int) {
					newCtx := ctx.WithValue(id, id)
					Expect(newCtx.Value(id)).To(Equal(id))
					done <- true
				}(i)
			}

			for i := 0; i < goroutines; i++ {
				<-done
			}
		})
	})
})
