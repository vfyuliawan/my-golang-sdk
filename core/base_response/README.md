# BaseResponse Package

## Overview
The `base_response` package provides a standardized structure for API responses, ensuring consistency in response formatting across different endpoints.

## Struct Definition
```go
// BaseResponse defines the structure for API responses.
type BaseResponse[T any] struct {
    SrcSystem    string `json:"srcSystem"`
    Code         string `json:"code"`
    CodeSystem   string `json:"codeSystem"`
    Message      string `json:"message"`
    MessageError string `json:"messageError"`
    Result       T      `json:"result"`
}
```

## Methods
### Success Response
```go
func Success[T any](result T) BaseResponse[T]
```
- **Description**: Creates a successful response with the given result and message.
- **Parameters**:
    - `result T`: The actual data to be returned.
- **Returns**: A `BaseResponse` object with status code `200`.

### Success With Message Response
```go
func SuccessWithMessage[T any](result T, message string) BaseResponse[T]
```
- **Description**: Creates a successful response with the given result and message.
- **Parameters**:
  - `result T`: The actual data to be returned.
  - `message string`: A success message.
- **Returns**: A `BaseResponse` object with status code `200`.

### Error Response
```go
func NewError(codeSystem string, code string, message string, httpStatus int) BaseResponse[T]
```
- **Description**: Creates an error response with the given error code and message.
- **Parameters**:
    - `errorCode string`: The error code.
    - `messageError string`: Detailed error message.
    - `message string`: A human-readable error message.
- **Returns**: A `BaseResponse` object containing error details.

## Usage Example
### Handling a Successful Response
```go
base_response.Success(userData, "User created successfully")
```

### Handling an Error Response
```go
base_response.NewError(
	constant.InvalidParameter.GetCode(),
    constant.InvalidParameter.GetDescription(),
    err.Error(),
	400,
)
```

## Conclusion
The `base_response` package helps streamline API response handling, making it easier to standardize responses across different services.

