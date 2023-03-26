package singleton

import (
	"sync"
	"testing"
)

const parCount = 60

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 == ins2 {
		t.Fatal("the same")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instance := [parCount]*Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instance[index] = GetInstance()
			wg.Done()
		}(i)

	}

	wg.Wait()
	for i := 1; i < parCount; i++ {
		t.Log(instance[i])
		if instance[i] != instance[i-1] {
			t.Fatal("not same")
		}
	}
}
