package main

import (
	"bufio"
	"fmt"
	"os"
)

type Participant struct {
	Login   string
	Solved  int
	Penalty int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	participants := make([]Participant, n)
	for i := 0; i < n; i++ {
		var login string
		var solved, penalty int
		fmt.Fscanln(reader, &login, &solved, &penalty)
		participants[i] = Participant{Login: login, Solved: solved, Penalty: penalty}
	}

	quickSort(participants, 0, len(participants)-1)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, participant := range participants {
		fmt.Fprintln(writer, participant.Login)
	}
}

func quickSort(arr []Participant, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []Participant, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if compareParticipants(arr[j], pivot) <= 0 {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func compareParticipants(p1, p2 Participant) int {
	if p1.Solved > p2.Solved {
		return -1
	} else if p1.Solved < p2.Solved {
		return 1
	} else {
		if p1.Penalty < p2.Penalty {
			return -1
		} else if p1.Penalty > p2.Penalty {
			return 1
		} else {
			return compareLogins(p1.Login, p2.Login)
		}
	}
}

func compareLogins(login1, login2 string) int {
	if login1 < login2 {
		return -1
	} else if login1 > login2 {
		return 1
	} else {
		return 0
	}
}
