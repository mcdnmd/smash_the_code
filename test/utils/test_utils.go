package utils

import (
	"fmt"
	"smash_the_code/bot"
)

func PrintMap(m bot.Map) {
	fmt.Println("########## MAP")
	for row := range m.Grid {
		for col := range m.Grid[row] {
			fmt.Print(m.Grid[row][col].Color, m.Grid[row][col].Position, " ")
		}
		fmt.Println()
	}
	fmt.Println("##########")
}
