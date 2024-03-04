package exercises

import (
	"reflect"
	"testing"
)

func TestNumbers_Init(t *testing.T) {
	n := Numbers{}
	n.Init()

	if n.Value == nil {
		t.Error("Expected numbers to be initialized")
	}

	if !reflect.DeepEqual(n.Value, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected numbers to be initialized with 1 to 10 but got %v", n.Value)
	}
}

func TestNumbers_Ascending(t *testing.T) {
	n := Numbers{}
	n.Init()

	ascending := n.Ascending()

	if !reflect.DeepEqual(ascending, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected ascending order to be 1 to 10 but got %v", ascending)
	}
}

func TestNumbers_Descending(t *testing.T) {
	n := Numbers{}
	n.Init()

	descending := n.Descending()

	if !reflect.DeepEqual(descending, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}) {
		t.Errorf("Expected descending order to be 10 to 1 but got %v", descending)
	}
}

// TODO - even and odd count sequences
