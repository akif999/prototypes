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
	ch     uint8
	id     uint16
	dir    string
	stat   string
	dlc    uint8
	data   []byte
	remain string
}

type FormatFunc func(time float64, ch uint8, id uint16, dir string, dlc uint8, data []byte) string

// NewCanLog は、新たにCanlogオブジェクトを作成する
func NewCanLog() *CanLog {
	return &CanLog{
		formater: formatLog,
	}
}

// NewLog は、新たにlogオブジェクトを作成する
func newLog() *log {
	return &log{}
}

// Parseは、ファイルを解析してCanlogオブジェクトを作成する
func (cl *CanLog) Parse(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		l := newLog()

		fs := strings.Fields(scanner.Text())
		if fs[1] == "1" || fs[1] == "2" {
			l.time, err = strconv.ParseFloat(fs[0], 32)
			if err != nil {
				return err
			}
			c, err := strconv.ParseUint(fs[1], 10, 32)
			if err != nil {
				return err
			}
			l.ch = uint8(c)
			i, err := strconv.ParseUint(fs[2], 16, 32)
			if err != nil {
				return err
			}
			l.id = uint16(i)
			l.dir = fs[3]
			l.stat = fs[4]
			d, err := strconv.Atoi(fs[5])
			if err != nil {
				return err
			}
			l.dlc = uint8(d)
			for _, f := range fs[6 : l.dlc+6] {
				d, _ := strconv.ParseUint(f, 16, 32)
				l.data = append(l.data, byte(d))
			}
			l.remain = strings.Join(fs[l.dlc+7:l.dlc+15], " ")
			cl.logs = append(cl.logs, l)
		}
	}
	return scanner.Err()
}

// PickRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードのみを抽出したCanlogオブジェクトを返す
func PickRecord(c *CanLog, ids []uint16) *CanLog {
	nc := NewCanLog()
	for _, r := range c.logs {
		if isContains(r.id, ids) {
			nc.logs = append(nc.logs, r)
		}
	}
	return nc
}

// DelRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードを削除したCanlogオブジェクトを返す
func DelRecord(c *CanLog, ids []uint16) *CanLog {
	nc := NewCanLog()

	for _, r := range c.logs {
		if !isContains(r.id, ids) {
			nc.logs = append(nc.logs, r)
		}
	}
	return nc
}

// isContainsは、tgtが、配列arrに含まれているかをbooleanとして返す
func isContains(tgt uint16, arr []uint16) bool {
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

func formatLog(time float64, ch uint8, id uint16, dir string, dlc uint8, data []byte) string {
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
