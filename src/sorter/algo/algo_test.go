package algo

import "testing"

func TestBubble1(t *testing.T) {
	values := []int{5, 2, 1, 4, 3}
	BubboleSort(values)

	if string(values) != "12345" {
		t.Error("failed")
	}
}

func TestBubble2(t *testing.T) {
	values := []int{5, 5, 1, 4, 3}
	BubboleSort(values)

	if string(values) != "13455" {
		t.Error("failed")
	}
}
func TestQuick1(t *testing.T) {
	values := []int{5, 2, 1, 4, 3}
	QuickSort(values)

	if string(values) != "12345" {
		t.Error("failed")
	}
}

func TestQuick2(t *testing.T) {
	values := []int{5, 5, 1, 4, 3}
	QuickSort(values)

	if string(values) != "13455" {
		t.Error("failed")
	}
}
