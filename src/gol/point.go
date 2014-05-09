package gol

import "strings"
import "strconv"

type Point struct {
  x,y int
}

func NewPoint(x,y int) *Point {
  return &Point{x,y}
}

func (p *Point) ToString() string {
  return strings.Join([]string{strconv.Itoa(p.x), strconv.Itoa(p.y)}, ",")
}

func PointFromString(k string) *Point {
  el := strings.Split(k, ",")
  x, _ := strconv.Atoi(el[0])
  y, _ := strconv.Atoi(el[1])
  return NewPoint(x, y)
}

func (p *Point) Neighbors() []*Point {
  return []*Point{
    NewPoint(p.x-1, p.y-1),NewPoint(p.x, p.y-1),NewPoint(p.x+1, p.y-1),
    NewPoint(p.x-1, p.y),NewPoint(p.x+1, p.y),
    NewPoint(p.x-1, p.y+1),NewPoint(p.x, p.y+1),NewPoint(p.x+1, p.y+1),
  }
}
