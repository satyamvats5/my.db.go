package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func print_prompt() {
	fmt.Printf("db > ")
}

type (
	InputBuffer struct {
		buffer        string
		buffer_length int
	}
)

func read_input(scanner *bufio.Reader) InputBuffer {
	txt, err := scanner.ReadString('\n')
	var buffer InputBuffer
	if err == nil {
		buffer = InputBuffer{strings.TrimSpace(txt), len(txt)}
	} else {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if buffer.buffer_length < 0 {
		fmt.Fprintf(os.Stderr, "Error reading Input")
		os.Exit(1)
	}
	return buffer
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		print_prompt()
		switch buffer := read_input(scanner).buffer; buffer {
		case ".exit":
			fmt.Fprintf(os.Stdout, "Exitting Bye!")
			os.Exit(0)

		default:
			fmt.Fprintf(os.Stdout, "Unrecognised Command: %s.\n", buffer)
		}
	}
}
