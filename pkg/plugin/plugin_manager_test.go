package system_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/ebanfa/skeleton/pkg/common"
	"github.com/ebanfa/skeleton/pkg/mocks"
	system "github.com/ebanfa/skeleton/pkg/plugin"
	"github.com/ebanfa/skeleton/pkg/types"
)

var _ = Describe("PluginManager", func() {
	var (
		pluginManager types.PluginManagerInterface
		ctx           *common.Context
		mockPlugin    *mocks.PluginInterface
	)

	BeforeEach(func() {
		pluginManager = system.NewPluginManager()
		ctx = &common.Context{}
		mockPlugin = new(mocks.PluginInterface)
	})

	Describe("AddPlugin", func() {
		BeforeEach(func() {
			mockPlugin.On("Initialize", ctx, mock.Anything).Return(nil)
			mockPlugin.On("RegisterResources", ctx).Return(nil)
			mockPlugin.On("ID").Return("mock_plugin")
		})

		It("should add a plugin successfully", func() {
			err := pluginManager.AddPlugin(ctx, mockPlugin)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when adding a duplicate plugin", func() {
			BeforeEach(func() {
				err := pluginManager.AddPlugin(ctx, mockPlugin)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return an error", func() {
				err := pluginManager.AddPlugin(ctx, mockPlugin)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("already exists"))
			})
		})
	})

	Describe("RemovePlugin", func() {
		Context("when the plugin exists", func() {
			BeforeEach(func() {
				mockPlugin.On("Initialize", ctx, mock.Anything).Return(nil)
				mockPlugin.On("RegisterResources", ctx).Return(nil)
				mockPlugin.On("ID").Return("mock_plugin")
				err := pluginManager.AddPlugin(ctx, mockPlugin)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should remove the plugin successfully", func() {
				err := pluginManager.RemovePlugin(mockPlugin)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the plugin does not exist", func() {
			BeforeEach(func() {
				mockPlugin.On("ID").Return("non_existent_plugin")
			})

			It("should return an error", func() {
				err := pluginManager.RemovePlugin(mockPlugin)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("not found"))
			})
		})
	})

	Describe("GetPlugin", func() {
		BeforeEach(func() {
			mockPlugin.On("Initialize", ctx, mock.Anything).Return(nil)
			mockPlugin.On("RegisterResources", ctx).Return(nil)
			mockPlugin.On("ID").Return("mock_plugin")
			err := pluginManager.AddPlugin(ctx, mockPlugin)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should retrieve an existing plugin", func() {
			plugin, err := pluginManager.GetPlugin("mock_plugin")
			Expect(err).NotTo(HaveOccurred())
			Expect(plugin).To(Equal(mockPlugin))
		})

		It("should return an error for a non-existent plugin", func() {
			_, err := pluginManager.GetPlugin("non_existent_plugin")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("not found"))
		})
	})

	Describe("StartPlugins", func() {
		var mockPlugin1, mockPlugin2 *mocks.PluginInterface

		BeforeEach(func() {
			mockPlugin1 = new(mocks.PluginInterface)
			mockPlugin2 = new(mocks.PluginInterface)

			mockPlugin1.On("Initialize", ctx, mock.Anything).Return(nil)
			mockPlugin2.On("Initialize", ctx, mock.Anything).Return(nil)
			mockPlugin1.On("RegisterResources", ctx).Return(nil)
			mockPlugin2.On("RegisterResources", ctx).Return(nil)
			mockPlugin1.On("ID").Return("super_plugin1")
			mockPlugin2.On("ID").Return("super_plugin2")

			err := pluginManager.AddPlugin(ctx, mockPlugin1)
			Expect(err).NotTo(HaveOccurred())
			err = pluginManager.AddPlugin(ctx, mockPlugin2)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when all plugins start successfully", func() {
			BeforeEach(func() {
				mockPlugin1.On("Start", mock.Anything).Return(nil)
				mockPlugin2.On("Start", mock.Anything).Return(nil)
			})

			It("should start all plugins without error", func() {
				err := pluginManager.StartPlugins(ctx)
				Expect(err).NotTo(HaveOccurred())
				mockPlugin1.AssertCalled(GinkgoT(), "Start", mock.Anything)
				mockPlugin2.AssertCalled(GinkgoT(), "Start", mock.Anything)
			})
		})

		Context("when a plugin fails to start", func() {
			BeforeEach(func() {
				mockPlugin1.On("Start", mock.Anything).Return(errors.New("start error"))
				mockPlugin2.On("Start", mock.Anything).Return(nil)
			})

			It("should return an error", func() {
				err := pluginManager.StartPlugins(ctx)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("start error"))
				mockPlugin1.AssertCalled(GinkgoT(), "Start", mock.Anything)
			})
		})
	})

	Describe("StopPlugins", func() {
		var mockPlugin1, mockPlugin2 *mocks.PluginInterface

		BeforeEach(func() {
			mockPlugin1 = new(mocks.PluginInterface)
			mockPlugin2 = new(mocks.PluginInterface)

			mockPlugin1.On("Initialize", ctx, mock.Anything).Return(nil)
			mockPlugin2.On("Initialize", ctx, mock.Anything).Return(nil)
			mockPlugin1.On("RegisterResources", ctx).Return(nil)
			mockPlugin2.On("RegisterResources", ctx).Return(nil)
			mockPlugin1.On("ID").Return("super_plugin1")
			mockPlugin2.On("ID").Return("super_plugin2")
			mockPlugin1.On("Start", mock.Anything).Return(nil)
			mockPlugin2.On("Start", mock.Anything).Return(nil)

			err := pluginManager.AddPlugin(ctx, mockPlugin1)
			Expect(err).NotTo(HaveOccurred())
			err = pluginManager.AddPlugin(ctx, mockPlugin2)
			Expect(err).NotTo(HaveOccurred())
			err = pluginManager.StartPlugins(ctx)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when all plugins stop successfully", func() {
			BeforeEach(func() {
				mockPlugin1.On("Stop", ctx).Return(nil)
				mockPlugin2.On("Stop", ctx).Return(nil)
			})

			It("should stop all plugins without error", func() {
				err := pluginManager.StopPlugins(ctx)
				Expect(err).NotTo(HaveOccurred())
				mockPlugin1.AssertCalled(GinkgoT(), "Stop", ctx)
				mockPlugin2.AssertCalled(GinkgoT(), "Stop", ctx)
			})
		})

		Context("when a plugin fails to stop", func() {
			BeforeEach(func() {
				mockPlugin1.On("Stop", ctx).Return(errors.New("stop error"))
				mockPlugin2.On("Stop", ctx).Return(nil)
			})

			It("should return an error", func() {
				err := pluginManager.StopPlugins(ctx)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("stop error"))
				mockPlugin1.AssertCalled(GinkgoT(), "Stop", ctx)
			})
		})
	})
})
