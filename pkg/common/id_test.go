package common_test

import (
	"crypto/sha256"
	"encoding/hex"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ebanfa/skeleton/pkg/common"
)

// Replace with the actual import path

var _ = Describe("ID Generator and Hash Functions", func() {
	Describe("ProcessIDGenerator", func() {
		var generator *common.ProcessIDGenerator

		BeforeEach(func() {
			generator = common.NewProcessIDGenerator("test")
		})

		Describe("NewProcessIDGenerator", func() {
			It("should create a new ProcessIDGenerator with the given prefix", func() {
				Expect(generator).NotTo(BeNil())
			})
		})

		Describe("GenerateID", func() {
			It("should generate a unique process ID with the correct prefix", func() {
				id, err := generator.GenerateID()
				Expect(err).NotTo(HaveOccurred())
				Expect(id).To(HavePrefix("test-"))
				Expect(id).To(MatchRegexp(`^test-\d+$`))
			})

			It("should generate unique IDs for multiple calls", func() {
				id1, err := generator.GenerateID()
				Expect(err).NotTo(HaveOccurred())

				id2, err := generator.GenerateID()
				Expect(err).NotTo(HaveOccurred())

				Expect(id1).NotTo(Equal(id2))
			})

			It("should return an error when the prefix is empty", func() {
				emptyPrefixGenerator := common.NewProcessIDGenerator("")
				_, err := emptyPrefixGenerator.GenerateID()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("prefix cannot be empty"))
			})
		})

		Describe("Concurrent generation", func() {
			It("should generate unique IDs when called concurrently", func() {
				const goroutines = 100
				ids := make(chan string, goroutines)

				for i := 0; i < goroutines; i++ {
					go func() {
						id, err := generator.GenerateID()
						Expect(err).NotTo(HaveOccurred())
						ids <- id
					}()
				}

				generatedIDs := make(map[string]bool)
				for i := 0; i < goroutines; i++ {
					id := <-ids
					Expect(generatedIDs[id]).To(BeFalse(), "Duplicate ID generated: %s", id)
					generatedIDs[id] = true
				}
			})
		})
	})

	Describe("HashSHA256", func() {
		It("should generate the correct SHA-256 hash for a given input", func() {
			input := "test string"
			expectedHash := "d5579c46dfcc7f18207013e65b44e4cb4e2c2298f4ac457ba8f82743f31e930b"

			hash := common.HashSHA256(input)
			Expect(hash).To(Equal(expectedHash))
		})

		It("should generate different hashes for different inputs", func() {
			hash1 := common.HashSHA256("input1")
			hash2 := common.HashSHA256("input2")

			Expect(hash1).NotTo(Equal(hash2))
		})

		It("should generate the same hash for the same input", func() {
			input := "consistent input"
			hash1 := common.HashSHA256(input)
			hash2 := common.HashSHA256(input)

			Expect(hash1).To(Equal(hash2))
		})

		It("should generate a hash of the correct length", func() {
			hash := common.HashSHA256("any input")
			Expect(hash).To(HaveLen(64)) // SHA-256 produces a 32-byte (256-bit) hash, which is 64 characters when hex-encoded
		})

		It("should generate a valid hexadecimal string", func() {
			hash := common.HashSHA256("test input")
			Expect(hash).To(MatchRegexp("^[0-9a-f]{64}$"))
		})

		It("should match the output of crypto/sha256 package", func() {
			input := "test input for comparison"

			// Generate hash using the function under test
			hashFromFunction := common.HashSHA256(input)

			// Generate hash using crypto/sha256 directly
			hasher := sha256.New()
			hasher.Write([]byte(input))
			hashFromCrypto := hex.EncodeToString(hasher.Sum(nil))

			Expect(hashFromFunction).To(Equal(hashFromCrypto))
		})

		It("should handle empty input correctly", func() {
			hash := common.HashSHA256("")
			Expect(hash).NotTo(BeEmpty())
			Expect(hash).To(HaveLen(64))
		})

		It("should handle unicode input correctly", func() {
			hash := common.HashSHA256("こんにちは世界") // "Hello World" in Japanese
			Expect(hash).To(HaveLen(64))
			Expect(hash).To(MatchRegexp("^[0-9a-f]{64}$"))
		})
	})
})
