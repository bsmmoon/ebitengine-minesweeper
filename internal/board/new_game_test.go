package board

import "testing"

func TestNewGameFirstClickSafe(t *testing.T) {
  w, h := 5, 5
  mines := 5
  seed := int64(100)

  b := NewGame(w, h, mines, seed, 2, 2)
  if b.At(2, 2).Mine {
    t.Fatalf("first click (2,2) should never be a mine")
  }
  if b.Mines != mines {
    t.Fatalf("expected %d mines, got %d", mines, b.Mines)
  }
}

func TestNewGameDeterminismSafeSpot(t *testing.T) {
  w, h := 4, 4
  mines := 4
  seed := int64(7)

  b1 := NewGame(w, h, mines, seed, 0, 0)
  b2 := NewGame(w, h, mines, seed, 0, 0)

  for i := range b1.Cells {
    if b1.Cells[i].Mine != b2.Cells[i].Mine {
      t.Fatalf("determinism mismatch at idx %d", i)
    }
  }
}
