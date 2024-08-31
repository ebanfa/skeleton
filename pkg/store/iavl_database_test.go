package store_test

import (
	"github.com/cosmos/iavl"
	"github.com/ebanfa/skeleton/pkg/store"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"cosmossdk.io/log"
	iavldb "github.com/cosmos/iavl/db"
)

var _ = Describe("IAVLDatabase", func() {
	var (
		iavlTree *iavl.MutableTree
		mockDB   *store.IAVLDatabase
	)

	BeforeEach(func() {
		iavlTree = iavl.NewMutableTree(iavldb.NewMemDB(), 100, false, log.NewNopLogger())
		mockDB = store.NewIAVLDatabase(iavlTree)
	})

	Describe("Get", func() {
		Context("when retrieving a nonexistent key", func() {
			It("should return nil value without error", func() {
				value, err := mockDB.Get([]byte("nonexistent"))
				Expect(err).NotTo(HaveOccurred())
				Expect(value).To(BeNil())
			})
		})
	})

	Describe("Set", func() {
		It("should set a new key-value pair without error", func() {
			err := mockDB.Set([]byte("testKey"), []byte("testValue"))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Delete", func() {
		Context("when attempting to delete a nonexistent key", func() {
			It("should not return an error", func() {
				err := mockDB.Delete([]byte("nonexistent"))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("Has", func() {
		It("should correctly check for the existence of a key", func() {
			key := []byte("testKey")
			value := []byte("testValue")
			err := mockDB.Set(key, value)
			Expect(err).NotTo(HaveOccurred())

			hasKey, err := mockDB.Has(key)
			Expect(err).NotTo(HaveOccurred())
			Expect(hasKey).To(BeTrue())
		})
	})

	Describe("Iterate", func() {
		It("should iterate over key-value pairs without error", func() {
			mockDB.Set([]byte("key1"), []byte("value1"))
			mockDB.Set([]byte("key2"), []byte("value2"))

			err := mockDB.Iterate(func(key, value []byte) bool {
				return false
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("IterateRange", func() {
		It("should iterate over a range of key-value pairs", func() {
			mockDB.Set([]byte("key1"), []byte("value1"))
			mockDB.Set([]byte("key2"), []byte("value2"))

			var keys [][]byte
			var values [][]byte
			err := mockDB.IterateRange(nil, nil, true, func(key, value []byte) bool {
				keys = append(keys, key)
				values = append(values, value)
				return false // Continue iteration
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(keys).To(Equal([][]byte{[]byte("key1"), []byte("key2")}))
			Expect(values).To(Equal([][]byte{[]byte("value1"), []byte("value2")}))
		})
	})

	Describe("Hash", func() {
		It("should return a non-nil hash", func() {
			hash := mockDB.Hash()
			Expect(hash).NotTo(BeNil())
		})
	})

	Describe("Version", func() {
		It("should return the initial version as 0", func() {
			version := mockDB.Version()
			Expect(version).To(Equal(int64(0)))
		})
	})

	Describe("Load", func() {
		It("should load the latest versioned tree without error", func() {
			version, err := mockDB.Load()
			Expect(err).NotTo(HaveOccurred())
			Expect(version).To(Equal(int64(0)))
		})
	})

	Describe("SaveVersion and Rollback", func() {
		It("should save a new version and rollback correctly", func() {
			// Set initial data
			Expect(mockDB.Set([]byte("key1"), []byte("value1"))).To(Succeed())
			Expect(mockDB.Set([]byte("key2"), []byte("value2"))).To(Succeed())

			// Save initial version
			hash1, version1, err := mockDB.SaveVersion()
			Expect(err).NotTo(HaveOccurred())
			Expect(hash1).NotTo(BeNil())
			Expect(version1).To(Equal(int64(1)))

			// Make additional changes
			Expect(mockDB.Set([]byte("key3"), []byte("value3"))).To(Succeed())

			// Rollback to previous version
			mockDB.Rollback()

			// Check data after rollback
			value1, _ := mockDB.Get([]byte("key1"))
			value2, _ := mockDB.Get([]byte("key2"))
			value3, _ := mockDB.Get([]byte("key3"))

			Expect(value1).NotTo(BeNil())
			Expect(value2).NotTo(BeNil())
			Expect(value3).To(BeNil())
		})
	})

	Describe("Close", func() {
		It("should close the tree without error", func() {
			err := mockDB.Close()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("String", func() {
		It("should return a non-empty string representation", func() {
			str, err := mockDB.String()
			Expect(err).NotTo(HaveOccurred())
			Expect(str).NotTo(BeEmpty())
		})
	})

	Describe("WorkingVersion", func() {
		It("should return the initial working version as 1", func() {
			workingVersion := mockDB.WorkingVersion()
			Expect(workingVersion).To(Equal(int64(1)))
		})
	})

	Describe("WorkingHash", func() {
		It("should return a non-nil working hash", func() {
			workingHash := mockDB.WorkingHash()
			Expect(workingHash).NotTo(BeNil())
		})
	})

	Describe("AvailableVersions", func() {
		It("should return available versions after saving", func() {
			_, _, _ = mockDB.SaveVersion() // Save a version to make it available
			versions := mockDB.AvailableVersions()
			Expect(versions).NotTo(BeEmpty())
		})
	})

	Describe("IsEmpty", func() {
		It("should return true for an empty database", func() {
			empty := mockDB.IsEmpty()
			Expect(empty).To(BeTrue())
		})
	})
})
