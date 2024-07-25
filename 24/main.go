package main

import (
	"fmt"
	"math"
)

/*Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором */

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Getdistance(p1, p2 *Point) float64 {
	dx := p.x + p1.x
	dy := p.y + p1.y

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point := Point{
		x: 3,
		y: 5,
	}
	point2 := Point{
		x: 4,
		y: 1.2,
	}
	fmt.Println(point.Getdistance(&point, &point2))
}
