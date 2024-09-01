package system_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/mocks"
	"github.com/ebanfa/skeleton/pkg/system"
	"github.com/ebanfa/skeleton/pkg/types"
)

var _ = Describe("BaseSystemComponent", func() {
	var (
		baseComponent *system.BaseSystemComponent
		ctx           *common.Context
		mockSystem    *mocks.SystemInterface
	)

	BeforeEach(func() {
		baseComponent = system.NewBaseSystemComponent("test-id", "Test Component", "A test component")
		ctx = &common.Context{}
		mockSystem = mocks.NewSystemInterface(GinkgoT())
	})

	Describe("NewBaseSystemComponent", func() {
		It("creates a new BaseSystemComponent with the correct attributes", func() {
			Expect(baseComponent.Id).To(Equal("test-id"))
			Expect(baseComponent.Nm).To(Equal("Test Component"))
			Expect(baseComponent.Desc).To(Equal("A test component"))
		})
	})

	Describe("Type", func() {
		It("returns the correct component type", func() {
			Expect(baseComponent.Type()).To(Equal(types.SystemComponentType))
		})
	})

	Describe("Initialize", func() {
		It("initializes the component with the provided system", func() {
			err := baseComponent.Initialize(ctx, mockSystem)

			Expect(err).NotTo(HaveOccurred())
			Expect(baseComponent.System).To(Equal(mockSystem))
		})
	})
})

// MockSystemInterface is a mock implementation of SystemInterface
type MockSystemInterface struct{}

func NewMockSystemInterface() *MockSystemInterface {
	return &MockSystemInterface{}
}

// Implement the methods of SystemInterface for MockSystemInterface if needed
