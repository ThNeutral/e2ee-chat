package raylib

import (
	"chat/raylib/runner"
	"chat/shared"
	"log"
)

type Config struct {
	Runner *runner.Runner

	Echo Echo
}
type Raylib struct {
	runner *runner.Runner
	echo   Echo
}

func New(cfg Config) *Raylib {
	return &Raylib{
		runner: cfg.Runner,
		echo:   cfg.Echo,
	}
}

func (r *Raylib) Run() {
	err := r.runner.Init()
	if err != nil {
		log.Println(err)
		return
	}
	defer shared.Close(r.runner)

	err = r.runner.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
