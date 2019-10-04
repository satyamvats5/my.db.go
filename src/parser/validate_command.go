package parser

import (
	"MY_DB/utils"
	"strings"
)

// IsSQLCommand Will validate whether a givencommand is SQL COmmand or not.
func IsSQLCommand(inputbuffer utils.InputBuffer) utils.SQLCommandResults {
	if strings.Compare(inputbuffer.Buffer, "insert") == 0 {
		return utils.SQLCommandRecognised
	}
	if strings.Compare(inputbuffer.Buffer, "select") == 0 {
		return utils.SQLCommandRecognised
	}

	return utils.SQLCommandUnrecognised
}
