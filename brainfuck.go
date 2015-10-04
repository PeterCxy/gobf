// Brainfuck implementation in Golang
package brainfuck

import "errors"

const (
	ShiftR = '>'
	ShiftL = '<'
	OnePlus = '+'
	OneMinus = '-'
	LoopL = '['
	LoopR = ']'
	Print = '.'
	Read = ','
)

type Brainfuck struct {
	stack []byte
	cursor int
	buffer string
	input func(string) string
	interrupter func() bool
}

// Create a new instance of Brainfuck
func New() *Brainfuck {
	return &Brainfuck{
		stack: make([]byte, 1),
		cursor: 0,
	}
}

// Set an input source
// Everytime a input is received, it will be stored to the buffer
// When reading, the buffer will be poped out byte by byte
// If the buffer has run out, the input function will be called again.
// Everytime an input is required, the output buffer will be cleared
func (this *Brainfuck) SetInput(i func(string) string) *Brainfuck {
	this.input = i
	return this
}

// A Interrupter
// When returning true, the current execution will be interrupted
func (this *Brainfuck) SetInterrupter(i func() bool) *Brainfuck {
	this.interrupter = i
	return this
}

// Execute Brainfuck code
func (this *Brainfuck) Exec(code string) (out string, err error) {
	for i := 0; i < len(code); i++ {
		if (this.interrupter != nil) && this.interrupter() {
			err = errors.New("Interrupted")
			return
		}

		if (this.cursor >= len(this.stack)) || (this.cursor < 0) {
			err = errors.New("Pointer out of range.")
			return
		}

		char := code[i]
		switch char {
			case ShiftR:
				this.cursor++
				this.realloc()
			case ShiftL:
				this.cursor--
				this.realloc()
			case OnePlus:
				this.stack[this.cursor]++
			case OneMinus:
				this.stack[this.cursor]--
			case LoopL:
				if this.stack[this.cursor] == 0 {
					i = this.findLoopR(code, i, 0)

					if i < 0 {
						err = errors.New("Loop mismatch")
						return
					}
				}
			case LoopR:
				if this.stack[this.cursor] != 0 {
					i = this.findLoopL(code, i, 0)

					if i < 0 {
						err = errors.New("Loop mismatch")
						return
					}
				}
			case Print:
				out += string(this.stack[this.cursor])
			case Read:
				this.stack[this.cursor] = this.readInput(&out)
		}
	}

	return
}

func (this *Brainfuck) realloc() {
	if this.cursor >= 30000 {
		this.cursor = this.cursor - 30000
	} else if this.cursor < 0 {
		this.cursor = 30000 + this.cursor
	}

	if this.cursor >= len(this.stack) {
		for {
			this.stack = append(this.stack, 0)
			if this.cursor < len(this.stack) {
				break
			}
		}
	}
}

func (this *Brainfuck) readInput(out *string) byte {
	if len(this.buffer) == 0 {
		if this.input != nil {
			this.buffer = this.input(*out)
			*out = ""
		} else {
			return 0
		}
	}

	return this.bufferPop()
}

func (this *Brainfuck) bufferPop() (b byte) {
	b = this.buffer[0]
	this.buffer = this.buffer[1:]
	return
}

func (this *Brainfuck) findLoopR(code string, i int, count int) int {
	if count >= 1024 {
		return -1
	}

	for i = i - 1; i < len(code); i++ {
		if i < 0 {
			return -1
		}

		char := code[i]
		switch char {
			case LoopL:
				i = this.findLoopR(code, i, count + 1)

				if i == -1 {
					return -1
				}
			case LoopR:
				return i
		}
	}

	return -1
}

func (this *Brainfuck) findLoopL(code string, i int, count int) int {
	if count >= 1024 {
		return -1
	}

	for i = i - 1; i >= 0; i-- {
		char := code[i]
		switch char {
			case LoopR:
				i = this.findLoopL(code, i, count + 1)

				if i == -1 {
					return -1
				}
			case LoopL:
				return i
		}
	}

	return -1
}
