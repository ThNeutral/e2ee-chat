package errs

import "fmt"

type Error struct {
	Message string
	Cause   error
	Code    int
}

func (err Error) Error() string {
	if err.Message == "" {
		return fmt.Sprint("%w", err.Cause)
	}

	return fmt.Sprint("%s: %w", err.Message, err.Cause)
}
