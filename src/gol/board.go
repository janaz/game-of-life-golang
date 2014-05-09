package gol

import "fmt"

type Board struct {
  board map[string]*Cell
}

func NewBoard() *Board {
  empty := map[string]*Cell {}
  return &Board{empty}
}

func (b *Board) SetAlive(p *Point) {
  b.board[p.ToString()] = NewCell(true)
}

func (b *Board) Transfer(next *Board, p *Point) {
  c:= b.GetCell(p).Next(b.AliveNeighbors(p))

  if (c.IsAlive()) {
    next.SetAlive(p)
  }

}
func (b *Board) Next() *Board {
  next := NewBoard()
  for k,_ := range ( b.board ) {
    p := PointFromString(k)
    go b.Transfer(next, p)
    for _,n := range p.Neighbors() {
      go b.Transfer(next, n)
    }

  }
  return next
}

func (b *Board) AliveNeighbors(p *Point) int {
  total :=0
  for _,n := range p.Neighbors() {
    total += b.GetCell(n).Value()
  }
  return total
}

func (b *Board) GetCell(p *Point) *Cell {
  c := b.board[p.ToString()]
  if (c != nil) {
    return c
  } else {
    return NewCell(false)
  }
}

func (b *Board) Print(w,h int) {
  fmt.Print("-----\n")
  for y:= 0 ; y<h; y++ {
    for x:= 0 ; x<w; x++ {
      p := NewPoint(x,y)
      if b.GetCell(p).IsAlive() {
        fmt.Print("X")
      } else {
        fmt.Print(".")
      }
    }
    fmt.Print("\n")
  }
}
