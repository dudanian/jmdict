// Actually obtaining and updating the jmdict file
package jmdict

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const jmdictFile = "JMdict_e.gz"
const jmdictBaseUrl = "http://ftp.monash.edu/pub/nihongo/"

var jmdictPath = downloadPath()

func downloadPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// check if there is a downloads folder in home
	dirs := []string{"Downloads", "downloads"}
	for _, dir := range dirs {
		path := filepath.Join(home, dir)

		if fi, err := os.Stat(path); !os.IsNotExist(err) {
			fmt.Println(path)
			if fi.IsDir() {
				home = path
				break
			}
		}
	}

	return filepath.Join(home, jmdictFile)
}

func readRaw() []byte {
	if _, err := os.Stat(jmdictPath); os.IsNotExist(err) {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(jmdictPath)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func Download() {
	url := jmdictBaseUrl + jmdictFile
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(jmdictPath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func NewFileReader() (*os.File, error) {
	return os.Open(jmdictPath)
}

func NewGzipReader() (*gzip.Reader, error) {
	fi, err := NewFileReader()
	if err != nil {
		return nil, err
	}

	return gzip.NewReader(fi)
}
