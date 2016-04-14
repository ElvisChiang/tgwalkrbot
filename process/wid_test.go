package process

import (
	"fmt"
	"testing"
)

func TestGetWidURL(t *testing.T) {
	cases := []struct {
		in    string
		found bool
	}{
		{"    xoCTJMzNAyz\numWyuC7stYx\nj3rzxDoTUMk    \n", true},
		{"q    xoCTJMzNAyz\numWyuC7stYx\nj3rzxDoTUMk    \n", false},
		{"   xoCTJMzNAyz  umWyuC7stYx\nj3rzxDoTUMk    \n", true},
	}

	fmt.Println("----------------")
	for _, c := range cases {
		_, found := GetWidURL(c.in)
		if found != c.found {
			t.Errorf("test fail for `%s`", c.in)
		}
	}
}
