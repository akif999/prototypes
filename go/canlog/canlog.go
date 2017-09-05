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

// Canlog型はCanlogに合わせたデータ構造を持ったオブジェクト
type Canlog struct {
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

type Canlogs []Canlog

// Newは、新たにCanlogオブジェクトを作成する
func New() Canlogs {
	var c Canlogs
	return c
}

// Parseは、ファイルを解析してCanlogオブジェクトを作成する
func (records *Canlogs) Parse(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	p := 0.000000
	for scanner.Scan() {
		c := new(Canlog)

		fs := strings.Fields(scanner.Text())
		if fs[1] == "1" || fs[1] == "2" {
			c.Prevtime = p
			c.Crnttime, _ = strconv.ParseFloat(fs[0], 32)
			c.Difftime = c.Crnttime - p
			c.Ch = fs[1]
			c.Id = fs[2]
			c.Dir = fs[3]
			c.Stat = fs[4]
			c.Dlc, _ = strconv.Atoi(fs[5])
			for _, f := range fs[6 : c.Dlc+6] {
				d, _ := strconv.ParseInt(f, 16, 32)
				c.Data = append(c.Data, d)
			}
			c.Remain = strings.Join(fs[c.Dlc+7:c.Dlc+15], " ")
			*records = append(*records, *c)
			p = c.Crnttime
		}
	}
	return scanner.Err()
}

// PickRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードのみを抽出したCanlogオブジェクトを返す
func PickRecord(records Canlogs, ids []string) Canlogs {
	var nc Canlogs
	for _, r := range records {
		if isContains(r.Id, ids) {
			nc = append(nc, r)
		}
	}
	return nc
}

// DelRecordは、引数で渡したCanlogオブジェクトからidsで渡したIDが存在するレコードを削除したCanlogオブジェクトを返す
func DelRecord(records Canlogs, ids []string) Canlogs {
	var nc Canlogs

	for _, r := range records {
		if !isContains(r.Id, ids) {
			nc = append(nc, r)
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

// PrintLogは、Canlogsオブジェクトを1レコードずつフォーマットして標準出力へ出力する
func (records *Canlogs) PrintLog(opt int) {
	for _, r := range *records {
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
func convertDataToString(record Canlog, sep string) string {
	var data []string

	for _, d := range record.Data {
		data = append(data, fmt.Sprintf("%02X", d))
	}
	return strings.Join(data, sep)
}
