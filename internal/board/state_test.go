package board

import "testing"

func TestLoseOnMineReveal(t *testing.T) {
	b := New(2, 1)
	b.SetMine(1, 0)
	ok := b.Reveal(1, 0)
	if ok {
		t.Fatalf("expected false when revealing a mine")
	}
	if !b.Lost() || b.State() != StateLost {
		t.Fatalf("expected lost state")
	}
	if b.Won() {
		t.Fatalf("should not be won")
	}
}

func TestWinWhenAllSafeRevealed(t *testing.T) {
	b := New(2, 2)
	b.SetMine(1, 1) // one mine -> 3 safe cells

	// Reveal the three safe cells
	if !b.Reveal(0, 0) {
		t.Fatal("unexpected mine")
	}
	if !b.Reveal(1, 0) {
		t.Fatal("unexpected mine")
	}
	if !b.Reveal(0, 1) {
		t.Fatal("unexpected mine")
	}

	if !b.Won() || b.State() != StateWon {
		t.Fatalf("expected won state; got state=%v", b.State())
	}
	if b.Lost() {
		t.Fatalf("should not be lost")
	}
}

func TestPlayingStateMidGame(t *testing.T) {
	b := New(4, 1)
	b.SetMine(2, 0)
	if !b.Reveal(0, 0) {
		t.Fatal("unexpected mine")
	}
	if b.State() != StatePlaying {
		t.Fatalf("expected playing state, got %v", b.State())
	}
}
