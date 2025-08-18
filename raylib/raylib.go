package raylib

import (
	"chat/shared/utils"
	"log"
)

type Config struct {
	Runner Runner
	Echo   Echo
	GUI    GUI
}
type Raylib struct {
	runner Runner
	echo   Echo
	gui    GUI
}

func New(cfg Config) *Raylib {
	return &Raylib{
		runner: cfg.Runner,
		echo:   cfg.Echo,
		gui:    cfg.GUI,
	}
}

func (r *Raylib) InitLayout() {
	r.gui.SetRootComponent(rootComponent(r))
}

func (r *Raylib) Run() {
	err := r.runner.Init()
	if err != nil {
		log.Println(err)
		return
	}
	defer utils.Close(r.runner)

	err = r.runner.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
