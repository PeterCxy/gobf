package brainfuck

import (
	"testing"
	"fmt"
)

func TestHelloWorld(t *testing.T) {
	res, err := New().Exec(`++++++++++[>+++++++>++++++++++>+++>+<<<<-]
		>++.>+.+++++++..+++.>++.<<+++++++++++++++.
		>.+++.------.--------.>+.>.`)

	fmt.Println(res)

	if err != nil {
		t.Fail()
	}
}

func TestPlus(t *testing.T) {
	// 3 + 4
	res, err := New().Exec(`+++>++++[<+>-]++++++[<++++++++>-]<.`)

	fmt.Println(res)

	if err != nil {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	// 3 * 3
	res, err := New().Exec(`+++>+++[[<+>>+<-]>[<+>-]<-]++++++
		[<++++++++>-]<.`)

	fmt.Println(res)

	if err != nil {
		t.Fail()
	}
}
