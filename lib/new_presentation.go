package lib

import (
	"os"
	"path/filepath"
)

func Presentation(version, name string, cache bool) error {

	err := os.Mkdir(name, 0700)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	newPath := filepath.Join(wd, name)

	return FetchSlides(version, newPath, cache)
}
