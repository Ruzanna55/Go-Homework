package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}

	return nil
}

func main() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("test1.zip", flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}

	defer file.Close()

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	directory, err := os.Open("./folderOfFiles")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer directory.Close()

	files, _ := directory.Readdirnames(0) // 0 to read  all files and folders

	for _, filename := range files {
		path := "./folderOfFiles/" + filename
		fmt.Println(path)
		if err := appendFiles(path, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
	}

}
