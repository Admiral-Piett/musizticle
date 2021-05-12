package utils

// FIXME - there has to be a more elegant object for this
type LogFieldStruct = struct {
	ErrorMessage string
	RequestBody string
	StackContext string
}

var LogFields = LogFieldStruct{
	ErrorMessage: "error_message",
	RequestBody: "request_body",
	StackContext: "stack_context",
}