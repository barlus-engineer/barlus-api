// Package logger provides utility functions for logging messages with different severity levels (info, warning, alert, fatal).
// It includes support for formatted and unformatted messages, color-coded output, and timestamped logs.
// Fatal functions terminate the program execution immediately after logging the message.

package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/barlus-engineer/barlus-api/pkg/colors"
)

/*
=====================================================================================================
= Helper Functions
=====================================================================================================
*/

// Sinfo generates a string formatted for informational logging.
// It includes a timestamp, "info" label, and color-coded output.
//
// Parameters:
// - values: A variadic parameter containing the values to log.
//
// Returns:
// - A string formatted for informational logging.
func Sinfo(values ...any) string {
	value := ValuesJoin(values...)
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	result := fmt.Sprintf("%s [%sinfo%s] %s", currentTime, colors.Blue, colors.Reset, value)
	return result
}

// Swarning generates a string formatted for warning logging.
// It includes a timestamp, "yellow" label, and color-coded output.
//
// Parameters:
// - values: A variadic parameter containing the values to log.
//
// Returns:
// - A string formatted for warning logging.
func Swarning(values ...any) string {
	value := ValuesJoin(values...)
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	result := fmt.Sprintf("%s [%swarning%s] %s", currentTime, colors.Yellow, colors.Reset, value)
	return result
}

// Salert generates a string formatted for alert logging.
// It includes a timestamp, "alert" label, and color-coded output.
//
// Parameters:
// - values: A variadic parameter containing the values to log.
//
// Returns:
// - A string formatted for alert logging.
func Salert(values ...any) string {
	value := ValuesJoin(values...)
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	result := fmt.Sprintf("%s [%salert%s] %s", currentTime, colors.Red, colors.Reset, value)
	return result
}

// sCrash generates a string formatted for fatal logging.
// It includes a timestamp, "Crash" label, and color-coded output.
//
// Parameters:
// - values: A variadic parameter containing the values to log.
//
// Returns:
// - A string formatted for fatal logging.
func sCrash(values ...any) string {
	value := ValuesJoin(values...)
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	result := fmt.Sprintf("%s [%scrash%s] %s", currentTime, colors.Red, colors.Reset, value)
	return result
}

/*
=====================================================================================================
= Logging Functions
=====================================================================================================
*/

// Info prints an info log to the console.
//
// Parameters:
// - values: A variadic parameter containing the values to log.
//
// Example:
//
//	logger.Info("Server started successfully")
func Info(values ...any) {
	text := Sinfo(values...)
	fmt.Println(text)
}

// Warning prints a warning log to the console.
//
// Parameters:
// - values: A variadic parameter containing the values to log.
//
// Example:
//
//	logger.Warning("Memory usage is high")
func Warning(values ...any) {
	text := Swarning(values...)
	fmt.Println(text)
}

// Alert prints an alert log to the console.
//
// Parameters:
// - values: A variadic parameter containing the values to log.
//
// Example:
//
//	logger.Alert("Database connection lost")
func Alert(values ...any) {
	text := Salert(values...)
	fmt.Println(text)
}

// Fatal logs a critical error message to the console and immediately terminates the program.
//
// Parameters:
// - values: A variadic parameter containing the values to log.
//
// Example:
//
//	logger.Crash("Unable to connect to database")
func Crash(values ...any) {
	text := sCrash(values...)
	fmt.Println(text)
	os.Exit(1)
}

/*
=====================================================================================================
= Formatted Logging Functions
=====================================================================================================
*/

// Sinfof generates a formatted string for informational logging.
// It includes a timestamp, "info" label, and color-coded output.
//
// Parameters:
// - text: A string containing the format for the log message.
// - values: A variadic parameter containing arguments to replace placeholders in the format string.
//
// Returns:
// - A formatted string for informational logging.
func Sinfof(text string, values ...any) string {
	value := ValuesJoinf(text, values...)
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	result := fmt.Sprintf("%s [%sinfo%s] %s", currentTime, colors.Blue, colors.Reset, value)
	return result
}

// Swarningf generates a formatted string for warning logging.
// It includes a timestamp, "yellow" label, and color-coded output.
//
// Parameters:
// - text: A string containing the format for the log message.
// - values: A variadic parameter containing arguments to replace placeholders in the format string.
//
// Returns:
// - A formatted string for warning logging.
func Swarningf(text string, values ...any) string {
	value := ValuesJoinf(text, values...)
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	result := fmt.Sprintf("%s [%swarning%s] %s", currentTime, colors.Yellow, colors.Reset, value)
	return result
}

// Salertf generates a formatted string for alert logging.
// It includes a timestamp, "alert" label, and color-coded output.
//
// Parameters:
// - text: A string containing the format for the log message.
// - values: A variadic parameter containing arguments to replace placeholders in the format string.
//
// Returns:
// - A formatted string for alert logging.
func Salertf(text string, values ...any) string {
	value := ValuesJoinf(text, values...)
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	result := fmt.Sprintf("%s [%salert%s] %s", currentTime, colors.Red, colors.Reset, value)
	return result
}

// sCrashf generates a formatted string for fatal logging.
// It includes a timestamp, "Crash" label, and color-coded output.
//
// Parameters:
// - text: A string containing the format for the log message.
// - values: A variadic parameter containing arguments to replace placeholders in the format string.
//
// Returns:
// - A formatted string for fatal logging.
func sCrashf(text string, values ...any) string {
	value := ValuesJoinf(text, values...)
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	result := fmt.Sprintf("%s [%scrash%s] %s", currentTime, colors.Red, colors.Reset, value)
	return result
}

// Infof prints a formatted info log to the console.
//
// Parameters:
// - text: A string containing the format for the log message.
// - values: A variadic parameter containing arguments to replace placeholders in the format string.
//
// Example:
//
//	logger.Infof("Server started on port %d", port)
func Infof(text string, values ...any) {
	result := Sinfof(text, values...)
	fmt.Println(result)
}

// Warningf prints a formatted warning log to the console.
//
// Parameters:
// - text: A string containing the format for the log message.
// - values: A variadic parameter containing arguments to replace placeholders in the format string.
//
// Example:
//
//	logger.Warningf("Memory usage is at %d%%", usage)
func Warningf(text string, values ...any) {
	result := Swarningf(text, values...)
	fmt.Println(result)
}

// Alertf prints a formatted alert log to the console.
//
// Parameters:
// - text: A string containing the format for the log message.
// - values: A variadic parameter containing arguments to replace placeholders in the format string.
//
// Example:
//
//	logger.Alertf("Database connection lost: %s", err.Error())
func Alertf(text string, values ...any) {
	result := Salertf(text, values...)
	fmt.Println(result)
}

// Fatalf logs a critical error message with formatted text to the console and immediately terminates the program.
//
// Parameters:
// - text: A string containing the format for the log message.
// - values: A variadic parameter containing arguments to replace placeholders in the format string.
//
// Example:
//
//	logger.Crashf("Failed to load configuration: %s", err.Error())
func Crashf(text string, values ...any) {
	result := sCrashf(text, values...)
	fmt.Println(result)
	os.Exit(1)
}
