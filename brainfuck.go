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
}

// Create a new instance of Brainfuck
func New() *Brainfuck {
	return &Brainfuck{
		stack: make([]byte, 1),
		cursor: 0,
	}
}

// Execute Brainfuck code
func (this *Brainfuck) Exec(code string) (out string, err error) {
	for i := 0; i < len(code); i++ {
		char := code[i]
		switch char {
			case ShiftR:
				this.cursor++
				this.realloc()
			case ShiftL:
				this.cursor--
			case OnePlus:
				this.stack[this.cursor]++
			case OneMinus:
				this.stack[this.cursor]--
			case LoopL:
				if this.stack[this.cursor] == 0 {
					i = this.findLoopR(code, i)

					if i < 0 {
						err = errors.New("Loop mismatch")
						return
					}
				}
			case LoopR:
				if this.stack[this.cursor] != 0 {
					i = this.findLoopL(code, i)

					if i < 0 {
						err = errors.New("Loop mismatch")
						return
					}
				}
			case Print:
				out += string(this.stack[this.cursor])
			case Read:
				// Not implemented
		}
	}

	return
}

func (this *Brainfuck) realloc() {
	if this.cursor >= len(this.stack) {
		for {
			this.stack = append(this.stack, 0)
			if this.cursor < len(this.stack) {
				break
			}
		}
	}
}

func (this *Brainfuck) findLoopR(code string, i int) int {
	for i = i - 1; i < len(code); i++ {
		if i < 0 {
			return -1
		}

		char := code[i]
		switch char {
			case LoopL:
				i = this.findLoopR(code, i)
			case LoopR:
				return i
		}
	}

	return -1
}

func (this *Brainfuck) findLoopL(code string, i int) int {
	for i = i - 1; i > 0; i-- {
		char := code[i]
		switch char {
			case LoopR:
				i = this.findLoopL(code, i)
			case LoopL:
				return i
		}
	}

	return -1
}
