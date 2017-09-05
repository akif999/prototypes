package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestScanLinesCustom(t *testing.T) {
	type tp struct {
		input string
		want  string
		got   string
	}
	tests := []tp{
		// 1~4 : no lb or one at end
		tp{
			input: "abcdefghijkl",
			want:  "abcdefghijkl\n",
		},
		tp{
			input: "abcdefghijkl\r\n",
			want:  "abcdefghijkl\n",
		},
		tp{
			input: "abcdefghijkl\n",
			want:  "abcdefghijkl\n",
		},
		tp{
			input: "abcdefghijkl\r",
			want:  "abcdefghijkl\n",
		},
		// 5~7 : top
		tp{
			input: "\r\nabcdefghijkl",
			want:  "\nabcdefghijkl\n",
		},
		tp{
			input: "\nabcdefghijkl",
			want:  "\nabcdefghijkl\n",
		},
		tp{
			input: "\rabcdefghijkl",
			want:  "\nabcdefghijkl\n",
		},
		// 8~11 : top and buttom
		tp{
			input: "\r\nabcdefghijkl\r\n",
			want:  "\nabcdefghijkl\n",
		},
		tp{
			input: "\nabcdefghijkl\n",
			want:  "\nabcdefghijkl\n",
		},
		tp{
			input: "\rabcdefghijkl\r",
			want:  "\nabcdefghijkl\n",
		},
		tp{
			input: "\r\nabcdefghijkl\n",
			want:  "\nabcdefghijkl\n",
		},
		// 12 : only crlf
		tp{
			input: "abc\r\ndef\r\nghi\r\njkl\r\n",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 13: only lf
		tp{
			input: "abc\ndef\nghi\njkl\n",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 14 : only cr
		tp{
			input: "abc\rdef\rghi\rjkl\r",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 15 :  lf in crlf
		tp{
			input: "abc\r\ndef\nghi\r\njkl\r\n",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 16 : cr in crlf
		tp{
			input: "abc\r\ndef\rghi\r\njkl\r\n",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 17 : crlf in lf
		tp{
			input: "abc\ndef\nghi\r\njkl\n",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 18 : cr in lf
		tp{
			input: "abc\ndef\nghi\rjkl\n",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 19 : crlf in cr
		tp{
			input: "abc\rdef\rghi\r\njkl\r",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 20 : lf in cr
		tp{
			input: "abc\rdef\rghi\njkl\r",
			want:  "abc\ndef\nghi\njkl\n",
		},
		// 21 : crlf duplicate
		tp{
			input: "abc\r\ndef\r\nghi\r\n\r\njkl\r\n",
			want:  "abc\ndef\nghi\n\njkl\n",
		},
		// 22 : lf duplicate
		tp{
			input: "abc\ndef\nghi\n\njkl\n",
			want:  "abc\ndef\nghi\n\njkl\n",
		},
		// 23 : cr duplicate
		tp{
			input: "abc\rdef\rghi\r\rjkl\r",
			want:  "abc\ndef\nghi\n\njkl\n",
		},

		// 24 : cr duplicate in crlf
		tp{
			input: "abc\r\ndef\r\nghi\r\r\njkl\r\n",
			want:  "abc\ndef\nghi\n\njkl\n",
		},
		// 25 : lf duplicate in crlf
		tp{
			input: "abc\r\ndef\r\nghi\r\n\njkl\r\n",
			want:  "abc\ndef\nghi\n\njkl\n",
		},
		// 26 : crlf duplicate in crlf
		tp{
			input: "abc\r\ndef\r\nghi\r\n\r\njkl\r\n",
			want:  "abc\ndef\nghi\n\njkl\n",
		},
	}
	for i, e := range tests {
		e.got = ""
		scanner := bufio.NewScanner(strings.NewReader(e.input))
		scanner.Split(scanLinesCustom)
		for scanner.Scan() {
			e.got += scanner.Text() + "\n"
		}
		if e.got != e.want {
			t.Errorf("case %d got :\n%s\nwant :\n%s\n", i+1, e.got, e.want)
		}
	}
}
