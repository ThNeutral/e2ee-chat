package utils

import (
	"chat/shared/errs"
	"log"
)

type closer interface {
	Close() error
}

func Close(c closer) {
	err := c.Close()
	if err != nil {
		log.Println(err)
	}
}

func CloseWithEB(c closer, eb *errs.ErrorBuilder) {
	err := c.Close()
	if err != nil {
		log.Println(eb.Cause(err).Err())
	}
}
