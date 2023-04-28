package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	kStr, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(kStr))

	m := make(map[rune]int)
	for i := 0; i < 20; i++ {
		c, _, _ := reader.ReadRune()
		if c == '.' || c == '\n' {
			continue
		} else {
			key := c
			if m[key] != 0 {
				m[key]++
			} else {
				m[key]++
			}
		}
	}
	var res int
	for _, value := range m {
		if value <= k+k {
			res++
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(res))
	writer.Flush()
}