package component_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/component"
	"github.com/ebanfa/skeleton/pkg/types"
)

var _ = Describe("BaseComponent", func() {
	var (
		baseComponent *component.BaseComponent
		id            string
		name          string
		description   string
	)

	BeforeEach(func() {
		id = "test-id"
		name = "Test Component"
		description = "A test component for unit testing"
		baseComponent = component.NewComponentImpl(id, name, description)
	})

	Describe("NewComponentImpl", func() {
		It("creates a new BaseComponent with the correct attributes", func() {
			Expect(baseComponent).NotTo(BeNil())
			Expect(baseComponent.Id).To(Equal(id))
			Expect(baseComponent.Nm).To(Equal(name))
			Expect(baseComponent.Desc).To(Equal(description))
		})
	})

	Describe("ID", func() {
		It("returns the correct ID", func() {
			Expect(baseComponent.ID()).To(Equal(id))
		})

		It("returns the same value as the Id field", func() {
			Expect(baseComponent.ID()).To(Equal(baseComponent.Id))
		})
	})

	Describe("Name", func() {
		It("returns the correct name", func() {
			Expect(baseComponent.Name()).To(Equal(name))
		})

		It("returns the same value as the Nm field", func() {
			Expect(baseComponent.Name()).To(Equal(baseComponent.Nm))
		})
	})

	Describe("Type", func() {
		It("returns the BasicComponentType", func() {
			Expect(baseComponent.Type()).To(Equal(types.BasicComponentType))
		})
	})

	Describe("Description", func() {
		It("returns the correct description", func() {
			Expect(baseComponent.Description()).To(Equal(description))
		})

		It("returns the same value as the Desc field", func() {
			Expect(baseComponent.Description()).To(Equal(baseComponent.Desc))
		})
	})

	Describe("ComponentInterface", func() {
		It("implements the ComponentInterface", func() {
			var _ types.ComponentInterface = &component.BaseComponent{}
		})
	})

	Context("when creating multiple components", func() {
		var anotherComponent *component.BaseComponent

		BeforeEach(func() {
			anotherComponent = component.NewComponentImpl("another-id", "Another Component", "Another test component")
		})

		It("creates distinct components with different attributes", func() {
			Expect(baseComponent.ID()).NotTo(Equal(anotherComponent.ID()))
			Expect(baseComponent.Name()).NotTo(Equal(anotherComponent.Name()))
			Expect(baseComponent.Description()).NotTo(Equal(anotherComponent.Description()))
		})

		It("returns the same type for both components", func() {
			Expect(baseComponent.Type()).To(Equal(anotherComponent.Type()))
		})
	})
})
