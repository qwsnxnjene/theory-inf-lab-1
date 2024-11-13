package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Диагностика диабета")

	window.Resize(fyne.NewSize(800, 800))

	window.SetContent(Init())
	window.ShowAndRun()
}
