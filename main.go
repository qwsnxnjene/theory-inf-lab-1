package main

import "fmt"

// "fyne.io/fyne/v2/app"
// "fyne.io/fyne/v2/container"
// "fyne.io/fyne/v2/widget"

func getStarted() {

}

func main() {
	lst := SortDataByIndex(GenerateData(), 16.0, 40.0)
	// for i := range lst {
	// 	fmt.Printf("H: %d, W: %d, Suspended: %v\n", int(lst[i].Height), int(lst[i].Weight), lst[i].Suspended)
	// }

	lst = CalcGlucoseIndex(lst, 2.0)
	lst = MarkDiabesePeople(lst, 3.0)
	for i := range lst {
		fmt.Printf("H: %d, W: %d, Diabetes: %v\n", int(lst[i].Height), int(lst[i].Weight), lst[i].Diabetes)
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
