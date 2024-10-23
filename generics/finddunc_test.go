package main

import "testing"

// определите функцию Find
// ...
func Find[T comparable](s []T, toFind T) int {
	for i, e := range s {
		if e == toFind {
			return i
		}
	}
	return -1
}

func TestFind(t *testing.T) {
	if Find([]string{"bob", "alice", "mike"}, "alice") != 1 {
		t.Errorf("wrong Find (string)")
	}
	if Find([]float64{4.56, 7.89, 2.28}, 4.76) != -1 {
		t.Errorf("wrong Find (float)")
	}
	if Find([]int{34, 53, 90, 100, 34}, 34) != 0 {
		t.Errorf("wrong Find (int)")
	}
}
