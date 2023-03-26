package adapter

import "testing"

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
	adapter := NewAdaptee()
	target := NewAdapter(adapter)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect:%s,actual:%s", expect, res)
	}
}
