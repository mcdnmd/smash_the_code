package bot

import (
	"smash_the_code/lib"
)

type Block struct {
	Cell1    Cell
	Cell2    Cell
	Rotation int
	Place    int
}

func (b *Block) ApplyRotation(rotation int) {
	if b.Cell1.Position == nil {
		panic("position1 not provided")
	}
	if rotation < 0 || rotation > 3 {
		panic("impossible rotation type")
	}
	b.Rotation = rotation
	switch rotation {
	case 0:
		b.Cell2.Position = &lib.V{b.Cell1.Position.X + 1, b.Cell1.Position.Y}
		b.Place = b.Cell1.Position.X
	case 1:
		b.Cell2.Position = &lib.V{b.Cell1.Position.X, b.Cell1.Position.Y - 1}
		b.Place = b.Cell1.Position.X
	case 2:
		b.Cell2.Position = &lib.V{b.Cell1.Position.X - 1, b.Cell1.Position.Y}
		b.Place = b.Cell1.Position.X
	case 3:
		b.Cell2.Position = &lib.V{b.Cell1.Position.X, b.Cell1.Position.Y + 1}
		b.Place = b.Cell1.Position.X
	}
}

func (b *Block) Clone() *Block {
	return &Block{
		Cell1:    *b.Cell1.Clone(),
		Cell2:    *b.Cell2.Clone(),
		Rotation: b.Rotation,
	}
}
