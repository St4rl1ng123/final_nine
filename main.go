package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	// если размер слайса 0, вернёт пустой слайс
	if size <= 0 {
		return nil
	}

	//инициализация генератора слуачайных чисел
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// создание слайса заданной длины и заполнение его рандомными намберами
	data := make([]int, size)
	// заполняем слайс
	for i := range data {
		data[i] = localRand.Int()
		data[i] %= SIZE
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	// проверка что слайс не пуст
	if len(data) == 0 {
		return 0
	}

	// делаем макс как первый элемент слайса
	max := data[0]

	// идём дальше и обновляем максимум при необходимости
	for _, val := range data[1:] {
		if val > max {
			max = val
		}
	}

	// возвращаем найденный макс. и ошибки нет
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	// если слайс пуст
	if len(data) == 0 {
		return 0
	}

	// для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// размер каждого чанка-(среза)
	chunkSize := len(data) / CHUNKS

	// слайс для хранения максимумов из каждого чанка
	maxValues := make([]int, CHUNKS)

	//кол-во задач (8 горутин)
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		// вычисляем начальный и конечный индекс чанка
		start := i * chunkSize
		end := start + chunkSize

		// для ласт чанка нужно вять оставшиеся элементы
		if i == CHUNKS-1 {
			end = len(data)
		}

		// запуск горутины
		go func(chunkIndex int, chunkData []int) {
			defer wg.Done() // лучшее применение

			// находим максимум
			max := maximum(chunkData)

			// save this
			maxValues[chunkIndex] = max
		}(i, data[start:end])
	}

	wg.Wait() // ждём завершение

	// находим макс среди максов
	finalMax := maxValues[0]
	for _, val := range maxValues[1:] {
		if val > finalMax {
			finalMax = val
		}
	}

	return finalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	if len(data) == 0 {
		fmt.Println("Слайс пуст")
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	startSeq := time.Now() // замеряем время начала
	maxSeq := maximum(data)
	elapsed := time.Since(startSeq) // вычисляем затраченное время

	// переводим в микросекунды
	elapsedSeq := elapsed.Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSeq, elapsedSeq)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	startChunks := time.Now() // замеряем время для горутин
	maxCh := maxChunks(data)
	elapsedCh := time.Since(startChunks)
	elapsedChunks := elapsedCh.Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxCh, elapsedChunks)
}
