package gol

import "fmt"
import "sync"

type Board struct {
  mu sync.Mutex
  board map[string]*Cell
}

func NewBoard() *Board {
  empty := map[string]*Cell {}
  m := sync.Mutex{}
  return &Board{m, empty}
}

func (b *Board) SetAlive(p *Point) {
  b.mu.Lock()
  b.board[p.ToString()] = NewCell(true)
  b.mu.Unlock()
}

func (b *Board) Transfer(nr int, next *Board, cp chan Point, w *sync.WaitGroup) {
  for pv := range cp{
    p := &pv
    //fmt.Printf("%v thread reporting %v\n",nr,p)
    c:= b.GetCell(p).Next(b.AliveNeighbors(p))
    if (c.IsAlive()) {
      next.SetAlive(p)
    }
  }
  w.Done()
}

func (b *Board) Next() *Board {
  next := NewBoard()
  w := &sync.WaitGroup{}
  w.Add(4)
  points_channel := make(chan Point)
  //4 threads
  go b.Transfer(1, next, points_channel, w)
  go b.Transfer(2, next, points_channel, w)
  go b.Transfer(3, next, points_channel, w)
  go b.Transfer(4, next, points_channel, w)

  for k,_ := range ( b.board ) {
    p:=PointFromString(k)
    points_channel <- *p
    for _,n := range p.Neighbors() {
      points_channel <- *n
    }
  }
  close(points_channel)
  w.Wait()
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
