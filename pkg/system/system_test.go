package system_test

import (
	"errors"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/mocks"
	systemApi "github.com/ebanfa/skeleton/pkg/system"
	"github.com/ebanfa/skeleton/pkg/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("System", func() {
	var (
		ctx                    *common.Context
		logger                 *mocks.LoggerInterface
		sys                    types.SystemInterface
		registrar              *mocks.ComponentRegistrarInterface
		serviceFactory         *mocks.ComponentFactoryInterface
		operationFactory       *mocks.ComponentFactoryInterface
		mockServiceComponent   *mocks.SystemServiceInterface
		mockOperationComponent *mocks.SystemOperationInterface
		mockPluginManager      *mocks.PluginManagerInterface
		mockMultiStore         *mocks.MultiStore
		configuration          *types.Configuration
	)

	BeforeEach(func() {
		ctx = common.Background()
		logger = &mocks.LoggerInterface{}
		eventBus := &mocks.EventBusInterface{}
		registrar = &mocks.ComponentRegistrarInterface{}
		serviceFactory = &mocks.ComponentFactoryInterface{}
		operationFactory = &mocks.ComponentFactoryInterface{}
		mockServiceComponent = &mocks.SystemServiceInterface{}
		mockOperationComponent = &mocks.SystemOperationInterface{}
		mockPluginManager = &mocks.PluginManagerInterface{}
		mockMultiStore = &mocks.MultiStore{}

		configuration = &types.Configuration{
			Services: []*types.ServiceConfiguration{
				{
					ComponentConfig: types.ComponentConfig{
						ID:   "Service1_ID",
						Name: "Service1",
					},
				},
			},
			Operations: []*types.OperationConfiguration{
				{
					ComponentConfig: types.ComponentConfig{
						ID:   "Operation1_ID",
						Name: "Operation1",
					},
				},
			},
		}

		sys = systemApi.NewSystem(logger, eventBus, configuration, mockPluginManager, registrar, mockMultiStore)

		mockServiceComponent.On("Type").Return(types.ServiceType)
		mockServiceComponent.On("Initialize", ctx, mock.Anything).Return(nil)

		mockOperationComponent.On("Type").Return(types.OperationType)
		mockOperationComponent.On("Initialize", ctx, mock.Anything).Return(nil)

		serviceFactory.On("CreateComponent", mock.Anything).Return(mockServiceComponent, nil)
		operationFactory.On("CreateComponent", mock.Anything).Return(mockOperationComponent, nil)

		mockPluginManager.On("Initialize", ctx, mock.Anything).Return(nil)
		mockPluginManager.On("StartPlugins", ctx).Return(nil)
	})

	Describe("Initialize", func() {
		Context("when initialization is successful", func() {
			BeforeEach(func() {
				registrar.On("GetComponentFactory", "Service1Factory").Return(serviceFactory, nil)
				registrar.On("GetComponentFactory", "Operation1Factory").Return(operationFactory, nil)
			})

			It("should initialize without error", func() {
				err := sys.Initialize(ctx)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when initialization fails", func() {
			BeforeEach(func() {

				errPluginManager := mocks.NewPluginManagerInterface(GinkgoT())
				errPluginManager.On("Initialize", mock.Anything, mock.Anything).Return(errors.New("Plugin initializatin error"))

				sys = systemApi.NewSystem(nil, nil, configuration, errPluginManager, registrar, mockMultiStore)
			})

			It("should return an error", func() {
				err := sys.Initialize(ctx)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Start", func() {
		Context("when start is successful", func() {
			BeforeEach(func() {
				registrar.On("GetComponentFactory", "Service1Factory").Return(serviceFactory, nil)
				registrar.On("GetComponentFactory", "Operation1Factory").Return(operationFactory, nil)
				registrar.On("GetComponentsByType", types.ServiceType).Return([]types.ComponentInterface{mockServiceComponent}, nil)
				registrar.On("GetComponentsByType", types.OperationType).Return([]types.ComponentInterface{mockOperationComponent}, nil)
				mockServiceComponent.On("Start", ctx).Return(nil)
			})

			It("should start without error", func() {

				err := sys.Initialize(ctx)
				Expect(err).NotTo(HaveOccurred())

				err = sys.Start(ctx)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when start fails", func() {
			BeforeEach(func() {
				registrar.On("GetComponentFactory", "Service1Factory").Return(serviceFactory, nil)
				registrar.On("GetComponentFactory", "Operation1Factory").Return(operationFactory, nil)
				registrar.On("GetComponentsByType", types.ServiceType).Return([]types.ComponentInterface{mockServiceComponent}, nil)
				registrar.On("GetComponentsByType", types.OperationType).Return([]types.ComponentInterface{mockOperationComponent}, nil)
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, registrar, mockMultiStore)
			})

			It("should return an error", func() {
				err := sys.Start(ctx)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Stop", func() {
		Context("when stop is successful", func() {
			BeforeEach(func() {
				registrar.On("GetComponentFactory", "testFactoryInitializeServiceSuccess").Return(serviceFactory, nil)
				registrar.On("GetComponentsByType", types.ServiceType).Return([]types.ComponentInterface{}, nil)
				mockServiceComponent.On("Stop", ctx).Return(nil)
			})

			It("should stop without error", func() {
				err := sys.Initialize(ctx)
				Expect(err).NotTo(HaveOccurred())

				err = sys.Start(ctx)
				Expect(err).NotTo(HaveOccurred())

				err = sys.Stop(ctx)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when stop fails", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should return an error", func() {
				err := sys.Stop(ctx)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("ExecuteOperation", func() {
		Context("when execution is successful", func() {
			BeforeEach(func() {
				mockOperation := &mocks.SystemOperationInterface{}
				operationInput := &types.SystemOperationInput{}
				expectedOutput := &types.SystemOperationOutput{}

				registrar.On("GetComponent", "Operation1_ID").Return(mockOperation, nil)
				mockOperation.On("Execute", ctx, operationInput).Return(expectedOutput, nil)
			})

			It("should execute without error", func() {
				_, err := sys.ExecuteOperation(ctx, "Operation1_ID", &types.SystemOperationInput{})
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when component is not found", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				componentReg.On("GetComponent", "operation_id").Return(mockOperationComponent, types.ErrComponentNotFound)
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should return an error", func() {
				output, err := sys.ExecuteOperation(ctx, "operation_id", &types.SystemOperationInput{})
				Expect(err).To(HaveOccurred())
				Expect(output).To(BeNil())
			})
		})

		Context("when component is not an operation", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				componentReg.On("GetComponent", "operation_id").Return(mockServiceComponent, nil)
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should return an error", func() {
				output, err := sys.ExecuteOperation(ctx, "operation_id", &types.SystemOperationInput{})
				Expect(err).To(HaveOccurred())
				Expect(output).To(BeNil())
			})
		})
	})

	Describe("StartService", func() {
		Context("when start is successful", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}

				mockServiceComponent.On("Start", mock.Anything).Return(nil)
				componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)

				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should start without error", func() {
				err := sys.StartService(ctx, "service_id")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when component is not found", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, types.ErrComponentNotFound)
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should return an error", func() {
				err := sys.StartService(ctx, "service_id")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("StopService", func() {
		Context("when stop is successful", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				mockServiceComponent.On("Stop", mock.Anything).Return(nil)
				componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should stop without error", func() {
				err := sys.StopService(ctx, "service_id")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when component is not found", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, types.ErrComponentNotFound)
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should return an error", func() {
				err := sys.StopService(ctx, "service_id")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("RestartService", func() {
		Context("when restart is successful", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				mockServiceComponent.On("Stop", mock.Anything).Return(nil)
				mockServiceComponent.On("Start", mock.Anything).Return(nil)
				componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should restart without error", func() {
				err := sys.RestartService(ctx, "service_id")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when stop fails during restart", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)
				mockServiceComponent.On("Stop", ctx).Return(errors.New("Error stopping service"))
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should return an error", func() {
				err := sys.RestartService(ctx, "service_id")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when start fails during restart", func() {
			BeforeEach(func() {
				componentReg := &mocks.ComponentRegistrarInterface{}
				componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)
				mockServiceComponent.On("Stop", ctx).Return(nil)
				mockServiceComponent.On("Start", ctx).Return(errors.New("Error starting service"))
				sys = systemApi.NewSystem(nil, nil, &types.Configuration{}, mockPluginManager, componentReg, mockMultiStore)
			})

			It("should return an error", func() {
				err := sys.RestartService(ctx, "service_id")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
