package cli

import (
	"MY_DB/parser"
	"MY_DB/utils"
	"bufio"
	"fmt"
	"os"
)

// LoadCli is loading the Command Line INterface for out DB.
func LoadCli() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		printPrompt()
		inputBuffer := readInput(scanner)
		if inputBuffer.Buffer[0] == '.' {
			switch result := isMetaCommand(inputBuffer); result {
			case utils.MetaCommandRecognised:
				continue
			case utils.MetaCommandUnrecognised:
				fmt.Fprintf(os.Stdout, "Unrecognised Meta Command: %s.\n", inputBuffer.Buffer)
				continue
			}
		}

		switch isSQL := parser.IsSQLCommand(inputBuffer); isSQL {
		case utils.SQLCommandRecognised:
			fmt.Println("Executing SQL Command")
			break
		case utils.SQLCommandUnrecognised:
			fmt.Fprintf(os.Stdout, "Unrecognised SQL Command: %s.\n", inputBuffer.Buffer)
			continue
		}

	}
}
