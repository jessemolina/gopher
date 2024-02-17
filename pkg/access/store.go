package access

import (
	"os"
)


// KeyTofile writes a Key as bytes to the given path.
func KeyToFile(key []byte, path string) error {
	os.WriteFile(path, key, 0644)

	return nil
}

// TODO Build logic for storing and reading keys from vault.
func KeyToVault() {
}
