package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const configDir = ".go-slides"

func MkConfigDir() (bool, error) {
	home := os.Getenv("HOME")
	path := filepath.Join(home, configDir)
	if err := os.Mkdir(path, 0755); err != nil {
		log.Println(err.Error())
		return false, err
	}

	fmt.Println(fmt.Sprintf("Created new config directory @ %s", path))
	return true, nil
}
