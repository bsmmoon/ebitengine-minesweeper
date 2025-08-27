package board

import "testing"

func TestRevealMine(t *testing.T) {
  b := New(2, 1)
  b.SetMine(0, 0)
  safe := b.Reveal(0, 0)
  if safe {
    t.Fatalf("expected false when revealing a mine")
  }
}

func TestRevealSafeCell(t *testing.T) {
  b := New(2, 2)
  b.SetMine(1, 1)
  safe := b.Reveal(0, 0)
  if !safe {
    t.Fatalf("expected safe reveal")
  }
  if !b.At(0, 0).Revealed {
    t.Fatalf("expected cell to be revealed")
  }
}

func TestFloodFill(t *testing.T) {
  b := New(3, 3)
  // No mines at all
  b.Reveal(1, 1)
  // Every cell should be revealed
  for y := 0; y < 3; y++ {
    for x := 0; x < 3; x++ {
      if !b.At(x, y).Revealed {
        t.Fatalf("expected flood fill, but (%d,%d) not revealed", x, y)
      }
    }
  }
}

func TestToggleFlag(t *testing.T) {
  b := New(2, 1)
  b.ToggleFlag(0, 0)
  if !b.At(0, 0).Flagged {
    t.Fatalf("expected flagged")
  }
  b.ToggleFlag(0, 0)
  if b.At(0, 0).Flagged {
    t.Fatalf("expected unflagged after toggle")
  }
}

func TestFlagDoesNotReveal(t *testing.T) {
  b := New(2, 1)
  b.ToggleFlag(0, 0)
  b.Reveal(0, 0) // should do nothing
  if b.At(0, 0).Revealed {
    t.Fatalf("flagged cell should not be revealed")
  }
}
