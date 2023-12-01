package helpers

import (
	"fmt"
	"os"
	"strings"
)

func ReadLines(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("cannot read the file: %s", fileName))
	}

	return strings.Split(string(file), "\n")
}
