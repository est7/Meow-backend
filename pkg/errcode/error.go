package errcode

import (
	"fmt"
	"net/http"
	"sync"
)

// Error represents a general error with code, message, and optional cause
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   error  `json:"cause,omitempty"`
}

// CustomError represents a domain-specific error with additional details
type CustomError struct {
	code    int      `json:"code"`
	message string   `json:"message"`
	details []string `json:"details,omitempty"`
}

var (
	errorCodes = make(map[int]struct{})
	toStatus   = sync.Map{}
	mu         sync.RWMutex
)

// NewError creates a new Error
func NewError(code int, message string, cause error) *Error {
	return &Error{Code: code, Message: message, Cause: cause}
}

// NewCustomError creates a new CustomError
func NewCustomError(code int, message string) *CustomError {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := errorCodes[code]; ok {
		panic(fmt.Sprintf("code %d already exists, please use a different one", code))
	}
	errorCodes[code] = struct{}{}
	return &CustomError{code: code, message: message}
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("code: %d, message: %s, cause: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.code, e.message)
}

func (e *CustomError) Code() int {
	return e.code
}

func (e *CustomError) Message() string {
	return e.message
}
func (e *CustomError) Details() []string {
	return e.details
}

func (e *CustomError) WithDetails(details ...string) *CustomError {
	newError := *e
	newError.details = append(newError.details, details...)
	return &newError
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.Code(), Success.Message()
	}

	switch typed := err.(type) {
	case *Error:
		return typed.Code, typed.Message
	case *CustomError:
		return typed.Code(), typed.Message()
	default:
	}

	return ErrInternalServer.Code(), err.Error()
}

func SetHTTPStatusCode(err *CustomError, status int) {
	toStatus.Store(err.Code(), status)
}

func ToHTTPStatusCode(code int) int {
	if status, ok := toStatus.Load(code); ok {
		return status.(int)
	}
	return http.StatusBadRequest
}

func initToStatus() {
	statusMap := map[int]int{
		Success.Code():               http.StatusOK,
		ErrInternalServer.Code():     http.StatusInternalServerError,
		ErrNotFound.Code():           http.StatusNotFound,
		ErrInvalidParam.Code():       http.StatusBadRequest,
		ErrToken.Code():              http.StatusUnauthorized,
		ErrInvalidToken.Code():       http.StatusUnauthorized,
		ErrTokenTimeout.Code():       http.StatusUnauthorized,
		ErrTooManyRequests.Code():    http.StatusTooManyRequests,
		ErrServiceUnavailable.Code(): http.StatusServiceUnavailable,
	}
	for code, status := range statusMap {
		SetHTTPStatusCode(NewCustomError(code, ""), status)
	}
}

func init() {
	initToStatus()
}
