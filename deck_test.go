package main
import "testing"
func TestNewDeck(t *testing.T){
	d := newDeck()
	if len(d) != 16{
		t.Errorf("Expected length of 20, but got ", len(d))
	}
}