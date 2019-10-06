package main

import "fmt"

func main() {
	input := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	splitter := []uint32{2, 4, 7, 9}
	fmt.Println(splitBy(splitter, input))
}

func splitBy(splitter, src []uint32) (ret [][]uint32) {
	last := uint32(0)
	for _, s := range splitter {
		ret = append(ret, src[last:s])
		last = s
	}
	ret = append(ret, src[last:])
	return ret
}
