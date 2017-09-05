package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	GetVal1FirstFig = iota
	GetVal1st
	GetOpeByZero
	GetVal2FirstFig
	GetVal2nd
	GetEndByZero
)

type formula struct {
	val1 uint32
	ope  rune
	val2 uint32
}

func newFormula() *formula {
	return &formula{}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("require input")
		return
	}
	val, err := calcByFormula(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

func calcByFormula(str string) (uint32, error) {
	f := newFormula()
	f.parseFormula(str)
	if f.ope == '+' {
		return f.val1 + f.val2, nil
	} else if f.ope == '-' {
		return f.val1 - f.val2, nil
	} else {
		return 0, fmt.Errorf("invalid formula")
	}
}

func (f *formula) parseFormula(str string) error {
	var val1, val2 string
	var ope rune

	state := GetVal1FirstFig
	for _, r := range str {
		switch state {
		case GetVal1FirstFig:
			if '0' <= r && r <= '9' {
				if r == '0' {
					val1 = string(r)
					state = GetOpeByZero
				} else {
					val1 = string(r)
					state = GetVal1st
				}
			} else {
				return fmt.Errorf("invalid formula")
			}
		case GetVal1st:
			if '0' <= r && r <= '9' {
				val1 += string(r)
			} else if r == '+' || r == '-' {
				ope = r
				state = GetVal2FirstFig
			} else {
				return fmt.Errorf("invalid formula")
			}
		case GetOpeByZero:
			if r == '+' || r == '-' {
				ope = r
				state = GetVal2FirstFig
			} else {
				return fmt.Errorf("invalid formula")
			}
		case GetVal2FirstFig:
			if '0' <= r && r <= '9' {
				if r == '0' {
					val2 = string(r)
					state = GetEndByZero
				} else {
					val2 = string(r)
					state = GetVal2nd
				}
			} else {
				return fmt.Errorf("invalid formula")
			}
		case GetVal2nd:
			if '0' <= r && r <= '9' {
				val2 += string(r)
			} else {
				return fmt.Errorf("invalid formula")
			}
		case GetEndByZero:
			// GetEndByZeroとなった後は、あらゆる文字を受け付けない
			return fmt.Errorf("invalid formula")
		}
	}
	v, err := strconv.ParseUint(val1, 10, 32)
	if err != nil {
		return err
	}
	f.val1 = uint32(v)
	f.ope = ope
	v, err = strconv.ParseUint(val2, 10, 32)
	if err != nil {
		return err
	}
	f.val2 = uint32(v)

	return nil
}
