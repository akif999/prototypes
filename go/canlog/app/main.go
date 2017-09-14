package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akif999/prototypes/go/canlog"
)

func main() {

	if len(os.Args) < 2 {
		panic("require argument")
	}
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
	c.Format(formatFuncCustom)
	s = c.String()
	fmt.Print(s)
}

func formatFuncCustom(time float64, ch, id, dir string, dlc int, data []byte) string {
	return fmt.Sprintf("%06f %03s\n", time, id)
}
