package build

import "testing"

func TestBulider1(t *testing.T) {
	builder := &Builder1{}
	direct := NewDirector(builder)
	direct.Construct()
	res := builder.GetResult()
	if res != "123" {
		t.Fatalf("builder1 result is not string %s", res)
	}
}

func TestBulider2(t *testing.T) {
	builder := &Builder2{}
	direct := NewDirector(builder)
	direct.Construct()
	res := builder.GetResult()
	if res != 6 {
		t.Fatalf("builder1 result is not int %d", res)
	}
}
