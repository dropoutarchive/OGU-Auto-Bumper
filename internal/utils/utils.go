package utils

import "os"

func ReadFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	return file
}
