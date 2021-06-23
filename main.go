package main

import (
  "fmt"
  "math"
  "math/rand"
  "time"
)

type geom interface {
  calc()
  init()
}

type point struct {
  x, y float64
}

type rectangle struct {
  points[] point
}

func (r *rectangle) calc() {
  fline :=  Line(r.points[0].x, r.points[1].x, r.points[0].y, r.points[1].x)
  sline := Line(r.points[1].x, r.points[2].x, r.points[1].y, r.points[2].x)
  tline := Line(r.points[2].x, r.points[0].x, r.points[2].y, r.points[0].x)
  fmt.Println("First Line: ", fline)
  fmt.Println("Second Line: ", sline)
  fmt.Println("Third Line: ", tline)
  p := (fline + sline + tline) / 2
  area := math.Sqrt(p * (p - fline) * (p - sline) * (p - tline))
  fmt.Println("Area: ", area)
}

func (r *rectangle) init() {
  for i := 0; i < 3; i++ {
    r.points[i] = point{rand.Float64(), rand.Float64()}
  }
}

type circle struct {
  radius float64
}

func (c *circle) init() {
  c.radius = rand.Float64()
}

func (c *circle) calc() {
  fmt.Println("Diametr: ", c.radius * 2)
  area := math.Pi * c.radius * c.radius
  fmt.Println("Area: ", area)
  len := 2 * math.Pi * c.radius
  fmt.Println("Circle length: ", len)
}

type piramid struct {
  points[] point
}

func (p *piramid) init() {
  for i := 0; i < 4; i++ {
    p.points[i] = point{rand.Float64(), rand.Float64()}
  }
}

func (p *piramid) calc() {
  fline :=  Line(p.points[0].x, p.points[1].x, p.points[0].y, p.points[1].x)
  sline := Line(p.points[1].x, p.points[2].x, p.points[1].y, p.points[2].x)
  tline := Line(p.points[2].x, p.points[3].x, p.points[2].y, p.points[3].x)
  fourthline := Line(p.points[3].x, p.points[0].x, p.points[3].y, p.points[0].x)
  fmt.Println("Perim: ", fline + sline + tline + fourthline)
}

func Line(X1, X2, Y1, Y2 float64) float64 {
  return math.Sqrt(math.Pow((X1 - X2), 2) + math.Pow((Y1 - Y2), 2))
}

func measure (g geom) {
  fmt.Println(g)
  g.calc()
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("---rectangle---")
  r := rectangle{make([]point, 3)}
  r.init()
  r.calc()
  fmt.Println("---circle---")
  c := circle{0}
  c.init()
  c.calc()
  fmt.Println("---piramid---")
  p := piramid{make([]point, 4)}
  p.init()
  p.calc()
}
