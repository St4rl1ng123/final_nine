package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	// проверка генерации с положительным числом
	size := 10
	data := generateRandomElements(size)

	if len(data) != size {
		t.Errorf("Error!") // вывод ошибки
	}

	// все элементы - целые числа?
	for _, val := range data {
		if val < 0 || val > 999999 {
			t.Errorf("Error!")
		}
	}

	// размер слайса равен нулю?
	emptyData := generateRandomElements(0)
	if len(emptyData) != 0 {
		t.Errorf("Error!") // вывод ошибки
	}

	// проверка крайнего случая - отрицательный размер слайса
	negativeData := generateRandomElements(-5)
	if len(negativeData) != 0 {
		t.Errorf("Error!")
	}
}

func TestMaximum(t *testing.T) {
	// несколько элементов
	testData1 := []int{2, 7, 4, 5, 9}
	max := maximum(testData1)
	if max != 9 {
		t.Errorf("ожидаемый максимум %d неверный", max)
	}

	// один элемент
	testData2 := []int{33}
	max = maximum(testData2)
	if max != 33 {
		t.Errorf("ожидаемый максимум %d неверный", max)
	}

	// отрицательные числа
	testData3 := []int{-2, -5, -12}
	max = maximum(testData3)
	if max != -2 {
		t.Errorf("ожидаемый максимум %d неверный", max)
	}

	// пустой слайс
	testData4 := []int{}
	max = maximum(testData4)
	if max != 0 {
		t.Errorf("тест работает неисправно, возвращая результат %d", max)
	}

	// одинаковые элементы
	testData5 := []int{4, 4, 4, 4}
	max = maximum(testData5)
	if max != 4 {
		t.Errorf("ожидаемый максимум %d неверный", max)
	}
}
