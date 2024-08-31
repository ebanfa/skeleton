package store_test

import (
	"fmt"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/mocks"
	"github.com/ebanfa/skeleton/pkg/store"
)

var _ = Describe("StoreFactory", func() {
	var (
		mockDbFactory *mocks.DatabaseFactory
		storeFactory  store.StoreFactory
		databasesDir  string
	)

	BeforeEach(func() {
		mockDbFactory = &mocks.DatabaseFactory{}
		databasesDir = "/tmp/test_databases"
		storeFactory = store.NewStoreFactory(databasesDir, mockDbFactory)
	})

	Describe("NewStoreFactory", func() {
		It("creates a new StoreFactory", func() {
			Expect(storeFactory).NotTo(BeNil())
		})
	})

	Describe("CreateStore", func() {
		var (
			storeName    string
			mockDatabase *mocks.Database
		)

		BeforeEach(func() {
			storeName = "testStore"
			mockDatabase = &mocks.Database{}
		})

		Context("when database creation is successful", func() {
			BeforeEach(func() {
				expectedDatabaseID, expectedDatabasePath := store.GenererateStorageInfo(storeName, databasesDir)
				mockDbFactory.On("CreateDatabase", expectedDatabaseID, expectedDatabasePath).Return(mockDatabase, nil)
			})

			It("creates a new store successfully", func() {
				createdStore, err := storeFactory.CreateStore(storeName)

				Expect(err).NotTo(HaveOccurred())
				Expect(createdStore).NotTo(BeNil())

				mockDbFactory.AssertExpectations(GinkgoT())
			})

			It("generates the correct storage info", func() {
				_, err := storeFactory.CreateStore(storeName)

				Expect(err).NotTo(HaveOccurred())

				expectedDatabaseID, expectedDatabasePath := store.GenererateStorageInfo(storeName, databasesDir)
				Expect(filepath.Base(expectedDatabasePath)).To(Equal(fmt.Sprintf("%s.db", expectedDatabaseID)))
				Expect(filepath.Dir(expectedDatabasePath)).To(Equal(databasesDir))
			})
		})

		Context("when database creation fails", func() {
			BeforeEach(func() {
				expectedDatabaseID, expectedDatabasePath := store.GenererateStorageInfo(storeName, databasesDir)
				mockDbFactory.On("CreateDatabase", expectedDatabaseID, expectedDatabasePath).Return(nil, fmt.Errorf("database creation failed"))
			})

			It("returns an error", func() {
				createdStore, err := storeFactory.CreateStore(storeName)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("failed to create database"))
				Expect(createdStore).To(BeNil())

				mockDbFactory.AssertExpectations(GinkgoT())
			})
		})
	})
})
