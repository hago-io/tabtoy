package compiler

import (
	"github.com/ahmetb/go-linq"
	"github.com/davyxu/tabtoy/v3/helper"
	"github.com/davyxu/tabtoy/v3/model"
	"github.com/davyxu/tabtoy/v3/report"
)

func LoadTypeTable(typeTab *model.TypeTable, indexGetter helper.FileGetter, fileName string) error {

	tabs, err := LoadDataTable(indexGetter, fileName, "TypeDefine", "TypeDefine", typeTab)

	if err != nil {
		return err
	}

	for _, tab := range tabs {

		//resolveHeaderFields(tab, "TypeDefine", typeTab)

		for row := 1; row < len(tab.Rows); row++ {

			var objType model.TypeDefine

			if !ParseRow(&objType, tab, row, typeTab) {
				continue
			}

			if typeTab.FieldByName(objType.ObjectType, objType.FieldName) != nil {

				cell := tab.GetValueByName(row, "字段名")

				if cell != nil {
					report.Error("DuplicateTypeFieldName", cell.String(), objType.ObjectType, objType.FieldName)
				} else {
					report.Error("InvalidTypeTable", objType.ObjectType, objType.FieldName, tab.FileName)
				}

			}

			typeTab.AddField(&objType, tab, row)
		}

	}

	return nil
}

func typeTableCheckEnumValueEmpty(typeTab *model.TypeTable) {
	linq.From(typeTab.Raw()).WhereT(func(td *model.TypeData) bool {

		return td.Define.Kind == model.TypeUsage_Enum && td.Define.Value == ""
	}).ForEachT(func(td *model.TypeData) {

		cell := td.Tab.GetValueByName(td.Row, "值")

		report.Error("EnumValueEmpty", cell.String())
	})
}

func typeTableCheckDuplicateEnumValue(typeTab *model.TypeTable) {

	type NameValuePair struct {
		Name  string
		Value string
	}

	checker := map[NameValuePair]*model.TypeData{}

	for _, td := range typeTab.Raw() {

		if td.Define.IsBuiltin || td.Define.Kind != model.TypeUsage_Enum {
			continue
		}

		key := NameValuePair{td.Define.ObjectType, td.Define.Value}

		if _, ok := checker[key]; ok {

			cell := td.Tab.GetValueByName(td.Row, "值")

			report.Error("DuplicateEnumValue", cell.String())
		}

		checker[key] = td
	}
}

func CheckTypeTable(typeTab *model.TypeTable) {

	typeTableCheckEnumValueEmpty(typeTab)

	typeTableCheckDuplicateEnumValue(typeTab)
}
