package reader

import (
	"os"
)

// ReadFile - read go.mod file
func ReadFile(path string) (data []byte, err error) {
	data, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}
