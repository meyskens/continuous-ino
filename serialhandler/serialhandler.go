package serialhandler

import "context"

// HandlerInterface is the interface for a SerialHandler
type HandlerInterface interface {
	Write()
	Output() string
	Errors() []string
}

// SerialHandler handles serial output from the Arduino to generate results
type SerialHandler struct {
	errors     []string
	output     string
	cancelFunc context.CancelFunc
}

// New generates a new SerialHandler
func New(cancel context.CancelFunc) SerialHandler {
	return SerialHandler{
		cancelFunc: cancel,
	}
}

func (s *SerialHandler) Write(p []byte) (n int, err error) {
	n = len(p)
	s.output += string(p)

	for _, c := range p {
		if c == 0x07 { // C89 alert bell
			s.cancelFunc()
		}
	}

	return
}

// Output gives a string of the current sent output
func (s *SerialHandler) Output() string {
	return s.output
}

// Errors gives back the errors found
func (s *SerialHandler) Errors() []string {
	return s.errors
}
