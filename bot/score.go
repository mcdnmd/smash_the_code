package bot

import (
	"math"
)

func (s *State) UpdateScore(chainsPtr *[][]Cell) {
	cellGroups := FindColorGroups(s)
	notEmptyCellsNumber := 0
	for i := range cellGroups {
		group := cellGroups[i]
		notEmptyCellsNumber += len(group)
	}

	chains := *chainsPtr
	removedChainsNumber := len(chains)

	//colorCount := 0
	//for i := range colors {
	//	if colors[i] {
	//		colorCount++
	//	}
	//}

	//emptyCellsNumber := MapWidth*MapHeight - notEmptyCellsNumber

	s.Score += 53 * removedChainsNumber
}

func CalcScore(chains [][]Cell) int {
	b := 0
	cp := 0
	cb := 0
	gb := 0

	doneSteps := len(chains)
	switch doneSteps {
	case 0:
		cp = 0
	case 1:
		cp = 8
	default:
		cp = 8 * int(math.Pow(2, float64(doneSteps-1)))
	}

	diffColors := 0
	colors := make([]bool, 6)
	for i := range chains {
		chain := chains[i]
		b += len(chain)
		for j := range chain {
			cell := chain[j]
			colors[cell.Color] = true
		}
	}
	for i := range colors {
		if colors[i] {
			diffColors++
		}
	}

	if diffColors > 1 {
		cb = int(math.Pow(float64(2), float64(diffColors)))
	}

	switch b {
	case 4:
		gb = 0
	case 5:
		gb = 1
	case 6:
		gb = 2
	case 7:
		gb = 3
	case 8:
		gb = 4
	case 9:
		gb = 5
	case 10:
		gb = 6
	default:
		gb = 8
	}

	return (10 * b) * (cp + cb + gb)
}
