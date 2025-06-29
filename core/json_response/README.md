# json_response

`json_response` is a Go package that provides helper functions to simplify API responses in applications utilizing the [Gin Web Framework](https://github.com/gin-gonic/gin). The package is designed to format success and error responses consistently by wrapping Gin's `c.JSON` method with additional formatting.

---

## Features
- **Consistent Success Responses**: Return a standardized JSON format for successful API calls.
- **Error Handling**: Simplify formatting for error-related responses while following provided error code structures.

---

## Installation

To install the package, use Go modules:

```bash
go get <module-path>
```

---

## Usage
The package provides two main functions: `Success` and `Error`.

### `Success`

Use the `Success` function to return a success response. It wraps `c.JSON` to create well-structured JSON responses for API success scenarios.

**Function Signature**:
```go
func Success[T any](c *gin.Context, result T, message string) { 
	//
}
func SuccessWithMessage[T any](c *gin.Context, result T, message string) {
	//
}
```

**Parameters**:
- `c *gin.Context`: The Gin context.
- `result T`: The actual data or result to be returned.
- `message string`: An optional message to accompany the response.

**Example**:
```go
import (
	"repo.pegadaian.co.id/tring/ist-commons-util/json_response"
	"github.com/gin-gonic/gin"
)

router.GET("/success", func(c *gin.Context) {
    result := map[string]interface{}{
        "id":   1,
        "name": "sample",
    }
    message := "Operation completed successfully"
    json_response.Success(c, 200, result, message)
})
```

### `Error`

Use the `Error` function to handle error responses effectively. It helps format error details in a clear and consistent manner.

**Function Signature**:
```go
func Error[T any](c *gin.Context, statusCode int, errorCode *constant.ErrorCode, err error)
```

**Parameters**:
- `c *gin.Context`: The Gin context.
- `statusCode int`: HTTP status code for the response (e.g., 500 for server error).
- `errorCode *constant.ErrorCode`: An object encapsulating error details.
- `err error`: The actual error instance.

**Example**:
```go
import "github.com/gin-gonic/gin"

router.GET("/error", func(c *gin.Context) {
    var customError = errors.New("something went wrong")
    var errorCode = &constant.ErrorCode{
        Code:        "ERR001",
        Description: "An error occurred while processing your request",
    }
    json_response.Error(c, 500, errorCode, customError)
})
```

---

## Dependencies

This package relies on the following libraries:
- [Gin Web Framework](https://github.com/gin-gonic/gin)

---
erms of the [MIT License](LICENSE). See the LICENSE file for more details.