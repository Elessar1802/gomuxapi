package err

import (
  "net/http" 
  enc "github.com/Elessar1802/api/src/v1/internal/encoder"
)

/* Request Error */
func BadRequestResponse(msg ...string) (enc.Response) {
  code := http.StatusBadRequest
  var m string = "Bad Request"
  // type of variadic function argument is []string {[]type}
  // so we can use len()
  if hasMsg := len(msg) > 0; hasMsg {
    m = msg[0]
  }
  return enc.Response{
    Code: code,
    Error: &Error{
      Code: code,
      Message: m,
    },
  }
}

func UnauthorizedAccessResponse(msg ...string) (enc.Response) {
  code := http.StatusUnauthorized
  var m string = "Unauthorized Access!"
  // type of variadic function argument is []string {[]type}
  // so we can use len()
  if hasMsg := len(msg) > 0; hasMsg {
    m = msg[0]
  }
  return enc.Response{
    Code: code,
    Error: &Error{
      Code: code,
      Message: m,
    },
  }
}

func ForbiddenErrorResponse(msg ...string) (enc.Response) {
  code := http.StatusForbidden
  var m string = "Requested resource is forbidden"
  if hasMsg := len(msg) > 0; hasMsg {
    m = msg[0]
  }
  return enc.Response{
    Code: code,
    Error: &Error{
      Code: code,
      Message: m,
    },
  }
}

func NotFoundErrorResponse(msg ...string) (enc.Response) {
  code := http.StatusNotFound
  var m string = "Resource not found"
  if hasMsg := len(msg) > 0; hasMsg {
    m = msg[0]
  }
  return enc.Response{
    Code: code,
    Error: &Error{
      Code: code,
      Message: m,
    },
  }
}

func MethodNotAllowedErrorResponse(msg ...string) (enc.Response) {
  code := http.StatusMethodNotAllowed
  var m string = "Method not allowed!"
  if hasMsg := len(msg) > 0; hasMsg {
    m = msg[0]
  }
  return enc.Response{
    Code: code,
    Error: &Error{
      Code: code,
      Message: m,
    },
  }
}

/* Internal Error */
func InternalServerErrorResponse(msg ...string) (enc.Response) {
  code := http.StatusInternalServerError
  var m string = "Internal Server Error"
  if hasMsg := len(msg) > 0; hasMsg {
    m = msg[0]
  }
  return enc.Response{
    Code: code,
    Error: &Error{
      Code: code,
      Message: m,
    },
  }
}
