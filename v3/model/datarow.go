package model

type DataRow struct {
	row   int
	cells []*Cell
	tab   *DataTable
}

func (row *DataRow) Cells() []*Cell {
	return row.cells
}

func (row *DataRow) Cell(col int) *Cell {
	return row.cells[col]
}

func (row *DataRow) AddCell() (ret *Cell) {

	ret = &Cell{
		Col:   len(row.cells),
		Row:   row.row,
		Table: row.tab,
	}

	row.cells = append(row.cells, ret)
	return
}

func (row *DataRow) IsEmpty() bool {
	return len(row.cells) == 0
}

func newDataRow(row int, tab *DataTable) *DataRow {
	return &DataRow{
		row: row,
		tab: tab,
	}
}
