package serialhandler

import "context"
import "strings"

var arduinoUnitErrorWords = []string{
	"Failed",
}

// HandlerInterface is the interface for a SerialHandler
type HandlerInterface interface {
	Write()
	Output() string
	Errors() []string
	DisableErrorCheck()
	SetErrorWords(l []string)
}

// SerialHandler handles serial output from the Arduino to generate results
type SerialHandler struct {
	errors         []string
	output         string
	cancelFunc     context.CancelFunc
	checkForErrors bool
	errorWords     []string
}

// New generates a new SerialHandler
func New(cancel context.CancelFunc) SerialHandler {
	return SerialHandler{
		cancelFunc:     cancel,
		checkForErrors: true,
		errorWords:     arduinoUnitErrorWords,
	}
}

func (s *SerialHandler) Write(p []byte) (n int, err error) {
	n = len(p)
	s.output += string(p)

	for _, c := range p {
		if c == 0x07 { // C89 alert bell
			if s.checkForErrors {
				s.errors = s.checkErrors()
			}
			s.cancelFunc()
		}
	}

	return
}

func (s *SerialHandler) checkErrors() []string {
	out := []string{}

	lines := strings.Split(s.Output(), "\n")
	for _, line := range lines {
	L:
		for _, word := range s.errorWords {
			if strings.Contains(line, word) {
				out = append(out, line)
				break L
			}
		}
	}

	return out
}

// Output gives a string of the current sent output
func (s *SerialHandler) Output() string {
	return s.output
}

// Errors gives back the errors found
func (s *SerialHandler) Errors() []string {
	return s.errors
}

// DisableErrorCheck turns of checking errors on \a
func (s *SerialHandler) DisableErrorCheck() {
	s.checkForErrors = false
}

// SetErrorWords allows to replace the words that trigger errors
func (s *SerialHandler) SetErrorWords(l []string) {
	s.errorWords = l
}
