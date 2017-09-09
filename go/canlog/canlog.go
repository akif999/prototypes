package canlog

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CanLog struct {
	logs     []*log
	formater FormatFunc
}

// Canlog型はCanlogに合わせたデータ構造を持ったオブジェクト
type log struct {
	time   float64
	ch     string
	id     string
	dir    string
	stat   string
	dlc    int
	data   []byte
	remain string
}

type FormatFunc func(time float64, ch, id, dir string, dlc int, data []byte) string

// Newは、新たにCanlogオブジェクトを作成する
func NewCanLog() *CanLog {
	return &CanLog{
		formater: formatLog,
	}
}

// Parseは、ファイルを解析してCanlogオブジェクトを作成する
func (c *CanLog) Parse(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		l := new(log)

		fs := strings.Fields(scanner.Text())
		if fs[1] == "1" || fs[1] == "2" {
			l.time, _ = strconv.ParseFloat(fs[0], 32)
			l.ch = fs[1]
			l.id = fs[2]
			l.dir = fs[3]
			l.stat = fs[4]
			l.dlc, _ = strconv.Atoi(fs[5])
			for _, f := range fs[6 : l.dlc+6] {
				d, _ := strconv.ParseUint(f, 16, 32)
				l.data = append(l.data, byte(d))
			}
			l.remain = strings.Join(fs[l.dlc+7:l.dlc+15], " ")
			c.logs = append(c.logs, l)
		}
	}
	return scanner.Err()
}

// PickRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードのみを抽出したCanlogオブジェクトを返す
func PickRecord(c *CanLog, ids []string) *CanLog {
	nc := NewCanLog()
	for _, r := range c.logs {
		if isContains(r.id, ids) {
			nc.logs = append(nc.logs, r)
		}
	}
	return nc
}

// DelRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードを削除したCanlogオブジェクトを返す
func DelRecord(c *CanLog, ids []string) *CanLog {
	nc := NewCanLog()

	for _, r := range c.logs {
		if !isContains(r.id, ids) {
			nc.logs = append(nc.logs, r)
		}
	}
	return nc
}

// isContainsは、文字列tgtが、配列arrに含まれているかをbooleanとして返す
func isContains(tgt string, arr []string) bool {
	for _, e := range arr {
		if e == tgt {
			return true
		}
	}
	return false
}

// PrintLogは、CanLogオブジェクトを1レコードずつフォーマットして標準出力へ出力する
func (c *CanLog) String() string {
	s := ""
	for _, l := range c.logs {
		s += c.formater(l.time, l.ch, l.id, l.dir, l.dlc, l.data)
	}
	return s
}

func formatLog(time float64, ch, id, dir string, dlc int, data []byte) string {
	s := bytesToSeparatedString(data, " ")
	return fmt.Sprintf("%06f %s %03s %s %X %s\n", time, ch, id, dir, dlc, s)
}

// bytesToSeparatedStringは、byte sliceをsepで区切った文字列へ変換する
func bytesToSeparatedString(bytes []byte, sep string) string {
	var data []string

	for _, b := range bytes {
		data = append(data, fmt.Sprintf("%02X", b))
	}
	return strings.Join(data, sep)
}

func (c *CanLog) Format(formater FormatFunc) {
	c.formater = formater
}
