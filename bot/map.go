package bot

import (
	"smash_the_code/algorithms"
	"smash_the_code/lib"
	"strconv"
)

const MapWidth = 6
const MapHeight = 12
const EmptyColor = -1
const SkullColor = 0

type Map struct {
	Grid [][]Cell
}

func (m *Map) Init() {
	m.Grid = make([][]Cell, MapHeight)
	for y := 0; y < MapHeight; y++ {
		m.Grid[y] = make([]Cell, MapWidth)
		for x := 0; x < MapWidth; x++ {
			m.Grid[y][x] = Cell{Color: EmptyColor, Position: &lib.V{X: x, Y: y}}
		}
	}
}

func (m *Map) FindEmptyCellsInColumns() []Cell {
	emptyPositions := make([]Cell, MapWidth)
	for x := 0; x < MapWidth; x++ {
		for y := MapHeight - 1; y >= 0; y-- {
			if m.Grid[y][x].Color == EmptyColor {
				emptyPositions[x] = m.Grid[y][x]
				break
			}
		}
	}
	return emptyPositions
}

// SetBlock FIXME исправить баг, когда полностью заполняется столбец
func (m *Map) SetBlock(b *Block) bool {
	emptyCells := m.FindEmptyCellsInColumns()

	if !b.Cell1.Position.InMap(MapWidth, MapHeight) || !b.Cell2.Position.InMap(MapWidth, MapHeight) {
		return false
	}

	if emptyCells[b.Cell1.Position.X].Position == nil || emptyCells[b.Cell2.Position.X].Position == nil {
		return false
	}

	position1 := new(lib.V)
	position2 := new(lib.V)
	switch b.Rotation {
	case 0:
		position1 = emptyCells[b.Cell1.Position.X].Position
		position2 = emptyCells[b.Cell2.Position.X].Position
	case 1:
		position1 = emptyCells[b.Cell1.Position.X].Position
		position2 = &lib.V{X: position1.X, Y: position1.Y - 1}
	case 2:
		position1 = emptyCells[b.Cell1.Position.X].Position
		position2 = emptyCells[b.Cell2.Position.X].Position
	case 3:
		position2 = emptyCells[b.Cell2.Position.X].Position
		position1 = &lib.V{X: position2.X, Y: position2.Y - 1}
	}
	//fmt.Fprintln(os.Stderr, "before set", b.Cell1.Position, b.Cell2.Position)
	if !position1.InMap(MapWidth, MapHeight) || !position2.InMap(MapWidth, MapHeight) {
		return false
	}
	b.Cell1.Position = position1
	b.Cell2.Position = position2
	m.Grid[position1.Y][position1.X].Color = b.Cell1.Color
	m.Grid[position2.Y][position2.X].Color = b.Cell2.Color
	return true
}

func (m *Map) SetCell(x int, y int, val uint8) {
	if val != '.' {
		n, _ := strconv.ParseInt(string(val), 10, 64)
		m.Grid[y][x] = Cell{int(n), &lib.V{X: x, Y: y}}
	} else {
		m.Grid[y][x] = Cell{Color: EmptyColor, Position: &lib.V{X: x, Y: y}}
	}
}

func (m *Map) RemoveChain(chain []Cell) {
	for i := range chain {
		cell := chain[i]
		m.Grid[cell.Position.Y][cell.Position.X].Color = EmptyColor
	}
}

func (m *Map) BallsFreeFall() {
	for x := 0; x < MapWidth; x++ {
		for y := MapHeight - 1; y >= 0; y-- {
			if m.Grid[y][x].Color == EmptyColor {
				for i := y - 1; i >= 0; i-- {
					if m.Grid[i][x].Color != EmptyColor {
						m.Grid[y][x].Color = m.Grid[i][x].Color
						m.Grid[i][x].Color = EmptyColor
						break
					}
				}
			}
		}
	}
}

func (m *Map) Clone() *Map {
	newMap := new(Map)
	newMap.Init()
	for y := 0; y < MapHeight; y++ {
		for x := 0; x < MapWidth; x++ {
			newMap.Grid[y][x] = *m.Grid[y][x].Clone()
		}
	}
	return newMap
}

func GetHashPosition(v lib.V) int {
	return v.Y*MapWidth + v.X
}

func FindColorGroups(state *State) [][]Cell {
	grid := state.Map.Grid
	visited := make([]bool, MapHeight*MapWidth)

	cellGroups := make([][]Cell, 0)

	for y := 0; y < MapHeight; y++ {
		for x := 0; x < MapWidth; x++ {
			position := lib.V{X: x, Y: y}
			if !visited[GetHashPosition(position)] {
				chain := algorithms.BFS(&grid, grid[position.Y][position.X], &visited)
				if len(chain) > 2 {
					tmpCellChain := make([]Cell, len(chain))
					for i, cell := range chain {
						tmpCellChain[i] = Cell{Color: cell.Color, Position: cell.Position.Clone()}
					}
					cellGroups = append(cellGroups, tmpCellChain)
				}
				visited[GetHashPosition(position)] = true
			}
		}
	}
	return cellGroups
}
