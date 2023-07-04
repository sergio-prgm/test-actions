package pkg

import (
	"strings"
	"testing"
)

func TestReturnNumberTrue (t *testing.T) {
	input := true
	result, err := ReturnNumber(input)		

	if err != nil || result <= 0 {
		
		t.Fatalf("ReturnNumber(%t) -> %d, %v", input, result, err)	
	}
}

func TestGreetingName(t *testing.T)  {
	input := "John"	
	greeting := GreetPerson(input)

	if !strings.Contains(greeting, "Hello John") {
		t.Fatalf("GreetPerson(%s) -> %s", input, greeting)
	}
}

func TestGreetingEmpty(t *testing.T) {
	input := ""	
	greeting := GreetPerson(input)

	if !strings.Contains(greeting, "but I'm gonna call you Bob") {
		t.Fatalf("GreetPerson(%s) -> %s", input, greeting)
	}
}
