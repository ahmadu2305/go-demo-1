package mascot

import "testing"

func TestBestMascot(t *testing.T) {
	expected := "Go Gopher"
	result := BestMascot()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
