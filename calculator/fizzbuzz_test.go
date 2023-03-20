package calculator_test

import (
	"GO_workshop/calculator"
	"testing"
)

func TestFizzBuzz(t *testing.T){
	
	type testcase struct {
		give int 
		want string
	}
	tests := []testcase {
		{give:1, want:"1"},
		{give:2, want:"2"},
		{give:3, want:"Fizz"},
		{give:6, want:"Fizz"},
		{give:9, want:"Fizz"},
		{give:5, want:"Buzz"},
		{give:10, want:"Buzz"},
		{give:15, want:"FizzBuzz"},
		{give:30, want:"FizzBuzz"},
	}


	for _, tests := range tests {
       got := calculator.FizzBuzz(tests.give)		
			 
			 if got != tests.want {
				t.Errorf("give %v want %v but got %v", tests.give , tests.want, got)
			}
	}

	

	
}