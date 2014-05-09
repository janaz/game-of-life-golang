package main

import "testing"
import "gol"


func TestAliveCell(t *testing.T) {
    cell := gol.NewCell(true)
    if (cell.Next(1).IsAlive()) {
      t.Fatal("Alive cell with 1 neighbour should die")
    }
    if (!cell.Next(2).IsAlive()) {
      t.Fatal("Alive cell with 2 neighbour should not die")
    }
    if (!cell.Next(3).IsAlive()) {
      t.Fatal("Alive cell with 3 neighbour should not die")
    }
}

func TestDeadCell(t *testing.T) {
    cell := gol.NewCell(false)
    if (cell.Next(1).IsAlive()) {
      t.Fatal("Dead cell with 1 neighbour should stay dead")
    }
    if (cell.Next(4).IsAlive()) {
      t.Fatal("Dead cell with 4 neighbour should stay dead")
    }
    if (cell.Next(5).IsAlive()) {
      t.Fatal("Dead cell with 5 neighbour should stay dead")
    }
    if (!cell.Next(3).IsAlive()) {
      t.Fatal("Dead cell with 3 neighbour should become alive")
    }
}

func TestBoard(t *testing.T) {
    board := gol.NewBoard()
    board.SetAlive(gol.NewPoint(2,3))
    board.SetAlive(gol.NewPoint(0,3))
    if (!board.GetCell(gol.NewPoint(2,3)).IsAlive()) {
      t.Fatal("Expected alive cell on 2,3")
    }
    if (board.AliveNeighbors(gol.NewPoint(1,3)) != 2) {
      t.Fatal("Expected neighbors for 1,3 is 2")
    }
    if (board.AliveNeighbors(gol.NewPoint(3,3)) != 1) {
      t.Fatal("Expected neighbors for 3,3 is 1")
    }
}

func TestNextBoard(t *testing.T) {
  board := gol.NewBoard()
  board.SetAlive(gol.NewPoint(1,2))
  board.SetAlive(gol.NewPoint(2,2))
  board.SetAlive(gol.NewPoint(3,2))
  next := board.Next()
  if (!next.GetCell(gol.NewPoint(2,1)).IsAlive()) {
    t.Fatal("Expected alive cell on 2,1")
  }
  if (!next.GetCell(gol.NewPoint(2,2)).IsAlive()) {
    t.Fatal("Expected alive cell on 2,2")
  }
  if (!next.GetCell(gol.NewPoint(2,3)).IsAlive()) {
    t.Fatal("Expected alive cell on 2,3")
  }

}
