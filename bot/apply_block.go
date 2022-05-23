package bot

import (
	"smash_the_code/algorithms"
	"smash_the_code/lib"
)

func ApplyBlock(state *State, block *Block) (bool, *[][]Cell) {
	status := state.Map.SetBlock(block)
	if !status {
		return false, nil
	}

	visited := make([]bool, MapWidth*MapHeight)
	removedCellChains := make([][]Cell, 0)
	colors := make([]bool, 8)

	if block.Cell1.Color != block.Cell2.Color {
		chain1 := algorithms.BFS(&state.Map.Grid, block.Cell1, &visited)
		if len(chain1) > 0 {
			removedCellChains = append(removedCellChains, chain1)
			colors[chain1[0].Color] = true
		}
		chain2 := algorithms.BFS(&state.Map.Grid, block.Cell1, &visited)
		if len(chain2) > 0 {
			removedCellChains = append(removedCellChains, chain2)
			colors[chain2[0].Color] = true
		}
		state.Map.RemoveChain(chain1)
		state.Map.RemoveChain(chain2)
	} else {
		chain := algorithms.BFS(&state.Map.Grid, block.Cell1, &visited)
		if len(chain) > 0 {
			removedCellChains = append(removedCellChains, chain)
			colors[chain[0].Color] = true
		}
		state.Map.RemoveChain(chain)
	}
	state.BallsFreeFall()

	ApplyAllMapMutations(state, &colors, &removedCellChains)

	//state.UpdateScore(colors, &removedCellChains)
	return true, &removedCellChains
}

func ApplyAllMapMutations(state *State, colorsPtr *[]bool, removedCellChainsPtr *[][]Cell) {
	mapGridHasChanges := true
	removedChains := *removedCellChainsPtr
	colors := *colorsPtr

	for mapGridHasChanges {
		visited := make([]bool, MapWidth*MapHeight)
		mapGridHasChanges = false
		shouldFall := false

		cellChains := FindAllCellChains(&state.Map.Grid, &visited, &colors)

		for i := range cellChains {
			if len(cellChains[i]) > 0 {
				state.RemoveChain(cellChains[i])
				shouldFall = true
				removedChains = append(removedChains, cellChains[i])
			}
		}
		if shouldFall {
			state.BallsFreeFall()
		}
	}
}

func FindAllCellChains(gridPtr *[][]Cell, visitedPtr *[]bool, colorsPtr *[]bool) (cellChains [][]Cell) {
	grid := *gridPtr
	visited := *visitedPtr
	colors := *colorsPtr

	for y := MapHeight - 1; y >= 0; y-- {
		for x := 0; x < MapWidth; x++ {
			if grid[y][x].Color != EmptyColor && grid[y][x].Color != SkullColor {
				cell := Cell{grid[y][x].Color, &lib.V{X: x, Y: y}}
				if !visited[GetHashPosition(*cell.Position)] {
					chain := algorithms.BFS(gridPtr, cell, visitedPtr)
					if len(chain) > 0 {
						colors[chain[0].Color] = true
						cellChains = append(cellChains, chain)
					}
				}
			}
		}
	}
	return cellChains
}
