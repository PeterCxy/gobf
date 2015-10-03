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

func TestInput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	} else {
		res, err := New().SetInput(func() (str string) {
			str = "test"
			return
		}).Exec(`++++[>,>++++++++[<---->-]<.<-]`)

		fmt.Println(res)

		if err != nil {
			t.Fail()
		}
	}
}
