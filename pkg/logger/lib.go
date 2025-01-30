package logger

import "fmt"

// Logger Lib
// This package provides utility functions for joining and formatting values into strings.
// It is typically used to prepare log messages for consistent formatting and output.

// ValuesJoin concatenates a variadic list of values into a single string.
//
// Parameters:
// - values: A variadic parameter containing any type of values to join.
//
// Returns:
// - A single string that is the concatenation of all input values.
//
// Example:
//    ValuesJoin("Hello, ", "World!") -> "Hello, World!"
func ValuesJoin(values ...any) string {
	var result string
	for _, text := range values {
		result += fmt.Sprint(text) // Convert each value to a string and append to the result.
	}
	return result
}

// ValuesJoinf formats a string with the provided values using a format string.
//
// Parameters:
// - text: A format string specifying the desired output format (e.g., "Hello, %s!").
// - values: A variadic parameter containing the values to replace placeholders in the format string.
//
// Returns:
// - A formatted string with placeholders replaced by the provided values.
//
// Example:
//    ValuesJoinf("Hello, %s!", "World") -> "Hello, World!"
func ValuesJoinf(text string, values ...any) string {
	result := fmt.Sprintf(text, values...) // Format the string using fmt.Sprintf.
	return result
}
