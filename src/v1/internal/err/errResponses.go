package err

func BadRequestResponse() (interface{}, *Error) {
  return nil, &Error{Code: 400, Message: "Bad Request!"}
}

func UnauthorizedAccessResponse () (interface{}, *Error) {
  return nil, &Error{Code: 405, Message:"Unauthorized Access!"}
}

func InternalErrorResponse () (interface{}, *Error) {
  return nil, &Error{Code: 500, Message:"Internal server error!"}
}
