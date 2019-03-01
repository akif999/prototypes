package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req, err := http.NewRequest("GET", "http://localhost:18888/basic", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("user", "pass")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(getContent(resp)))
}

func getContent(resp *http.Response) []byte {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
