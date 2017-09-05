package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
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

	err := gzipByFile(files)
	if err != nil {
		log.Fatal(err)
	}
}

func gzipByFile(files []string) error {
	cpus := runtime.NumCPU()
	s := make(chan uint, cpus)
	f := make(chan string, len(files))
	e := make(chan error, len(files))
	for i := 0; i < len(files); i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			s <- 1
			file := <-f
			b := new(bytes.Buffer)
			w := gzip.NewWriter(b)
			body, err := ioutil.ReadFile(file)
			if err != nil {
				e <- err
			}
			<-s
			w.Write(body)
			w.Close()

			s := strings.Split(file, "/")
			err = save(b, string(s[1]))
			if err != nil {
				e <- err
			}
			e <- nil
			fmt.Println(file)
		}()
	}
	go func() {
		for _, file := range files {
			f <- file
		}
	}()
	wait.Wait()
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
