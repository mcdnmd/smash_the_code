package algorithms

import (
	"math/rand"
	"smash_the_code/bot"
	"time"
)

// Score state in the end of loop

const GreedyLookup = 1

func GreedySearch(baseState *bot.State) ([]bot.Block, int) {
	state := baseState

	startTime := time.Now()
	duration := time.Since(startTime)
	scoredSteps := make(map[int][]bot.Block)
	best := 0
	c := 0

	for duration.Milliseconds() < 65 {

		greedyState := state.Clone()

		tmpState := greedyState.Clone()
		sequence := make([]bot.Block, GreedyLookup)

		for i := range baseState.QueueBlocks[:GreedyLookup] {
			cloneState := tmpState
			block := baseState.QueueBlocks[i]
			possibleBlocks := bot.GetAllPossibleBlocks(*block)
			possibleStates := make([]bot.State, 0)

			for j := range possibleBlocks {
				cloneState2 := cloneState.Clone()
				b := possibleBlocks[j]
				status, _ := bot.ApplyBlock(cloneState2, b)
				if !status {
					continue
				}
				cloneState2.LastStep = *b
				possibleStates = append(possibleStates, *cloneState2)
			}
			idx := rand.Intn(len(possibleStates))
			tmpState = &possibleStates[idx]
			sequence[i] = tmpState.LastStep
		}

		_, chain := bot.ApplyBlock(greedyState, &sequence[0])
		score := bot.CalcScore(*chain)

		if score > best {
			best = score
		}

		scoredSteps[score] = sequence
		duration = time.Since(startTime)
		c++
	}
	//fmt.Fprintln(os.Stderr, "Best Score", best, scoredSteps)
	return scoredSteps[best], c
}
