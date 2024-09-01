package store_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/mocks"
	"github.com/ebanfa/skeleton/pkg/store"
	"github.com/ebanfa/skeleton/pkg/types"
)

const (
	MockDbName = "MockDb"
	MockDbPath = "MockDbPath"
)

var _ = Describe("StoreImpl", func() {
	Describe("NewStoreImpl", func() {
		Context("when creating a new StoreImpl instance", func() {
			It("should create a new StoreImpl successfully", func() {
				// Arrange
				mockDatabase := &mocks.Database{}

				expectedStore := &store.StoreImpl{
					Database: mockDatabase,
				}

				// Act
				createdStore, err := store.NewStoreImpl(MockDbName, MockDbPath, mockDatabase)

				// Assert
				Expect(err).NotTo(HaveOccurred())
				Expect(createdStore).NotTo(BeNil())
				Expect(createdStore.Database).To(Equal(expectedStore.Database))
			})

			It("should return an error when database is nil", func() {
				// Arrange
				var nilDb types.Database

				// Act
				createdStore, err := store.NewStoreImpl(MockDbName, MockDbPath, nilDb)

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(createdStore).To(BeNil())
			})
		})
	})
})
