package board

import "math/rand"

type Cell struct {
	Mine     bool
	Revealed bool
	Flagged  bool
	Adj      int
}

type State int

const (
	StatePlaying State = iota
	StateWon
	StateLost
)

type Board struct {
	W, H         int
	Cells        []Cell
	Mines        int
	RevealedSafe int  // number of revealed NON-mine cells
	Exploded     bool // true if a mine was revealed
}

func New(w, h int) *Board {
	return &Board{
		W: w, H: h,
		Cells: make([]Cell, w*h),
	}
}

func (b *Board) InBounds(x, y int) bool { return x >= 0 && y >= 0 && x < b.W && y < b.H }
func (b *Board) Idx(x, y int) int       { return y*b.W + x }

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

// Reveal attempts to reveal the cell at (x,y).
// Returns true if safe, false if a mine was hit.
func (b *Board) Reveal(x, y int) bool {
	c := b.At(x, y)
	if c == nil || c.Revealed || c.Flagged {
		return true
	}
	c.Revealed = true
	if c.Mine {
		b.Exploded = true
		return false
	}
	b.RevealedSafe++
	if c.Adj == 0 {
		b.forEachNeighbor(x, y, func(nx, ny int, nc *Cell) {
			if !nc.Revealed && !nc.Mine {
				b.Reveal(nx, ny)
			}
		})
	}
	return true
}

// ToggleFlag toggles the flagged state of a cell (no effect if already revealed).
func (b *Board) ToggleFlag(x, y int) {
	c := b.At(x, y)
	if c == nil || c.Revealed {
		return
	}
	c.Flagged = !c.Flagged
}

// --- Game state helpers ---

func (b *Board) safeCellsTotal() int {
	return b.W*b.H - b.Mines
}

func (b *Board) Won() bool {
	return !b.Exploded && b.RevealedSafe == b.safeCellsTotal()
}

func (b *Board) Lost() bool {
	return b.Exploded
}

func (b *Board) State() State {
	if b.Lost() {
		return StateLost
	}
	if b.Won() {
		return StateWon
	}
	return StatePlaying
}

// Generate resets the board and places `mines` mines deterministically using `seed`.
// Returns false if `mines` is invalid (mines < 0 or mines >= w*h).
func (b *Board) Generate(mines int, seed int64) bool {
	area := b.W * b.H
	if mines < 0 || mines >= area {
		return false
	}

	// Full reset
	for i := range b.Cells {
		b.Cells[i] = Cell{}
	}
	b.Mines = 0
	b.RevealedSafe = 0
	b.Exploded = false

	// Build index list and shuffle deterministically
	idxs := make([]int, area)
	for i := 0; i < area; i++ {
		idxs[i] = i
	}
	rng := rand.New(rand.NewSource(seed))
	rng.Shuffle(area, func(i, j int) { idxs[i], idxs[j] = idxs[j], idxs[i] })

	// Place first N indices as mines (SetMine updates adjacency)
	for i := 0; i < mines; i++ {
		id := idxs[i]
		x := id % b.W
		y := id / b.W
		b.SetMine(x, y)
	}
	return true
}
