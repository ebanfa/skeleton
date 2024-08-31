package store

import (
	"path/filepath"
	"unicode"

	"github.com/ebanfa/skeleton/pkg/common"
)

func GenererateStorageInfo(name string, databasesDir string) (string, string) {
	databaseID := GenerateStoreId(name)

	databasePath := filepath.Join(databasesDir, databaseID+".db")
	return databaseID, databasePath
}

func GenerateStoreId(name string) string {
	return common.HashSHA256(name)
}

func CreateMultiStore(name, databasesDir string, storeFactory StoreFactory) (MultiStore, error) {

	internalStore, err := storeFactory.CreateStore(name)

	if err != nil {
		return nil, err
	}

	multiStore, err := NewMultiStore(internalStore, storeFactory)
	if err != nil {
		return nil, err
	}

	return multiStore, nil
}

func IsValidStoreName(s string) bool {
	// Check if the string is empty
	if len(s) == 0 {
		return false
	}

	// Check if the string consists only of whitespace or punctuation characters
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			// If the character is not a letter or number, return false
			return false
		}
	}

	// If the string passes all checks, return true
	return true
}
