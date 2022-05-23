package test

//
//import (
//	"fmt"
//	"smash_the_code/bot"
//	"smash_the_code/lib"
//	"smash_the_code/test/utils"
//	"testing"
//)
//
//func TestRemoveChainBlock(t *testing.T) {
//	state := new(bot.State)
//	state.Init()
//	state.Map = [][]int{
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, 3, -1, -1, -1, -1},
//		{3, 1, -1, -1, -1, -1},
//		{3, 1, -1, -1, -1, -1},
//		{1, 3, -1, -1, -1, -1},
//		{1, 3, 3, -1, -1, 3},
//	}
//
//	fmt.Println("MAP CREATED")
//	utils.PrintMap(*state)
//
//	block := bot.Block{Color1: 3, Color2: 3}
//	block.Position1 = &lib.V{X: 3, Y: 0}
//	block.ApplyRotation(1)
//
//	fmt.Println("BLOCK SET")
//	utils.PrintMap(*state)
//
//	bot.ApplyBlock(state, &block)
//
//	fmt.Println("CLEAR ALL COMBOS")
//	utils.PrintMap(*state)
//
//	fmt.Println(state.Score)
//	if false {
//		t.Error("Wrong")
//	}
//}
//
//func TestFindBestSolution(t *testing.T) {
//	state := new(bot.State)
//	state.Init()
//	state.Map = [][]int{
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, 3, -1, -1, -1, -1},
//		{3, 1, -1, -1, -1, -1},
//		{3, 1, -1, -1, -1, -1},
//		{1, 3, -1, -1, -1, -1},
//		{1, 3, 3, -1, -1, 3},
//	}
//
//	fmt.Println("MAP CREATED")
//	utils.PrintMap(*state)
//
//	baseBlock := bot.Block{Color1: 3, Color2: 3}
//	blocks := bot.GetAllPossibleBlocks(state, baseBlock)
//
//	scores := make(map[int]bot.Block)
//	best := 0
//
//	for j := range blocks {
//		clone := state.Clone()
//		block := blocks[j]
//		status := bot.ApplyBlock(clone, block)
//		if !status {
//			continue
//		}
//		fmt.Println("ROUND PLAYED")
//		utils.PrintMap(*clone)
//		scores[clone.Score] = *block
//		if clone.Score > best {
//			best = clone.Score
//		}
//	}
//
//	fmt.Println(scores[best])
//
//	if best != 1560 {
//		t.Error("Wrong")
//	}
//}
//
//func TestFindBestSequence(t *testing.T) {
//	state := new(bot.State)
//	state.Init()
//	state.Map = [][]int{
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, 3, -1, -1, -1, -1},
//		{3, 1, -1, -1, -1, -1},
//		{3, 1, -1, -1, -1, -1},
//		{1, 3, -1, -1, -1, -1},
//		{1, 3, 3, -1, -1, 3},
//	}
//
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 3, Color2: 3})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 3, Color2: 3})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 3, Color2: 3})
//
//	sequence := bot.FindBestSequence(state)
//	for _, block := range sequence {
//		fmt.Println(block.Position1, block.Position2, block.Rotation)
//	}
//
//	if false {
//		t.Error("Wrong")
//	}
//}
