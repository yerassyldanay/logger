package utils

import (
	"fmt"
)

type LogMessage interface {
	LogExtended() string
	LogMini() string
}

type FunctionCallSuccess struct {
	Host 				string
	FunctionName		string
	Message				string
}

type FunctionCallError struct {
	FunctionCallSuccess
}

func (err *FunctionCallError) LogExtended() string {
	return fmt.Sprintf("[%v] Error occured while calling a function. Host: %v | Message: %v",
				err.FunctionName, err.Host, err.Message)
}

func (err *FunctionCallError) LogMini() string {
	return fmt.Sprintf("[%v] Error occured while calling a function. Message: %v",
				err.FunctionName, err.Message)
}

func (success *FunctionCallSuccess) LogExtended() string {
	return fmt.Sprintf("[%v] Operation has been done successfully. Host: %v | Message: %v",
				success.FunctionName, success.Host, success.Message)
}

func (success *FunctionCallSuccess) LogMini() string {
	return fmt.Sprintf("[%v] Operation has been done successfully. Message: %v",
				success.FunctionName, success.Message)
}

func LogExtended(message LogMessage) {
	GetLogger().LogMessageToFile(message.LogExtended())
}

func LogMini(message LogMessage) {
	GetLogger().LogMessageToFile(message.LogMini())
}
