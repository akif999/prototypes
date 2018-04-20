package main

import (
	"html/template"
	"os"
)

type TemplateSections struct {
	Sec1 string
	Sec2 string
	Sec3 string
	Sec4 string
}

func main() {
	tmpl := template.Must(template.ParseFiles("test.tmpl"))

	ts := &TemplateSections{
		"hoge",
		"fuga",
		"fuma",
		"funa",
	}

	err := tmpl.Execute(os.Stdout, ts)
	if err != nil {
		panic(err)
	}
}
