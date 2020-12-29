package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func (t *testing.T) {
		numbers := [5]int {1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got '%d' expected '%d' given '%v'", got, expected, numbers)
		}
	})
	
	t.Run("collection of any numbers", func (t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum2(numbers)
		expected := 10

		if got != expected {
			t.Errorf("got '%d' expected '%d' given '%v'", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int {3, 7}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got '%v' expected '%v'", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func (t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		expected := []int {2, 4}
	
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got '%v' expected '%v'", got, expected)
		}
	})
	
	t.Run("safely sum empty slices", func (t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4})
		expected := []int {0, 4}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got '%v' expected '%v'", got, expected)
		}
	})
	
}