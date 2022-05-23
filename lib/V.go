package lib

type V struct {
	X int
	Y int
}

func (v *V) Area4() []*V {
	result := make([]*V, 4)
	result[0] = &V{v.X, v.Y + 1}
	result[1] = &V{v.X, v.Y - 1}
	result[2] = &V{v.X + 1, v.Y}
	result[3] = &V{v.X - 1, v.Y}
	return result
}

func (v *V) InMap(MaxWidth int, MaxHeight int) bool {
	return (v.X >= 0 && v.X < MaxWidth) && (v.Y >= 0 && v.Y < MaxHeight)
}

func (v *V) Clone() *V {
	if v != nil {
		return &V{X: v.X, Y: v.Y}
	}
	return nil
}
