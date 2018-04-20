package main

import (
	"html/template"
	"os"
)

type Greeting struct {
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("test.tmpl"))

	g := &Greeting{"Hello, World."}

	err := tmpl.Execute(os.Stdout, g)
	if err != nil {
		panic(err)
	}
}
