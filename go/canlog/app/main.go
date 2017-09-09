package main

import (
	"github.com/akif999/prototypes/go/canlog"
	"log"
	"os"
)

func main() {

	filename := os.Args[1]
	ids := []string{"1F3", "710", "1C8"}

	c := canlog.New()
	err := c.Parse(filename)
	if err != nil {
		log.Fatal(err)
	}

	ca := canlog.PickRecord(c, ids)
	ca.PrintLog(canlog.WHOLE)

	c = canlog.DelRecord(c, ids)
	c.PrintLog(canlog.WHOLE)
	c.PrintLog(canlog.DATA)

}
