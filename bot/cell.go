package bot

import "smash_the_code/lib"

type Cell struct {
	Color    int
	Position *lib.V
}

func (c *Cell) Clone() *Cell {
	return &Cell{c.Color, c.Position.Clone()}
}
