package db_test

import (
	"os"
	"path/filepath"

	"github.com/ebanfa/skeleton/pkg/db"
	"github.com/ebanfa/skeleton/pkg/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("IAVLDatabaseFactory", func() {
	var (
		factory *db.IAVLDatabaseFactory
		mockDbm *mocks.Database
	)

	BeforeEach(func() {
		mockDbm = &mocks.Database{}
		mockDbm.On("Get", mock.Anything).Return([]byte{}, nil)
		mockDbm.On("NewBatchWithSize", mock.Anything).Return(&mocks.Database{}, nil)
	})

	Describe("CreateDatabase", func() {
		Context("when creation is successful", func() {
			BeforeEach(func() {
				factory = db.NewIAVLDatabaseFactory()
			})

			It("should create a database without error", func() {
				dbPath := filepath.Join(os.TempDir(), "testDb")
				database, err := factory.CreateDatabase("test", dbPath)

				Expect(err).NotTo(HaveOccurred())
				Expect(database).NotTo(BeNil())
				// Additional expectations specific to IAVL database can be added here if necessary
			})
		})

		Context("when creation fails", func() {

			It("should return an error and nil database", func() {
				database, err := factory.CreateDatabase("test", "/path/to/db")

				Expect(err).To(HaveOccurred())
				Expect(database).To(BeNil())
				//Expect(err).To(MatchError("failed to initialize database: mkdir /path: permission denied"))
			})
		})
	})
})
