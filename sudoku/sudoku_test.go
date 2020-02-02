package sudoku

import "testing"

func TestNewCellGet(t *testing.T) {
	c := &cell{}
	want := 0
	if got := c.get(); got != want {
		t.Errorf("&cell{}.get() == %d, expected %d", got, want)
	}
}

func TestCellSetInBounds(t *testing.T) {
	c := &cell{}
	want := 1
	c.set(want)
	if got := c.get(); got != want {
		t.Errorf("cell.set(1) == %d, expected %d", got, want)
	}
}

func TestCellSetOutOfBounds(t *testing.T) {
	c := &cell{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic on out of bounds value")
		}
	}()

	c.set(0)
}

func TestNewCellNotes(t *testing.T) {
	c := &cell{}
	if got := c.notes(); got != nil {
		t.Errorf("&cell{}.notes() == %d, expected nil", got)
	}
}

func TestNewCellString(t *testing.T) {
	c := &cell{}
	want := "{}"
	if got := c.String(); got != want {
		t.Errorf("NewCell().String() == %q, expected %q", got, want)
	}
}

func TestNewGridCells(t *testing.T) {
	n := 9
	g := NewGrid(n)
	if height := len(g.cells); height != n {
		t.Errorf("Grid height == %d, expected %d", height, n)
	}
	if capacity := cap(g.cells); capacity != n {
		t.Errorf("Grid capacity == %d, expected %d", capacity, n)
	}
	for _, row := range g.cells {
		if width := len(row); width != n {
			t.Errorf("Grid width == %d, expected %d", width, n)
		}
		if capacity := cap(row); capacity != n {
			t.Errorf("Grid row capacity == %d, expected %d", capacity, n)
		}
		for _, cell := range row {
			if value := cell.get(); value != 0 {
				t.Errorf("Cell value == %d, expected 0", value)
			}
			if notes := cell.notes(); notes != nil {
				t.Errorf("Cell notes == %d, expected nil", notes)
			}
		}
	}
}

func TestGridSetValueOutOfBounds(t *testing.T) {
	n := 9
	g := NewGrid(n)
	v := n + 1

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic on out of bounds value")
		}
	}()

	g.Set(0, 0, v)
}

func TestGridSetIndexColumnOutOfBounds(t *testing.T) {

	test := func(n, c, r int) {
		g := NewGrid(n)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Did not panic on out of bounds index")
			}
		}()

		g.Set(c, r, n)
	}

	for n := 3; n < 100; n *= n {
		for i := -100; i < 0; i += 10 {
			test(n, i, 0)
			test(n, i, n-1)
			test(n, 0, i)
			test(n, n-1, i)
		}
		for i := n; i < 100; i += 10 {
			test(n, i, 0)
			test(n, i, n-1)
			test(n, 0, i)
			test(n, n-1, i)
		}
	}
}
