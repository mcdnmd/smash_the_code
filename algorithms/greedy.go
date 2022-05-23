package algorithms

import (
	"math/rand"
	"smash_the_code/bot"
)

func GreedySearch(baseState *bot.State) []bot.Block {
	state := baseState
	sequence := make([]bot.Block, 8)

	for i := range baseState.QueueBlocks {
		block := baseState.QueueBlocks[i]
		possibleBlocks := bot.GetAllPossibleBlocks(*block)
		scoredStates := make(map[int][]bot.State)

		best := state.Score

		//fmt.Println(i, " BEST befor ", best)
		for j := range possibleBlocks {

			cloneState := state.Clone()
			b := possibleBlocks[j]
			status := bot.ApplyBlock(cloneState, b)
			if !status {
				continue
			}
			cloneState.LastStep = *b
			scoredStates[cloneState.Score] = append(scoredStates[cloneState.Score], *cloneState)
			if cloneState.Score > best {
				best = cloneState.Score
			}
		}

		bestStates := scoredStates[best]
		idx := rand.Intn(len(bestStates))
		state = &bestStates[idx]
		sequence[i] = state.LastStep
		//fmt.Println(i, " BEST after ", best)
	}
	return sequence
}
