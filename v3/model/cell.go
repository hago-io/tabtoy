package model

import (
	"fmt"

	"github.com/davyxu/tabtoy/util"
)

type Cell struct {
	Value string
	Row   int // base 0
	Col   int // base 0
	Table *DataTable
}

// 全拷贝
func (cell *Cell) CopyFrom(c *Cell) {
	cell.Value = c.Value
	cell.Row = c.Row
	cell.Col = c.Col
	cell.Table = c.Table
}

func (cell *Cell) String() string {

	var file, sheet string
	if cell.Table != nil {
		file = cell.Table.FileName
		sheet = cell.Table.SheetName
	}

	return fmt.Sprintf("'%s' @%s|%s(%s)", cell.Value, file, sheet, util.R1C1ToA1(cell.Row+1, cell.Col+1))
}
