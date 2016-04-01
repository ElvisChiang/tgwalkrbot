package process

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	cases := []struct {
		in     string
		result bool
	}{
		//		{"/wn @elvisfb"},
		{"/wn \\T'Elvis Chiang' tg name \\W'Nicole Lai' walkr name @elvisfb codename", true},
		{"/wn \\W'Elvis' walkr name", true},
		{"/wp 1", true},
		{"/wp 地球", true},
		{"/wp 瓦肯", false},
	}

	fmt.Println("----------------")
	for _, c := range cases {
		output, ok := Command(c.in)
		if ok != c.result {
			t.Errorf("cannot process " + c.in)
			continue
		}
		fmt.Printf("%s -> `%s`\n", c.in, output)
		fmt.Println("----------------")
	}
}
