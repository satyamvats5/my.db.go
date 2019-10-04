package cli

import (
	"MY_DB/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(scanner *bufio.Reader) utils.InputBuffer {
	txt, err := scanner.ReadString('\n')
	var buffer utils.InputBuffer
	if err == nil {
		buffer = utils.InputBuffer{strings.TrimSpace(txt), len(txt)}
	} else {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if buffer.BufferLength < 0 {
		fmt.Fprintf(os.Stderr, "Error reading Input")
		os.Exit(1)
	}
	return buffer
}
