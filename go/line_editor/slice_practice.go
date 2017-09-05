package main

import (
	"bufio"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type File struct {
	lines   []string
	lineNum uint
}

type Command struct {
	command string
	arg1    uint
	arg2    string
}

const ()

var (
	fileptr = kingpin.Arg("file", "Editing filename").Required().File()

	file File
	cmd  Command
)

func main() {

	err := Init()
	if err != nil {
		log.Fatal(err)
	}

	// Intaractive routine
mainloop:
	for {
		fmt.Printf("Cmd>")
		fmt.Scanf("%s%d%s", &cmd.command, &cmd.arg1, &cmd.arg2)

		switch cmd.command {
		case "out":
			file.OutBuf()
		case "line":
			file.DispLineNum()
		case "mov":
			file.MoveLine(cmd.arg1)
		case "ins":
			file.InsertStr(cmd.arg1, cmd.arg2)
		case "del":
			file.DeleteStr(cmd.arg1, cmd.arg2)
		case "save":
			file.SaveBuf()
		case "quit":
			break mainloop
		default:
			fmt.Println("command not found")
		}
		cmd.ClearCmd()
	}

}

func Init() error {
	kingpin.Parse()
	defer func() {
		(*fileptr).Close()
	}()

	scanner := bufio.NewScanner(*fileptr)
	for scanner.Scan() {
		file.lines = append(file.lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	file.lineNum = 0
	return nil
}

func (f *File) MoveLine(line uint) {
	if int(line) > len(file.lines) {
		fmt.Println("out of range")
		return
	}
	f.lineNum = line
}

func (f *File) DispLineNum() {
	fmt.Println(file.lineNum)
}

func (file *File) InsertStr(pos uint, exp string) {
	if int(pos) > len(file.lines[file.lineNum]) {
		fmt.Println("out of range")
		return
	}

	buf := []byte(file.lines[file.lineNum])
	ptr := pos
	str := exp
	for _, c := range []byte(str) {
		// 挿入のためのスライス操作
		buf = append(buf, 0)
		copy(buf[ptr+1:], buf[ptr:])
		buf[ptr] = c
		ptr++
	}
	file.lines[file.lineNum] = string(buf)
}

func (file *File) DeleteStr(pos uint, delnum string) {
	num, _ := strconv.Atoi(delnum)
	if int(pos)+num > len(file.lines[file.lineNum]) {
		fmt.Println("out of range")
		return
	}
	pos++
	buf := []byte(file.lines[file.lineNum])
	// 削除のためのスライス操作
	buf = append(buf[:pos-1], buf[pos+uint(num)-1:]...)
	file.lines[file.lineNum] = string(buf)
}

func (f *File) OutBuf() {
	for _, line := range f.lines {
		fmt.Println(line)
	}
}

func (f *File) SaveBuf() {
	var content []byte
	for _, line := range f.lines {
		for _, c := range []byte(line) {
			content = append(content, c)
		}
		content = append(content, byte('\n'))
	}
	ioutil.WriteFile("out.txt", content, os.ModePerm)
}

func (c *Command) ClearCmd() {
	c.command = ""
	c.arg1 = 0
	c.arg2 = ""
}
