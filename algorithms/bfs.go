package algorithms

import (
	"smash_the_code/bot"
	"smash_the_code/lib"
)

func BFS(gridPtr *[][]bot.Cell, startCell bot.Cell, visitedPtr *[]bool) (chain []bot.Cell) {
	color := startCell.Color
	if color < 1 {
		return chain
	}

	grid := *gridPtr
	visited := *visitedPtr

	queue := make([]lib.V, 0)
	queue = append(queue, *startCell.Position)

	for len(queue) > 0 {
		cellPosition := queue[0]
		queue = queue[1:]

		area4 := cellPosition.Area4()
		for i := range area4 {
			position := area4[i]
			if position.InMap(bot.MapWidth, bot.MapHeight) {
				hashPosition := bot.GetHashPosition(*position)
				if (grid[position.Y][position.X].Color == color || grid[position.Y][position.X].Color == bot.SkullColor) && !visited[hashPosition] {
					if grid[position.Y][position.X].Color != bot.SkullColor {
						queue = append(queue, *position)
					}
					if len(chain) > 0 || grid[position.Y][position.X].Color == color {
						chain = append(chain, *grid[position.Y][position.X].Clone())
					}
					visited[hashPosition] = true
				}
			}
		}
	}
	return chain
}
