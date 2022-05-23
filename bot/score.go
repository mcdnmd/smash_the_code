package bot

func (s *State) UpdateScore(colors []bool, chainsPtr *[][]Cell) {
	cellGroups := FindColorGroups(s)
	notEmptyCellsNumber := 0
	for i := range cellGroups {
		group := cellGroups[i]
		notEmptyCellsNumber += len(group)
	}

	chains := *chainsPtr
	removedChainsNumber := len(chains)

	colorCount := 0
	for i := range colors {
		if colors[i] {
			colorCount++
		}
	}

	emptyCellsNumber := MapWidth*MapHeight - notEmptyCellsNumber

	s.Score += removedChainsNumber*130 + colorCount*5 + emptyCellsNumber*13
	//s.Score += emptyCellsNumber*713 + removedChainsNumber*1337
}
