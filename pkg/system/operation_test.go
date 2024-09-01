package system_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/system"
	"github.com/ebanfa/skeleton/pkg/types"
)

var _ = Describe("BaseSystemOperation", func() {
	var operation *system.BaseSystemOperation

	BeforeEach(func() {
		operation = system.NewBaseSystemOperation("1", "Operation1", "Description1")
	})

	Describe("NewBaseSystemOperation", func() {
		It("creates a new BaseSystemOperation instance", func() {
			Expect(operation).NotTo(BeNil())
			Expect(operation.ID()).To(Equal("1"))
			Expect(operation.Name()).To(Equal("Operation1"))
			Expect(operation.Description()).To(Equal("Description1"))
		})
	})

	Describe("Type", func() {
		It("returns the correct component type", func() {
			Expect(operation.Type()).To(Equal(types.BasicComponentType))
		})
	})

	Describe("Interface Implementation", func() {
		It("implements the SystemOperationInterface", func() {
			var _ types.SystemOperationInterface = (*system.BaseSystemOperation)(nil)
		})
	})

	Describe("Execute", func() {
		It("returns an error when executed", func() {
			mockContext := &common.Context{}
			output, err := operation.Execute(mockContext, nil)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("operation not implemented"))
			Expect(output).To(BeNil())
		})
	})

	Describe("Initialize", func() {
		It("initializes without error", func() {
			mockContext := &common.Context{}
			err := operation.Initialize(mockContext, nil)

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
