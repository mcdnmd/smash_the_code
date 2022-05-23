package bot

import "smash_the_code/lib"

func GetAllPossibleBlocks(block Block) []*Block {
	blocks := make([]*Block, 24)
	c := 0
	t := 0
	for i := 0; i < len(blocks); i++ {
		blocks[i] = new(Block)
		blocks[i].Cell1 = *block.Cell1.Clone()
		blocks[i].Cell2 = *block.Cell2.Clone()
		blocks[i].Cell1.Position = &lib.V{X: i % 6, Y: 0}
		if c >= 6 {
			t += 1
			c = 0
		}
		blocks[i].ApplyRotation(t)
		c += 1
	}
	return blocks
}
