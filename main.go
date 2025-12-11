package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		return nil, errors.New("размер должен быть положительным")
	}
	if size > 100_000_000 {
		return nil, errors.New("размер больше максимально допустимого")
	}

	src := rand.NewSource(time.Now().Unix())
	data := make([]int, size)

	for i := range size {
		data[i] = int(src.Int63() % int64(SIZE))
	}

	return data, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("передан пустой слайс")
	}

	maxValue := data[0]
	for i := range len(data) {
		if data[i] > maxValue {
			maxValue = data[i]
		}
	}

	return maxValue, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	// ваш код здесь
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data, err := generateRandomElements(SIZE)

	if err != nil {
		fmt.Printf("Ошибка генерации: %v\n", err)
		return
	}

	// Однопоточный поиск
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max, err := maximum(data)

	if err != nil {
		fmt.Printf("Ошибка поиска максимума: %v\n", err)
		return
	}
	elapsed := time.Since(start).Microseconds() //?
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
