package cortana

import (
	"bufio"
	"fmt"
	"io"
)

// Sentient is a general writer object that allows easy
// control of output streams
type Sentient struct {
}

// New will return a new Sentient object
func New() *Sentient {
	return &Sentient{}
}

// Follow will continuelly querry the stream given
// and stop once the Reader has closed
func (s *Sentient) Follow(stream io.ReadCloser) {
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
