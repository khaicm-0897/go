package main

import (
	"fmt"
	"sync"
)

func main1() {
	s := "ぶん Golang"
	for _, v := range s {
		// v is of type rune
		fmt.Printf("Unicode code point : %U - character '%c' - binary %b - hex %X - Decimal %d\n ", v, v, v, v, v)
	}

	var cp ColoredPoint

	cp.X = 3.4
	cp.Point.X = 2
}

type Point struct{ X, Y float64 }

type ColoredPoint struct {
    // thuộc tính ẩn danh
    Point

    // thuộc tính bình thường
    Color string
}

type Mutex struct {}
func (m *Mutex) Lock()
func (m *Mutex) Unlock()

type Cache struct {
	m map[string]string
	sync.Mutex
}