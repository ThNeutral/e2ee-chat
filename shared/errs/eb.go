package errs

import "fmt"

type ErrorBuilder struct {
	message string
	cause   error
	code    int
}

func (b *ErrorBuilder) GetCode() int {
	return b.code
}

func B() *ErrorBuilder {
	return &ErrorBuilder{}
}

func (b *ErrorBuilder) Msg(msg string) *ErrorBuilder {
	b.message = msg
	return b
}

func (b *ErrorBuilder) Code(code int) *ErrorBuilder {
	b.code = code
	return b
}

func (b *ErrorBuilder) Cause(cause error) *ErrorBuilder {
	b.cause = cause
	return b
}

func (b *ErrorBuilder) Causef(format string, a ...any) *ErrorBuilder {
	return b.Cause(fmt.Errorf(format, a...))
}

func (b *ErrorBuilder) Err() error {
	return Error{
		Message: b.message,
		Cause:   b.cause,
		Code:    b.code,
	}
}
