package utils

import "testing"

func TestNewID(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		id := NewID()
		if m[id] {
			t.Errorf("ID %s is not unique", id)
		}
		m[id] = true
	}
}
