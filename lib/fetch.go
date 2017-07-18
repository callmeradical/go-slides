package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var webslidesURL = "https://github.com/webslides/WebSlides/releases/download/version/package"

func IsCached(version string) bool {
	home := os.Getenv("HOME")
	filename := fmt.Sprintf("webslides_%s", version)
	path := filepath.Join(home, ".go-slides", filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func ZipFileName(ver string) string {
	v := strings.Split(ver, ".")
	if len(v) < 3 {
		return "invalid version"
	}

	minor, _ := strconv.Atoi(v[1])

	if minor > 2 {
		return "webslides.zip"
	} else {
		return "release.zip"
	}
}

func GetDownloadLink(ver string) string {
	uri := strings.Replace(webslidesURL, "version", ver, -1)
	filename := ZipFileName(ver)
	return strings.Replace(uri, "package", filename, -1)

}
func FetchZip(ver string, cache bool) (string, error) {
	var filename string
	url := GetDownloadLink(ver)

	fmt.Println(cache)
	if !cache {
		filename = filepath.Join(os.TempDir(), fmt.Sprintf("webslides_%s.zip", ver))
	} else {
		filename = filepath.Join(os.Getenv("HOME"), ".go-slides", fmt.Sprintf("webslides_%s.zip", ver))
		if exists := IsCached(ver); exists {
			fmt.Println("File is cached... using cache")
			return filename, nil
		}
	}

	fmt.Println("Downloading file... skipping cache")

	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", err
	}

	return filename, nil

}

func FetchSlides(ver, path string, cache bool) error {
	tmp, err := FetchZip(ver, cache)

	newPath := filepath.Join(path, "webslides.zip")
	if !cache {
		err = os.Rename(tmp, newPath)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(newPath); os.IsExist(err) {
		tmp = newPath
	}

	err = unzip(tmp, path)
	if err != nil {
		return err
	}

	if !cache {
		return os.Remove(tmp)
	}
	return nil
}
