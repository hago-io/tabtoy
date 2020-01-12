package model

type DataTableList struct {
	data []*DataTable
}

func (tl *DataTableList) GetDataTable(headerType string) *DataTable {

	for _, tab := range tl.data {
		if tab.HeaderType == headerType {
			return tab
		}
	}

	return nil
}

func (tl *DataTableList) AddDataTable(t *DataTable) {
	tl.data = append(tl.data, t)
}
func (tl *DataTableList) AllTables() []*DataTable {
	return tl.data
}

func (tl *DataTableList) Count() int {
	return len(tl.data)
}
