package bot

//const MapWidth = 6
//const MapHeight = 12
//const EmptyColor = -1

type State struct {
	Map
	QueueBlocks []*Block
	//RemovedBlocks [][]*lib.V
	Score    int
	LastStep Block
}

//func (s *State) UpdateScore(colors []bool, chains [][]lib.V, doneSteps int) {
//	b := 0
//	cp := 1
//	cb := 0
//	gb := 0
//
//	switch doneSteps {
//	case 0:
//		cp = 0
//	case 1:
//		cp = 8
//	default:
//		cp = 8 * int(math.Pow(2, float64(doneSteps-1)))
//	}
//
//	diffColors := 0
//	for i := range chains {
//		chain := chains[i]
//		b += len(chain)
//	}
//	for i := range colors {
//		if colors[i] {
//			diffColors++
//		}
//	}
//	switch b {
//	case 4:
//		gb = 0
//	case 5:
//		gb = 1
//	case 6:
//		gb = 2
//	case 7:
//		gb = 3
//	case 8:
//		gb = 4
//	case 9:
//		gb = 5
//	case 10:
//		gb = 6
//	default:
//		gb = 8
//	}
//	if diffColors > 1 {
//		cb = int(math.Pow(float64(2), float64(diffColors)))
//	}
//	cellGroups := FindColorGroups(s)
//	if len(cellGroups) == 0 {
//		s.Score += (10 * b) * (cp + cb + gb)
//	} else {
//		s.Score += (10 * b) * (cp + cb + gb) * len(cellGroups)
//	}
//}

//func (s *State) BFS(v lib.V) *[]lib.V {
//	if s.bfsVisited[GetHashPosition(v)] {
//		chain := make([]lib.V, 0)
//		return &chain
//	}
//
//	color := s.Map[v.Y][v.X]
//
//	queue := make([]lib.V, 0)
//	queue = append(queue, v)
//
//	chain := make([]lib.V, 0)
//	chain = append(chain, v)
//
//	s.bfsVisited[GetHashPosition(v)] = true
//
//	for len(queue) > 0 {
//		p := queue[0]
//		queue = queue[1:]
//
//		area4 := p.Area4()
//		for i := range area4 {
//			newP := *area4[i]
//			hashPosition := GetHashPosition(newP)
//			if newP.InMap(MapWidth, MapHeight) {
//				if s.Map[newP.Y][newP.X] == color && !s.bfsVisited[hashPosition] {
//					queue = append(queue, newP)
//					chain = append(chain, newP)
//					s.bfsVisited[hashPosition] = true
//				}
//			}
//		}
//	}
//	if len(chain) < 4 {
//		chain = make([]lib.V, 0)
//	}
//	return &chain
//}

//func (s *State) DFS(p *lib.V, color int, chain *[]lib.V, visited *[][]bool) {
//	v := *visited
//	if p.InMap(MapWidth, MapHeight) && s.Map[p.Y][p.X] == color && !v[p.Y][p.X] {
//		*chain = append(*chain, *p)
//		v[p.Y][p.X] = true
//		*visited = v
//
//		s.DFS(&lib.V{p.X, p.Y + 1}, color, chain, visited)
//		s.DFS(&lib.V{p.X + 1, p.Y}, color, chain, visited)
//		s.DFS(&lib.V{p.X - 1, p.Y}, color, chain, visited)
//		s.DFS(&lib.V{p.X, p.Y - 1}, color, chain, visited)
//	}
//}

func (s *State) Clone() *State {
	state := new(State)
	state.Map = *s.Map.Clone()

	state.QueueBlocks = make([]*Block, len(s.QueueBlocks))
	for i := 0; i < len(s.QueueBlocks); i++ {
		state.QueueBlocks[i] = s.QueueBlocks[i].Clone()
	}
	state.Score = s.Score
	return state
}

func (s *State) Init() {
	s.Map.Init()
}
