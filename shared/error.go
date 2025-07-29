package shared

import "fmt"

type ErrorBuilder struct {
	message string
	cause   error
}

func NewErrorBuilder() *ErrorBuilder {
	return &ErrorBuilder{}
}

func (b *ErrorBuilder) Msg(msg string) *ErrorBuilder {
	b.message = msg
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
	if b.message == "" {
		return fmt.Errorf("%w", b.cause)
	}

	return fmt.Errorf("%s: %w", b.message, b.cause)
}
