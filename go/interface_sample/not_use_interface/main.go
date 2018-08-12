package main

import (
	"fmt"
	"math"
	"reflect"
)

type triangle struct {
	base   uint32
	height uint32
}

type rectangle struct {
	height uint32
	width  uint32
}

type circle struct {
	radius uint32
}

func main() {
	t := triangle{base: 10, height: 20}
	r := rectangle{height: 20, width: 10}
	c := circle{radius: 10}
	ta, _ := calcArea(t)
	ra, _ := calcArea(r)
	ca, _ := calcArea(c)
	fmt.Printf("area (triangle) : %d\n", ta)
	fmt.Printf("area (rectangle) : %d\n", ra)
	fmt.Printf("area (circle) : %d\n", ca)
}

func calcArea(s interface{}) (uint32, error) {
	t := reflect.TypeOf(s)
	if t == reflect.TypeOf(triangle{}) {
		return s.(triangle).base * s.(triangle).height, nil
	} else if t == reflect.TypeOf(rectangle{}) {
		return s.(rectangle).height * s.(rectangle).width, nil
	} else if t == reflect.TypeOf(circle{}) {
		return uint32(float32(s.(circle).radius) * math.Pi), nil
	} else {
		return 0, fmt.Errorf("invalid input type")
	}
}
