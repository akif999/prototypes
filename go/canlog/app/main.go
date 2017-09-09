package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akif999/prototypes/go/canlog"
)

func main() {

	filename := os.Args[1]
	ids := []string{"1F3", "710", "1C8"}

	c := canlog.NewCanLog()
	err := c.Parse(filename)
	if err != nil {
		log.Fatal(err)
	}

	ca := canlog.PickRecord(c, ids)
	s := ca.String()
	fmt.Print(s)

	fmt.Println()

	c = canlog.DelRecord(c, ids)
	s = c.String()
	fmt.Print(s)
}
