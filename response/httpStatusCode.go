package response

const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500
	StatusInvalidToken        = 40101
	StatusUserHasExist        = 50001 // User has already registered
)

// Message for http status code
var msg = map[int]string{
	StatusOK:                  "OK",
	StatusCreated:             "Created",
	StatusBadRequest:          "Bad Request",
	StatusUnauthorized:        "Unauthorized",
	StatusForbidden:           "Forbidden",
	StatusNotFound:            "Not Found",
	StatusInternalServerError: "Internal Server Error",
	StatusInvalidToken:        "Invalid Token",
	StatusUserHasExist:        "User has already registered",
}