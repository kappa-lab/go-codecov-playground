package main

import "fmt"

func main() {
	fmt.Println("Hello Universe")
}

func IsTrueByString(value string) bool {
	if value == "true" {
		return true
	}

	if value == "TRUE" {
		return true
	}

	if value == "True" {
		return true
	}

	return false
}
