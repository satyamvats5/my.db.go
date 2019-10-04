package utils

type (
	// InputBuffer is a structure used to store the command line input from user
	InputBuffer struct {
		Buffer       string
		BufferLength int
	}
	// MetaCommandResult is a boolean type used for storing Result of Meta Command Validation
	MetaCommandResult bool
	// SQLCommandResults is  a boolean type used for storing Result of SQL Command Validation
	SQLCommandResults bool
)

const (
	// MetaCommandRecognised will Recognise the valid Meta Commands
	MetaCommandRecognised MetaCommandResult = true
	// MetaCommandUnrecognised will Recognise the Invalid Meta Commands
	MetaCommandUnrecognised MetaCommandResult = false

	// SQLCommandRecognised will Recognise the valid SQL Commands
	SQLCommandRecognised SQLCommandResults = true
	// SQLCommandUnrecognised will Recognise the Invalid SQL Commands
	SQLCommandUnrecognised SQLCommandResults = false
)
