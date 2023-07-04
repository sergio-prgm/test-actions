package pkg

import (
	"errors"
	"fmt"
)

func ReturnNumber(correct bool) (int, error) {
	if correct {
		return 10, nil
	} else {
		return 0, errors.New("Incorrect")
	}
}

func GreetPerson(name string) string {
	if name == "" {
		return "I was expecting a name, but I'm gonna call you Bob"
	} else {
		return fmt.Sprintf("Hello %s, nice meeting you!", name)
	}
}
