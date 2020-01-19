package sudoku

type cell struct {
	value    int
	possible [9]bool
}

func (c cell) Set(value int) {
	c.value = value
}

type Segment struct {
}

type Grid struct {
	values [9][9]cell
}
