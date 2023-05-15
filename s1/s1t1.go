// https://contest.yandex.ru/contest/22450/run-report/86662738/

// Алгоритм имеет сложность O(n), где n - количество домов на улице.
// Алгоритм проходит по каждому дому на улице по два раза
// (справа налево в первом цикле и в обратную сторону обновляя значения во втором),
// выполняя константное количество операций для каждого дома.

// По памяти алгоритм требует O(n), так как используется массив расстояний длиной n для хранения результатов.

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	streetStr, _ := reader.ReadString('\n')
	streetStr = strings.TrimSpace(streetStr)
	street := strings.Split(streetStr, " ")

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	distance := make([]int, n)

	distanceNearestEmpty := 1000001
	for house := len(street) - 1; house >= 0; house-- {
		c, _ := strconv.Atoi(street[house])
		if c == 0 {
			distanceNearestEmpty = 0
		} else {
			distanceNearestEmpty++
		}
		distance[house] = distanceNearestEmpty
	}

	distanceNearestEmpty = 1000001
	for house := range street {
		c, _ := strconv.Atoi(street[house])
		if c == 0 {
			distanceNearestEmpty = 0
		} else {
			distanceNearestEmpty++
		}
		if distanceNearestEmpty < distance[house] {
			distance[house] = distanceNearestEmpty
		}
	}

	for house := range distance {
		writer.WriteString(strconv.Itoa(distance[house]) + " ")
	}
}