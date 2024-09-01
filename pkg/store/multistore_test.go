package store_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/ebanfa/skeleton/pkg/mocks"
	"github.com/ebanfa/skeleton/pkg/store"
	"github.com/ebanfa/skeleton/pkg/types"
)

const (
	MultiDbName = "MockDb"
	MultiDbPath = "MockDbPath"
)

var _ = Describe("MultiStore", func() {
	var (
		mockStore        *mocks.Store
		mockStoreFactory *mocks.StoreFactory
		ms               types.MultiStore
	)

	BeforeEach(func() {
		mockStore = &mocks.Store{}
		mockStoreFactory = &mocks.StoreFactory{}
	})

	Describe("NewMultiStore", func() {
		It("creates a new MultiStore successfully", func() {
			var err error
			ms, err = store.NewMultiStore(mockStore, mockStoreFactory)

			Expect(err).NotTo(HaveOccurred())
			Expect(ms).NotTo(BeNil())
			Expect(ms.GetStoreCount()).To(Equal(0))
		})
	})

	Describe("CreateStore", func() {
		BeforeEach(func() {
			var err error
			ms, err = store.NewMultiStore(mockStore, mockStoreFactory)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when creating a store successfully", func() {
			BeforeEach(func() {
				mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)
			})

			It("creates a new store", func() {
				createdStore, created, err := ms.CreateStore("mockNamespace")

				Expect(err).NotTo(HaveOccurred())
				Expect(created).To(BeTrue())
				Expect(createdStore).To(Equal(mockStore))

				mockStoreFactory.AssertCalled(GinkgoT(), "CreateStore", mock.Anything)
			})
		})

		Context("when creating a store with an invalid namespace", func() {
			It("returns an error", func() {
				createdStore, created, err := ms.CreateStore("")

				Expect(err).To(HaveOccurred())
				Expect(created).To(BeFalse())
				Expect(createdStore).To(BeNil())

				mockStoreFactory.AssertNotCalled(GinkgoT(), "CreateStore", mock.Anything)
			})
		})

		Context("when creating a new store in MultiStore", func() {
			BeforeEach(func() {
				mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)
			})

			It("creates a new store and increments the store count", func() {
				store, _, err := ms.CreateStore("test")

				Expect(err).NotTo(HaveOccurred())
				Expect(store).NotTo(BeNil())
				Expect(ms.GetStoreCount()).To(Equal(1))
			})
		})

		Context("when attempting to create an existing store", func() {
			BeforeEach(func() {
				mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)
				_, _, err := ms.CreateStore("test")
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns the existing store without creating a new one", func() {
				_, created, err := ms.CreateStore("test")

				Expect(err).NotTo(HaveOccurred())
				Expect(created).To(BeFalse())
				Expect(ms.GetStoreCount()).To(Equal(1))
			})
		})
	})

	Describe("GetStore", func() {
		BeforeEach(func() {
			var err error
			ms, err = store.NewMultiStore(mockStore, mockStoreFactory)
			Expect(err).NotTo(HaveOccurred())
			mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)
		})

		Context("when retrieving an existing store", func() {
			BeforeEach(func() {
				_, _, err := ms.CreateStore("test")
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns the existing store", func() {
				testStore := ms.GetStore([]byte(store.GenerateStoreId("test")))

				Expect(testStore).NotTo(BeNil())
			})
		})

		Context("when retrieving a non-existing store", func() {
			It("returns nil", func() {
				testStore := ms.GetStore([]byte(store.GenerateStoreId("test")))

				Expect(testStore).To(BeNil())
			})
		})
	})

	Describe("GetStoreCount", func() {
		BeforeEach(func() {
			var err error
			ms, err = store.NewMultiStore(mockStore, mockStoreFactory)
			Expect(err).NotTo(HaveOccurred())
			mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)
		})

		It("returns the correct number of stores", func() {
			_, _, err := ms.CreateStore("test1")
			Expect(err).NotTo(HaveOccurred())

			_, _, err = ms.CreateStore("test2")
			Expect(err).NotTo(HaveOccurred())

			count := ms.GetStoreCount()
			Expect(count).To(Equal(2))
		})
	})

	Describe("Load", func() {
		BeforeEach(func() {
			mockStore.On("Load").Return(int64(1), nil)
			mockStore.On("Iterate", mock.Anything).Return(nil)
			mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

			var err error
			ms, err = store.NewMultiStore(mockStore, mockStoreFactory)
			Expect(err).NotTo(HaveOccurred())
		})

		It("loads the latest versioned database from disk", func() {
			version, err := ms.Load()

			Expect(err).NotTo(HaveOccurred())
			Expect(version).To(Equal(int64(1)))
		})
	})

	Describe("SaveVersion", func() {
		BeforeEach(func() {
			mockStore.On("SaveVersion").Return([]byte("data"), int64(1), nil)
			mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

			var err error
			ms, err = store.NewMultiStore(mockStore, mockStoreFactory)
			Expect(err).NotTo(HaveOccurred())
		})

		It("saves a new version of the database to disk", func() {
			data, version, err := ms.SaveVersion()

			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(Equal([]byte("data")))
			Expect(version).To(Equal(int64(1)))
		})
	})
})
