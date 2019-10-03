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

type meta_command_result bool

const (
	meta_command_success              meta_command_result = true
	meta_command_unrecognized_command meta_command_result = false
)

func is_meta_command(inputbuffer *InputBuffer) meta_command_result {
	if strings.Compare(inputbuffer.buffer, ".exit") == 0 {
		fmt.Fprintf(os.Stdout, "Exitting Bye!")
		os.Exit(0)
	}

	return meta_command_unrecognized_command

}

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
		input_buffer := read_input(scanner)
		if input_buffer.buffer[0] == '.' {
			switch result := is_meta_command(&input_buffer); result {
			case meta_command_success:
				continue
			case meta_command_unrecognized_command:
				fmt.Fprintf(os.Stdout, "Unrecognised Command: %s.\n", input_buffer.buffer)
				continue
			}
		}
	}
}
