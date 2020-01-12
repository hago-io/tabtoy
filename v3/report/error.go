package report

import (
	"fmt"
	"strings"
)

type TableError struct {
	ID string

	context []interface{}
}

func getErrorDesc(id string) string {

	if lan, ok := ErrorByID[id]; ok {
		return lan.CHS
	}

	return ""
}

func (err *TableError) Error() string {

	var sb strings.Builder

	sb.WriteString("TableError.")
	sb.WriteString(err.ID)
	sb.WriteString(" ")
	sb.WriteString(getErrorDesc(err.ID))
	sb.WriteString(" | ")

	for index, c := range err.context {
		if index > 0 {
			sb.WriteString(" ")
		}

		sb.WriteString(fmt.Sprintf("%+v", c))
	}

	return sb.String()
}

func Error(id string, context ...interface{}) {

	panic(&TableError{
		ID:      id,
		context: context,
	})
}
