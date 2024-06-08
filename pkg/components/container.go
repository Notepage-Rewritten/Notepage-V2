package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Container struct {
	container *fyne.Container
}

func NewContainer() Container {
	return Container{
		container: container.NewVBox(),
	}
}

func (c Container) GetContainer() *fyne.Container {
	return c.container
}

func (c Container) AppendToContainer(obj fyne.CanvasObject) {
	c.container.Add(obj)
}
