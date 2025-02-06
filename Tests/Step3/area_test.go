package step3

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	width, length := 3, 5
	expected := 15

	area, err := CalculateArea(width, length)

	if area != expected {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}

func TestCalculateAreaNegativeWidth(t *testing.T){
  var width = -4
  var height = 6
  
  result, err := CalculateArea(width, height)
  
  if (result != 0) {
    t.Fail()
  }
  if (err == nil) {
    t.Fail()
  }
}

func TestCalculateAreaNegativeLength(t *testing.T){
  var width = 5
  var height = -7
  result, err := CalculateArea(width, height)

  if (result != 0) {
   t.Errorf("wrong value for 'area', expected: %v, got %v", 0, result)
  }

  if(err == nil) {
    t.Error("expected an error")
  }
  
}
