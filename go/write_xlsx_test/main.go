package main

import (
	"github.com/tealeg/xlsx"
)

const ()

var ()

func main() {
	xlsxFile := xlsx.NewFile()
	xlsxFile.AddSheet("sheet1")
	xlsxFile.Save("./hoge.xlsx")
}
