package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string
	filepath.Walk("./input", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	err := gzipByFile(files)
	if err != nil {
		log.Fatal(err)
	}

}

func gzipByFile(files []string) error {
	for _, file := range files {
		b := new(bytes.Buffer)
		w := gzip.NewWriter(b)
		body, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		w.Write(body)
		w.Close()

		s := strings.Split(file, "/")
		err = save(b, string(s[1]))
		if err != nil {
			return err
		}
		fmt.Println(file)
	}
	return nil
}

func save(b *bytes.Buffer, filename string) error {
	zf, err := os.Create("./output/" + filename + ".zip")
	if err != nil {
		return err
	}
	zf.Write(b.Bytes())
	zf.Close()
	return nil
}
