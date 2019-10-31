package main

import "logger/utils"

func main() {
	/*
		Close the file at the end of the function
		to avoid a leakage of resources
	*/
	defer utils.GetLogFile().Close()
	var error_message = utils.FunctionCallError{utils.FunctionCallSuccess{
		FunctionName: "TEST_ERROR",
		Message:      "This is an error message. Checking",
	}}

	utils.LogMini(&error_message)

	var success_message = utils.FunctionCallSuccess{
		FunctionName: "TEST_ERROR",
		Message:      "This is a success message. Checking",
	}

	utils.LogMini(&success_message)
}