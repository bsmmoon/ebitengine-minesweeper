package board

import "testing"

func TestIdxAndBounds(t *testing.T) {
	b := New(3, 2) // width=3, height=2 -> 6 cells
	if !b.InBounds(0, 0) || !b.InBounds(2, 1) {
		t.Fatalf("expected corners to be in bounds")
	}
	if b.InBounds(-1, 0) || b.InBounds(3, 0) || b.InBounds(0, 2) {
		t.Fatalf("expected out-of-bounds to be false")
	}
	if b.Idx(2, 1) != 5 {
		t.Fatalf("Idx mismatch: got %d want 5", b.Idx(2, 1))
	}
}

func TestAdjacencyCounts(t *testing.T) {
	b := New(3, 3)
	if !b.SetMine(1, 1) {
		t.Fatalf("failed to set mine")
	}
	// All 8 neighbors should have Adj==1, center remains mine
	want := map[[2]int]int{
		{0, 0}: 1, {1, 0}: 1, {2, 0}: 1,
		{0, 1}: 1, {2, 1}: 1,
		{0, 2}: 1, {1, 2}: 1, {2, 2}: 1,
	}
	for pos, adj := range want {
		c := b.At(pos[0], pos[1])
		if c == nil || c.Adj != adj {
			t.Fatalf("adj mismatch at %v: got %v want %v", pos, c, adj)
		}
	}
	if !b.At(1, 1).Mine {
		t.Fatalf("center should be a mine")
	}
}
