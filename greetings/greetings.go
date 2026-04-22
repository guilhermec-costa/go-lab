package greetings

import (
	"fmt"
	"example.com/greetings/v3"
	"errors"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Name is empty");
	}
	var message string;
	message = fmt.Sprintf("Hi, %v. Welcome!", name);

	fmt.Println(message);

	end := "End of greetings modified";
	fmt.Println(end);

	v3.HelloV3(name);
	return message, nil;
}