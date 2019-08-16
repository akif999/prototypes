package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/akif999/prototypes/go/assets_life_exp/assets"
)

func main() {
	f, err := assets.Root.Open("/hoge.txt")
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	b, _ := ioutil.ReadAll(f)
	fmt.Println(string(b))
}
