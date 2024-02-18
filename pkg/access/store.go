package access

import (
	"os"
)

// TODO Build logic for reading pvt keys from file and vault.

func ReadKeyFromFile() {}

func ReadKeyFromVault() {}

// KeyTofile writes a Key as bytes to the given path.
func WriteToFile(data []byte, path string) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func WriteToStdout(data []byte) error {
	_, err := os.Stdout.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func WriteToVault() {}

func WriteTokenToFile(token, path string) {}
