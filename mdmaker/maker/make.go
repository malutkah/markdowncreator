package maker

import (
	"fmt"
	"os"
)

func CreateMarkdown(text string, fileName string) error {
	data := []byte(text)
	err := os.WriteFile(fileName+".md", data, 0644)
	if err != nil {
		err = fmt.Errorf("\nERROR: CreateMarkdown: %v\n\n", err)
		return err
	}

	return nil
}
