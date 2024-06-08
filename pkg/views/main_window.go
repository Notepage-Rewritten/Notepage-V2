package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Window struct {
	mainVbox *fyne.Container
}

func NewMainWidnow() Window {
	return Window{
		mainVbox: container.NewVBox(),
	}
}

func (w Window) GetWindow() *fyne.Container {
	return w.mainVbox
}

func (w Window) AddContentToWindow(object fyne.CanvasObject) {
	w.mainVbox.Objects = append(w.mainVbox.Objects, object)
}
