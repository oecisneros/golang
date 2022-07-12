package main

import (
	"errors"
	"fmt"
)

func toBoolean(s string) (bool, error) {
	switch s {
	case "true", "yes", "1":
		return true, nil
	case "false", "no", "0":
		return false, nil
	}
	return false, errors.New(fmt.Sprint("Invalid input:", s))
}

// go run input.go
func main() {
	fmt.Println(toBoolean("true"))
	fmt.Println(toBoolean("false"))
	fmt.Println(toBoolean("yes"))
	fmt.Println(toBoolean("no"))
	fmt.Println(toBoolean("1"))
	fmt.Println(toBoolean("0"))
	fmt.Println(toBoolean("True"))
	fmt.Println(toBoolean("False"))
	fmt.Println(toBoolean(""))
}
