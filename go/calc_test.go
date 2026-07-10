package calc

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(3, 4); got != 7 {
		t.Errorf("Add(3, 4) = %d; want 7", got)
	}
}

func TestSub4_3(t *testing.T) {
	if got := Sub(4, 3); got != 1 {
		t.Errorf("Sub(4, 3) = %d; want 1", got)
	}
}

func TestSub3_4(t *testing.T) {
	if got := Sub(3, 4); got != -1 {
		t.Errorf("Sub(3, 4) = %d; want -1", got)
	}
}
