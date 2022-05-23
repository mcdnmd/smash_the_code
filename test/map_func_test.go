package test

//func TestMapFindEmptyCells(t *testing.T) {
//	state := new(bot.State)
//	state.Init()
//	state.Map[11][0] = 1
//	state.Map[11][1] = 3
//	state.Map[10][0] = 1
//	state.Map[11][2] = 3
//
//	emptyCells := state.FindEmptyCellsInColumns()
//
//	answer := []lib.V{{0, 9}, {1, 10}, {2, 10}, {3, 11}, {4, 11}, {5, 11}}
//
//	for i := range answer {
//		if answer[i] != emptyCells[i] {
//			t.Error("Wrong")
//		}
//	}
//}
