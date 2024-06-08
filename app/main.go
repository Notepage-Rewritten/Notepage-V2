package main

import (
	"de.notepage-rewritten.app/pkg/views"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")

	mainWindow := views.NewMainWidnow()
	mainWindow.AddContentToWindow(hello)
	mainWindow.AddContentToWindow(
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	)

	w.SetContent(mainWindow.GetWindow())

	w.ShowAndRun()
}
