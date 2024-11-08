package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Диагностика диабета")
	window.Resize(fyne.NewSize(800, 800))

	lst := GenerateData(&window)
	var infoToPrint [3][501]string
	infoToPrint[0][0] = ""
	infoToPrint[1][0] = "Вес"
	infoToPrint[2][0] = "Рост"
	for i := 1; i < 501; i++ {
		infoToPrint[0][i] = fmt.Sprintf("%d", i)
		infoToPrint[1][i] = fmt.Sprintf("%d", int(lst[i-1].Weight))
		infoToPrint[2][i] = fmt.Sprintf("%d", int(lst[i-1].Height))
	}

	window.ShowAndRun()
}
