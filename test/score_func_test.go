package test

//func TestSimpleBestSequence(t *testing.T) {
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
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, -1, -1, -1, -1, -1},
//		{-1, 3, -1, -1, -1, -1},
//	}
//	utils.PrintMap(*state)
//
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 3, Color2: 1})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 1, Color2: 2})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 2, Color2: 2})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 4, Color2: 5})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 2, Color2: 5})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 3, Color2: 5})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 1, Color2: 1})
//	state.Blocks = append(state.Blocks, &bot.Block{Color1: 2, Color2: 1})
//
//	sequence := bot.FindBestSequence(state)
//	for _, block := range sequence {
//		bot.ApplyBlock(state, &block)
//		//fmt.Println(block.Position1, block.Position2, block.Rotation)
//		utils.PrintMap(*state)
//	}
//	if false {
//		t.Error("Wrong")
//	}
//}
