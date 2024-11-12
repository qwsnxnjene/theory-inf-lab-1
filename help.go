package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"

	"fyne.io/fyne/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	//"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	userNumber = 500 //кол-во людей в таблице

	//диапазон роста
	minHeight = 165
	maxHeight = 225

	//диапазон веса
	minWeight = 62
	maxWeight = 100

	//диапазон ИМТ
	minBMI = 18.0
	maxBMI = 30.0

	//граничный уровень глюкозы
	level = 1.5
)

var (
	data  [userNumber]UserDate
	table *widget.Table
)

type UserDate struct {
	Weight       float64
	Height       float64
	GlucoseIndex float64
	Diabetes     bool
	Suspended    bool
}

// roundFloat округляет число до двух цифр после запятой
func roundFloat(val float64) float64 {
	ratio := math.Pow(10, 2)
	return math.Round(val*ratio) / ratio
}

func BuildTable() {
	table = widget.NewTable(
		func() (int, int) { return 5, userNumber + 1 },
		func() fyne.CanvasObject { return widget.NewLabel("Ур. глюкозы") },
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			toSet := ""

			if i.Col == 0 {
				if i.Row == 0 {
					toSet = ""
				} else if i.Row == 1 {
					toSet = "Вес"
				} else if i.Row == 2 {
					toSet = "Рост"
				} else if i.Row == 3 {
					toSet = "Ур. глюкозы"
				} else if i.Row == 4 {
					toSet = "Диабет"
				}
			} else if i.Row == 0 {
				if data[i.Col-1].Suspended {
					toSet = "xxx"
				} else {
					toSet = fmt.Sprintf("%d", i.Col)
				}
			} else if i.Row == 1 {
				toSet = fmt.Sprintf("%d", int(data[i.Col-1].Weight))
			} else if i.Row == 2 {
				toSet = fmt.Sprintf("%d", int(data[i.Col-1].Height))
			} else if i.Row == 3 {
				if data[i.Col-1].Suspended {
					toSet = "xxx"
				} else {
					toSet = fmt.Sprintf("%.2f", data[i.Col-1].GlucoseIndex)
				}
			} else if i.Row == 4 {
				if data[i.Col-1].Suspended {
					toSet = "xxx"
				} else if data[i.Col-1].Diabetes {
					toSet = "+++"
				} else {
					toSet = "-"
				}
			}

			obj.(*widget.Label).SetText(toSet)
		},
	)
	table.Resize(fyne.NewSize(800, table.MinSize().Height*5.5))
	table.Move(fyne.NewPos(0, 0))
}

// Init инициализирует интерфейс программы
func Init() *fyne.Container {
	BuildTable()

	btnGenerate := widget.NewButton("Сгенерировать данные", func() { GenerateData() })
	btnGenerate.Resize(fyne.NewSize(600, 50))
	btnGenerate.Move(fyne.NewPos(100, table.MinSize().Height*5+50))

	btnSortByData := widget.NewButton("Исключить нелогичные данные", func() { SortDataByIndex(); VisualizeSortedData(); WeightHeightRatioPlot() })
	btnSortByData.Resize(fyne.NewSize(600, 50))
	btnSortByData.Move(fyne.NewPos(100, btnGenerate.Position().Y+80))

	btnGlucoseIndex := widget.NewButton("Вычислить уровень глюкозы", func() { CalcGlucoseIndex(1.0) })
	btnGlucoseIndex.Resize(fyne.NewSize(600, 50))
	btnGlucoseIndex.Move(fyne.NewPos(100, btnSortByData.Position().Y+80))

	btnMarkDiabetes := widget.NewButton("Выделить людей с диабетом", func() { MarkDiabetesePeople() })
	btnMarkDiabetes.Resize(fyne.NewSize(600, 50))
	btnMarkDiabetes.Move(fyne.NewPos(100, btnGlucoseIndex.Position().Y+80))

	c := container.NewWithoutLayout(table, btnGenerate, btnSortByData, btnGlucoseIndex, btnMarkDiabetes)

	return c
}

