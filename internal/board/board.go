package board

type Cell struct {
	Mine     bool
	Revealed bool
	Flagged  bool
	Adj      int // adjacent mine count
}

type Board struct {
	W, H  int
	Cells []Cell
	Mines int
}

func New(w, h int) *Board {
	return &Board{
		W: w, H: h,
		Cells: make([]Cell, w*h),
	}
}

func (b *Board) InBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < b.W && y < b.H
}

func (b *Board) Idx(x, y int) int {
	return y*b.W + x
}

func (b *Board) At(x, y int) *Cell {
	if !b.InBounds(x, y) {
		return nil
	}
	return &b.Cells[b.Idx(x, y)]
}

// SetMine places a mine at (x,y) and increments adjacency around it.
// Returns false if out-of-bounds or already a mine.
func (b *Board) SetMine(x, y int) bool {
	c := b.At(x, y)
	if c == nil || c.Mine {
		return false
	}
	c.Mine = true
	b.Mines++
	b.forEachNeighbor(x, y, func(nx, ny int, nc *Cell) {
		nc.Adj++
	})
	return true
}

func (b *Board) forEachNeighbor(x, y int, f func(nx, ny int, c *Cell)) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nc := b.At(nx, ny); nc != nil {
				f(nx, ny, nc)
			}
		}
	}
}
