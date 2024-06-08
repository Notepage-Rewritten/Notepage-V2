package main

import (
	"de.notepage-rewritten.app/pkg/components"
	"de.notepage-rewritten.app/pkg/views"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	button := components.NewButton("test")

	button.SetEventHandler(func() {
		hello.SetText("Hello, World!")
	})

	mainWindow := views.NewMainWidnow()

	container := components.NewContainer()
	container.AppendToContainer(hello)
	container.AppendToContainer(button.GetButton())

	mainWindow.AddContentToWindow(container.GetContainer())

	w.SetContent(mainWindow.GetWindow())
	w.ShowAndRun()
}
