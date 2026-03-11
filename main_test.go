package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	// Таблица тестовых случаев
	testCases := []struct {
		name           string // имя теста
		size           int    // входной размер слайса
		expectedLength int    // ожидаемая длина результата
		checkValues    bool   // проверка диапозона значений
	}{
		{"положительный размер (10)", 10, 10, true},
		{"нулевой размер 0", 0, 0, false},
		{"отрицательный размер (-5)", -5, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// создаём локальный объект утверждений
			assert := assert.New(t)

			// генерируем данные
			data := generateRandomElements(tc.size)

			// проверяем длину слайса
			assert.Equal(tc.expectedLength, len(data), "Длина слайса не совпадает с ожидаемой")

			// проверяем диапозон значений
			if tc.checkValues {
				for _, val := range data {
					assert.True(val >= 0 && val <= 100000000, "Значения выходят за границы")
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	// Таблица тестовых случаев
	testCases := []struct {
		name        string // имя теста
		testData    []int  // входные данные
		expectedMax int    // ожидаемый максимум
		description string // краткое описание кейса
	}{
		{"несколько элементов", []int{2, 7, 4, 5, 9}, 9, "поиск максимума в слайсе среди рандом чисел"},
		{"один элемент", []int{33}, 33, "максимум в слайсе с однмим элементом"},
		{"пустой слайс", []int{}, 0, "Обработка пустого слайса"},
		{"отрицательные числа", []int{-2, -5, -12}, -2, "поиск максимума среди отрицательных чисел"},
		{"одинаковые элементы", []int{4, 4, 4, 4}, 4, "поиск максимума среди одинаковых элементов"},
	}

	// перебираем все тестовые случаи
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// создаём локальный объект
			assert := assert.New(t)

			// вызов тестовой функции
			max := maximum(tc.testData)

			// проверяем результат
			assert.Equal(tc.expectedMax, max, tc.description)
		})
	}
}
