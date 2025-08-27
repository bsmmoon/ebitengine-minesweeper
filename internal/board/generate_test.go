package board

import "testing"

func TestGenerateCountAndDeterminism(t *testing.T) {
	b1 := New(5, 4)
	b2 := New(5, 4)

	if !b1.Generate(6, 42) {
		t.Fatal("generate failed on b1")
	}
	if !b2.Generate(6, 42) {
		t.Fatal("generate failed on b2")
	}

	if b1.Mines != 6 {
		t.Fatalf("mine count mismatch: got %d want 6", b1.Mines)
	}
	for i := range b1.Cells {
		if b1.Cells[i].Mine != b2.Cells[i].Mine {
			t.Fatalf("determinism mismatch at idx %d", i)
		}
	}
}

func TestGenerateInvalidInputs(t *testing.T) {
	b := New(3, 3)
	if b.Generate(9, 1) { // 9 == area -> invalid
		t.Fatalf("expected false when mines == area")
	}
	if b.Generate(-1, 1) { // negative -> invalid
		t.Fatalf("expected false when mines < 0")
	}
}
