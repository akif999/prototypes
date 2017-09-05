package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("require input")
	}
	pp, err := os.Create("scan_lines_cs.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(pp)
	fp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	scanner.Split(scanLinesCustom)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	pprof.StopCPUProfile()
}

// scanLinesCustom is a splitFunc corresponding to all line feed codes of CRLF, LF, CR
func scanLinesCustom(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	i, j := bytes.IndexByte(data, '\n'), bytes.IndexByte(data, '\r')
	if i < j {
		// if LF
		if i >= 0 {
			return i + 1, data[0:i], nil
		} else {
			// if CRLF
			if j < len(data)-1 && isLF(data[j+1]) {
				return j + 2, data[0:j], nil
			}
			// if CR
			return j + 1, data[0:j], nil
		}
	} else if j < i {
		if j >= 0 {
			// if CRLF
			if j < len(data)-1 && isLF(data[j+1]) {
				return j + 2, data[0:j], nil
			}
			// if CR
			return j + 1, data[0:j], nil
			// if LF
		} else {
			return i + 1, data[0:i], nil
		}
	} else {
		// this case is only "i == -1 && j == -1"
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// isLF returns weather b is LF
func isLF(b byte) bool {
	if b == '\n' {
		return true
	}
	return false
}
