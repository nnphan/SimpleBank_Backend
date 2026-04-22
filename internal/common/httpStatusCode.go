package common

// internal err code for front-end and back-end
const (
	ErrCodeSuccess = 20001 //Success , 
	ErrCodeParamInvalid = 20003 //  Param is Invalid,
)

//message
var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "Email is invalid",
}