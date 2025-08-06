package widgets

import (
	"chat/shared"
	"image/color"
	"sync/atomic"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WidgetCallback func() error

type RectangleButton struct {
	rl.RectangleInt32

	IdleColor color.RGBA

	LoadingColor color.RGBA

	ErrColor     color.RGBA
	ErrTextColor color.RGBA
	ErrTextSize  int32

	IsAsync bool

	CallbackFunc WidgetCallback

	err     shared.AtomicString
	loading atomic.Bool
}

func (rw *RectangleButton) Draw() {
	if err := rw.err.Load(); err != "" {
		rl.DrawRectangle(rw.X, rw.Y, rw.Width, rw.Height, rw.ErrColor)
		rl.DrawText(err, rw.X+rw.Width/2, rw.Y+rw.Height/2, rw.ErrTextSize, rw.ErrTextColor)
	} else if rw.loading.Load() {
		rl.DrawRectangle(rw.X, rw.Y, rw.Width, rw.Height, rw.LoadingColor)
	} else {
		rl.DrawRectangle(rw.X, rw.Y, rw.Width, rw.Height, rw.IdleColor)
	}
}

func (rw *RectangleButton) Contains(point rl.Vector2) bool {
	x := int32(point.X)
	y := int32(point.Y)

	return x >= rw.X && x < rw.X+rw.Width &&
		y >= rw.Y && y < rw.Y+rw.Height
}

func (rw *RectangleButton) HasCallback() bool {
	return rw.CallbackFunc != nil
}

func (rw *RectangleButton) Callback() {
	if !rw.HasCallback() {
		panic("callback is not set")
	}

	if rw.IsAsync {
		go rw.callback()
	} else {
		rw.callback()
	}
}

func (rw *RectangleButton) callback() {
	if rw.loading.Load() {
		return
	}

	rw.loading.Store(true)
	defer rw.loading.Store(false)

	err := rw.CallbackFunc()
	if err != nil {
		rw.err.Store(err.Error())
	}
}
