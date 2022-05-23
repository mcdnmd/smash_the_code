package main

import (
	"fmt"
	"math/rand"
	"os"
	"smash_the_code/algorithms"
	"smash_the_code/bot"
	"time"
)

// greedy search for one step

// HillClimbing - плохая идея с точки зрения мутации

// BeamSearch -
//		очки - плохая функция оценки
// 		* максимизация одного цвета -
//		* кол-во областей размера 3 с открытым боком для сбора комбо
//		*
//		*

// change to BFS using queue like in samegame

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		startTime := time.Now()
		blocks := make([]*bot.Block, 8)
		for i := 0; i < 8; i++ {
			// colorA: color of the first block
			// colorB: color of the attached block
			var colorA, colorB int
			fmt.Scan(&colorA, &colorB)
			blocks[i] = &bot.Block{Cell1: bot.Cell{Color: colorA}, Cell2: bot.Cell{Color: colorB}}
		}

		// Our map
		var score1 int
		fmt.Scan(&score1)

		MyState := new(bot.State)
		MyState.Init()
		MyState.QueueBlocks = blocks

		fmt.Fprintln(os.Stderr, "READ")

		for i := 0; i < 12; i++ {
			// row: One line of the map ('.' = empty, '0' = skull block, '1' to '5' = colored block)
			var row string
			fmt.Scan(&row)
			for j := range row {
				MyState.SetCell(j, i, row[j])
			}
		}

		// Enemy map
		var score2 int
		fmt.Scan(&score2)

		EnemyState := new(bot.State)
		EnemyState.Init()

		for i := 0; i < 12; i++ {
			var row string
			fmt.Scan(&row)
			for j := range row {
				EnemyState.SetCell(j, i, row[j])
			}
		}
		//place, rotation := Logic2(MyState)
		sequence := algorithms.GreedySearch(MyState)
		step := sequence[0]
		fmt.Println(step.Place, step.Rotation)
		fmt.Fprintln(os.Stderr, "Round duration:", time.Since(startTime).Milliseconds(), "ms")
	}
}

//func Logic2(state *bot.State) (int, int) {
//	scores := make(map[int]bot.Block)
//	possibleMoves := make([]bot.Block, 0)
//	baseBlock := state.Blocks[0]
//	blocks := bot.GetAllPossibleBlocks(state, *baseBlock)
//	best := 0
//	for j := range blocks {
//		clone := state.Clone()
//		block := blocks[j]
//		status := bot.ApplyBlock(clone, block)
//		if !status {
//			continue
//		}
//		scores[clone.Score] = *block
//		possibleMoves = append(possibleMoves, *block)
//		if clone.Score > best {
//			best = clone.Score
//		}
//	}
//	ans := scores[best]
//	if best == 0 {
//		for {
//			i := rand.Intn(len(possibleMoves))
//			ans = possibleMoves[i]
//			clone := state.Clone()
//			status := clone.SetBlock(&ans)
//			if status {
//				break
//			}
//		}
//	}
//	return ans.Position1.X, ans.Rotation
//}

//func Logic1(MyState *bot.State) int {
//	//color := MyState.Blocks[0].Color1
//	if MyState.Map[2][0] == -1 {
//		return 0
//	} else if MyState.Map[3][5] == -1 {
//		return 5
//	} else if MyState.Map[3][1] == -1 {
//		return 1
//	} else if MyState.Map[3][4] == -1 {
//		return 4
//	} else if MyState.Map[3][2] == -1 {
//		return 2
//	} else {
//		return 3
//	}
//}
