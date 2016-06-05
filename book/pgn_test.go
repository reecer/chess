package book

import (
	//"fmt"
	"fmt"
	"github.com/andrewbackes/chess/pgn"
	"strings"
	"testing"
)

/*
func TestFromPGN(t *testing.T) {
	input := `[Event "one"]
[Round "1"]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0

[Event "two"]
[Round "2"]
[Result "1/2-1/2"]

1. e2e4 e7e5 2. d1h5 e8e7 1/2-1/2
`
	pgns, _ := pgn.Open(strings.NewReader(input))
	book, _ := FromPGN(pgns, 4)
	fmt.Println(book)
}
*/

func TestCorrectNumber(t *testing.T) {
	input := `[Event "one"]
[Round "1"]
[Result "1-0"]

1. e2e4 e7e5 2. d1h5 e8e7 3. h5e5 1-0
`
	pgns, _ := pgn.Open(strings.NewReader(input))
	book, _ := FromPGN(pgns, 4)

	if len(book.Positions) != 4 {
		t.Log(book)
		t.Fail()
	}
	for _, m := range book.Positions {
		if len(m) != 1 {
			t.Log(book)
			t.Fail()
		}
	}
}
