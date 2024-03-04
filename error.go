package main

import (
	"fmt"
	"errors"
)

func main()  {
	err := errors.New("this is the new error")
	if err !=nil {
		fmt.Print(err)
	}
}