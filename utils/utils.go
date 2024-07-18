package utils

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var serverPath = "/var/instagram/products/"

func AddPhotos() []string {
	entries, err := os.ReadDir("/home/mahdi/product_documents")
	if err != nil {
		log.Fatal(err)
	}

	var photos []string

	for _, e := range entries {
		if strings.Contains(e.Name(), ".jpg") || strings.Contains(e.Name(), ".JPG") || strings.Contains(e.Name(), ".png") {
			fmt.Println(e.Name())
			id := uuid.New()
			var format = ""
			if strings.Contains(e.Name(), ".jpg") || strings.Contains(e.Name(), ".JPG") || strings.Contains(e.Name(), ".jpeg") {
				format = ".jpg"
			} else {
				format = ".png"
			}
			var name = serverPath + id.String() + format
			photos = append(photos, filepath.Base(name))
			copy("/home/mahdi/product_documents/"+e.Name(), name)
		}
	}

	return photos
}

func AddDatasheet() string {
	entries, err := os.ReadDir("/home/mahdi/product_documents")
	if err != nil {
		log.Fatal(err)
	}
	var datasheet = ""

	for _, e := range entries {
		if strings.Contains(e.Name(), ".pdf") {
			fmt.Println(e.Name())
			id := uuid.New()
			datasheet = serverPath + id.String() + ".pdf"
			copy("/home/mahdi/product_documents/"+e.Name(), datasheet)
		}
	}
	return filepath.Base(datasheet)
}

func copy(src, dst string) (int64, error) {

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
