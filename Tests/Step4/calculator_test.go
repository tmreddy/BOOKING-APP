package step4

import (
	"log"
	"os"
	"testing"
)

var c Calculator

func TestMain(m *testing.M){
  log.Println("Preparation")
  c = Calculator{}
  exitVal := m.Run()
  log.Println("Cleanup")

  os.Exit(exitVal)
}

func TestAdd(t *testing.T) {
    t.Run("Add", func(t *testing.T){
        c := Calculator{}
  
    	expected := 5.0
  
    	result := c.Add(2, 3)
  
    	if result != expected {
    		t.Fail()
    	}  
    })
}

func TestSubtract(t *testing.T) {
    t.Run("Subtract", func(t *testing.T){
  
    	c := Calculator{}
    
    	expected := 2.0
    
    	result := c.Subtract(5, 3)
    
    	if result != expected {
    		t.Fail()
    	}
      })
}

func TestMultiply(t *testing.T) {
    t.Run("Multiply", func(t *testing.T){
  
    	c := Calculator{}

    	expected := 6.0

    	result := c.Multiply(2, 3)

    	if result != expected {
    		t.Fail()
    	}
    })
}

func TestDivideValid(t *testing.T) {
    t.Run("DivideValid", func(t *testing.T){
  
    	c := Calculator{}

    	expected := 2.0

    	result, err := c.Divide(6, 3)

    	if err != nil {
    		t.Fail()
    	}
  
    	if result != expected {
    		t.Fail()
    	}
        t.Run("DivideByZero", func(t *testing.T){
  
    	c := Calculator{}

    	_, err := c.Divide(6, 0)

    	if err == nil {
    		t.Fail()
    	}
    })

    })
}

func TestDivideByZero(t *testing.T) {
    t.Run("DivideByZero", func(t *testing.T){
  
    	c := Calculator{}

    	_, err := c.Divide(6, 0)

    	if err == nil {
    		t.Fail()
    	}
    })
}
