package canlog

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	WHOLE = iota
	DATA
)

type CanLog struct {
	logs []*log
}

// Canlog型はCanlogに合わせたデータ構造を持ったオブジェクト
type log struct {
	time   float64
	ch     string
	id     string
	dir    string
	stat   string
	dlc    int
	data   []int64
	remain string
}

// Newは、新たにCanlogオブジェクトを作成する
func New() CanLog {
	var c CanLog
	return c
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
				d, _ := strconv.ParseInt(f, 16, 32)
				l.data = append(l.data, d)
			}
			l.remain = strings.Join(fs[l.dlc+7:l.dlc+15], " ")
			c.logs = append(c.logs, l)
		}
	}
	return scanner.Err()
}

// PickRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードのみを抽出したCanlogオブジェクトを返す
func PickRecord(c CanLog, ids []string) CanLog {
	var nc CanLog
	for _, r := range c.logs {
		if isContains(r.id, ids) {
			nc.logs = append(nc.logs, r)
		}
	}
	return nc
}

// DelRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードを削除したCanlogオブジェクトを返す
func DelRecord(c CanLog, ids []string) CanLog {
	var nc CanLog

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

type formatFunc func(time, ch, id, dir, dlc string, data []byte) string

// PrintLogは、CanLogオブジェクトを1レコードずつフォーマットして標準出力へ出力する
func (c *CanLog) PrintLog(opt int) {
	for _, r := range c.logs {
		switch opt {
		case WHOLE:
			s := DataToString(r, " ")
			fmt.Printf("%06f %s %03s %s %X %s\n", r.time, r.ch, r.id, r.dir, r.dlc, s)
		case DATA:
			s := DataToString(r, "")
			fmt.Printf("%s\n", s)
		default:
			s := DataToString(r, " ")
			fmt.Printf("%06f %s %03s %s %X %s\n", r.time, r.ch, r.id, r.dir, r.dlc, s)
		}
	}
}

// DataToStringは、レコードの配列データから、sepで区切った文字列へ変換する
func DataToString(record *log, sep string) string {
	var data []string

	for _, d := range record.data {
		data = append(data, fmt.Sprintf("%02X", d))
	}
	return strings.Join(data, sep)
}
