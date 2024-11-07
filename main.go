package main

import "fmt"

// "fyne.io/fyne/v2/app"
// "fyne.io/fyne/v2/container"
// "fyne.io/fyne/v2/widget"

func getStarted() {

}

func main() {
	lst := SortDataByIndex(GenerateData(), 16.0, 40.0)
	for i := range lst {
		fmt.Printf("H: %d, W: %d, Suspended: %v\n", lst[i].Height, lst[i].Weight, lst[i].Suspended)
	}

	// a := app.New()
	// w := a.NewWindow("Hello")

	// hello := widget.NewLabel("Hello Fyne!")
	// w.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// ))

	// w.ShowAndRun()
}
