package shared

import "log"

type closer interface {
	Close() error
}

func Close(c closer) {
	err := c.Close()
	if err != nil {
		log.Println(err)
	}
}

func CloseWithEB(c closer, eb *ErrorBuilder) {
	err := c.Close()
	if err != nil {
		log.Println(eb.Cause(err).Err())
	}
}
