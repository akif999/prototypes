package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type columnNumber int

const (
	firstColumn  columnNumber = 3
	secondColumn columnNumber = 2
	thirdColumn  columnNumber = 1
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid argument")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	values := []int{}
	groups := map[int][]int{}
	values, err = parse(file)
	if err != nil {
		log.Fatal(err)
	}

	groups, err = groupByColumn(values, firstColumn)
	if err != nil {
		log.Fatal(err)
	}
	values = sortByNumGroups(groups)
	groups, err = groupByColumn(values, secondColumn)
	if err != nil {
		log.Fatal(err)
	}
	values = sortByNumGroups(groups)
	fmt.Println(values)
}

func parse(input io.Reader) ([]int, error) {
	ret := []int{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		vs := strings.Split(scanner.Text(), ",")
		for _, v := range vs {
			s, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return []int{}, err
			}
			ret = append(ret, int(s))
		}
	}
	return ret, nil
}

func groupByColumn(vals []int, col columnNumber) (map[int][]int, error) {
	ret := map[int][]int{}
	for _, v := range vals {
		if col == firstColumn {
			c := ((v % 100) % 10)
			ret[c] = append(ret[c], v)
		} else if col == secondColumn {
			c := ((v % 100) / 10)
			ret[c] = append(ret[c], v)
		} else {
			return map[int][]int{}, fmt.Errorf("invalid column number")
		}
	}
	return ret, nil
}

func sortByNumGroups(groups map[int][]int) []int {
	keys := getKeys(groups)
	sort.Ints(keys)
	ret := []int{}
	for _, k := range keys {
		ret = append(ret, groups[k]...)
	}
	return ret
}

func getKeys(mp map[int][]int) []int {
	ret := []int{}
	for k, _ := range mp {
		ret = append(ret, k)
	}
	return ret
}
