package greetings

import (
	"errors"
	"fmt"
	"math/rand" 
)
//here math/rand is the function that is used generate the random values for the string that is being used here the 
//greetings 3-4 lines are being generated randomly whenever a new user will come

func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    // message := fmt.Sprintf(randomFormat(), name)
    message := fmt.Sprint(randomFormat())
    return message, nil
}
//func Hello( name string) {string}-> this datatype is of returning value 

func Hellos(names []string) (map[string]string,error) {
	messages := make(map[string]string)
	for _, name := range names {
		message , err := Hello(name)
		if err != nil {
			return nil ,err
		}
		messages[name] = message
	}
	return messages ,nil
}

func randomFormat() string {
	formats :=[] string {
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}
	
	return formats[rand.Intn(len(formats))]
}
//empty commit
//empty commit
// hello fucntion not working