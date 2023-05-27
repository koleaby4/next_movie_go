package fileutils

import (
	"log"
	"os"
)

func GetFileContents(filePath string) ([]byte, error) {
	contents, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatalf("unable to read file %v: %v", filePath, err)
    }

    return contents, nil
}
