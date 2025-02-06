package step2

import "testing"

func TestCalculateArea(t *testing.T){
  var width = 3
  var height = 5

  result, err := CalculateArea(width, height)
  if (result != 15){
    t.Fail()
  } 
  if (err != nil) {
    t.Fail()
  }
  
}
