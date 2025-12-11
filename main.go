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
		return nil, errors.New("размер массива должен быть положительным")
	}
	if size > SIZE {
		return nil, errors.New("размер массива больше максимально допустимого")
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
	// ваш код здесь
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
	}

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
