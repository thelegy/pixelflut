package pixelflut

// New returns an error that formats as the given text.
func newError(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
// It is here only to drop the errors dependency.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

var (
	// ParsingError can be returned from the parsing functions.
	ParsingError = newError("Error Parsing")
)
