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