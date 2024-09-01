package system_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/mocks"
	"github.com/ebanfa/skeleton/pkg/system"
	"github.com/ebanfa/skeleton/pkg/types"
)

var _ = Describe("System Utils", func() {
	var (
		ctx           *common.Context
		mockSystem    *mocks.SystemInterface
		mockConfig    *types.ComponentConfig
		mockComponent *mocks.SystemServiceInterface
		mockRegistrar *mocks.ComponentRegistrarInterface
	)

	BeforeEach(func() {
		ctx = &common.Context{}
		mockSystem = &mocks.SystemInterface{}
		mockConfig = &types.ComponentConfig{}
		mockComponent = &mocks.SystemServiceInterface{}
		mockRegistrar = &mocks.ComponentRegistrarInterface{}

		mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	})

	Describe("StartService", func() {
		Context("when the service starts successfully", func() {
			BeforeEach(func() {
				mockComponent.On("Start", ctx).Return(nil)
				mockComponent.On("ID").Return("mockService")
				mockComponent.On("Initialize", ctx, mockSystem).Return(nil)
				mockRegistrar.On("CreateComponent", ctx, mockConfig).Return(mockComponent, nil)
			})

			It("should not return an error", func() {
				err := system.StartService(ctx, mockSystem, mockConfig)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when creating component fails", func() {
			BeforeEach(func() {
				expectedErr := errors.New("create component error")
				mockRegistrar.On("CreateComponent", ctx, mockConfig).Return(mockComponent, expectedErr)
			})

			It("should return an error", func() {
				err := system.StartService(ctx, mockSystem, mockConfig)
				Expect(err).To(MatchError(fmt.Sprintf("failed to start service. Could not create component %s", mockConfig.ID)))
			})
		})

		Context("when service initialization fails", func() {
			expectedErr := errors.New("initialize error")
			BeforeEach(func() {
				mockComponent.On("ID").Return("mockService")
				mockComponent.On("Initialize", ctx, mockSystem).Return(expectedErr)
				mockRegistrar.On("CreateComponent", ctx, mockConfig).Return(mockComponent, nil)
			})

			It("should return an error", func() {
				err := system.StartService(ctx, mockSystem, mockConfig)
				Expect(err).To(MatchError(fmt.Sprintf("failed to initialize service: %s %v", mockComponent.ID(), expectedErr)))
			})
		})

		Context("when service start fails", func() {
			expectedErr := errors.New("start error")
			BeforeEach(func() {
				mockComponent.On("Start", ctx).Return(expectedErr)
				mockComponent.On("ID").Return("mockService")
				mockComponent.On("Initialize", ctx, mockSystem).Return(nil)
				mockRegistrar.On("CreateComponent", ctx, mockConfig).Return(mockComponent, nil)
			})

			It("should return an error", func() {
				err := system.StartService(ctx, mockSystem, mockConfig)
				Expect(err).To(MatchError(expectedErr))
			})
		})
	})

	Describe("StopService", func() {
		Context("when the service stops successfully", func() {
			BeforeEach(func() {
				mockComponent.On("Stop", ctx).Return(nil)
				mockComponent.On("ID").Return("mockService")
				mockRegistrar.On("GetComponent", "mockService").Return(mockComponent, nil)
			})

			It("should not return an error", func() {
				err := system.StopService(ctx, mockSystem, "mockService")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the component is not found", func() {
			BeforeEach(func() {
				mockRegistrar.On("GetComponent", "nonExistentService").Return(nil, errors.New("component not found"))
			})

			It("should return an error", func() {
				err := system.StopService(ctx, mockSystem, "nonExistentService")
				Expect(err).To(MatchError("failed to stop build service. Service not found: component not found"))
			})
		})

		Context("when the component does not implement SystemServiceInterface", func() {
			var mockNonSystemComponent *mocks.ComponentInterface

			BeforeEach(func() {
				mockNonSystemComponent = &mocks.ComponentInterface{}
				mockNonSystemComponent.On("ID").Return("mockService")
				mockRegistrar.On("GetComponent", "mockService").Return(mockNonSystemComponent, nil)
			})

			It("should return an error", func() {
				err := system.StopService(ctx, mockSystem, "mockService")
				Expect(err).To(MatchError("failed to stop service. Service component is not a system service"))
			})
		})
	})

	Describe("RegisterComponent", func() {
		var (
			mockLogger  *mocks.LoggerInterface
			mockFactory *mocks.ComponentFactoryInterface
		)

		BeforeEach(func() {
			mockLogger = &mocks.LoggerInterface{}
			mockFactory = &mocks.ComponentFactoryInterface{}
			mockSystem.On("Logger").Return(mockLogger)
			mockLogger.On("Logf", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mock.Anything)

			mockConfig = &types.ComponentConfig{
				ID:        "TestComponent",
				FactoryID: "TestFactory",
			}
		})

		Context("when registration is successful", func() {
			BeforeEach(func() {
				mockRegistrar.On("RegisterFactory", ctx, mockConfig.FactoryID, mockFactory).Return(nil)
				mockRegistrar.On("CreateComponent", ctx, mockConfig).Return(mockComponent, nil)
				mockComponent.On("Initialize", mock.Anything, mock.Anything).Return(nil)
			})

			It("should not return an error", func() {
				err := system.RegisterComponent(ctx, mockSystem, mockConfig, mockFactory)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when factory registration fails", func() {
			BeforeEach(func() {
				mockRegistrar.On("RegisterFactory", ctx, mockConfig.FactoryID, mockFactory).Return(errors.New("error registering factory"))
			})

			It("should return an error", func() {
				err := system.RegisterComponent(ctx, mockSystem, mockConfig, mockFactory)
				Expect(err).To(MatchError(ContainSubstring("failed to register component factory")))
			})
		})

		Context("when component creation fails", func() {
			BeforeEach(func() {
				mockRegistrar.On("RegisterFactory", ctx, mockConfig.FactoryID, mockFactory).Return(nil)
				mockRegistrar.On("CreateComponent", ctx, mockConfig).Return(nil, errors.New("error creating component"))
			})

			It("should return an error", func() {
				err := system.RegisterComponent(ctx, mockSystem, mockConfig, mockFactory)
				Expect(err).To(MatchError(ContainSubstring("failed to create component")))
			})
		})
	})
})
