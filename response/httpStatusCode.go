package response

const (
	ErrorCodeSuccess  = 20001 // Success
	ErrorCodeFail     = 20002 // Fail
	ErrorInvalidToken = 30001 // Invalid token
)

// message
var msg = map[int]string{
	ErrorCodeSuccess:  "Success",
	ErrorCodeFail:     "Fail",
	ErrorInvalidToken: "Invalid-token",
}
