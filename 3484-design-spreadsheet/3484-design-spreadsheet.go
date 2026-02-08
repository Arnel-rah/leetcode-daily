
type Spreadsheet struct {
	cells map[string]int
	rows  int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		cells: make(map[string]int),
		rows:  rows,
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.cells[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	delete(this.cells, cell)
}

func (this *Spreadsheet) GetValue(formula string) int {
	i := 1
	for formula[i] != '+' {
		i++
	}
	x := formula[1:i]
	y := formula[i+1:]
	return this.getVal(x) + this.getVal(y)
}

func (this *Spreadsheet) getVal(s string) int {
	if s[0] >= 'A' && s[0] <= 'Z' {
		if v, ok := this.cells[s]; ok {
			return v
		}
		return 0
	}
	v, _ := strconv.Atoi(s)
	return v
}
