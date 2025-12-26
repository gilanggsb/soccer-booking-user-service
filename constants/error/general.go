package constants

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")                   // 2 usages
	ErrSQLError            = errors.New("database server failed to execute query") // 1 usage
	ErrTooManyRequests     = errors.New("too many requests")                       // 1 usage
	ErrUnauthorized        = errors.New("unauthorized")                            // 1 usage
	ErrInvalidToken        = errors.New("invalid token")                           // 1 usage
	ErrForbidden           = errors.New("forbidden")                               // 1 usage
)

var GeneralErrors = []error{ // no usages
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
