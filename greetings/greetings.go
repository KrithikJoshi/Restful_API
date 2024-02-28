package greetings

import ("fmt"
		"errors"

)

func Hello( name string) (string, error){
	//no name given return error message
	if name == "" {
		return "" ,errors.New("empty name")
	}
	message:=fmt.Sprintf("Hi,%v.Welcome!",name)
	return message, nil
}
//func Hello( name string) {string}-> this datatype is of returning value 