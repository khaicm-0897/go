package main

import (
    "fmt"
    "math"
)

type hinh_hoc interface {
    dien_tich() float64
    chu_vi() float64
}

type chu_nhat struct {
    width, height float64
}
type hinh_tron struct {
    radius float64
}

func (r chu_nhat) dien_tich() float64 {
    return r.width * r.height
}
func (r chu_nhat) chu_vi() float64 {
    return 2*r.width + 2*r.height
}

func (c hinh_tron) dien_tich() float64 {
    return math.Pi * c.radius * c.radius
}
func (c hinh_tron) chu_vi() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g hinh_hoc) {
    fmt.Println(g)
    fmt.Println(g.dien_tich())
    fmt.Println(g.chu_vi())
}

func main3() {
    r := chu_nhat{width: 3, height: 4}
    c := hinh_tron{radius: 5}

    measure(r)
    measure(c)
}