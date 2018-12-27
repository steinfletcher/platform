package errors

import "net/http"

var InvalidRequestBody = New(
	"INVALID_REQUEST_BODY",
	"the request body is not valid",
	http.StatusBadRequest)
