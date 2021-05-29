package main

import (
	"sampleGoWebProject/tmp"
	"testing"
)

func TestSum(t *testing.T)  {
	sum := tmp.Sum(1, 2)
	if sum != 3 {
		t.Errorf("Expected {%d}, found {%d}", 3, sum)
	}
}
