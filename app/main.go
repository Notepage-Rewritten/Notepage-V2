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
	button := widget.NewButton("Hi!", func() {
		hello.SetText("Welcome :)")
	})
	mainWindow := views.NewMainWidnow()
	container := components.NewContainer()

	container.AppendToContainer(hello)
	container.AppendToContainer(button)

	mainWindow.AddContentToWindow(container.GetContainer())

	w.SetContent(mainWindow.GetWindow())

	w.ShowAndRun()
}
