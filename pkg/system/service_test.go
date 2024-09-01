package system_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/system"
	"github.com/ebanfa/skeleton/pkg/types"
)

var _ = Describe("BaseSystemService", func() {
	var (
		service *system.BaseSystemService
	)

	BeforeEach(func() {
		service = system.NewBaseSystemService("1", "Service1", "Description1")
	})

	Describe("NewBaseSystemService", func() {
		It("creates a new BaseSystemService instance", func() {
			Expect(service).NotTo(BeNil())
			Expect(service.ID()).To(Equal("1"))
			Expect(service.Name()).To(Equal("Service1"))
			Expect(service.Description()).To(Equal("Description1"))
		})
	})

	Describe("Type", func() {
		It("returns the correct component type", func() {
			Expect(service.Type()).To(Equal(types.ServiceType))
		})
	})

	Describe("Interface Implementation", func() {
		It("implements the StartableInterface", func() {
			var _ types.StartableInterface = (*system.BaseSystemService)(nil)
		})
	})

	Describe("Initialize", func() {
		It("initializes without error", func() {
			mockContext := &common.Context{}
			err := service.Initialize(mockContext, nil)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Start", func() {
		It("returns an error indicating service not implemented", func() {
			mockContext := &common.Context{}
			err := service.Start(mockContext)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("service not implemented"))
		})
	})

	Describe("Stop", func() {
		It("returns an error indicating service not implemented", func() {
			mockContext := &common.Context{}
			err := service.Stop(mockContext)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("service not implemented"))
		})
	})
})
