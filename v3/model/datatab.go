package model

import (
	"fmt"
	"strings"
)

// 表格的完整数据，表头有屏蔽时，对应行值为空
type DataTable struct {
	HeaderType         string // 表名，Index表里定义的类型
	OriginalHeaderType string // HeaderFields对应的ObjectType，KV表为TableField
	FileName           string
	SheetName          string
	Rows               []*DataRow // 0下标为表头数据
	Headers            []*HeaderField
}

// 重复列在表中数量, 重复列可以将数组拆在多个列中填写
func (table *DataTable) RepeatedFieldCount(field *HeaderField) (ret int) {
	for _, hf := range table.Headers {
		if hf.TypeInfo == field.TypeInfo {
			ret++
		}
	}

	return
}

// 重复列在表中的索引, 相对于重复列的数量
func (table *DataTable) RepeatedFieldIndex(field *HeaderField) (ret int) {
	for _, hf := range table.Headers {
		if hf.TypeInfo == field.TypeInfo {
			if hf == field {
				break
			}
			ret++
		}
	}

	return
}

// 模板用，排除表头的数据索引
func (table *DataTable) DataRowIndex() (ret []int) {
	numRows := len(table.Rows)
	if numRows == 0 {
		return
	}
	ret = make([]int, numRows-1)
	// 排除表头数据
	for i := 0; i < numRows-1; i++ {
		ret[i] = i + 1
	}

	return
}

func (table *DataTable) String() string {

	var sb strings.Builder
	sb.WriteString("====DataTable====\n")
	sb.WriteString(fmt.Sprintf("HeaderType: %s\n", table.HeaderType))
	sb.WriteString(fmt.Sprintf("OriginalHeaderType: %s\n", table.OriginalHeaderType))
	sb.WriteString(fmt.Sprintf("FileName: %s\n", table.FileName))
	sb.WriteString(fmt.Sprintf("SheetName: %s\n", table.SheetName))

	// 遍历所有行
	for row, rowData := range table.Rows {
		sb.WriteString(fmt.Sprintf("%d ", row))
		// 遍历一行中的所有列值
		for index, cell := range rowData.Cells() {
			if index > 0 {
				sb.WriteString("/")
			}
			sb.WriteString(cell.Value)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (table *DataTable) MustGetHeader(col int) (header *HeaderField) {
	for len(table.Headers) <= col {
		table.Headers = append(table.Headers, &HeaderField{
			Cell: &Cell{
				Col: len(table.Headers),
			},
		})
	}
	return table.HeaderByColumn(col)
}

func (table *DataTable) HeaderByColumn(col int) *HeaderField {
	if col >= len(table.Headers) {
		return nil
	}
	return table.Headers[col]
}

func (table *DataTable) HeaderByName(name string) *HeaderField {
	for _, header := range table.Headers {
		if header.TypeInfo == nil {
			continue
		}
		if header.TypeInfo.Name == name || header.TypeInfo.FieldName == name {
			return header
		}
	}
	return nil
}

func (table *DataTable) AddRow() (row int) {
	row = len(table.Rows)
	table.Rows = append(table.Rows, newDataRow(row, table))
	return
}

func (table *DataTable) AddCell(row int) *Cell {

	if row >= len(table.Rows) {
		return nil
	}

	rowData := table.Rows[row]

	return rowData.AddCell()
}

func (table *DataTable) MustGetCell(row, col int) *Cell {

	for len(table.Rows) <= row {
		table.AddRow()
	}

	rowData := table.Rows[row]
	for len(rowData.cells) <= col {
		rowData.AddCell()
	}

	return rowData.Cell(col)
}

// 代码生成专用
func (table *DataTable) GetCell(row, col int) *Cell {

	if row >= len(table.Rows) {
		return nil
	}

	rowData := table.Rows[row]

	if col >= len(rowData.cells) {
		return nil
	}

	return rowData.Cell(col)
}

// 根据列头找到该行对应的值
func (table *DataTable) GetValueByName(row int, name string) *Cell {

	header := table.HeaderByName(name)

	if header == nil {
		return nil
	}

	return table.GetCell(row, header.Cell.Col)
}

func NewDataTable() *DataTable {
	return &DataTable{}
}
