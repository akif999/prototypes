package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	wait sync.WaitGroup
)

func main() {
	var files []string
	filepath.Walk("./input", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	for _, f := range files {
		wait.Add(1)
		go gzipByFile(f)
	}
	wait.Wait()
}

func gzipByFile(file string) error {
	b := new(bytes.Buffer)
	w := gzip.NewWriter(b)
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	w.Write(body)
	w.Close()
	// zip済みファイルを保存する
	s := strings.Split(file, "/")
	err = save(b, string(s[1]))
	if err != nil {
		return err
	}
	fmt.Println(file)
	wait.Done()
	return nil
}

func save(b *bytes.Buffer, filename string) error {
	f, err := os.OpenFile("./output/"+filename+".zip", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	f.Write(b.Bytes())
	f.Close()
	return nil
}
