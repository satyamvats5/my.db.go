package cli

import (
	"MY_DB/utils"
	"fmt"
	"os"
	"strings"
)

// IsMetaCommand is used to validate that whether the given meta command is valid or not.
func isMetaCommand(inputbuffer utils.InputBuffer) utils.MetaCommandResult {
	if strings.Compare(inputbuffer.Buffer, ".exit") == 0 {
		fmt.Fprintf(os.Stdout, "Exitting Bye!")
		os.Exit(0)
	}

	return utils.MetaCommandUnrecognised
}
