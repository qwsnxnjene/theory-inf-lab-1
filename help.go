package main

import (
	"fmt"
	"math"
	"math/rand"

	"fyne.io/fyne/v2"
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

// GenerateData генерирует данные об указанном кол-ве пользователей
// в формате Вес, Рост
func GenerateData(w *fyne.Window) [userNumber]UserDate {
	var data [userNumber]UserDate

	for i := 0; i < userNumber; i++ {
		u := UserDate{
			Weight: roundFloat(float64(rand.Intn(maxWeight-minWeight) + minWeight)),
			Height: roundFloat(float64(rand.Intn(maxHeight-minHeight) + minHeight)),
		}

		data[i] = u
	}

	table := widget.NewTable(func() (int, int) { return 3, 501 },
		func() fyne.CanvasObject { return widget.NewLabel(".......") },
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			toSet := ""

			if i.Col == 0 {
				if i.Row == 0 {
					toSet = ""
				} else if i.Row == 1 {
					toSet = "Вес"
				} else if i.Row == 2 {
					toSet = "Рост"
				}
			} else if i.Row == 0 {
				toSet = fmt.Sprintf("%d", i.Col)
			} else if i.Row == 1 {
				toSet = fmt.Sprintf("%d", int(data[i.Col-1].Weight))
			} else if i.Row == 2 {
				toSet = fmt.Sprintf("%d", int(data[i.Col-1].Height))
			}

			obj.(*widget.Label).SetText(toSet)
		})

	(*w).SetContent(table)

	return data
}

// SortDataByIndex сортирует данные о пользователях, помечая некорректные данные меткой Suspended
func SortDataByIndex(data [userNumber]UserDate, minBMI, maxBMI float64) [userNumber]UserDate {
	for i := range data {
		data[i].checkGeneratedData(minBMI, maxBMI)
	}

	return data
}

// checkGeneratedData вычисляет массу тела пользователя и проверяет её на корректность
func (u *UserDate) checkGeneratedData(minBMI, maxBMI float64) {
	index := u.Weight / ((u.Height / 100.0) * (u.Height / 100.0))

	if index <= minBMI || index >= maxBMI {
		u.Suspended = true
	}
}

// VisualizeSortedData строит график с предварительно обработанными данными
func VisualizeSortedData(data [userNumber]UserDate) {
	//TODO
}

// WeightHeightRatioPlot вычисляет отношение веса и роста и строит гистограмму соотношений
func WeightHeightRatioPlot(data [userNumber]UserDate) {
	var ratioList []float64
	for _, user := range data {
		if user.Suspended {
			continue
		}
		ratio := roundFloat(user.Weight / user.Height)
		ratioList = append(ratioList, ratio)
	}

	//TODO
}

// CalcGlucoseIndex заполняет значение уровня глюкозы для всех пользователей
func CalcGlucoseIndex(data [userNumber]UserDate, sigma float64) [userNumber]UserDate {
	for i := range data {
		if data[i].Suspended {
			continue
		}
		data[i].calcGlucoseIndex(sigma * sigma)
	}

	return data
}

// CalcGlucoseIndex моделирует уровень глюкозы и сохраняет значение
func (u *UserDate) calcGlucoseIndex(sigma float64) {
	glucose := roundFloat(u.Weight/u.Height + sigma)
	u.GlucoseIndex = glucose
}

// MarkDiabetesePeople маркирует пользователей на наличие диабета
func MarkDiabetesePeople(data [userNumber]UserDate, level float64) [userNumber]UserDate {
	for i := range data {
		data[i].markDiabetesePeople(level)
	}

	return data
}

// markDiabetesePeople устанавливает значение поля Diabese в зависимости от уровня глюкозы
func (u *UserDate) markDiabetesePeople(level float64) {
	if u.GlucoseIndex >= level {
		u.Diabetes = true
	}
}

// VisualizeFinalData выводит итоговую информацию
func VisualizeFinalData() {
	//TODO
}
