package client

type GUI interface {
	Init() error
	Close() error
	Run() error
}
