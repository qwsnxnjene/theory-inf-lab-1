package main

import "math/rand"

const (
	userNumber = 500 //кол-во людей в таблице

	//диапазон роста
	minHeight = 165
	maxHeight = 225

	//диапазон веса
	minWeight = 65
	maxWeight = 100
)

type UserDate struct {
	Weight       int
	Height       int
	GlucoseIndex float64
	Diabetes     bool
	Suspended    bool
}

//GenerateData генерирует данные об указанном кол-ве пользователей
//в формате Вес, Рост
func GenerateData() [userNumber]UserDate {
	var data [userNumber]UserDate

	for i := 0; i < userNumber; i++ {
		u := UserDate{
			Weight: rand.Intn(maxWeight-minWeight) + minWeight,
			Height: rand.Intn(maxHeight-minHeight) + minHeight,
		}

		data[i] = u
	}

	return data
}

//SortDataByIndex сортирует данные о пользователях, помечая некорректные данные меткой Suspended
func SortDataByIndex(data [userNumber]UserDate, minBMI, maxBMI float64) [userNumber]UserDate {
	for i := range data {
		data[i].checkGeneratedData(minBMI, maxBMI)
	}

	return data
}

//checkGeneratedData вычисляет массу тела пользователя и проверяет её на корректность
func (u *UserDate) checkGeneratedData(minBMI, maxBMI float64) {
	index := float64(u.Weight) / ((float64(u.Height) / 100.0) * (float64(u.Height) / 100.0))

	if index <= minBMI || index >= maxBMI {
		u.Suspended = true
	}
}
