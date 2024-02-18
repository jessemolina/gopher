package access

import (
	"os"
)


// TODO Build logic for reading pvt keys from file and vault.

// KeyTofile writes a Key as bytes to the given path.
func KeyToFile(key []byte, path string) error {
	os.WriteFile(path, key, 0644)

	return nil
}

func KeyToVault() {
}
