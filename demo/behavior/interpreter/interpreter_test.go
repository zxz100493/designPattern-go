package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("5 + 2 - 3")
	fmt.Printf("--test--%s\n", p.prev)

	res := p.Result().Interpret()
	except := 1
	if res != except {
		t.Fatalf("expect %d, got %d", except, res)
	}
}
