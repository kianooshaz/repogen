package parser

// Define constants for various special characters.
const (
	minus          = rune(45) // ASCII value of '-'
	semicolon      = rune(59) // ASCII value of ';'
	tab            = rune(9)  // ASCII value of tab character
	newLine        = rune(10) // ASCII value of newline character
	carriageReturn = rune(13) // ASCII value of carriage return character
)

// isNewLine checks if the given rune represents a newline character.
func isNewLine(r rune) bool {
	return r == newLine || r == carriageReturn
}
