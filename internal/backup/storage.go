package backup

import (
	"fmt"
	"os"
)

const outputDir = "outputs"

func EnsureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
	return err
}

func SaveToLocal(filePath string, data []byte) error {
	fmt.Println("Saving to", filePath)
	err := os.WriteFile(filePath, data, 0644)
	return err
}
