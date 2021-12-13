package handlers

import (
	"fmt"
	"os"
)

// CheckError checks whether the error is nil or not, if so, prints the error and terminates the program.
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}
