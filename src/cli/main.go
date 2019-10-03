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

type (
	meta_command_result bool
	sql_command_results bool
)

const (
	// Meta commands
	meta_command_recognised   meta_command_result = true
	meta_command_unrecognized meta_command_result = false

	// SQL Commands
	sql_command_recognised   sql_command_results = true
	sql_command_unrecognised sql_command_results = false
)

func is_meta_command(inputbuffer *InputBuffer) meta_command_result {
	if strings.Compare(inputbuffer.buffer, ".exit") == 0 {
		fmt.Fprintf(os.Stdout, "Exitting Bye!")
		os.Exit(0)
	}

	return meta_command_unrecognized

}

func is_sql_command(inputbuffer *InputBuffer) sql_command_results {
	if strings.Compare(inputbuffer.buffer, "insert") == 0 {
		return sql_command_recognised
	}
	if strings.Compare(inputbuffer.buffer, "select") == 0 {
		return sql_command_recognised
	}

	return sql_command_unrecognised
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
			case meta_command_recognised:
				continue
			case meta_command_unrecognized:
				fmt.Fprintf(os.Stdout, "Unrecognised Meta Command: %s.\n", input_buffer.buffer)
				continue
			}
		}

		switch is_sql := is_sql_command(&input_buffer); is_sql {
		case sql_command_recognised:
			fmt.Println("Executing SQL Command")
			break
		case sql_command_unrecognised:
			fmt.Fprintf(os.Stdout, "Unrecognised SQL Command: %s.\n", input_buffer.buffer)
			continue
		}

	}
}
