//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how wrapping errors work with the stdlib.
package main

import (
	"errors"
	"fmt"
)

// AppError represents a custom error type.
type AppError struct {
	State int
}

// Error implements the error interface.
func (ae *AppError) Error() string {
	return fmt.Sprintf("App Error, State: %d", ae.State)
}

// IsAppError checks if an error of type AppError exists.
func IsAppError(err error) bool {
	var ae *AppError
	return errors.As(err, &ae)
}

// GetAppError returns a copy of the AppError pointer.
func GetAppError(err error) *AppError {
	var ae *AppError
	if !errors.As(err, &ae) {
		return nil
	}
	return ae
}

func main() {

	// Make the function call and validate the error.
	if err := firstCall(10); err != nil {

		// Check if the error is an AppError.
		if IsAppError(err) {
			ae := GetAppError(err)
			fmt.Printf("Is AppError, State: %d\n", ae.State)
		}

		fmt.Print("\n********************************\n\n")

		// Display the error using the implementation of
		// the error interface.
		fmt.Printf("%v\n", err)
	}
}

// firstCall makes a call to a second function and wraps any error.
func firstCall(i int) error {
	if err := secondCall(i); err != nil {
		return fmt.Errorf("firstCall->secondCall(%d) : %w", i, err)
	}
	return nil
}

// secondCall makes a call to a third function and wraps any error.
func secondCall(i int) error {
	if err := thirdCall(); err != nil {
		return fmt.Errorf("secondCall->thirdCall() : %w", err)
	}
	return nil
}

// thirdCall create an error value we will validate.
func thirdCall() error {
	return &AppError{99}
}
