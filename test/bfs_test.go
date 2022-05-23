package test

import (
	"fmt"
	"smash_the_code/algorithms"
	"smash_the_code/bot"
	"smash_the_code/lib"
	"smash_the_code/test/utils"
	"testing"
)

func TestCase5BallsChainBfs(t *testing.T) {
	state := new(bot.State)
	state.Init()
	state.Map.Grid[11][0] = bot.Cell{Color: 2, Position: &lib.V{X: 0, Y: 11}}
	state.Map.Grid[11][1] = bot.Cell{Color: 2, Position: &lib.V{X: 1, Y: 11}}
	state.Map.Grid[10][0] = bot.Cell{Color: 2, Position: &lib.V{X: 0, Y: 10}}
	state.Map.Grid[11][2] = bot.Cell{Color: 2, Position: &lib.V{X: 2, Y: 11}}
	state.Map.Grid[11][3] = bot.Cell{Color: 2, Position: &lib.V{X: 3, Y: 11}}

	utils.PrintMap(state.Map)

	visited := make([]bool, bot.MapWidth*bot.MapHeight)
	chain := algorithms.BFS(&state.Grid, state.Grid[11][0], &visited)
	fmt.Println("BFS Chain: ", chain)

	if len(chain) != 5 {
		t.Error("Wrong size")
	}
}
