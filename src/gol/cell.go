package gol

type Cell struct {
  alive bool
}

func NewCell(alive bool) *Cell {
  return &Cell{alive}
}

func (c *Cell) Next(count int) *Cell {
  return NewCell(count == 3 || (c.alive && count == 2))
}

func (c *Cell) Value() int {
  if (c.alive) {
    return 1
  } else {
    return 0
  }
}

func (c *Cell) IsAlive() bool {
  return c.alive
}
