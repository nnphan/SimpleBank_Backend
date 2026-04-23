package util

func CheckPaniceError(err error, message string) {
	if err != nil {
		panic(message + ": " + err.Error())
	}
}