// GenerateData генерирует данные об указанном кол-ве пользователей
// в формате Вес, Рост
func GenerateData() {
	for i := 0; i < userNumber; i++ {
		u := UserDate{
			Weight: roundFloat(float64(rand.Intn(maxWeight-minWeight) + minWeight)),
			Height: roundFloat(float64(rand.Intn(maxHeight-minHeight) + minHeight)),
		}

		data[i] = u
	}

	table.Refresh()
}

// SortDataByIndex сортирует данные о пользователях, помечая некорректные данные меткой Suspended
func SortDataByIndex() {
	for i := range data {
		data[i].checkGeneratedData()
	}

	table.Refresh()
}

// checkGeneratedData вычисляет массу тела пользователя и проверяет её на корректность
func (u *UserDate) checkGeneratedData() {
	index := u.Weight / ((u.Height / 100.0) * (u.Height / 100.0))

	if index <= minBMI || index >= maxBMI {
		u.Suspended = true
	}
}

// VisualizeSortedData строит график с предварительно обработанными данными
func VisualizeSortedData() {
	p := plot.New()
	p.Title.Text = "Соотношение веса и роста"
	p.X.Label.Text = "Рост [см]"
	p.Y.Label.Text = "Вес [кг]"
	p.X.Label.Position = draw.PosRight
	p.Y.Label.Position = draw.PosTop
	p.X.Min = minHeight - 5
	p.X.Max = maxHeight + 5
	p.Y.Min = minWeight - 5
	p.Y.Max = maxWeight + 5

	legend := plot.NewLegend()

	// red := exampleThumbnailer{Color: }
	// green := exampleThumbnailer{Color: color.NRGBA{G: 255, A: 255}}
	//blue := exampleThumbnailer{Color: color.NRGBA{B: 255, A: 255}}

	p.Legend = legend

	scatterData := make(plotter.XYs, 0)
	for _, user := range data {
		if user.Suspended {
			continue
		}
		scatterData = append(scatterData, plotter.XY{X: user.Height, Y: user.Weight})
	}

	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		log.Fatal(err)
	}

	s.Color = color.NRGBA{R: 255, A: 255}

	p.Add(s)

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "scatter.png"); err != nil {
		log.Fatal(err)
	}
}

// WeightHeightRatioPlot вычисляет отношение веса и роста и строит гистограмму соотношений
func WeightHeightRatioPlot() {
	var ratioList plotter.Values
	for _, user := range data {
		if user.Suspended {
			continue
		}
		ratio := roundFloat(user.Weight / user.Height)
		ratioList = append(ratioList, ratio)
	}

	p := plot.New()
	p.Title.Text = "Гистограмма отношения веса и роста"

	h, err := plotter.NewHist(ratioList, 10)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(h)

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "histogram.png"); err != nil {
		log.Fatal(err)
	}
	//TODO
}

// CalcGlucoseIndex заполняет значение уровня глюкозы для всех пользователей
func CalcGlucoseIndex(sigma float64) {
	for i := range data {
		if data[i].Suspended {
			continue
		}
		data[i].calcGlucoseIndex(sigma * sigma)
	}

	table.Refresh()
}

// CalcGlucoseIndex моделирует уровень глюкозы и сохраняет значение
func (u *UserDate) calcGlucoseIndex(sigma float64) {
	glucose := roundFloat(u.Weight/u.Height + sigma)
	u.GlucoseIndex = glucose
}

// MarkDiabetesePeople маркирует пользователей на наличие диабета
func MarkDiabetesePeople() {
	for i := range data {
		if data[i].Suspended {
			continue
		}
		data[i].markDiabetesePeople()
	}

	table.Refresh()
}

// markDiabetesePeople устанавливает значение поля Diabese в зависимости от уровня глюкозы
func (u *UserDate) markDiabetesePeople() {
	if u.GlucoseIndex >= level {
		u.Diabetes = true
	}
}

// VisualizeFinalData выводит итоговую информацию
func VisualizeFinalData() {
	//TODO
}
