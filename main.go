package main

import (
	"errors"
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
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		return nil, errors.New("размер должен быть положительным")
	}
	if size > 100_000_000 {
		return nil, errors.New("размер больше максимально допустимого")
	}

	data := make([]int, size)

	for i := range size {
		data[i] = rand.Intn(SIZE) + 1
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
	if len(data) == 0 {
		return 0, errors.New("передан пустой слайс")
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup
	chunkSize := len(data) / CHUNKS
	results := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(chunkIndex int) {
			defer wg.Done()

			startIndex := chunkIndex * chunkSize
			endIndex := startIndex + chunkSize

			if chunkIndex == CHUNKS-1 {
				endIndex = len(data)
			}
			if startIndex >= len(data) {
				return
			}
			chunk := data[startIndex:endIndex]
			if len(chunk) == 0 {
				return
			}
			chunkMax, err := maximum(chunk)
			if err != nil {
				return
			}
			results[chunkIndex] = chunkMax
		}(i)
	}
	wg.Wait()

	finalMax, err := maximum(results)
	if err != nil || finalMax == 0 {
		return 0, errors.New("не удалось обработать чанки")
	}
	return finalMax, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
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
		fmt.Printf("Ошибка однопоточного поиска максимума: %v\n", err)
		return
	}
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	// Многопоточный поиск
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max, err = maxChunks(data)
	if err != nil {
		fmt.Printf("Ошибка многопоточного поиска максимума: %v\n", err)
		return
	}
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
