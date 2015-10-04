package brainfuck

import (
	"testing"
	"fmt"
	"time"
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
		res, err := New().SetInput(func(out string) (str string) {
			str = "test"
			return
		}).Exec(`++++[>,>++++++++[<---->-]<.<-]`)

		fmt.Println(res)

		if err != nil {
			t.Fail()
		}
	}
}

func TestInterrupt(t *testing.T) {
	if testing.Short() {
		t.Skip()
	} else {
		end := time.Now().Add(5 * time.Second)
		res, err := New().SetInterrupter(func() bool {
			return time.Now().After(end)
		}).Exec(`+[].`)

		fmt.Println(res)

		if err != nil {
			fmt.Println(err)
		}
	}
}
