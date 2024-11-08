package main

import (
	"math"
	"math/rand"
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
func GenerateData() [userNumber]UserDate {
	var data [userNumber]UserDate

	for i := 0; i < userNumber; i++ {
		u := UserDate{
			Weight: roundFloat(float64(rand.Intn(maxWeight-minWeight) + minWeight)),
			Height: roundFloat(float64(rand.Intn(maxHeight-minHeight) + minHeight)),
		}

		data[i] = u
	}

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

// MarkDiabesePeople маркирует пользователей на наличие диабета
func MarkDiabesePeople(data [userNumber]UserDate, level float64) [userNumber]UserDate {
	for i := range data {
		data[i].markDiabesePeople(level)
	}

	return data
}

// markDiabesePeople устанавливает значение поля Diabese в зависимости от уровня глюкозы
func (u *UserDate) markDiabesePeople(level float64) {
	if u.GlucoseIndex >= level {
		u.Diabetes = true
	}
}

func VisualizeFinalData() {
	//TODO
}
