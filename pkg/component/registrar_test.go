package component_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/component"
	"github.com/ebanfa/skeleton/pkg/mocks"
	"github.com/ebanfa/skeleton/pkg/types"
)

var _ = Describe("ComponentRegistrar", func() {
	var (
		registrar *component.ComponentRegistrar
		ctx       *common.Context
	)

	BeforeEach(func() {
		registrar = component.NewComponentRegistrar()
		ctx = &common.Context{}
	})

	Describe("Factory Registration", func() {
		It("should register a factory successfully", func() {
			mockFactory := mocks.NewComponentFactoryInterface(GinkgoT())
			err := registrar.RegisterFactory(ctx, "mock-factory", mockFactory)
			Expect(err).NotTo(HaveOccurred())

			factory, err := registrar.GetFactory("mock-factory")
			Expect(err).NotTo(HaveOccurred())
			Expect(factory).To(Equal(mockFactory))
		})

		It("should return an error when registering a duplicate factory", func() {
			mockFactory := mocks.NewComponentFactoryInterface(GinkgoT())
			err := registrar.RegisterFactory(ctx, "mock-factory", mockFactory)
			Expect(err).NotTo(HaveOccurred())

			err = registrar.RegisterFactory(ctx, "mock-factory", mockFactory)
			Expect(err).To(MatchError("factory with ID mock-factory already exists"))
		})
	})

	Describe("Component Creation and Retrieval", func() {
		var mockFactory *mocks.ComponentFactoryInterface
		var mockComponent *mocks.ComponentInterface

		BeforeEach(func() {
			mockFactory = mocks.NewComponentFactoryInterface(GinkgoT())
			mockComponent = mocks.NewComponentInterface(GinkgoT())
			err := registrar.RegisterFactory(ctx, "mock-factory", mockFactory)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should create and retrieve a component successfully", func() {
			config := &types.ComponentConfig{
				ID:        "test-component",
				FactoryID: "mock-factory",
			}

			mockComponent.On("ID").Return(config.ID)
			mockFactory.On("CreateComponent", mock.Anything).Return(mockComponent, nil)

			component, err := registrar.CreateComponent(ctx, config)
			Expect(err).NotTo(HaveOccurred())
			Expect(component).NotTo(BeNil())

			retrievedComponent, err := registrar.GetComponent(config.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(retrievedComponent).To(Equal(component))
		})

		It("should return an error when creating a component with non-existent factory", func() {
			config := &types.ComponentConfig{
				ID:        "test-component",
				FactoryID: "non-existent-factory",
			}
			_, err := registrar.CreateComponent(ctx, config)
			Expect(err).To(MatchError("factory with ID non-existent-factory not found"))
		})

		It("should return an error when retrieving a non-existent component", func() {
			_, err := registrar.GetComponent("non-existent-component")
			Expect(err).To(MatchError("component with ID non-existent-component not found"))
		})
	})

	Describe("GetComponentsByType", func() {
		BeforeEach(func() {
			mockFactory := mocks.NewComponentFactoryInterface(GinkgoT())
			mockServiceComponent := mocks.NewSystemServiceInterface(GinkgoT())
			mockOperationComponent := mocks.NewSystemOperationInterface(GinkgoT())

			err := registrar.RegisterFactory(ctx, "mock-factory", mockFactory)
			Expect(err).NotTo(HaveOccurred())

			configs := []*types.ComponentConfig{
				{ID: "component1", FactoryID: "mock-factory"},
				{ID: "component2", FactoryID: "mock-factory"},
			}

			mockServiceComponent.On("ID").Return(configs[0].ID)
			mockOperationComponent.On("ID").Return(configs[1].ID)

			mockServiceComponent.On("Type").Return(types.ServiceType)
			mockOperationComponent.On("Type").Return(types.OperationType)

			mockFactory.On("CreateComponent", mock.Anything).Return(mockServiceComponent, nil).Once()
			mockFactory.On("CreateComponent", mock.Anything).Return(mockOperationComponent, nil).Once()

			for _, config := range configs {
				_, err := registrar.CreateComponent(ctx, config)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("should return components of the specified type", func() {
			components := registrar.GetComponentsByType(types.ServiceType)
			Expect(components).To(HaveLen(1))
			Expect(components[0].ID()).To(Equal("component1"))

			components = registrar.GetComponentsByType(types.OperationType)
			Expect(components).To(HaveLen(1))
			Expect(components[0].ID()).To(Equal("component2"))
		})

		It("should return an empty slice for a type with no components", func() {
			components := registrar.GetComponentsByType(types.BasicComponentType)
			Expect(components).To(BeEmpty())
		})
	})

	Describe("GetAllComponents and GetAllFactories", func() {
		BeforeEach(func() {
			configs := []*types.ComponentConfig{
				{ID: "component1", FactoryID: "mock-factory1"},
				{ID: "component2", FactoryID: "mock-factory2"},
			}

			mockFactory1 := mocks.NewComponentFactoryInterface(GinkgoT())
			mockFactory2 := mocks.NewComponentFactoryInterface(GinkgoT())

			mockServiceComponent := mocks.NewSystemServiceInterface(GinkgoT())
			mockOperationComponent := mocks.NewSystemOperationInterface(GinkgoT())

			Expect(registrar.RegisterFactory(ctx, "mock-factory1", mockFactory1)).To(Succeed())
			Expect(registrar.RegisterFactory(ctx, "mock-factory2", mockFactory2)).To(Succeed())

			mockServiceComponent.On("ID").Return(configs[0].ID)
			mockOperationComponent.On("ID").Return(configs[1].ID)

			mockFactory1.On("CreateComponent", mock.Anything).Return(mockServiceComponent, nil).Once()
			mockFactory2.On("CreateComponent", mock.Anything).Return(mockOperationComponent, nil).Once()

			for _, config := range configs {
				_, err := registrar.CreateComponent(ctx, config)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("should return all registered components", func() {
			components := registrar.GetAllComponents()
			Expect(components).To(HaveLen(2))
			Expect(components[0].ID()).To(Or(Equal("component1"), Equal("component2")))
			Expect(components[1].ID()).To(Or(Equal("component1"), Equal("component2")))
		})

		It("should return all registered factories", func() {
			factories := registrar.GetAllFactories()
			Expect(factories).To(HaveLen(2))
		})
	})

	Describe("Unregister Component and Factory", func() {
		var mockComponent *mocks.ComponentInterface
		var mockFactory *mocks.ComponentFactoryInterface

		BeforeEach(func() {
			config := &types.ComponentConfig{
				ID:        "test-component",
				FactoryID: "mock-factory",
			}

			mockComponent = mocks.NewComponentInterface(GinkgoT())
			mockFactory = mocks.NewComponentFactoryInterface(GinkgoT())

			err := registrar.RegisterFactory(ctx, "mock-factory", mockFactory)
			Expect(err).NotTo(HaveOccurred())

			mockComponent.On("ID").Return(config.ID)
			mockFactory.On("CreateComponent", mock.Anything).Return(mockComponent, nil)

			_, err = registrar.CreateComponent(ctx, config)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should unregister a component successfully", func() {
			err := registrar.UnregisterComponent(ctx, "test-component")
			Expect(err).NotTo(HaveOccurred())

			_, err = registrar.GetComponent("test-component")
			Expect(err).To(MatchError("component with ID test-component not found"))
		})

		It("should return an error when unregistering a non-existent component", func() {
			err := registrar.UnregisterComponent(ctx, "non-existent-component")
			Expect(err).To(MatchError("component with ID non-existent-component not found"))
		})

		It("should unregister a factory successfully", func() {
			err := registrar.UnregisterFactory(ctx, "mock-factory")
			Expect(err).NotTo(HaveOccurred())

			_, err = registrar.GetFactory("mock-factory")
			Expect(err).To(MatchError("factory with ID mock-factory not found"))

			/* _, err = registrar.GetComponent("test-component")
			Expect(err).To(MatchError("component with ID test-component not found")) */
		})

		It("should return an error when unregistering a non-existent factory", func() {
			err := registrar.UnregisterFactory(ctx, "non-existent-factory")
			Expect(err).To(MatchError("factory with ID non-existent-factory not found"))
		})
	})
})
