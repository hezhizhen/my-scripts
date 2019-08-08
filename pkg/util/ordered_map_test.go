package util

import "testing"

func TestOrderedMap(t *testing.T) {
	m := New()
	if err := m.Set("a", "1"); err != nil {
		t.Error(err)
	}
	m.Pretty()
}
