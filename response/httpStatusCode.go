package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeFail         = 20002 // Fail
	ErrorCodeParamInvalid = 20003 // Invalid
)

// message
var msg = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeFail:         "Fail",
	ErrorCodeParamInvalid: "Invalid",
}
