// https://contest.yandex.ru/contest/22450/run-report/86674416/

// Алгоритм имеет сложность O(n), где n - количество элементов в матрице 4x4.
// В алгоритме используется два вложенных цикла по 4 итерации каждый,
// что даёт общее количество операций равное 4*4=16.
// Далее один проход по массиву счетчиков длиной 10, что занимает константное время.

// Алгоритм использует фиксированное количество памяти для хранения счетчика и доски,
// поэтому сложность по памяти постоянна и равна O(1).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var k int
	fmt.Fscanln(reader, &k)

	board := make([][]rune, 4)
	for i := range board {
		board[i] = make([]rune, 4)
	}

	for i := 0; i < 4; i++ {
		inputStr, _ := reader.ReadString('\n')
		for j, r := range inputStr[:len(inputStr)-1] {
			board[i][j] = r
		}
	}

	points := 0
	counter := make([]int, 10)
	for _, row := range board {
		for _, value := range row {
			if value != '.' {
				counter[value-'0']++
			}
		}
	}

	for _, value := range counter {
		if value > 0 && value <= k+k {
			points++
		}
	}

	fmt.Fprintln(writer, points)
}