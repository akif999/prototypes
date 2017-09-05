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
	WITHDIFFTIME
	DATA
)

type CanLog struct {
	logs []*log
}

// Canlog型はCanlogに合わせたデータ構造を持ったオブジェクト
type log struct {
	Prevtime float64
	Crnttime float64
	Difftime float64
	Ch       string
	Id       string
	Dir      string
	Stat     string
	Dlc      int
	Data     []int64
	Remain   string
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

	p := 0.000000
	for scanner.Scan() {
		l := new(log)

		fs := strings.Fields(scanner.Text())
		if fs[1] == "1" || fs[1] == "2" {
			l.Prevtime = p
			l.Crnttime, _ = strconv.ParseFloat(fs[0], 32)
			l.Difftime = l.Crnttime - p
			l.Ch = fs[1]
			l.Id = fs[2]
			l.Dir = fs[3]
			l.Stat = fs[4]
			l.Dlc, _ = strconv.Atoi(fs[5])
			for _, f := range fs[6 : l.Dlc+6] {
				d, _ := strconv.ParseInt(f, 16, 32)
				l.Data = append(l.Data, d)
			}
			l.Remain = strings.Join(fs[l.Dlc+7:l.Dlc+15], " ")
			c.logs = append(c.logs, l)
			p = l.Crnttime
		}
	}
	return scanner.Err()
}

// PickRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードのみを抽出したCanlogオブジェクトを返す
func PickRecord(c CanLog, ids []string) CanLog {
	var nc CanLog
	for _, r := range c.logs {
		if isContains(r.Id, ids) {
			nc.logs = append(nc.logs, r)
		}
	}
	return nc
}

// DelRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードを削除したCanlogオブジェクトを返す
func DelRecord(c CanLog, ids []string) CanLog {
	var nc CanLog

	for _, r := range c.logs {
		if !isContains(r.Id, ids) {
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
func (c *CanLog) PrintLog(opt int) {
	for _, r := range c.logs {
		switch opt {
		case WHOLE:
			s := convertDataToString(r, " ")
			fmt.Printf("%06f %s %03s %s %X %s\n", r.Crnttime, r.Ch, r.Id, r.Dir, r.Dlc, s)
		case WITHDIFFTIME:
			s := convertDataToString(r, " ")
			fmt.Printf("%06f %06f %s %03s %s %X %s\n", r.Difftime, r.Crnttime, r.Ch, r.Id, r.Dir, r.Dlc, s)
		case DATA:
			s := convertDataToString(r, "")
			fmt.Printf("%s\n", s)
		default:
			s := convertDataToString(r, " ")
			fmt.Printf("%06f %s %03s %s %X %s\n", r.Crnttime, r.Ch, r.Id, r.Dir, r.Dlc, s)
		}
	}
}

// convertDataToStringは、レコードの配列データから、sepで区切った文字列へ変換する
func convertDataToString(record *log, sep string) string {
	var data []string

	for _, d := range record.Data {
		data = append(data, fmt.Sprintf("%02X", d))
	}
	return strings.Join(data, sep)
}
